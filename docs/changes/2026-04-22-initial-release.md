# 2026-04-22 Initial Release

## 배경

원격 저장소가 비어 있었고, Apple Silicon MacBook 을 때렸을 때 소리가 나는 slap 감지 도구가 필요했습니다.

## 문제 또는 목표

- 실제 slap 감지를 할 수 있는 최소 제품을 만든다.
- 설치와 상시 실행 경로를 Homebrew 중심으로 단순화한다.
- 하드웨어와 권한 제약을 문서에서 숨기지 않는다.

## 변경 내용

- Go 기반 `slap-mac-replica` CLI 를 추가했습니다.
- `doctor` 명령으로 센서 존재 여부를 점검할 수 있게 했습니다.
- `run` 명령으로 root 권한 slap 감지 루프와 사운드 재생을 구현했습니다.
- Homebrew formula 와 `brew services` service 정의를 추가했습니다.
- GitHub release 바이너리 자산 기준 sha256 과 tap 설치 흐름을 맞췄습니다.
- 한국어 README 와 AI 작업 지침용 `AGENTS.md` 를 작성했습니다.
- macOS GitHub Actions CI 를 추가했습니다.

## 설계 이유

- GUI 보다 CLI + Homebrew service 가 구현과 배포가 단순합니다.
- 센서 접근은 undocumented IOKit HID 경로라서 root 요구사항을 우회하지 않고 그대로 드러냈습니다.
- 오디오 자산을 별도 포함하지 않고 macOS 내장 사운드를 기본값으로 사용해 초기 릴리즈를 단순화했습니다.

## 영향 범위

- Apple Silicon MacBook 사용자 설치/실행 흐름
- GitHub Actions CI
- Homebrew 설치 경로
- 문서와 온보딩

## 검증 방법

- `go test ./...`
- `go vet ./...`
- `go build ./cmd/slap-mac-replica`
- `./slap-mac-replica doctor`
- `./slap-mac-replica run` 으로 root 필요 에러 메시지 확인

## 남아 있는 한계

- 현재 환경에는 관리자 비밀번호가 없어 실제 root slap 감지 루프를 자동 실행하지 못했습니다.
- GUI, menubar, privileged helper 는 아직 없습니다.

## 후속 과제

- 관리자 프롬프트가 포함된 root helper 설치 흐름
- 사용자 지정 설정 파일
- 다중 사운드 프로필
- 실제 하드웨어별 민감도 튜닝 데이터 축적
