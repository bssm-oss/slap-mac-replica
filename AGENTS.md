# AGENTS.md

## 프로젝트 목적

Apple Silicon MacBook 의 숨겨진 가속도계를 이용해 노트북을 칠 때 소리를 재생하는 macOS 도구를 유지보수합니다.

## 빠른 시작 명령

```bash
go test ./...
go vet ./...
go build ./cmd/slap-mac-replica
./slap-mac-replica doctor
sudo ./slap-mac-replica run
```

## 설치 / 실행 / 테스트 명령

- 설치:
  `brew tap bssm-oss/slap-mac-replica https://github.com/bssm-oss/slap-mac-replica`
  `brew install bssm-oss/slap-mac-replica/slap-mac-replica`
- 환경 점검: `slap-mac-replica doctor`
- 실행: `sudo slap-mac-replica run`
- 서비스 실행: `sudo brew services start slap-mac-replica`
- 서비스 중지: `sudo brew services stop slap-mac-replica`
- 테스트: `go test ./...`
- 정적 검사: `go vet ./...`
- 빌드: `go build ./cmd/slap-mac-replica`

## 기본 작업 순서

1. `README.md`, `AGENTS.md`, `docs/` 를 먼저 읽습니다.
2. `go test ./...`, `go vet ./...`, `go build ./cmd/slap-mac-replica` 로 현재 상태를 확인합니다.
3. 변경 범위를 최소화합니다.
4. 코드 변경과 테스트 변경을 같이 합니다.
5. 문서를 업데이트합니다.
6. 최종적으로 같은 검증 명령을 다시 실행합니다.

## 정의된 완료 조건

- 요청한 동작이 코드에 반영됨
- 변경한 로직에 대응하는 테스트가 있음
- `go test ./...` 통과
- `go vet ./...` 통과
- `go build ./cmd/slap-mac-replica` 통과
- README / AGENTS / docs 가 최신 상태
- root 권한, 하드웨어 제약, 미해결 항목이 문서에 명시됨

## 코드 스타일 원칙

- 가능한 표준 라이브러리를 우선 사용합니다.
- 런타임 제약은 에러 메시지로 직접 드러냅니다.
- undocumented macOS 센서 제약은 숨기지 않습니다.
- 순수 로직은 테스트 가능한 작은 함수로 분리합니다.

## 파일 구조 원칙

- `cmd/` 에 엔트리포인트를 둡니다.
- `internal/` 에 실행 로직, 플랫폼 코드, 오디오, 설정 파서를 분리합니다.
- `Formula/` 는 Homebrew 관련 파일만 둡니다.
- `docs/changes/` 에 작업 단위별 변경 기록을 남깁니다.

## 문서화 원칙

- 실행 방법과 실제 제약이 어긋나면 안 됩니다.
- root 권한 필요 여부를 모든 설치/실행 문서에 명확히 적습니다.
- Apple Silicon 제약은 README 와 AGENTS 모두에 남깁니다.

## 테스트 원칙

- 순수 로직은 단위 테스트로 보호합니다.
- 하드웨어/루트 의존 동작은 자동화 가능 범위까지 검증하고, 막힌 이유를 문서화합니다.
- 수동 검증이 불가능하면 왜 불가능한지와 대체 검증을 기록합니다.

## 브랜치 / 커밋 / PR 규칙

- 기본 브랜치에 직접 기능 작업하지 않습니다.
- 커밋은 구현 / 테스트 / 문서 단위로 나눕니다.
- PR 본문에는 배경, 변경 요약, 테스트 결과, 수동 검증 결과, 리스크를 포함합니다.
- 예외: 빈 원격 저장소에서는 PR 베이스 생성을 위해 초기화 커밋이 필요할 수 있습니다.

## 민감한 경로 / 수정 주의 경로

- `internal/app/app.go`: root 권한, 센서 루프, shared memory 제어
- `Formula/slap-mac-replica.rb`: Homebrew 설치 및 서비스 동작
- `.github/workflows/ci.yml`: 릴리즈 품질 검증 경로

## 작업 전 체크리스트

- 현재 브랜치 확인
- 현재 변경사항 확인
- 하드웨어/권한 제약 확인
- 관련 문서 확인

## 작업 후 체크리스트

- 포맷팅 반영
- 테스트 / 정적 검사 / 빌드 실행
- README / AGENTS / docs 업데이트
- git diff 재검토

## 절대 하면 안 되는 것

- 실행하지 않은 테스트를 통과했다고 쓰지 말 것
- root 가 필요한 동작을 root 없이 검증했다고 말하지 말 것
- Apple Silicon 제약을 무시하고 범용 macOS 도구처럼 문서화하지 말 것
- 문서와 실제 명령을 어긋나게 두지 말 것
