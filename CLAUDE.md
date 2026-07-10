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

## 행동 규칙

- 클러스터/저장소 상태를 변경하기 전에 항상 현재 상태를 먼저 확인한다.
- 삭제, force-push 등 되돌리기 어려운 작업은 실행 전 반드시 확인받는다.
- kubectl 명령에는 항상 `--context gke-sysnet4admin_book_gitaiops`를 지정한다 (다른 클러스터를 잘못 대상으로 하지 않도록).
