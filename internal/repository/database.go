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
		FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE )`,
	)
	if err != nil {
		return err
	}
	postCategoryBridge.Exec()

	comment, err := db.Prepare(`CREATE TABLE IF NOT EXISTS comments(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		content TEXT, 
		post_id INTEGER, 
		creator_id INTEGER DEFAULT 0, 
		create_time DATETIME DEFAULT CURRENT_TIMESTAMP,  
		update_time	DATETIME DEFAULT CURRENT_TIMESTAMP, 
		count_like INTEGER DEFAULT 0, 
		count_dislike  INTEGER DEFAULT 0, 
		FOREIGN KEY(creator_id) REFERENCES users(id) ON DELETE CASCADE
		FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE )`,
	)
	if err != nil {
		return err
	}
	comment.Exec()

	post, err := db.Prepare(`CREATE TABLE IF NOT EXISTS posts(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		thread TEXT  ,
		content TEXT ,
		creator_id INTEGER,
		create_time DATETIME DEFAULT CURRENT_TIMESTAMP, 
		update_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		image BLOB,
		count_like INTEGER DEFAULT 0,
		count_dislike INTEGER DEFAULT 0 ,
		CHECK (length(content >= 20))
		FOREIGN KEY(creator_id) REFERENCES users(id) ON DELETE CASCADE
		)`,
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
		image BLOB )`,
	)
	if err != nil {
		return err
	}
	user.Exec()
	//UNIQUE(post_id, comment_id),
	votes, err := db.Prepare(`CREATE TABLE IF NOT EXISTS votes(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		user_id INTEGER, 
		post_id INTEGER, 
		comment_id INTEGER,
		like_state BOOLEAN DEFAULT FALSE,
		dislike_state BOOLEAN DEFAULT FALSE, 
		UNIQUE(post_id, user_id),
		FOREIGN KEY(post_id) REFERENCES posts(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
		)`,
	)
	if err != nil {
		return err
	}
	votes.Exec()

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
		room TEXT NOT NULL UNIQUE,
		FOREIGN KEY(user_id1) REFERENCES users(id),
		FOREIGN KEY(user_id2) REFERENCES users(id)
		)`,
	)
	if err != nil {
		return err
	}
	chat.Exec()

	chatusers, err := db.Prepare(`CREATE TABLE IF NOT EXISTS chatusers(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		sender_id INTEGER,
		receiver_id INTEGER,
		chat_id INTEGER,
		FOREIGN KEY(sender_id) REFERENCES users(id), 
		FOREIGN KEY(receiver_id) REFERENCES users(id), 
		FOREIGN KEY(chat_id) REFERENCES chats(id)
	)`)
	if err != nil {
		return err
	}
	chatusers.Exec()

	message, err := db.Prepare(`CREATE TABLE IF NOT EXISTS messages(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		content TEXT NOT NULL DEFAULT "",
		room TEXT,
		user_id INTEGER,
		sender_id INTEGER,
		receiver_id INTEGER,
		name TEXT,
		sent_time DATETIME,
		is_read  BOOLEAN DEFAULT FALSE,
		count_unread INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
		)`,
	)
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
