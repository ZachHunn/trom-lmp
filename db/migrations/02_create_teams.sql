-- Teams Table

-- DROP TABLE IF EXISTS teams

CREATE TABLE IF NOT EXISTS teams(
	team_id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description VARCHAR(255),
	org_id INT,
	CONSTRAINT fk_org_id
		FOREIGN KEY (org_id)
		REFERENCES organization (org_id)
		ON UPDATE NO ACTION
		ON DELETE CASCADE
);


