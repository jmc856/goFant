package transactions

const returnTxText = `RETURNING id, created_at, action_date, user_team_send_id user_team_receive_id,
								COALESCE(user_team_winner_id, 0) as user_team_winner_id,
						        COALESCE(user_team_loser_id, 0) as user_team_loser_id, amount, odds, status;`


const GET_USER_TEAM_ID = `SELECT id FROM user_teams WHERE user_id=$1 AND team_id=$2`

const selectUserByGuid = `SELECT u.* FROM users u JOIN user_profile up ON u.id=up.user_id WHERE up.guid=$1`

const GET_TEAM_BY_TEAM_KEY = `SELECT * FROM teams WHERE team_key=$1`

const createTx = `INSERT INTO tx
				   (user_team_send_id, user_team_receive_id, amount, odds)
				   VALUES ($1, $2, $3, $4)
				   RETURNING id, created_at, action_date, user_team_send_id user_team_receive_id,
								COALESCE(user_team_winner_id, 0) as user_team_winner_id,
						        COALESCE(user_team_loser_id, 0) as user_team_loser_id, amount, odds, status;`

const createTxDetail = `INSERT INTO tx_detail
						  (tx_id, description, week, player_id_send, player_name_send, player_id_receive, player_name_receive, action)
						  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
						  RETURNING *;`

const getTx = `SELECT * FROM tx WHERE id=$1`

const getUserTeamReceiveId = `SELECT user_team_receive_id FROM tx WHERE id=$1;`

const getUserTeamsFromUserId = `SELECT id FROM user_teams WHERE user_id=$1`

const getTxDetail = `SELECT created_at, amount, action, odds, week, player_performs_action_id
                       FROM tx_detail WHERE tx_id=$1;`

const getTxForUser = `SELECT tx.id, tx.created_at, tx.action_date, tx.user_team_receive_id, tx.user_team_send_id, COALESCE(user_team_winner_id, 0) as user_team_winner_id, COALESCE(user_team_loser_id, 0) as user_team_loser_id, tx.amount, tx.odds, tx.status,
                         txd.description, txd.week, txd.player_id_send, txd.player_name_send, txd.player_id_receive, txd.player_name_receive, txd.action FROM tx tx
                         JOIN user_teams u ON (u.id=tx.user_team_send_id OR u.id=tx.user_team_receive_id)
                         JOIN tx_detail txd ON txd.tx_id=tx.id
                         WHERE u.user_id=$1;`

const acceptTx = `UPDATE tx SET status='PENDING' WHERE id=$1 ` + returnTxText

const selectUserTeamByTeamKeyAndId = `SELECT ut.id FROM user_teams ut
											   JOIN teams t ON ut.team_id=t.id
											   WHERE t.team_key=$1 AND ut.user_id=$2;`