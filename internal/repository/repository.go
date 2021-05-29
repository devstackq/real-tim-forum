package repository

type User interface {
	CreateUser(*models.User) (int64, error)
}

type Repository struct {
	User
}

func NewRepository(db *sql.DB) *Repository{
	return &Repository{
		User: NewUserRepository(db)
	}
}
// Model -> CreateUser func, Repo - conn Db, -> service -> handler