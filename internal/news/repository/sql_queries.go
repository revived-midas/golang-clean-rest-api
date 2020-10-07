package repository

const (
	createUser = `INSERT INTO news (author_id, title, content, image_url, category, created_at) 
					VALUES ($1, $2, $3, NULLIF($4, ''), NULLIF($4, ''), now()) 
					RETURNING *`

	updateUser = `UPDATE news 
					SET title = COALESCE(NULLIF($1, ''), title),
						content = COALESCE(NULLIF($2, ''), content), 
					    image_url = COALESCE(NULLIF($3, ''), image_url), 
					    category = COALESCE(NULLIF($4, ''), category), 
					    updated_at = now() 
					RETURNING *`

	getNewsByID = `SELECT news_id, author_id, title, content, image_url, category, updated_at FROM news WHERE news_id = $1`
)