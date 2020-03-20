CREATE TABLE app_user (
    user_id serial PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password char(60) NOT NULL
);

CREATE TABLE pin (
    pin_id serial PRIMARY KEY,
    image_url varchar(255) NOT NULL,
    username varchar(255) NOT NULL REFERENCES app_user (username)
);

