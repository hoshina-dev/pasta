DROP INDEX IF EXISTS idx_part_categories_category_id;
DROP INDEX IF EXISTS idx_part_categories_part_id;
DROP INDEX IF EXISTS idx_parts_created_at;
DROP INDEX IF EXISTS idx_parts_available;
DROP INDEX IF EXISTS idx_parts_name;
DROP INDEX IF EXISTS idx_parts_user_id;
DROP INDEX IF EXISTS idx_parts_organization_id;

DROP TABLE IF EXISTS part_categories;
DROP TABLE IF EXISTS parts;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS manufacturers;