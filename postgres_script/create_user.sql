CREATE OR REPLACE FUNCTION auth.create_user(_username VARCHAR, _password VARCHAR) RETURNS SETOF auth.user AS $$
BEGIN
	IF EXISTS (SELECT 1 FROM auth.user u where u.username = _username) THEN
		RAISE EXCEPTION 'Username existed!';
	END IF;
	INSERT INTO auth.user(username, password) VALUES (_username, _password); 
	RETURN QUERY SELECT * FROM auth.user u where u.username = _username;
END
$$ LANGUAGE plpgsql;

