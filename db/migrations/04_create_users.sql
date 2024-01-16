-- User table

-- DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
	user_id SERIAL PRIMARY KEY,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	username VARCHAR(255),
	team_id INT,
	org_id INT,
	role_id INT,

	CONSTRAINT fk_org_id
		FOREIGN KEY (org_id)
		REFERENCES organizations (org_id)
		ON UPDATE NO ACTION
		ON DELETE CASCADE,
	
	CONSTRAINT fk_team_id
		FOREIGN KEY (team_id)
		REFERENCES teams (team_id)
		ON UPDATE NO ACTION
		ON DELETE CASCADE,

	CONSTRAINT fk_role_id
		FOREIGN KEY (role_id)
		REFERENCES roles (role_id)
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);



