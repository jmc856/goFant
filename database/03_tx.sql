/*
 many to many: many users have many teams
*/

CREATE TABLE IF NOT EXISTS user_teams (
  id serial,
  created_at timestamp NOT NULL DEFAULT NOW(),
  user_id int NOT NULL,
  team_id int NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (user_id, team_id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE,
  FOREIGN KEY (team_id) REFERENCES teams(id) ON UPDATE CASCADE
);

/*
 one to many: tx has many user_teams
 one to one:  tx has one tx_detail
*/

CREATE TABLE IF NOT EXISTS tx (
  id serial,
  created_at timestamp NOT NULL DEFAULT NOW(),
  action_date timestamp DEFAULT '1970-1-1'::TIMESTAMP,
  user_team_send_id INT NOT NULL,
  user_team_receive_id INT NOT NULL,
  user_team_winner_id INT DEFAULT NULL,
  user_team_loser_id INT DEFAULT NULL,
  amount VARCHAR(15) NOT NULL,
  odds VARCHAR(15) NOT NULL,
  amount_won INT NOT NULL DEFAULT 0,
  status VARCHAR(20) DEFAULT 'waiting',

  PRIMARY KEY (id),
  FOREIGN KEY (user_team_send_id) REFERENCES user_teams (id) ON UPDATE CASCADE,
  FOREIGN KEY (user_team_receive_id) REFERENCES user_teams (id) ON UPDATE CASCADE,
  FOREIGN KEY (user_team_winner_id) REFERENCES user_teams (id) ON UPDATE CASCADE,
  FOREIGN KEY (user_team_loser_id) REFERENCES user_teams (id) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS tx_detail (
  id serial,
  tx_id INT NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  description VARCHAR(255) NOT NULL DEFAULT '',
  week VARCHAR(15) NOT NULL,
  player_id_send VARCHAR(15) NOT NULL,
  player_name_send VARCHAR(60) NOT NULL,
  player_id_receive VARCHAR(15) NOT NULL,
  player_name_receive VARCHAR(60) NOT NULL,
  action VARCHAR(30) NOT NULL,
  UNIQUE (tx_id),
  PRIMARY KEY (id),
  FOREIGN KEY (tx_id) REFERENCES tx(id) ON UPDATE CASCADE
);
