CREATE EXTENSION "uuid-ossp";
CREATE TYPE lang AS ENUM ('en_US', 'en_GB', 'fr_FR');

-- This table should be accessible only from auth & profile microservices
-- and be very restricted
CREATE TABLE users (
    user_id UUID NOT NULL DEFAULT uuid_generate_v4(),
    login TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    locked BOOL NOT NULL DEFAULT('f'),
    salt1 TEXT NOT NULL,
    salt2 TEXT NOT NULL,
    PRIMARY KEY (user_id)
);

CREATE TABLE user_profiles (
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    user_language lang NOT NULL,
    PRIMARY KEY (user_id)
);

CREATE OR REPLACE FUNCTION user_exists(u UUID) RETURNS BOOLEAN AS $$
    BEGIN
        RETURN EXISTS (SELECT 1 FROM users WHERE user_id=u);
    END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION login_exists(l TEXT) RETURNS BOOLEAN AS $$
    BEGIN
        RETURN EXISTS (SELECT 1 FROM users WHERE login=l);
    END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION email_exists(e TEXT) RETURNS BOOLEAN AS $$
    BEGIN
        RETURN EXISTS (SELECT 1 FROM users WHERE email=e);
    END;
$$ LANGUAGE plpgsql;

-- Reset DB functions

CREATE OR REPLACE FUNCTION reset_userdata() RETURNS void AS $$
    BEGIN
        TRUNCATE TABLE user_achievements;
        TRUNCATE TABLE user_achievement_progress;
        TRUNCATE TABLE user_configurations;
    END;
$$ LANGUAGE plpgsql;


-- CREATE USER uc_auth;
-- ALTER USER uc_auth WITH PASSWORD 'uc_auth_dev';

-- CREATE USER_uc_auth_adm;
-- ALTER USER uc_auth_adm WITH PASSWORD 'uc_auth_adm';
