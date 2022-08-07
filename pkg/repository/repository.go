package repository

type Repository struct {
}

func NewRepositrory(db PostgresDB) *Repository {
	return &Repository{}
}
