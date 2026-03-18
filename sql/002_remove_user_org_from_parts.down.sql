ALTER TABLE parts 
    ADD COLUMN user_id UUID NOT NULL DEFAULT gen_random_uuid(),
    ADD COLUMN organization_id UUID NOT NULL DEFAULT gen_random_uuid();

CREATE INDEX IF NOT EXISTS idx_parts_organization_id ON parts(organization_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_parts_user_id ON parts(user_id) WHERE deleted_at IS NULL;

DROP TABLE IF EXISTS inventory;