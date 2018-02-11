package rosters

const GET_ROSTER_POSITIONS = `SELECT * FROM roster_positions;`

const GET_POSITION_TYPES = `SELECT * FROM position_types;`

const INSERT_POSITION_TYPE = `INSERT INTO position_types (game_key, game_id, season, type, display_name)
							  VALUES ($1, $2, $3, $4, $5);`

const GET_STAT_CATEGORIES = `SELECT * FROM stat_categories;`

const GET_POSITION_TYPE_ID = `SELECT id FROM position_types
							  WHERE season=$1 AND game_key=$2 AND type=$3`

const INSERT_STAT_CATEGORY = `INSERT INTO stat_categories (game_key, game_id, season, stat_id, name, display_name, sort_order, position_type_id)
							  VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`