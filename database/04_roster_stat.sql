CREATE TABLE IF NOT EXISTS position_types (
  id serial,
  created_at timestamp NOT NULL DEFAULT NOW(),
  game_key VARCHAR(10) NOT NULL,
  game_id VARCHAR(10) NOT NULL,
  season VARCHAR(10) NOT NULL,
  type VARCHAR(10) NOT NULL,
  display_name VARCHAR(50) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (game_key, season, type)
);

CREATE TABLE IF NOT EXISTS roster_positions (
  id serial,
  created_at timestamp NOT NULL DEFAULT NOW(),
  game_key VARCHAR(10) NOT NULL,
  game_id VARCHAR(10) NOT NULL,
  season VARCHAR(10) NOT NULL,
  position VARCHAR(50) NOT NULL,
  abbreviation VARCHAR(10) NOT NULL,
  display_name VARCHAR(50) NOT NULL,
  position_type_id INT NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_position_type_id FOREIGN KEY (position_type_id) REFERENCES position_types (id),
  UNIQUE (game_key, season, position, position_type_id)
);

CREATE TABLE IF NOT EXISTS stat_categories (
  id serial,
  created_at timestamp NOT NULL DEFAULT NOW(),
  game_key VARCHAR(10) NOT NULL,
  game_id VARCHAR(10) NOT NULL,
  season VARCHAR(10) NOT NULL,
  stat_id VARCHAR(10) NOT NULL,
  name VARCHAR(50) NOT NULL,
  display_name VARCHAR(50) NOT NULL,
  sort_order INT NOT NULL,
  position_type_id INT NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_position_type_id FOREIGN KEY (position_type_id) REFERENCES position_types (id),
  UNIQUE (game_key, season, stat_id, position_type_id)
);