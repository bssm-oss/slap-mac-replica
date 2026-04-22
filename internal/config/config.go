package config

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"
	"time"
)

var (
	// ErrHelpRequested asks the caller to print usage and exit zero.
	ErrHelpRequested = errors.New("help requested")
	// ErrVersionRequested asks the caller to print version and exit zero.
	ErrVersionRequested = errors.New("version requested")
)

// Config is the parsed command-line configuration.
type Config struct {
	Command string
	Run     RunConfig
}

// RunConfig holds runtime settings for slap detection.
type RunConfig struct {
	Threshold float64
	Cooldown  time.Duration
	Sound     string
}

const (
	defaultThreshold = 0.05
	defaultCooldown  = 750 * time.Millisecond
	defaultSound     = "Glass"
)

// Parse parses command-line arguments.
func Parse(args []string) (Config, error) {
	if len(args) == 0 {
		return parseRun(nil)
	}

	switch args[0] {
	case "run":
		return parseRun(args[1:])
	case "doctor":
		if hasHelp(args[1:]) {
			return Config{}, ErrHelpRequested
		}
		return Config{Command: "doctor"}, nil
	case "help", "-h", "--help":
		return Config{}, ErrHelpRequested
	case "version", "-v", "--version":
		return Config{}, ErrVersionRequested
	default:
		if strings.HasPrefix(args[0], "-") {
			return parseRun(args)
		}
		return Config{}, fmt.Errorf("unknown command %q", args[0])
	}
}

// Usage returns the command help text.
func Usage(version string) string {
	return fmt.Sprintf(`slap-mac-replica %s

Apple Silicon MacBook 의 숨겨진 가속도계를 읽어 노트북을 칠 때 소리를 재생합니다.

사용법:
  slap-mac-replica run [--threshold 0.05] [--cooldown 750ms] [--sound Glass]
  slap-mac-replica doctor
  slap-mac-replica help
  slap-mac-replica version

설명:
  run     root 권한으로 센서를 읽고 slap 감지를 시작합니다.
  doctor  현재 Mac 이 실행 가능한 환경인지 빠르게 점검합니다.

run 옵션:
  --threshold float    감지 임계값 (기본값: %.2f)
  --cooldown duration  재생 쿨다운 (기본값: %s)
  --sound value        내장 사운드 이름 또는 사용자 파일 경로 (기본값: %s)

예시:
  sudo slap-mac-replica run
  sudo slap-mac-replica run --sound Sosumi
  sudo slap-mac-replica run --sound /path/to/custom.wav
  slap-mac-replica doctor
`, version, defaultThreshold, defaultCooldown, defaultSound)
}

func parseRun(args []string) (Config, error) {
	runCfg := RunConfig{
		Threshold: defaultThreshold,
		Cooldown:  defaultCooldown,
		Sound:     defaultSound,
	}

	fs := flag.NewFlagSet("run", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Float64Var(&runCfg.Threshold, "threshold", runCfg.Threshold, "minimum amplitude threshold in g")
	fs.DurationVar(&runCfg.Cooldown, "cooldown", runCfg.Cooldown, "cooldown between sound triggers")
	fs.StringVar(&runCfg.Sound, "sound", runCfg.Sound, "built-in sound name or custom file path")

	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}

	if runCfg.Threshold <= 0 {
		return Config{}, errors.New("--threshold must be greater than 0")
	}
	if runCfg.Cooldown <= 0 {
		return Config{}, errors.New("--cooldown must be greater than 0")
	}

	return Config{
		Command: "run",
		Run:     runCfg,
	}, nil
}

func hasHelp(args []string) bool {
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			return true
		}
	}
	return false
}
