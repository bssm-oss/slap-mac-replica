# 2026-04-23 Local Effect Presets

## 배경

`/Users/heodongun/Desktop/효과음` 폴더 안의 음원을 쉽게 선택할 수 있는 프리셋 기능이 필요했습니다.

## 문제 또는 목표

- 효과음 폴더의 오디오 파일을 이름으로 선택할 수 있게 한다.
- 프리셋 목록을 CLI 로 확인할 수 있게 한다.
- 음원 파일은 저장소나 릴리즈에 포함하지 않는다.

## 변경 내용

- `slap-mac-replica presets` 명령을 추가했습니다.
- `run --preset <name>` 옵션을 추가했습니다.
- `run --preset random` 으로 효과음 폴더에서 랜덤 선택할 수 있게 했습니다.
- `--preset-dir` 로 프리셋 폴더를 바꿀 수 있게 했습니다.

## 설계 이유

- 로컬 음원을 저장소에 포함하지 않으면서도 쉽게 선택하려면 경로 대신 이름 기반 프리셋이 가장 단순합니다.
- `afplay` 가 재생 가능한 파일 확장자만 프리셋으로 노출합니다.

## 영향 범위

- CLI 옵션
- README 사용법
- 프리셋 탐색 로직

## 검증 방법

- `go test ./...`
- `go vet ./...`
- `go build ./cmd/slap-mac-replica`
- `slap-mac-replica presets`
- `slap-mac-replica run --preset op-gangnam-style`

## 남아 있는 한계

- 실제 slap 감지는 여전히 root 권한이 필요합니다.
- 효과음 폴더가 없는 다른 Mac 에서는 `--preset-dir` 를 지정해야 합니다.
