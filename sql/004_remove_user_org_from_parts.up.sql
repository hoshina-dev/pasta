DROP INDEX IF EXISTS idx_parts_user_id;
DROP INDEX IF EXISTS idx_parts_organization_id;

ALTER TABLE parts DROP COLUMN IF EXISTS user_id;
ALTER TABLE parts DROP COLUMN IF EXISTS organization_id;
ALTER TABLE parts DROP COLUMN IF EXISTS is_available;

CREATE TABLE parts_inventory (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    part_id UUID REFERENCES parts(id),
    serial_number TEXT NOT NULL UNIQUE,
    is_available BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

