package repository

import (
	"database/sql"
	"log"
)

func CreateDB(driver, path string) (*sql.DB, error) {
	db, err := sql.Open(driver, path)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(100)
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if err = createTables(db); err != nil {
		return nil, err
	}
	return db, nil
}

func createTables(db *sql.DB) error {

	db.Exec("PRAGMA foreign_keys=ON")

	postCategoryBridge, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS post_cat_bridge(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER,
		category_id INTEGER,
		FOREIGN KEY(category_id) REFERENCES category(id), 
		FOREIGN KEY(post_id) REFERENCES posts(id) 
		ON DELETE CASCADE )`,
	)
	if err != nil {
		return err
	}
	postCategoryBridge.Exec()

	comment, err := db.Prepare(`CREATE TABLE IF NOT EXISTS comments(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		parent_id INTEGER DEFAULT 0, 
		content TEXT, 
		post_id INTEGER, 
		creator_id INTEGER DEFAULT 0, 
		toWho INTEGER DEFAULT 0, 
		fromWho INTEGER DEFAULT 0, 
		create_time DATETIME,  
		update_time	DATETIME DEFAULT CURRENT_TIMESTAMP, 
		count_like INTEGER DEFAULT 0, 
		count_dislike  INTEGER DEFAULT 0, 
		CONSTRAINT fk_key_post_comment 
		FOREIGN KEY(post_id) REFERENCES posts(id) 
		ON DELETE CASCADE )`,
	)
	if err != nil {
		return err
	}
	comment.Exec()

	post, err := db.Prepare(`CREATE TABLE IF NOT EXISTS posts(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		thread TEXT,
		content TEXT,
		creator_id INTEGER,
		create_time DATETIME, 
		update_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		image BLOB,
		count_like INTEGER DEFAULT 0,
		count_dislike INTEGER DEFAULT 0, 
		FOREIGN KEY(creator_id) REFERENCES users(id) 
		ON DELETE CASCADE )`,
	)
	if err != nil {
		return err
	}
	post.Exec()

	session, err := db.Prepare(`CREATE TABLE IF NOT EXISTS session(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		uuid	TEXT, user_id 
		INTEGER UNIQUE, 
		cookie_time DATETIME)`,
	)
	if err != nil {
		return err
	}
	session.Exec()

	user, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		full_name TEXT NOT NULL, 
		email TEXT NOT NULL UNIQUE, 
		username TEXT NOT NULL UNIQUE, 
		password TEXT, 
		isAdmin INTEGER DEFAULT 0, 
		age INTEGER, 
		sex TEXT, 
		created_time DATETIME, 
		last_seen DATETIME, 
		city TEXT, 
		image BLOB)`,
	)
	if err != nil {
		return err
	}
	user.Exec()

	voteState, err := db.Prepare(`CREATE TABLE IF NOT EXISTS voteState(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		user_id INTEGER, post_id INTEGER, 
		comment_id INTEGER, 
		like_state INTEGER DEFAULT 0,
		dislike_state INTEGER DEFAULT 0, 
		unique(post_id, user_id), 
		FOREIGN KEY(comment_id) REFERENCES comments(id), 
		FOREIGN KEY(post_id) REFERENCES posts(id))`,
	)
	if err != nil {
		return err
	}
	voteState.Exec()

	notify, err := db.Prepare(`CREATE TABLE IF NOT EXISTS notify(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		post_id INTEGER,  
		current_user_id INTEGER, 
		voteState INTEGER DEFAULT 0, 
		created_time DATETIME, 
		to_whom INTEGER, 
		comment_id INTEGER )`)
	if err != nil {
		return err
	}
	notify.Exec()

	category, err := db.Prepare(`CREATE TABLE IF NOT EXISTS category(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		name TEXT UNIQUE)`,
	)

	if err != nil {
		return err
	}
	category.Exec()

	putCategoriesInDb(db)
	return nil
}

func putCategoriesInDb(db *sql.DB) {

	//count := utils.GetCountTable("category", db)
	count := 0
	if count != 3 {
		categories := []string{"science", "love", "sapid"}
		for i := 0; i < count; i++ {
			categoryPrepare, err := db.Prepare(`INSERT INTO category(name) VALUES(?)`)
			if err != nil {
				log.Println(err)
			}
			_, err = categoryPrepare.Exec(categories[i])
			if err != nil {
				log.Println(err)
			}
			defer categoryPrepare.Close()
		}
	}
}
