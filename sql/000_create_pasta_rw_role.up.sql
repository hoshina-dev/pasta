-- Create pasta_rw role with limited privileges
-- This role can SELECT, INSERT, UPDATE but NOT DELETE (soft delete only)
-- It cannot modify schema (no CREATE, DROP, ALTER)

-- Create the role if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'pasta_rw') THEN
        CREATE ROLE pasta_rw WITH LOGIN PASSWORD 'pasta_rw_password';
    END IF;
END
$$;

-- Allow the role to connect to the current database
GRANT CONNECT ON DATABASE pasta TO pasta_rw;

-- Grant usage on the public schema (required to see objects, but not modify schema)
GRANT USAGE ON SCHEMA public TO pasta_rw;

-- Grant SELECT, INSERT, UPDATE on all existing tables (no DELETE)
GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA public TO pasta_rw;

-- Grant USAGE on all existing sequences (needed for uuid_generate_v4 defaults, etc.)
GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO pasta_rw;

-- Grant EXECUTE on all existing functions/procedures
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA public TO pasta_rw;

-- Alter default privileges so future objects created by the current role (postgres)
-- automatically grant the same permissions to pasta_rw
ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT SELECT, INSERT, UPDATE ON TABLES TO pasta_rw;

ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT USAGE ON SEQUENCES TO pasta_rw;

ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT EXECUTE ON FUNCTIONS TO pasta_rw;
