-- Drop indexes first
DROP INDEX IF EXISTS codes_codeable_id_index;
DROP INDEX IF EXISTS codes_type_index;
DROP INDEX IF EXISTS codes_expired_at_index;

-- Then drop the table
DROP TABLE IF EXISTS codes;

-- Finally drop the enum type
DROP TYPE IF EXISTS code_type_enum;