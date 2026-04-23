# slap-mac-replica

`slap-mac-replica` 는 Apple Silicon MacBook 의 숨겨진 `AppleSPUHIDDevice` 가속도계를 읽어서, 노트북을 실제로 쳤을 때 소리를 재생하는 macOS CLI 도구입니다.

이 저장소는 "Slap Mac" 류의 앱이 하는 핵심 동작을 오픈소스 기준으로 재현하는 것을 목표로 합니다. 현재 버전은 GUI 대신 단일 바이너리와 Homebrew 서비스 흐름에 집중합니다.

## 무엇을 해결하나

- Apple Silicon MacBook 에서 실제 물리 충격을 감지하고 싶다
- 외부 앱 없이, 설치와 실행 경로가 단순한 slap 감지 도구가 필요하다
- `brew install` 과 `brew services` 기반으로 쉽게 올리고 내리고 싶다

## 핵심 기능

- Apple Silicon MacBook 의 숨겨진 가속도계 존재 여부 점검 (`doctor`)
- 루트 권한으로 slap 감지 루프 실행 (`run`)
- 기본 한국어 음성 반응 또는 사용자 지정 오디오 파일 재생
- GitHub 릴리즈/Homebrew 설치 시 효과음 프리셋 함께 제공
- `brew services` 로 root LaunchDaemon 형태로 상시 실행 가능

## 기술 스택

- Go 1.26
- `github.com/taigrr/apple-silicon-accelerometer`
- macOS IOKit HID
- GitHub Actions
- Homebrew Formula

## 요구 환경 / 사전 준비사항

- macOS
- Apple Silicon MacBook
- 관리자 권한
- Homebrew (선택 사항이지만 권장)

## 빠른 다운로드 / 실행

터미널에 아래 한 줄을 붙여 넣으면 GitHub 최신 릴리즈 바이너리를 바로 설치합니다.

```bash
curl -fsSL https://raw.githubusercontent.com/bssm-oss/slap-mac-replica/main/script/install.sh | bash
```

설치 확인:

```bash
slap-mac-replica doctor
```

바로 실행:

```bash
sudo slap-mac-replica run
```

기본 모드에서는 짧게 치면 `오빠 강남스타일`, 짧은 시간 안에 연속으로 3회 이상 치면 `예~~~` 를 말합니다.

직접 준비한 오디오 클립을 쓰려면:

```bash
sudo slap-mac-replica run \
  --short-sound /absolute/path/to/oppa.wav \
  --rapid-sound /absolute/path/to/yeah.wav
```

`--short-sound` 는 짧게 칠 때 나오는 파일이고, `--rapid-sound` 는 연속으로 칠 때 나오는 파일입니다. `afplay` 가 재생할 수 있는 `wav`, `aiff`, `mp3` 등을 사용할 수 있습니다.

기본 포함 프리셋을 쓰려면:

```bash
slap-mac-replica presets
sudo slap-mac-replica run --preset op-gangnam-style
sudo slap-mac-replica run --preset random
```

## 설치 방법

### 1. GitHub 에서 바로 설치

가장 쉬운 방법입니다. GitHub 최신 릴리즈에서 바이너리를 바로 내려받아 `/usr/local/bin/slap-mac-replica` 에 설치합니다.

```bash
curl -fsSL https://raw.githubusercontent.com/bssm-oss/slap-mac-replica/main/script/install.sh | bash
```

설치 뒤 확인:

```bash
slap-mac-replica doctor
```

실행:

```bash
sudo slap-mac-replica run
```

수동으로 직접 다운로드하려면 아래 파일을 받으면 됩니다.

```text
https://github.com/bssm-oss/slap-mac-replica/releases/latest/download/slap-mac-replica_darwin_arm64.tar.gz
```

GitHub 웹사이트에서 직접 받으려면:

1. <https://github.com/bssm-oss/slap-mac-replica/releases/latest> 로 이동합니다.
2. `Assets` 를 펼칩니다.
3. `slap-mac-replica_darwin_arm64.tar.gz` 를 다운로드합니다.
4. 압축을 풀고 `slap-mac-replica` 바이너리를 실행합니다.

### 2. Homebrew 로 설치

릴리즈가 올라가 있으면 아래 명령으로 설치할 수 있습니다.

```bash
brew tap bssm-oss/slap-mac-replica https://github.com/bssm-oss/slap-mac-replica
brew install bssm-oss/slap-mac-replica/slap-mac-replica
```

설치 뒤 환경 점검:

```bash
slap-mac-replica doctor
```

즉시 실행:

```bash
sudo slap-mac-replica run
```

부팅 후에도 계속 실행:

```bash
sudo brew services start slap-mac-replica
```

중지:

```bash
sudo brew services stop slap-mac-replica
```

### 3. 소스에서 직접 실행

```bash
go build ./cmd/slap-mac-replica
./slap-mac-replica doctor
sudo ./slap-mac-replica run
```

## 환경변수 설명

현재 필수 환경변수는 없습니다.

## 로컬 실행 방법

환경 점검:

```bash
slap-mac-replica doctor
```

기본 음성 모드로 실행:

```bash
sudo slap-mac-replica run
```

기본 모드에서는 짧게 치면 `오빠 강남스타일`, 짧은 시간 안에 연속으로 3회 이상 치면 `예~~~` 를 말합니다.

다른 내장 사운드 사용:

```bash
sudo slap-mac-replica run --sound Sosumi
```

사용자 지정 오디오 파일 사용:

```bash
sudo slap-mac-replica run --sound /absolute/path/to/custom.wav
```

