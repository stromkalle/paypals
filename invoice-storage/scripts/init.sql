CREATE TABLE IF NOT EXISTS invoices (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    invoice_id uuid NOT NULL REFERENCES invoices(id),
    amount DECIMAL(10, 2) NOT NULL,
    date TIMESTAMP NOT NULL,
    buyer VARCHAR(255) NOT NULL,
    vendor VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);