-- Seed User table
DO $$

	DECLARE
		tronOrgId INT;
		tronDevTeamId INT;
		tronDesignTeamId INT;
		tronPMTeamId INT;
		adminId INT;
		trainerId INT;
		traineeId INT;

	BEGIN
		BEGIN
			SELECT org_id INTO tronOrgId FROM organizations WHERE org_name = 'tron';
		EXCEPTION 
			WHEN NO_DATA_FOUND THEN
				RAISE EXCEPTION 'Organization found';
		END;

		BEGIN
			SELECT team_id INTO tronDevTeamId FROM teams WHERE team_name ILIKE 'developers';
			SELECT team_id INTO tronDesignTeamId FROM teams WHERE team_name ILIKE 'ui/ux design';
			SELECT team_id INTO tronPMTeamId FROM teams WHERE team_name ILIKE 'product management';
		EXCEPTION
			WHEN NO_DATA_FOUND THEN
				RAISE EXCEPTION 'Team not found';
		END;

		BEGIN
			SELECT role_id INTO adminId FROM roles WHERE name = 'admin';
			SELECT role_id INTO trainerId FROM roles WHERE name = 'trainer';
			SELECT role_id INTO traineeId FROM roles WHERE name = 'trainee';
		EXCEPTION
			WHEN NO_DATA_FOUND THEN
				RAISE EXCEPTION 'Role not found';
		END;
		
	INSERT INTO users(first_name, last_name, email, username, org_id, team_id, role_id)
		VALUES
			('John', 'Doe', 'john.doe@example.com', 'john_doe', tronOrgId, tronDevTeamId, adminId),
			('Jane', 'Smith', 'jane.smith@example.com', 'jane_smith', tronOrgId, tronDesignTeamId, trainerId),
		        ('Bob', 'Johnson', 'bob.johnson@example.com', 'bob_johnson', tronOrgId, tronDevTeamId, traineeId),
			('Alice', 'Williams', 'alice.williams@example.com', 'alice_williams', tronOrgId, tronPMTeamId, traineeId);
	END $$

