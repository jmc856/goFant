CREATE TABLE IF NOT EXISTS leagues (
  id serial,
  created_at timestamp NOT NULL DEFAULT NOW(),
  name VARCHAR(50) NOT NULL,
  league_key VARCHAR(100) NOT NULL,
  current_week INTEGER NOT NULL,
  draft_status VARCHAR(50) NOT NULL,
  league_type  VARCHAR(50) NOT NULL,
  num_teams    INTEGER NOT NULL,
  season VARCHAR(10) NOT NULL,
  game_id VARCHAR(30) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (league_key)
);

CREATE TABLE IF NOT EXISTS teams (
  id serial,
  created_at timestamp NOT NULL DEFAULT NOW(),
  name VARCHAR(50) NOT NULL,
  league_id INT NOT NULL,
  team_id INT NOT NULL,
  team_key VARCHAR(30) NOT NULL,
  -- manager_id INT NOT NULL,
  game_id VARCHAR(30) NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_league_id FOREIGN KEY (league_id) REFERENCES leagues (id),
  UNIQUE (team_key)
);