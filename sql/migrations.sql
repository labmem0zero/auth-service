CREATE SCHEMA IF NOT EXISTS auth;
CREATE TABLE IF NOT EXISTS auth.users(
    user_id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL ,
    password TEXT NOT NULL
);

CREATE OR REPLACE FUNCTION auth.user_set(
    _user_id INT,
    _username TEXT,
    _password TEXT
)
    RETURNS INT
    SECURITY DEFINER
    LANGUAGE plpgsql
    STRICT
    AS $$
DECLARE
    res_ INT;
BEGIN
    IF _user_id = 0 THEN
        IF EXISTS (SELECT 1 FROM auth.users WHERE username = _username) THEN
            RAISE EXCEPTION 'Имя пользователя уже занято %', _username;
        END IF;

        INSERT INTO auth.users (username, password)
        VALUES (_username, _password)
        RETURNING user_id INTO res_;
    ELSE
        IF NOT EXISTS (SELECT 1 FROM auth.users WHERE user_id = _user_id) THEN
            RAISE EXCEPTION 'Пользователь не найден %', _user_id;
        END IF;

        IF EXISTS (
            SELECT 1
            FROM auth.users
            WHERE username = _username AND user_id <> _user_id
        ) THEN
            RAISE EXCEPTION 'Имя пользователя уже занято %', _username;
        END IF;

        UPDATE auth.users
        SET username = _username,
            password = COALESCE(NULLIF(_password, ''),password)
        WHERE user_id = _user_id
        RETURNING user_id INTO res_;
    END IF;

    RETURN res_;
END
$$;

CREATE OR REPLACE FUNCTION auth.user_get(
    _user_id INT,
    _username TEXT
)
    RETURNS TABLE (
                      user_id INT,
                      username TEXT
                  )
    SECURITY DEFINER
    LANGUAGE plpgsql
    STRICT
    AS $$
BEGIN
    IF _user_id = 0 AND _username = '' THEN
        RAISE EXCEPTION 'Не указаны данные для поиска';
    END IF;

    RETURN QUERY
        SELECT user_id, username
        FROM auth.users
        WHERE user_id = _user_id OR username = _username
        LIMIT 1;
END;
$$;