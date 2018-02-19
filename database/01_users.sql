CREATE TABLE IF NOT EXISTS users (
  id serial,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  last_login timestamp NOT NULL DEFAULT NOW(),
  username VARCHAR(25) NOT NULL,
  password VARCHAR(100) NOT NULL,
  email    VARCHAR(100) NOT NULL DEFAULT '',
  state    VARCHAR(50) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (username),
  UNIQUE (state)
);


/*
 one to one: User has one user_profile
*/

CREATE TABLE IF NOT EXISTS user_profile (
  id serial,
  user_id int NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  first_name VARCHAR(30) NOT NULL DEFAULT '',
  last_name VARCHAR(30) NOT NULL DEFAULT '',
  guid VARCHAR(40) NOT NULL,
  nickname VARCHAR(30) NOT NULL DEFAULT '',
  UNIQUE (guid),
  PRIMARY KEY (user_id),
  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

/*
 one to one: User has one user_credentials
*/

CREATE TABLE IF NOT EXISTS user_credentials (
  id serial,
  user_id int NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  access_token VARCHAR(1500) NOT NULL,
  refresh_token VARCHAR(1500) NOT NULL,
  expiration timestamp NOT NULL DEFAULT NOW(),
  type VARCHAR(40),
  PRIMARY KEY (user_id),
  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id),
  UNIQUE (access_token),
  UNIQUE (refresh_token)
);