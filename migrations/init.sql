CREATE TABLE patients (
    id SERIAL PRIMARY KEY,
    resource_id VARCHAR(255) UNIQUE NOT NULL,
    active BOOLEAN,
    gender VARCHAR(50),
    birth_date DATE,
    deceased BOOLEAN,
    deceased_date TIMESTAMP WITHOUT TIME ZONE,
    marital_status_code VARCHAR(255),
    marital_status_text VARCHAR(255),
    -- Consider JSONB for complex nested structures like name, address, telecom
    name JSONB,
    address JSONB,
    telecom JSONB
);
