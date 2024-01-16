-- Organization table

-- DROP TABLE IF EXISTS organizations

CREATE TABLE IF NOT EXISTS organizations (
	org_id SERIAL PRIMARY KEY,
	org_name VARCHAR(255) NOT NULL,
	short_name VARCHAR(255),
	parent_org INT,
	child_org INT
)
