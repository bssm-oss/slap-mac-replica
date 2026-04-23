# 2026-04-23 Easy GitHub Download

## 배경

Homebrew 없이 GitHub 릴리즈에서 바로 내려받아 설치하는 경로가 필요했습니다.

## 문제 또는 목표

- 사용자가 GitHub 에서 최신 바이너리를 쉽게 받을 수 있게 한다.
- 버전 번호를 몰라도 항상 최신 릴리즈 자산을 다운로드할 수 있게 한다.
- README 에 설치 방법을 명확히 적는다.

## 변경 내용

- `script/install.sh` 원라인 설치 스크립트를 추가했습니다.
- 릴리즈 워크플로가 버전 포함 자산과 함께 고정 이름 자산 `slap-mac-replica_darwin_arm64.tar.gz` 도 업로드하도록 변경했습니다.
- README 에 GitHub 직접 설치와 직접 다운로드 URL 을 추가했습니다.

## 설계 이유

- `releases/latest/download/...` URL 은 버전이 바뀌어도 최신 릴리즈를 가리키므로 사용자가 가장 쉽게 설치할 수 있습니다.
- 설치 스크립트는 macOS/arm64 를 먼저 확인하고 `/usr/local/bin` 에 설치합니다.

## 영향 범위

- GitHub 릴리즈 자산
- README 설치 안내
- 로컬 설치 편의성

## 검증 방법

- `bash -n script/install.sh`
- `tar -tzf` 로 릴리즈 tar 구조 확인
- 새 릴리즈 후 `curl -fsSL .../releases/latest/download/slap-mac-replica_darwin_arm64.tar.gz` 로 다운로드 확인

## 남아 있는 한계

- 실제 slap 감지는 여전히 root 권한이 필요합니다.
- installer 는 Apple Silicon macOS 만 지원합니다.
