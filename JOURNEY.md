# Notiflex 여정 기록

이 파일은 독자가 실제로 진행한 내용을 기록한다. AI가 각 챕터 완료 시 자동으로 업데이트한다.

## 진행 현황

| 챕터 | 서브챕터 | 상태 | 완료일 | 비고 |
|------|---------|------|--------|------|
| ch2 | 2.2 설치 확인 | ✅ | 2026-07-10 | Claude Code 2.1.201, jq 설치, statusline 구성 |
| ch2 | 2.3 gcloud 설정 | ✅ | 2026-07-10 | 프로젝트 kihyun-gitaiops-project, 리전 asia-northeast3 |
| ch2 | 2.4 GitHub 저장소 | ✅ | 2026-07-10 | github.com/KKH0208/notiflex-platform (private) |
| ch2 | 2.5 GKE 클러스터 | ✅ | 2026-07-10 | notiflex-cluster, Standard/Zonal, e2-medium x2 Spot |
| ch2 | 2.6 빌드/배포 | ✅ | 2026-07-10 | notiflex-api v0.1.0, /health, /id 엔드포인트 |
| ch2 | 2.7 첫 커밋 | ✅ | 2026-07-10 | |
| ch3 | 3.2 GitOps 도구 | ✅ | 2026-07-11 | ArgoCD 설치, notiflex-smb Application 생성 (public repo, 별도 인증 불필요) |
| ch3 | 3.3 기능 추가 | ✅ | 2026-07-11 | /version 엔드포인트 추가, v0.1.1 롤링 업데이트 |
| ch3 | 3.4 CI | ✅ | 2026-07-14 | GitHub Actions 방식 A(docker build+push), notiflex-ci SA 생성, GCP_SA_KEY/GCP_PROJECT_ID secret 등록 |
| ch3 | 3.5 CI-CD 연결 | ✅ | 2026-07-14 | CI가 이미지 빌드 후 deployment.yaml 태그를 sha 기반으로 자동 커밋/푸시, ArgoCD가 감지해 배포. 엔드투엔드(/ping) 검증 완료 |
| ch4 | 4.2 메트릭 모니터링 | ⬜ | | |
| ch4 | 4.3 로그 수집 | ⬜ | | |
| ch4 | 4.4 알림 | ⬜ | | |
| ch5 | 5.2 트래픽 관리 | ⬜ | | |
| ch5 | 5.3 무중단 배포 | ⬜ | | |
| ch6 | 6.1 캐시 | ⬜ | | |
| ch6 | 6.2 시크릿 관리 | ⬜ | | |
| ch6 | 6.3 Canary 전환 | ⬜ | | |
| ch7 | 7.2 멀티 노드풀 | ⬜ | | |
| ch7 | 7.3 App of Apps | ⬜ | | |
| ch7 | 7.4 멀티테넌시 | ⬜ | | |
| ch8 | 8.1 메시징 | ⬜ | | |
| ch8 | 8.2 트레이싱 | ⬜ | | |
| ch8 | 8.3 CronJob | ⬜ | | |
| ch9 | 9.1 저장소 분석 | ⬜ | | |
| ch9 | 9.2 회고 | ⬜ | | |
| ch9 | 9.3 온보딩 문서 | ⬜ | | |
| ch9 | 9.4 GitAIOps 분석 | ⬜ | | |
| ch9 | 9.5 마무리 | ⬜ | | |

## 도구 선택 기록

독자가 3-프롬프트 패턴(탐색→비교→실행)에서 실제로 선택한 도구와 이유를 기록한다.

| 영역 | 선택 | 검토한 대안 | 선택 이유 |
|------|------|-----------|----------|
| GitOps 도구 (ch3.2) | ArgoCD | Flux, Jenkins X, Spinnaker | Web UI로 배포 상태를 시각적으로 확인 가능, e2-medium 2노드에서 리소스(~500MB) 감당 가능. Flux는 가볍지만 UI 없어서 학습 목적에 부적합 |
| CI 도구 (ch3.4) | GitHub Actions | Cloud Build, GitLab CI, Jenkins | GitHub 네이티브라 별도 서버 불필요, YAML 선언적 파이프라인, 프라이빗도 월 2,000분 무료. Cloud Build는 GitHub 트리거 설정이 별도 필요, Jenkins는 서버 운영 부담 |

## 현재 버전

| 컴포넌트 | 버전 | 변경 이력 |
|---------|------|----------|
| Go | 1.25 | 2026-07-10 최초 설정 (ch6 valkey-go, ch8 OTel SDK 요구사항 대비 처음부터 1.25로 시작) |
| Notiflex 이미지 | sha-c870e08 | 2026-07-10 v0.1.0 → 2026-07-11 v0.1.1(/version) → 2026-07-14 CI-CD 연결 후부터 이미지 태그는 git SHA 기반(sha-xxxxxxx)으로 전환, /ping 엔드포인트 추가 |
| ArgoCD | v3.4.5 | 2026-07-11 최초 설치 |
| Kafka | | |
| OTel SDK | | |

## 현재 리소스

| 노드풀 | 머신 타입 | 노드 수 | 주요 워크로드 |
|--------|----------|---------|-------------|
| default-pool | e2-medium (Spot) | 2 | notiflex-api (replicas: 2) |

## 트러블슈팅 이력

독자가 겪은 문제와 해결 방법을 기록한다. 같은 문제를 다시 겪지 않도록 한다.

| 챕터 | 문제 | 해결 |
|------|------|------|
| ch2.6 | `gcloud builds submit` 실행 시 `PERMISSION_DENIED` (Cloud Build API를 막 활성화한 직후) | IAM 전파 지연 문제. 약 1분 대기 후 재시도하니 정상 성공 |
| ch3.4 | `.github/workflows/` 파일 push 시 `refusing to allow an OAuth App to create or update workflow` 거부 | gh 인증 토큰에 `workflow` 스코프 없음 → `gh auth refresh -h github.com -s workflow`로 스코프 추가 |
| ch3.5 | CI가 생성한 커밋이 push되지 않고 실패 우려 | 저장소 `Settings → Actions → Workflow permissions`가 기본 `Read`였음 → `default_workflow_permissions=write`로 변경 |
