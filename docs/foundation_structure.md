# Foundation Structure Tracker

This document tracks the backend scaffold for the Giftbox loyalty platform.

## Rules

- Application-specific code lives under `internal/`.
- `pkg/` is intentionally absent until there is a public SDK or package intended for other Go modules.
- Domain logic must not depend on HTTP routers, PostgreSQL implementation details, provider payloads, Temporal internals, or AWS infrastructure.
- Handlers parse requests and call use cases; services own business rules; repositories abstract persistence.
- Goose migrations live under `db/migrations`.
- SQL query files for future `sqlc` generation live under `db/queries`.

## Dependency Direction

```text
HTTP Handlers / Webhooks / Workers
            ↓
        Use Cases
            ↓
      Domain Services
            ↓
 Repository Interfaces / Provider Interfaces
            ↑
 PostgreSQL, M-Pesa, WhatsApp, Temporal implementations
```

## Implementation Sequence

1. Platform infrastructure
2. Tenants, authentication, and RBAC
3. Customers, identity, and consent
4. Transactions and M-Pesa adapter
5. Loyalty ledger, rewards, and tiers
6. Outbox, notifications, and WhatsApp
7. Segments and campaign workflows
8. Experimentation, analytics, and profitability

## Status

- [x] Command layout: `api`, `webhook-gateway`, `worker`, `migrate`
- [x] Internal package layout
- [x] Platform and shared package boundaries
- [x] Goose migration location
- [x] SQLC folder scaffold
- [x] Test folder scaffold
- [ ] Real PostgreSQL schema
- [ ] Tenant context and RBAC middleware
- [ ] First vertical slice: tenant + branch + customer + consent
- [ ] Transaction ingestion
- [ ] Loyalty ledger posting
- [ ] Outbox dispatcher
