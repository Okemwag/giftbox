# Giftbox Loyalty & Retention Platform

Giftbox is a multi-tenant loyalty and retention platform for real-world merchants. It connects customer enrolment, consent, purchases, M-Pesa transaction ingestion, WhatsApp engagement, auditable loyalty ledgers, reward automation, campaign workflows, and profitability reporting into one merchant operating system.

The first version is intentionally designed as a **Go modular monolith** with independently deployable API, webhook, and worker processes. This keeps delivery fast while preserving clean domain boundaries that can later be extracted into services when volume and team size justify it.

> A transaction occurs -> customer identity is resolved -> loyalty balance changes -> customer behaviour is evaluated -> an appropriate message or reward is triggered -> the next purchase measures campaign effectiveness.

## Table of Contents

- [Why Giftbox Exists](#why-giftbox-exists)
- [Architecture Summary](#architecture-summary)
- [System Diagrams](#system-diagrams)
- [Product Capabilities](#product-capabilities)
- [Repository Structure](#repository-structure)
- [Deployable Applications](#deployable-applications)
- [Core Domain Modules](#core-domain-modules)
- [Technology Direction](#technology-direction)
- [Local Development](#local-development)
- [Configuration](#configuration)
- [API Surface](#api-surface)
- [Data Architecture](#data-architecture)
- [Event-Driven Processing](#event-driven-processing)
- [Security and Privacy](#security-and-privacy)
- [Reliability Requirements](#reliability-requirements)
- [Testing](#testing)
- [Deployment Direction](#deployment-direction)
- [Product Roadmap](#product-roadmap)
- [Engineering Standards](#engineering-standards)

## Why Giftbox Exists

Many merchants can record sales, but very few can reliably answer:

- Which customers are new, returning, at risk, or VIP?
- Which transactions should earn loyalty points or rewards?
- Which rewards actually produce profitable repeat purchases?
- Which customers consented to receive WhatsApp marketing?
- Which staff action changed a loyalty balance, reward, campaign, or integration?
- Which payment callback was missed, duplicated, delayed, or reconciled later?

Giftbox solves this by becoming a transaction-linked retention layer for merchants. The platform does not begin as a payment processor or a broad consumer marketplace. It starts as a practical merchant SaaS product:

1. Capture purchases from M-Pesa, CSV, POS, e-commerce, or manual entry.
2. Identify customers with consent.
3. Maintain an auditable points and rewards ledger.
4. Automate retention campaigns through WhatsApp.
5. Measure incremental profitability, not vanity redemptions.

## Architecture Summary

The architecture uses a modular monolith in Go with three deployables:

| Deployable | Responsibility |
| --- | --- |
| `api` | Merchant dashboard APIs, tenant management, customer enrolment, loyalty configuration, reporting |
| `webhook-gateway` | Daraja/M-Pesa callbacks, WhatsApp webhooks, future POS/e-commerce callbacks |
| `worker` | Outbox processing, campaign execution, reconciliation, reward expiry, notification delivery, reporting jobs |

All deployables share domain packages and PostgreSQL storage, but can scale independently.

The system is built around these principles:

- Payments cannot be lost or duplicated.
- Points must be auditable.
- WhatsApp delivery must be asynchronous and retryable.
- Campaigns can run for days or weeks and need durable workflows.
- Merchant data must remain tenant-scoped.
- Customer identity and purchase behaviour must be protected by consent and encryption.
- The MVP must remain simple enough for a small team to ship.

## System Diagrams

### High-Level Architecture

![Giftbox high-level architecture](./giftbox_architecture.png)

### Merchant Onboarding Journey

![Merchant onboarding journey](./merchant_onboarding_journey.png)

### Customer Registration and Consent Journey

![Customer registration journey](./customers_registration_journey.png)

### Payment Ingestion Flow

![Payment ingestion flow](./payment_ingestion_flow.png)

### WhatsApp Message Architecture

![Message architecture](./message_architecture.png)

### Transactional Outbox Pattern

![Transactional outbox pattern](./transactional_outbox_pattern.png)

## Product Capabilities

### Merchant Management

- Merchant tenants
- Branches and outlets
- Staff users
- Role-based access control
- Integration settings
- Audit trail for sensitive actions

### Customer Identity and Consent

- Customer profiles
- Verified phone identities
- WhatsApp identity linkage
- Merchant memberships
- Consent records by purpose
- Opt-out and preference management
- Customer merge rules for duplicate identities

### Transaction Ingestion

- M-Pesa payment callbacks
- M-Pesa reconciliation jobs
- CSV imports for pilot merchants
- Manual transaction entry with approval/audit controls
- Future POS and e-commerce integrations

### Loyalty

- Points programs
- Stamp cards
- VIP tiers
- Reward catalogue
- Redemption rules
- Append-only loyalty ledger
- Cached balances for dashboard performance
- Reversal and expiry handling

### Campaigns

- Customer segments
- First-return campaigns
- Win-back campaigns
- Reward expiry reminders
- VIP milestone campaigns
- Holdout/control groups
- Campaign profitability reporting

### Communications

- WhatsApp templates
- Message personalization
- Delivery status tracking
- Inbound customer commands
- Opt-out handling
- Retry and failure queues

### Analytics

- Repeat purchase rate
- Returning customer revenue
- Customer lifetime value directionally
- Reward liability
- Campaign conversion
- Campaign incremental lift
- Campaign net impact

## Repository Structure

```text
.
├── cmd/
│   ├── api/
│   ├── migrate/
│   ├── webhook-gateway/
│   └── worker/
├── internal/
│   ├── analytics/
│   ├── app/
│   ├── audit/
│   ├── auth/
│   ├── branches/
│   ├── campaigns/
│   ├── consent/
│   ├── customers/
│   ├── experimentation/
│   ├── identity/
│   ├── jobs/
│   ├── loyalty/
│   ├── mpesa/
│   ├── notifications/
│   ├── outbox/
│   ├── platform/
│   ├── rbac/
│   ├── rewards/
│   ├── segments/
│   ├── shared/
│   ├── tenants/
│   ├── tiers/
│   ├── transactions/
│   └── whatsapp/
├── db/
│   ├── migrations/
│   ├── queries/
│   ├── seeds/
│   └── sqlc/
├── web/
│   └── merchant-dashboard/
├── deployments/
│   ├── docker/
│   └── terraform/
├── api/
│   ├── generated/
│   └── openapi.yaml
└── tests/
    ├── contract/
    ├── e2e/
    ├── fixtures/
    ├── helpers/
    └── integration/
```

All application-specific business logic lives under `internal/`. The repository intentionally avoids a public `pkg/` directory until there is a real external package or SDK that another Go module should import.

## Deployable Applications

### API

Path: [`cmd/api`](./cmd/api)

Primary merchant-facing and platform API. It will own:

- Merchant onboarding
- Customer registration
- Loyalty program configuration
- Rewards and redemptions
- Campaign setup
- Dashboard analytics
- Staff and role management

Run locally:

```sh
go run ./cmd/api
```

Default address: `:8080`

### Webhook Gateway

Path: [`cmd/webhook-gateway`](./cmd/webhook-gateway)

External integration ingress for callback-heavy providers. It should acknowledge provider callbacks quickly, persist raw payloads, normalize events, and hand durable processing to the outbox/worker layer.

Initial providers:

- Safaricom Daraja/M-Pesa
- WhatsApp Business Platform

Run locally:

```sh
go run ./cmd/webhook-gateway
```

Default address: `:8081`

### Worker

Path: [`cmd/worker`](./cmd/worker)

Background job processor for:

- Outbox event processing
- Points posting
- Reward expiry
- Campaign execution
- WhatsApp message sending
- M-Pesa reconciliation
- Reporting refresh jobs

Run locally:

```sh
go run ./cmd/worker
```

### Migrations

Path: [`cmd/migrate`](./cmd/migrate)

Migration runner entrypoint backed by Goose.

```sh
DATABASE_DSN=postgres://user:pass@localhost:5432/giftbox?sslmode=disable go run ./cmd/migrate status
DATABASE_DSN=postgres://user:pass@localhost:5432/giftbox?sslmode=disable go run ./cmd/migrate up
```

## Core Domain Modules

| Module | Purpose |
| --- | --- |
| `internal/app` | Application construction and deployable wiring |
| `internal/auth` | Merchant authentication integration and session-facing routes |
| `internal/rbac` | Roles, permissions, and authorization policy |
| `internal/tenants` | Merchant tenants, branches, staff membership, roles |
| `internal/branches` | Merchant branches, outlets, branch-specific access, and reporting |
| `internal/customers` | Customer profiles, identities, memberships, merge rules |
| `internal/identity` | Phone, WhatsApp, QR, email, and POS identity matching |
| `internal/consent` | Consent purposes, opt-ins, opt-outs, history |
| `internal/transactions` | Purchases, refunds, imports, reconciliation state |
| `internal/mpesa` | Daraja credentials, callbacks, normalization, reconciliation |
| `internal/whatsapp` | WhatsApp messages, templates, inbound commands, delivery callbacks |
| `internal/loyalty` | Programs, rules, memberships, tiers, points ledger |
| `internal/rewards` | Reward catalogue, issuance, redemption, expiry |
| `internal/tiers` | VIP tier thresholds, upgrades, downgrades, and tier events |
| `internal/campaigns` | Campaign definitions, workflows, control groups, outcomes |
| `internal/experimentation` | Holdout groups, treatment allocation, and lift measurement |
| `internal/segments` | Audience selection and customer grouping |
| `internal/analytics` | Retention, campaign, revenue, and reward liability reporting |
| `internal/notifications` | Notification orchestration across customer and merchant channels |
| `internal/audit` | Immutable administrative and security event history |
| `internal/outbox` | Transactional outbox and reliable async event delivery |
| `internal/jobs` | Scheduled operational jobs such as reconciliation and expiry |
| `internal/platform` | Infrastructure implementations such as config, database, server, observability, cache, storage, and Temporal wiring |
| `internal/shared` | Small internal primitives such as money, validation, pagination, secure tokens, and typed errors |

## Technology Direction

The current scaffold is intentionally dependency-light so the repository compiles immediately. The intended production stack is:

| Layer | Direction |
| --- | --- |
| Backend | Go |
| HTTP | `net/http` initially; Chi is a good next step for route ergonomics |
| Database | PostgreSQL |
| SQL access | `pgx` and `sqlc` |
| Cache/rate limits | Redis |
| Workflows | Temporal Cloud with Go SDK |
| Dashboard | Next.js, TypeScript, Tailwind CSS, shadcn/ui |
| Object storage | Amazon S3 |
| Auth | AWS Cognito initially, Auth0 also acceptable |
| Cloud | AWS, starting in Africa Cape Town `af-south-1` after latency validation |
| Runtime | ECS Fargate |
| Observability | OpenTelemetry, CloudWatch, Grafana |
| Secrets | AWS Secrets Manager and KMS |
| Later event backbone | NATS JetStream |

## Local Development

### Prerequisites

- Go `1.25.1` or compatible with the version in [`go.mod`](./go.mod)
- Docker, when running PostgreSQL locally through Compose
- Node.js and a package manager once the dashboard dependencies are added

### Clone and Verify

```sh
go test ./...
```

### Run the API

```sh
go run ./cmd/api
```

Health checks:

```sh
curl http://localhost:8080/healthz
curl http://localhost:8080/readyz
```

Example module discovery route:

```sh
curl http://localhost:8080/v1/customers
```

### Run the Webhook Gateway

```sh
go run ./cmd/webhook-gateway
```

Example route:

```sh
curl http://localhost:8081/v1/webhooks/mpesa
```

### Run the Worker

```sh
go run ./cmd/worker
```

The scaffolded worker currently logs polling activity. It will later process durable outbox events and workflow activities.

### Run with Docker Compose

```sh
docker compose -f deployments/docker/docker-compose.yml up --build
```

The Compose file provides:

- API container
- PostgreSQL container

Redis, Temporal dev server, and dashboard services will be added as the platform implementation expands.

## Configuration

| Variable | Default | Description |
| --- | --- | --- |
| `APP_ENV` | `local` | Runtime environment name |
| `HTTP_ADDR` | `:8080` | API listen address |
| `WEBHOOK_ADDR` | `:8081` | Webhook gateway listen address |
| `DATABASE_DSN` | empty | PostgreSQL connection string |
| `WORKER_POLL_INTERVAL` | `5s` | Worker polling interval |

Future configuration groups:

- Daraja consumer key and secret
- Daraja shortcode mappings
- WhatsApp Business Account settings
- WhatsApp webhook verify token
- S3 bucket names
- Redis connection string
- Temporal namespace and task queues
- Cognito/Auth0 tenant settings
- KMS key IDs

## API Surface

The OpenAPI contract lives at [`api/openapi.yaml`](./api/openapi.yaml).

Expected endpoint groups:

```text
POST   /v1/tenants
POST   /v1/tenants/{tenant_id}/programs
POST   /v1/tenants/{tenant_id}/branches
POST   /v1/tenants/{tenant_id}/customers/import
GET    /v1/tenants/{tenant_id}/customers
GET    /v1/tenants/{tenant_id}/customers/{id}/ledger
POST   /v1/tenants/{tenant_id}/campaigns
POST   /v1/tenants/{tenant_id}/campaigns/{id}/launch
GET    /v1/tenants/{tenant_id}/reports/retention
GET    /v1/tenants/{tenant_id}/reports/campaign-profitability

POST   /webhooks/mpesa/{tenant_connection_id}
GET    /webhooks/whatsapp/verify
POST   /webhooks/whatsapp

POST   /customer/v1/enrol/{merchant_slug}
GET    /customer/v1/rewards
POST   /customer/v1/rewards/{id}/redeem
```

The scaffold currently exposes module discovery routes under `/v1/{module}` and health routes under `/healthz` and `/readyz`.

## Data Architecture

Giftbox starts with a shared PostgreSQL database and strict tenant partitioning.

Every business-facing table should include:

```sql
tenant_id UUID NOT NULL
```

Tenant enforcement happens at several layers:

1. Request authentication resolves the user.
2. Tenant membership and role are validated.
3. Repository queries are tenant-scoped.
4. PostgreSQL Row-Level Security should be added for defense in depth.
5. Audit logs record sensitive access and changes.

### Core Tables

Planned core tables:

| Table | Purpose |
| --- | --- |
| `tenants` | Merchant businesses |
| `branches` | Store locations or outlets |
| `staff_users` | Merchant team members |
| `customers` | Platform-level customer records |
| `customer_identities` | Phone, QR, merchant references, email |
| `customer_memberships` | Customer membership in a merchant program |
| `consent_records` | Marketing and processing permission history |
| `transactions` | Purchases and refunds |
| `transaction_webhooks` | Raw external callback storage |
| `loyalty_programs` | Merchant program configuration |
| `loyalty_rules` | Versioned earning and redemption rules |
| `loyalty_ledger_entries` | Auditable points and reward movements |
| `tiers` | VIP status levels |
| `rewards` | Redeemable offers |
| `reward_redemptions` | Reward usage |
| `segments` | Customer audience logic |
| `campaigns` | Promotion definitions |
| `campaign_audiences` | Customers selected for campaigns |
| `messages` | WhatsApp message history |
| `integration_connections` | M-Pesa, WhatsApp, POS credentials/config |
| `audit_logs` | Administrative and security actions |
| `outbox_events` | Reliable domain event publishing |

### Loyalty Ledger Rule

The platform must not rely only on a mutable `points_balance` field.

Correct model:

```text
Ledger as source of truth
Cached membership balance for reads
Periodic balance verification job
```

Example ledger event types:

- `POINTS_EARNED`
- `POINTS_BONUS`
- `POINTS_REDEEMED`
- `POINTS_EXPIRED`
- `POINTS_REVERSED`
- `ADMIN_ADJUSTMENT`

### Idempotency Requirements

Every provider transaction must have a stable uniqueness key:

```sql
UNIQUE (tenant_id, payment_provider, provider_transaction_reference)
```

Every points posting must have an idempotency key:

```text
tenant_id + transaction_id + loyalty_rule_version + ledger_action
```

Every reward issuance must reference the triggering event.

## Migrations

Giftbox uses Goose for database migrations. Migration files live in [`db/migrations`](./db/migrations) and follow Goose's annotated SQL format:

```sql
-- +goose Up
CREATE TABLE example_records (
    id TEXT PRIMARY KEY
);

-- +goose Down
DROP TABLE IF EXISTS example_records;
```

Run migrations through the project command:

```sh
DATABASE_DSN=postgres://giftbox:giftbox@localhost:5432/giftbox?sslmode=disable go run ./cmd/migrate status
DATABASE_DSN=postgres://giftbox:giftbox@localhost:5432/giftbox?sslmode=disable go run ./cmd/migrate up
```

See [`docs/migrations.md`](./docs/migrations.md) for the migration conventions and supported commands.

## Event-Driven Processing

Giftbox is an event-driven retention system, even while implemented as a modular monolith.

Important domain events:

```text
CustomerEnrolled
ConsentGranted
ConsentWithdrawn
TransactionReceived
TransactionConfirmed
TransactionRefunded
PointsEarned
TierUpgraded
RewardIssued
RewardRedeemed
RewardExpired
CampaignStarted
CampaignAudienceSelected
MessageRequested
MessageDelivered
CustomerReturned
CampaignConverted
CampaignBudgetExceeded
```

### Transactional Outbox

The transactional outbox is mandatory for payment and loyalty correctness.

Process:

1. Begin PostgreSQL transaction.
2. Write the domain change.
3. Write the outbox event in the same transaction.
4. Commit.
5. Worker reads unpublished events.
6. Worker delivers the event.
7. Worker marks the event published.

This prevents a transaction from being stored while the points-processing event is lost.

## Security and Privacy

Giftbox handles phone numbers, purchase behaviour, marketing preferences, merchant revenue data, campaign segments, and integration credentials. Security and privacy are product requirements, not follow-up tasks.

### Required Controls

| Risk | Control |
| --- | --- |
| Cross-merchant leakage | Tenant-scoped authorization and Row-Level Security |
| Duplicate payment callbacks | Idempotency keys and raw webhook archive |
| Customer phone leakage | Field-level encryption or tokenization |
| Reward fraud | Signed single-use redemption codes and audit logs |
| Staff misuse | RBAC and immutable audit records |
| Marketing without consent | Consent checks before audience activation and message sending |
| Provider credential exposure | Secrets Manager and KMS |
| Webhook forgery | Provider verification, HTTPS, replay protection |
| Account takeover | MFA for owners and managers |

### Consent Purposes

Track consent separately for:

- `LOYALTY_ENROLMENT`
- `TRANSACTION_LINKING`
- `WHATSAPP_SERVICE_MESSAGES`
- `WHATSAPP_MARKETING_MESSAGES`
- `CROSS_MERCHANT_REWARDS`
- `ANALYTICS_PERSONALIZATION`

Customers must be able to opt out of marketing without losing access to earned rewards, subject to merchant program rules.

### Audit Events

Record at minimum:

- Merchant changes reward rules.
- Staff manually adjusts points.
- Campaign is approved or launched.
- Customer withdraws consent.
- Integration credentials are updated.
- Reward is redeemed or reversed.
- Data export or deletion request is processed.

## Reliability Requirements

| Component | Requirement |
| --- | --- |
| Incoming webhooks | Persist raw event before processing |
| Payment processing | Idempotency and reconciliation |
| Loyalty updates | Database transaction plus ledger |
| Campaign delivery | Durable workflows, retries, and status tracking |
| Database | Backups and Multi-AZ production deployment |
| Integrations | Rate limits, circuit breakers, retries, dead-letter review |
| Reports | Rebuildable from transactions and ledger |

Expected failure behaviours:

- Duplicate payment callback: store raw callback, do not double-award points.
- Worker crash: resume from pending outbox event and idempotency key.
- WhatsApp failure: retry message, do not duplicate reward issuance.
- Customer opts out mid-campaign: cancel future marketing messages.
- Reconciliation finds missed transaction: issue points once and notify only if permitted.
- Merchant changes points rules: old earnings keep old rule version.

## Testing

Run all Go tests:

```sh
go test ./...
```

Current test areas:

- [`tests/integration`](./tests/integration): service-level integration tests
- [`tests/contract`](./tests/contract): API contract existence and compatibility checks
- [`tests/e2e`](./tests/e2e): end-to-end flows to be added as product surfaces mature

Expected future coverage:

- Unit tests for domain rules
- Ledger idempotency tests
- Tenant isolation tests
- Consent enforcement tests
- Webhook replay tests
- Migration tests
- API contract tests
- Worker retry tests
- Temporal workflow tests

## Deployment Direction

The MVP deployment target is AWS ECS Fargate.

Expected production components:

- CloudFront
- AWS WAF
- Application Load Balancer
- ECS Fargate services for API, webhook gateway, worker, and dashboard
- RDS PostgreSQL Multi-AZ
- ElastiCache Redis
- S3
- AWS Secrets Manager
- KMS
- CloudWatch
- OpenTelemetry collector/exporters
- Temporal Cloud

Container strategy:

```text
giftbox-api
giftbox-webhook-gateway
giftbox-worker
giftbox-merchant-dashboard
```

CI/CD direction:

```text
Pull Request
  -> tests
  -> security/dependency scan
  -> container build
  -> push to ECR
  -> deploy staging
  -> smoke and contract tests
  -> production approval
  -> deploy ECS services
  -> monitor
```

## Product Roadmap

### Phase 1: Foundation

- Multi-tenant merchant accounts
- Authentication and roles
- Business onboarding
- Loyalty program configuration
- Customer profiles and consent
- Points ledger
- QR enrolment
- CSV transaction import
- Basic dashboard

Outcome: a merchant can register customers, upload purchases, and award points.

### Phase 2: WhatsApp Activation

- WhatsApp Business integration
- Template management
- Welcome flow
- Balance messages
- Reward notifications
- Opt-out handling
- Delivery status dashboard

Outcome: customers can join and interact with loyalty through WhatsApp.

### Phase 3: M-Pesa Transaction Integration

- Daraja merchant connection
- Payment callback ingestion
- Raw webhook store
- Transaction normalization
- Idempotent points awarding
- Reconciliation worker
- Transaction error dashboard

Outcome: eligible M-Pesa purchases automatically generate loyalty rewards.

### Phase 4: Campaign Automation

- Customer segments
- First-return campaign
- Win-back campaign
- Reward expiry nudges
- Temporal workflow integration
- Campaign reporting
- Holdout group experimentation

Outcome: merchants can automatically bring customers back and see conversion.

### Phase 5: Profitability Reporting

- Reward cost configuration
- Gross margin configuration
- Repeat purchase reports
- Campaign net impact reports
- Rules-based campaign recommendations
- AI-assisted campaign copy within strict controls

Outcome: merchants understand whether loyalty is improving business profit.

## Engineering Standards

### Go

- Keep domain modules under `internal/`.
- Keep deployable entrypoints under `cmd/`.
- Prefer explicit SQL and transaction boundaries for ledger and payment logic.
- Keep provider-specific logic behind integration interfaces.
- Use context-aware functions for I/O and external calls.
- Use structured logging.
- Treat idempotency as part of the domain model, not middleware decoration.

### Database

- Every tenant-owned table must include `tenant_id`.
- Every webhook payload must be durably stored before processing.
- Every points movement must be represented in the ledger.
- Every reward redemption must be auditable.
- Migrations must be reversible when practical.
- Use database constraints for uniqueness and invariants that must never drift.

### API

- External APIs are REST.
- Internal processing is event-driven.
- Errors should be structured and safe to show to clients.
- Tenant ID, user ID, request ID, and correlation IDs should be traceable.
- Webhook handlers must respond quickly and defer slow work.

### Frontend

- Merchant dashboard is operational software, not a marketing site.
- Prioritize clear tables, filters, drilldowns, and action states.
- Make campaign profitability and customer status easy to scan.
- Treat mobile responsiveness as required for merchant staff workflows.

### Documentation

- Keep README accurate as architecture decisions change.
- Keep OpenAPI updated with external contract changes.
- Add ADRs for major decisions such as workflow engine, auth provider, and event backbone.
- Document provider assumptions for Daraja and WhatsApp integrations.

## Current Status

This repository currently contains the initial Go scaffold:

- Compilable API, webhook gateway, worker, and migrations entrypoints
- Domain package placeholders with module discovery routes
- Basic middleware, observability, validation, logging, encryption, and database packages
- Initial SQL migration
- Docker and Terraform deployment placeholders
- OpenAPI starter contract
- Merchant dashboard placeholder
- Starter contract and integration tests

Verify current state:

```sh
go test ./...
```

## License

License details are not yet defined. Add a license before public distribution or external contribution.
