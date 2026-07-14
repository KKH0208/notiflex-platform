# Notiflex Platform

## 프로젝트 개요

Notiflex — B2B 알림 SaaS 플랫폼. 「AI 시대에 개발자가 알아야 하는 인프라 구성 배포 with 클로드 코드」 실습으로 구축하는 운영 저장소.

## 기술 스택

- **언어**: Go 표준 라이브러리 (외부 프레임워크 없음)
- **컨테이너**: scratch 베이스 이미지

## GCP 설정

- **프로젝트 ID**: `kihyun-gitaiops-project`
- **리전**: `asia-northeast3` (서울)
- **존**: `asia-northeast3-a`

## Artifact Registry

- 이미지 경로: `asia-northeast3-docker.pkg.dev/kihyun-gitaiops-project/notiflex`

## GitOps / CI-CD

- **배포**: ArgoCD Application `notiflex-smb` (namespace: argocd)가 `k8s/smb` 경로를 `notiflex` 네임스페이스에 automated + prune + selfHeal로 동기화한다. `kubectl apply`로 클러스터를 직접 수정하지 않는다.
- **이미지 빌드/배포 자동화**: `.github/workflows/ci.yaml`이 `app/**` 변경 push 시 이미지를 빌드/푸시하고, `k8s/smb/deployment.yaml`의 이미지 태그를 git SHA 기준(`sha-xxxxxxx`)으로 자동 커밋한다. **이미지 태그는 CI가 관리하므로 deployment.yaml의 이미지 태그를 수동으로 편집하지 않는다** — 다음 CI 실행 시 충돌(merge conflict)이 발생할 수 있다.
- CI 인증: 전용 서비스 계정 `notiflex-ci`(`roles/artifactregistry.writer`만 부여) 키를 GitHub Secrets(`GCP_SA_KEY`, `GCP_PROJECT_ID`)에 저장. 저장소 Actions 권한은 `default_workflow_permissions=write`로 설정되어 있음 (CI가 매니페스트를 커밋/푸시하기 위해 필요).

## 관측 가능성

- **메트릭**: `monitoring` 네임스페이스에 kube-prometheus-stack (Prometheus + Grafana + Alertmanager) 설치. CPU requests는 `helm-values/kube-prometheus.yaml`에서 축소 설정됨 — **6장 진입 전 추가로 5m 수준까지 낮춰야 함** (CSI DaemonSet 240m 추가 대비).
- **로그**: Loki(SingleBinary) + Fluent Bit(DaemonSet)로 클러스터 전체(시스템 네임스페이스 포함) 로그 수집. Grafana에 `Loki` 데이터소스 등록됨. 조회는 `{namespace="notiflex"}` 같은 라벨 필터로 좁혀서 사용.
  - Loki는 `replication_factor: 1`, `chunksCache`/`resultsCache` 비활성화 상태(리소스 제약) — 나중에 차트 업그레이드 시 이 값들이 초기화되지 않았는지 확인할 것.
- **알림**: `k8s/monitoring/pod-restart-alert.yaml`의 `PrometheusRule`로 파드 재시작 과다 감지. **Slack 등 외부 Contact Point는 연결하지 않음** — Alertmanager 내부에서만 알림이 발생·표시된다(의도적 결정).
- 새 알림 규칙 추가 시 `metadata.labels.release: kube-prometheus`를 반드시 포함해야 Prometheus가 규칙을 로드한다.

## 행동 규칙

- 클러스터/저장소 상태를 변경하기 전에 항상 현재 상태를 먼저 확인한다.
- 삭제, force-push 등 되돌리기 어려운 작업은 실행 전 반드시 확인받는다.
- kubectl 명령에는 항상 `--context gke-sysnet4admin_book_gitaiops`를 지정한다 (다른 클러스터를 잘못 대상으로 하지 않도록).
