CREATE TABLE tenants (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE customers (
    id TEXT PRIMARY KEY,
    tenant_id TEXT NOT NULL REFERENCES tenants(id),
    phone TEXT NOT NULL,
    email TEXT,
    first_name TEXT,
    last_name TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (tenant_id, phone)
);

CREATE TABLE consent_records (
    id TEXT PRIMARY KEY,
    tenant_id TEXT NOT NULL REFERENCES tenants(id),
    customer_id TEXT NOT NULL REFERENCES customers(id),
    channel TEXT NOT NULL,
    status TEXT NOT NULL,
    captured_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE transactions (
    id TEXT PRIMARY KEY,
    tenant_id TEXT NOT NULL REFERENCES tenants(id),
    customer_id TEXT REFERENCES customers(id),
    provider TEXT NOT NULL,
    provider_reference TEXT NOT NULL,
    amount_cents BIGINT NOT NULL,
    currency TEXT NOT NULL DEFAULT 'KES',
    occurred_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (tenant_id, provider, provider_reference)
);

CREATE TABLE loyalty_accounts (
    id TEXT PRIMARY KEY,
    tenant_id TEXT NOT NULL REFERENCES tenants(id),
    customer_id TEXT NOT NULL REFERENCES customers(id),
    points_balance BIGINT NOT NULL DEFAULT 0,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (tenant_id, customer_id)
);

CREATE TABLE rewards (
    id TEXT PRIMARY KEY,
    tenant_id TEXT NOT NULL REFERENCES tenants(id),
    name TEXT NOT NULL,
    points_cost BIGINT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE outbox_messages (
    id TEXT PRIMARY KEY,
    tenant_id TEXT REFERENCES tenants(id),
    topic TEXT NOT NULL,
    payload JSONB NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending',
    attempts INT NOT NULL DEFAULT 0,
    available_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_customers_tenant_created_at ON customers (tenant_id, created_at DESC);
CREATE INDEX idx_transactions_tenant_occurred_at ON transactions (tenant_id, occurred_at DESC);
CREATE INDEX idx_outbox_messages_status_available_at ON outbox_messages (status, available_at);
