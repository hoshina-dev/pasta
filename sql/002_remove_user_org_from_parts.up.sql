DROP INDEX IF EXISTS idx_parts_user_id;
DROP INDEX IF EXISTS idx_parts_organization_id;

ALTER TABLE parts DROP COLUMN user_id;
ALTER TABLE parts DROP COLUMN organization_id;
