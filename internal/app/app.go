package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"github.com/bssm-oss/slap-mac-replica/internal/audio"
	"github.com/bssm-oss/slap-mac-replica/internal/config"
	"github.com/bssm-oss/slap-mac-replica/internal/platform"
	"github.com/taigrr/apple-silicon-accelerometer/detector"
	"github.com/taigrr/apple-silicon-accelerometer/sensor"
	"github.com/taigrr/apple-silicon-accelerometer/shm"
)

const (
	sensorStartupDelay = 100 * time.Millisecond
	pollInterval       = 10 * time.Millisecond
	maxBatchSize       = 200
)

// Execute runs the requested command.
func Execute(ctx context.Context, cfg config.Config, stdout, stderr io.Writer) error {
	switch cfg.Command {
	case "doctor":
		return runDoctor(stdout)
	case "run":
		return runDetector(ctx, cfg.Run, stdout, stderr)
	default:
		return fmt.Errorf("unsupported command %q", cfg.Command)
	}
}

func runDoctor(stdout io.Writer) error {
	status, err := platform.DetectSensor()
	if err != nil {
		return err
	}

	player, err := audio.NewPlayer("")
	if err != nil {
		return err
	}

	fmt.Fprintf(stdout, "platform: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Fprintf(stdout, "root: %t\n", os.Geteuid() == 0)
	fmt.Fprintf(stdout, "sensor: %t (%s)\n", status.Present, status.Summary)
	fmt.Fprintf(stdout, "default_sound: %s (%s)\n", player.Description(), player.Path())

	switch {
	case runtime.GOOS != "darwin":
		fmt.Fprintln(stdout, "ready: no (macOS 전용 도구입니다.)")
	case runtime.GOARCH != "arm64":
		fmt.Fprintln(stdout, "ready: no (Apple Silicon Mac 에서만 동작합니다.)")
	case !status.Present:
		fmt.Fprintln(stdout, "ready: no (AppleSPUHIDDevice 센서를 찾지 못했습니다.)")
	case os.Geteuid() != 0:
		fmt.Fprintln(stdout, "ready: partial (하드웨어는 준비됐지만 slap 감지는 root 권한으로만 실행됩니다.)")
	default:
		fmt.Fprintln(stdout, "ready: yes")
	}

	return nil
}

func runDetector(ctx context.Context, cfg config.RunConfig, stdout, stderr io.Writer) error {
	if runtime.GOOS != "darwin" {
		return errors.New("slap-mac-replica 는 macOS 전용입니다")
	}
	if runtime.GOARCH != "arm64" {
		return errors.New("slap-mac-replica 는 Apple Silicon Mac 에서만 동작합니다")
	}
	if os.Geteuid() != 0 {
		return errors.New("accelerometer access requires root; run `sudo slap-mac-replica run` or `sudo brew services start slap-mac-replica`")
	}

	status, err := platform.DetectSensor()
	if err != nil {
		return err
	}
	if !status.Present {
		return errors.New("AppleSPUHIDDevice sensor not found; this Mac does not appear to expose the required accelerometer")
	}

	player, err := audio.NewPlayer(cfg.Sound)
	if err != nil {
		return err
	}

	accelRing, err := shm.CreateRing(shm.NameAccel)
	if err != nil {
		return fmt.Errorf("create accelerometer shared memory: %w", err)
	}
	defer accelRing.Close()
	defer accelRing.Unlink()

	sensorErr := make(chan error, 1)
	go func() {
		if runErr := sensor.Run(sensor.Config{
			AccelRing: accelRing,
			Restarts:  0,
		}); runErr != nil {
			sensorErr <- runErr
		}
	}()

	time.Sleep(sensorStartupDelay)

	det := detector.New()
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	var lastTotal uint64
	var lastEventTime time.Time
	var lastPlayback time.Time

	fmt.Fprintf(stdout, "listening for slaps with threshold=%.3fg cooldown=%s sound=%s\n", cfg.Threshold, cfg.Cooldown, player.Description())

	for {
		select {
		case <-ctx.Done():
			fmt.Fprintln(stdout, "shutting down")
			return nil
		case err := <-sensorErr:
			return fmt.Errorf("sensor worker failed: %w", err)
		case <-ticker.C:
		}

		samples, newTotal := accelRing.ReadNew(lastTotal, shm.AccelScale)
		lastTotal = newTotal
		if len(samples) == 0 {
			continue
		}
		if len(samples) > maxBatchSize {
			samples = samples[len(samples)-maxBatchSize:]
		}

		now := time.Now()
		tNow := float64(now.UnixNano()) / 1e9
		for idx, sample := range samples {
			tSample := tNow - float64(len(samples)-idx-1)/float64(det.FS)
			det.Process(sample.X, sample.Y, sample.Z, tSample)
		}

		if len(det.Events) == 0 {
			continue
		}

		event := det.Events[len(det.Events)-1]
		if event.Time.Equal(lastEventTime) {
			continue
		}
		lastEventTime = event.Time

		if event.Amplitude < cfg.Threshold {
			continue
		}
		if since := time.Since(lastPlayback); since < cfg.Cooldown {
			continue
		}

		lastPlayback = now
		fmt.Fprintf(stdout, "slap detected: amplitude=%.5fg severity=%s sound=%s\n", event.Amplitude, event.Severity, player.Description())
		go func() {
			playCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if playErr := player.Play(playCtx); playErr != nil {
				fmt.Fprintf(stderr, "playback failed: %v\n", playErr)
			}
		}()
	}
}