기본 강남 음성 모드를 명시적으로 다시 켜기:

```bash
sudo slap-mac-replica run --sound gangnam
```

권리가 있는 로컬 오디오 클립을 짧은 slap/연속 slap 에 각각 쓰기:

```bash
sudo slap-mac-replica run \
  --short-sound /absolute/path/to/oppa.wav \
  --rapid-sound /absolute/path/to/yeah.wav
```

저작권이 있는 원곡에서 무단 추출한 클립을 저장소에 포함하거나 배포하지 마십시오. 직접 녹음했거나 사용 권리가 있는 파일만 지정해야 합니다.

포함된 효과음 프리셋 선택하기:

```bash
slap-mac-replica presets
sudo slap-mac-replica run --preset op-gangnam-style
sudo slap-mac-replica run --preset gopgopgop
sudo slap-mac-replica run --preset random
```

GitHub 릴리즈와 Homebrew 설치본에는 프리셋 MP3가 함께 들어갑니다. 다른 폴더를 쓰려면 `--preset-dir` 를 지정합니다.

```bash
slap-mac-replica presets --preset-dir /absolute/path/to/effects
sudo slap-mac-replica run --preset random --preset-dir /absolute/path/to/effects
```

예를 들어 직접 준비한 두 클립이 다운로드 폴더에 있다면:

```bash
sudo slap-mac-replica run \
  --short-sound "$HOME/Downloads/oppa.wav" \
  --rapid-sound "$HOME/Downloads/yeah.wav"
```

임계값과 쿨다운 조정:

```bash
sudo slap-mac-replica run --threshold 0.08 --cooldown 1s
```

## 테스트 실행 방법

```bash
go test ./...
go vet ./...
go build ./cmd/slap-mac-replica
```

## 주요 스크립트 / 명령 설명

- `slap-mac-replica doctor`: 현재 Mac 이 slap 감지 가능한 하드웨어인지 점검
- `slap-mac-replica run`: slap 감지 루프 시작
- `brew services start slap-mac-replica`: root LaunchDaemon 으로 등록해 상시 실행

## 기본 음성 반응

- 짧은 slap: `오빠 강남스타일`
- 짧은 시간 안의 연속 slap 3회 이상: `예~~~`
- 권리가 있는 로컬 클립을 쓰려면 `--short-sound`, `--rapid-sound` 를 지정합니다.
- 데스크톱 `효과음` 폴더 안의 파일은 `--preset <이름>` 으로 쉽게 선택할 수 있습니다.

## 지원하는 내장 사운드

- `Basso`
- `Blow`
- `Bottle`
- `Frog`
- `Funk`
- `Glass`
- `Hero`
- `Morse`
- `Ping`
- `Pop`
- `Purr`
- `Sosumi`
- `Submarine`
- `Tink`

## 폴더 구조

```text
cmd/slap-mac-replica/      메인 엔트리포인트
internal/app/              실행 루프와 doctor 명령
internal/audio/            사운드 경로 해석과 재생
internal/config/           CLI 파싱
internal/platform/         하드웨어 점검 유틸리티
Formula/                   Homebrew formula
.github/workflows/         CI
docs/changes/              변경 기록
```

## 아키텍처 개요

1. `doctor` 는 `ioreg` 를 읽어 `AppleSPUHIDDevice` 존재 여부를 점검합니다.
2. `run` 은 root 권한으로 `sensor.Run(...)` 을 띄워 IOKit HID 센서 데이터를 shared memory 로 받습니다.
3. 이벤트 루프는 shared memory 의 새로운 가속도계 샘플을 읽습니다.
4. `detector.New()` 기반 감지기가 slap 이벤트와 amplitude 를 계산합니다.
5. amplitude 가 임계값을 넘으면 기본 모드에서는 `say`, 다른 사운드 모드에서는 `afplay` 로 선택한 오디오를 재생합니다.

## 개발 원칙

- 구현보다 검증을 우선합니다.
- 실제 하드웨어 제약을 문서로 숨기지 않습니다.
- 루트 권한이 필요한 이유를 코드와 문서에 명시합니다.
- macOS 전용 제약은 런타임과 문서 양쪽에서 모두 드러냅니다.

## 기여 방법

1. 브랜치를 만듭니다.
2. 테스트를 추가하거나 수정합니다.
3. `go test ./...`, `go vet ./...`, `go build ./cmd/slap-mac-replica` 를 실행합니다.
4. 문서가 실제 동작과 일치하는지 확인합니다.
5. Pull Request 를 엽니다.

## CI 개요

GitHub Actions 가 macOS 러너에서 아래 항목을 검증합니다.

- `go test ./...`
- `go vet ./...`
- `go build ./cmd/slap-mac-replica`

## 알려진 제한 사항

- Apple Silicon MacBook 에서만 동작합니다.
- 센서 접근이 문서화되지 않은 IOKit HID 경로라서 root 권한이 필요합니다.
- 현재는 GUI 가 아니라 CLI + Homebrew 서비스 모델입니다.
- 이 환경에서는 관리자 비밀번호가 없어 실제 root slap 감지를 자동 검증할 수 없었습니다.
- `brew services` 로 slap 감지를 켜려면 `sudo brew services start slap-mac-replica` 처럼 root 로 실행해야 합니다.

## 향후 계획 / 로드맵

- 관리자 프롬프트를 동반한 root helper 설치 흐름
- 다중 사운드 프로필과 사용자 설정 파일
- 더 정교한 민감도 보정
- GUI 래퍼 또는 menubar 앱
