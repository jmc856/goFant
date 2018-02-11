package leagues

const CREATE_OR_UPDATE_TEAM = `INSERT INTO teams
						  (name, league_id, team_id, team_key, game_id)
						  VALUES ($1, $2, $3, $4, $5)
						  ON CONFLICT ON CONSTRAINT teams_team_key_key
						  DO UPDATE SET name=$1, league_id=$2, team_id=$3, game_id=$5 WHERE teams.team_key=$4
						  RETURNING *;`


const CREATE_OR_UPDATE_LEAGUE =  `INSERT INTO leagues
					    (name, league_key, current_week, draft_status, league_type, num_teams, season, game_id)
					    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
						ON CONFLICT ON CONSTRAINT leagues_league_key_key
						DO UPDATE SET name=$1, current_week=$3, draft_status=$4, league_type=$5, num_teams=$6, season=$7, game_id=$8 WHERE leagues.league_key=$2
						RETURNING id, name, league_key, current_week, draft_status, league_type, num_teams, season, game_id;`

const CREATE_OR_UPDATE_USER_TEAM = `INSERT INTO user_teams
                                    (user_id, team_id)
								    VALUES ($1, $2)
                                    ON CONFLICT ON CONSTRAINT user_teams_user_id_team_id_key
									DO NOTHING;`

const GET_LEAGUES = `SELECT * FROM leagues WHERE league_key=$1 LIMIT 1;`