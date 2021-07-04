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
	CREATE TABLE IF NOT EXISTS post_category_bridges(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER,
		category_id INTEGER,
		FOREIGN KEY(category_id) REFERENCES categories(id), 
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
		create_time DATETIME DEFAULT CURRENT_TIMESTAMP,  
		update_time	DATETIME, 
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
		category TEXT,
		create_time DATETIME DEFAULT CURRENT_TIMESTAMP, 
		update_time DATETIME,
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

	session, err := db.Prepare(`CREATE TABLE IF NOT EXISTS sessions(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		uuid TEXT, user_id 
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
		user_name TEXT NOT NULL UNIQUE, 
		password TEXT,
		isAdmin INTEGER DEFAULT 0, 
		age INTEGER, 
		sex TEXT,
		created_time DATETIME DEFAULT CURRENT_TIMESTAMP, 
		last_seen DATETIME DEFAULT CURRENT_TIMESTAMP, 
		city TEXT, 
		image BLOB)`,
	)
	if err != nil {
		return err
	}
	user.Exec()

	votes, err := db.Prepare(`CREATE TABLE IF NOT EXISTS votes(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		user_id INTEGER, post_id INTEGER, 
		comment_id INTEGER,
		like_state BOOLEAN DEFAULT FALSE,
		dislike_state BOOLEAN DEFAULT FALSE, 
		UNIQUE(post_id, user_id), 
		FOREIGN KEY(comment_id) REFERENCES comments(id), 
		FOREIGN KEY(post_id) REFERENCES posts(id))`,
	)
	if err != nil {
		return err
	}
	votes.Exec()

	notify, err := db.Prepare(`CREATE TABLE IF NOT EXISTS notifies(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		post_id INTEGER,  
		current_user_id INTEGER, 
		voteState INTEGER DEFAULT 0, 
		created_time DATETIME DEFAULT CURRENT_TIMESTAMP, 
		to_whom INTEGER, 
		comment_id INTEGER )`)
	if err != nil {
		return err
	}
	notify.Exec()

	category, err := db.Prepare(`CREATE TABLE IF NOT EXISTS categories(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		name TEXT UNIQUE)`,
	)
	if err != nil {
		return err
	}
	category.Exec()

	chat, err := db.Prepare(`CREATE TABLE IF NOT EXISTS chats(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		user_id1 INTEGER,
		user_id2 INTEGER,
		room TEXT NOT NULL UNIQUE
		)`,
	)
	if err != nil {
		return err
	}
	chat.Exec()

	message, err := db.Prepare(`CREATE TABLE IF NOT EXISTS messages(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		content TEXT,
		room TEXT,
		user_id INTEGER,
		name TEXT,
		message_sent_time DATETIME
		)`,
	)
	// FOREIGN KEY(room) REFERENCES chats(room)
	if err != nil {
		return err
	}
	message.Exec()

	putCategoriesInDb(db)
	log.Println(" db created")
	return nil
}

func putCategoriesInDb(db *sql.DB) {

	count := 0

	err := db.QueryRow("SELECT count(*) FROM categories").Scan(&count)
	if err != nil {
		log.Println(err)
		return
	}

	if count == 0 {
		categories := []string{"science", "love", "nature"}
		for i := 0; i < 3; i++ {
			categoryPrepare, err := db.Prepare(`INSERT INTO categories(name) VALUES(?)`)
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
