# 2026-04-23 Bundled Effect Presets

## 배경

효과음 프리셋을 로컬 전용이 아니라 GitHub 다운로드 사용자도 바로 쓸 수 있게 해야 했습니다.

## 문제 또는 목표

- 효과음 MP3를 릴리즈 자산에 포함한다.
- 설치 스크립트가 프리셋도 함께 설치한다.
- Homebrew 설치도 프리셋을 함께 설치한다.
- `slap-mac-replica presets` 가 설치된 프리셋을 기본으로 찾게 한다.

## 변경 내용

- `assets/presets/` 에 효과음 MP3 프리셋을 추가했습니다.
- 릴리즈 tarball 이 `slap-mac-replica` 바이너리와 `presets/` 폴더를 함께 포함하도록 변경했습니다.
- `script/install.sh` 가 프리셋을 `/Library/Application Support/slap-mac-replica/presets` 로 설치하도록 변경했습니다.
- Homebrew formula 가 프리셋을 `share/slap-mac-replica/presets` 로 설치하도록 변경했습니다.
- 앱의 기본 프리셋 탐색 경로를 설치 방식별로 확장했습니다.

## 설계 이유

- 직접 다운로드, Homebrew, 로컬 개발 환경 모두 같은 `presets` 명령으로 동작해야 합니다.
- root 로 실행해도 읽을 수 있는 시스템 경로를 기본 설치 경로로 사용했습니다.

## 영향 범위

- GitHub 릴리즈 자산 크기 증가
- Homebrew 설치 파일 목록
- 프리셋 탐색 경로

## 검증 방법

- `go test ./...`
- `go vet ./...`
- `go build ./cmd/slap-mac-replica`
- `tar -tzf` 로 릴리즈 tarball 에 `presets/` 포함 확인
- 설치 스크립트로 임시 디렉터리에 설치 후 `presets` 실행 확인

## 남아 있는 한계

- 실제 slap 감지는 여전히 root 권한이 필요합니다.
