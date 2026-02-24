-- SQL migration script for creating parts table
-- No foreign key constraints to users and organizations table
-- Enforces relations in application layer

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";

CREATE TABLE IF NOT EXISTS manufacturers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    name TEXT NOT NULL,
    country_of_origin CHAR(3), -- ISO codes same as gid_0 from GAPI

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT check_iso_format
        CHECK (country_of_origin IS NULL OR country_of_origin ~ '^[A-Z]{3}$'),

    CONSTRAINT manufacturers_name_country_unique
        UNIQUE (name, country_of_origin)
);

-- Part Categories lookup table
CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    name TEXT NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Parts table
CREATE TABLE IF NOT EXISTS parts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    name TEXT NOT NULL,
    part_number TEXT NOT NULL,
    manufacturer_id UUID NOT NULL REFERENCES manufacturers(id) ON DELETE RESTRICT,
    description TEXT,

    condition TEXT NOT NULL, -- Possible value handle in app layer
    temperature_stage TEXT, -- Possible value handle in app layer
    is_available BOOLEAN NOT NULL DEFAULT TRUE,

    user_id UUID NOT NULL,
    organization_id UUID NOT NULL,

    images TEXT[] NOT NULL DEFAULT '{}',

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT unique_manufacturer_part
        UNIQUE (manufacturer_id, part_number)
);

-- Table for parts-categories many-to-many relation
CREATE TABLE IF NOT EXISTS part_categories (
    part_id UUID NOT NULL REFERENCES parts(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (part_id, category_id)
);

-- Foreign key indexes
CREATE INDEX IF NOT EXISTS idx_parts_organization_id ON parts(organization_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_parts_user_id ON parts(user_id) WHERE deleted_at IS NULL;

-- Query indexes. change to trigram
CREATE INDEX IF NOT EXISTS idx_parts_name ON parts USING GIN (name gin_trgm_ops) WHERE deleted_at IS NULL;

-- Active + Available parts index
CREATE INDEX IF NOT EXISTS idx_parts_available ON parts(is_available) 
    WHERE deleted_at IS NULL AND is_available = TRUE;

-- Timestamp sorting index
CREATE INDEX idx_parts_created_at ON parts(created_at DESC) WHERE deleted_at IS NULL;

-- part_categories indexes
CREATE INDEX IF NOT EXISTS idx_part_categories_part_id
    ON part_categories(part_id);
CREATE INDEX IF NOT EXISTS idx_part_categories_category_id
    ON part_categories(category_id);