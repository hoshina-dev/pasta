ALTER TABLE parts 
    ADD COLUMN user_id UUID,
    ADD COLUMN organization_id UUID,
    ADD COLUMN is_available BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN condition TEXT NOT NULL;


CREATE INDEX IF NOT EXISTS idx_parts_organization_id ON parts(organization_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_parts_user_id ON parts(user_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_parts_available ON parts(is_available) WHERE deleted_at IS NULL;

DROP TABLE IF EXISTS parts_inventory;