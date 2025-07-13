-- Drop the trigger first (must be removed before dropping the table)
DROP TRIGGER IF EXISTS set_updated_at ON products;

-- Drop the trigger function
DROP FUNCTION IF EXISTS update_updated_at_column ();

-- Drop the products table
DROP TABLE IF EXISTS products;

-- Optionally remove the extension
DROP EXTENSION IF EXISTS pgcrypto;