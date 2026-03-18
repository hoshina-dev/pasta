DROP INDEX IF EXISTS idx_parts_user_id;
DROP INDEX IF EXISTS idx_parts_organization_id;

ALTER TABLE parts DROP COLUMN user_id;
ALTER TABLE parts DROP COLUMN organization_id;
ALTER TABLE parts DROP COLUMN is_available;

CREATE TYPE inventory_type AS ENUM ('part', 'object');

CREATE TABLE inventory (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    part_id UUID REFERENCES parts(id),
    object_id UUID REFERENCES objects(id),
    serial_number TEXT UNIQUE,

    type inventory_type NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);