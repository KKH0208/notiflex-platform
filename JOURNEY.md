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
| ch4 | 4.2 메트릭 모니터링 | ⬜ | | (진행 중) kube-prometheus-stack 설치 및 Grafana UI 접속 확인 완료. 남은 것: Prometheus 수집 확인, Notiflex 대시보드 생성 |
| ch4 | 4.3 로그 수집 | ✅ | 2026-07-15 | Loki(SingleBinary) + Fluent Bit 설치, Grafana Loki 데이터소스 등록, notiflex 네임스페이스 로그 실제 수집 확인 |
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
| 로그 수집 도구 (ch4.3) | Loki + Fluent Bit | ELK Stack, Google Cloud Logging | 경량(Loki 128Mi, Fluent Bit 64Mi)이라 e2-medium에서 감당 가능, Grafana에 이미 통합. ELK는 최소 2Gi 필요해 불가능 |

## 현재 버전

| 컴포넌트 | 버전 | 변경 이력 |
|---------|------|----------|
| Go | 1.25 | 2026-07-10 최초 설정 (ch6 valkey-go, ch8 OTel SDK 요구사항 대비 처음부터 1.25로 시작) |
| Notiflex 이미지 | sha-c870e08 | 2026-07-10 v0.1.0 → 2026-07-11 v0.1.1(/version) → 2026-07-14 CI-CD 연결 후부터 이미지 태그는 git SHA 기반(sha-xxxxxxx)으로 전환, /ping 엔드포인트 추가 |
| ArgoCD | v3.4.5 | 2026-07-11 최초 설치 |
| kube-prometheus-stack | 87.15.2 (app v0.92.1) | 2026-07-14 최초 설치 |
| Loki | chart 7.0.0 (app v3.6.7) | 2026-07-15 최초 설치, SingleBinary 모드, replication_factor=1 |
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
| ch4.3 | `helm install loki` 실패: "Cannot run scalable targets... without object storage backend" | 최신 loki 차트는 `deploymentMode: SingleBinary`만으로 backend/read/write replicas가 자동으로 0이 되지 않음. `singleBinary.replicas: 1`과 `backend/read/write.replicas: 0`을 명시적으로 지정 |
| ch4.3 | `loki-chunks-cache`, `loki-results-cache` 파드가 계속 Pending (CPU/메모리 부족) | 최신 loki 차트가 기본으로 memcached 캐시 2종을 활성화함(문서에 언급 없음). `chunksCache.enabled: false`, `resultsCache.enabled: false`로 비활성화 |
| ch4.3 | Loki 쿼리 시 "too many unhealthy instances in the ring" | 기본 `replication_factor: 3`인데 인스턴스가 1개뿐이라 쿼럼 부족. `loki.commonConfig.replication_factor: 1`로 낮춤 |
| ch4.3 | Fluent Bit → Loki 로그 전송 안 됨, `dial tcp: lookup fluent-bit-loki: no such host` | `grafana/fluent-bit` 차트는 `config.outputs`가 아니라 전용 `loki.serviceName/servicePort/servicePath` 값으로 output을 설정해야 함(값이 조용히 무시됨). `loki.serviceName: loki-gateway`, `servicePort: 80`, `servicePath: /loki/api/v1/push`로 수정 |
| ch4.3 | Grafana에 Loki 데이터소스가 자동 등록되지 않음 | loki 차트의 `grafana.datasource` 값은 차트 자체 번들 Grafana 대상이라 외부 kube-prometheus-stack Grafana에는 적용 안 됨. `grafana_datasource: "1"` 라벨을 가진 ConfigMap(`k8s/monitoring/loki-datasource.yaml`)을 직접 만들어 사이드카가 인식하게 함 |
