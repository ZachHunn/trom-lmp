-- Seed teams table

DO $$

DECLARE 
	tronOrgId INT;
	p1OrgId INT;

BEGIN
	BEGIN

		SELECT org_id INTO tronOrgId FROM organizations WHERE org_name = 'tron';
		SELECT org_id INTO p1OrgId FROM organizations WHERE org_name = 'platform one';
	EXCEPTION
		WHEN NO_DATA_FOUND THEN
		     RAISE EXCEPTION 'Organization not found.';
	END;

INSERT INTO teams(name, description, org_id)
	VALUES
		('product management', 'This is for the PM team for Tron', tronOrgId),
		('UI/UX design', 'Tron UI/UX team', tronOrgId ),
		('Developers', 'Tron software development team', tronOrgId),
		('P1 Product Team', 'PM team for P1', p1OrgId), 
		('P1 dev team', 'developer team for P1', p1OrgId);
	COMMIT;
END $$
