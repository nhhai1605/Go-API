CREATE OR REPLACE FUNCTION auth.login(_username VARCHAR, _password VARCHAR) 
RETURNS SETOF auth.user AS $$
DECLARE
    temp_user auth.user; 
BEGIN
    -- Attempt to fetch the user with the given credentials
    SELECT * INTO temp_user
    FROM auth.user
    WHERE username = _username AND password = _password;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'Invalid username or password!';
    END IF;

    RETURN NEXT temp_user;
END;
$$ LANGUAGE plpgsql;