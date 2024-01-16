-- Create roles table

-- ALTER TABLE other_table DROP CONSTRAINT fk_constaint_name

-- DROP TABLE IF EXISTS public.roles

CREATE TABLE IF NOT EXISTS roles (
	role_id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description VARCHAR(255),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); 


