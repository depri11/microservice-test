package auth

import (
	"database/sql"

	"github.com/depri11/lolipad/src/models"
)

type repository struct {
	*sql.DB
}

func newRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user *models.User) (int64, error) {
	query := `INSERT INTO public."user" (uuid, msdisn, "name", username, "password") VALUES($1, $2, $3, $4, $5) RETURNING id;`
	sql, err := r.DB.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer sql.Close()

	result, err := sql.Exec(user.ID, user.Msisdn, user.Name, user.Username, user.Password)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
