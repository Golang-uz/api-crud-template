package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/realtemirov/api-crud-template/models"
	"github.com/rs/zerolog"
)

const (
	fieldsOfUser          = "id, created_at, updated_at, deleted_at, first_name, last_name, user_name, email, password"
	fieldsOfUserWithoutID = "created_at, updated_at, deleted_at, first_name, last_name, user_name, email, password"
)

type userRepo struct {
	db  *sql.DB
	log zerolog.Logger
}

// newUserRepo constructor
func newUserRepo(db *sql.DB, log zerolog.Logger) *userRepo {
	return &userRepo{
		db:  db,
		log: log,
	}
}

// CreateUser implements storage.UserI
func (u *userRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	// response result
	var result models.User

	// query
	query := `
		INSERT INTO
			users (` + fieldsOfUserWithoutID + `)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING ` + fieldsOfUser

	// execute query and scan result
	err := u.db.QueryRowContext(ctx, query,
		time.Now().Unix(),
		user.UpdatedAt,
		user.DeletedAt,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.Email,
		user.Password,
	).Scan(
		&result.ID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
		&result.FirstName,
		&result.LastName,
		&result.UserName,
		&result.Email,
		&result.Password,
	)

	// check error
	if err != nil {

		// log error
		u.log.Info().Msg("Method CreateUser: error while creating user: " + err.Error())

		// return error
		return nil, err
	}

	// return result
	return &result, nil
}

// DeleteUser implements storage.UserI
func (u *userRepo) DeleteUser(ctx context.Context, id int) (*models.User, error) {

	// response result
	var result models.User

	// query
	query := `
		UPDATE
			users
		SET
			deleted_at = $1
		WHERE
			id = $2 AND deleted_at = 0
		RETURNING ` + fieldsOfUser

	// execute query and scan result
	err := u.db.QueryRowContext(ctx, query,
		time.Now().Unix(),
		id,
	).Scan(
		&result.ID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
		&result.FirstName,
		&result.LastName,
		&result.UserName,
		&result.Email,
		&result.Password,
	)

	// check error
	if err != nil {

		// log error
		u.log.Info().Msg("Method DeleteUser: error while deleting user: " + err.Error())

		// return error
		return nil, err
	}

	// return result
	return &result, nil
}

// GetUserByEmail implements storage.UserI
func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	// response result
	var result models.User

	// query
	query := `
		SELECT
			` + fieldsOfUser + `
		FROM
			users
		WHERE
			email = $1
		LIMIT 1`

	// execute query and scan result
	err := u.db.QueryRowContext(ctx, query, email).Scan(
		&result.ID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
		&result.FirstName,
		&result.LastName,
		&result.UserName,
		&result.Email,
		&result.Password,
	)

	// check error
	if err != nil {

		// log error
		u.log.Info().Msg("Method GetUserByEmail: error while getting user by email: " + err.Error())

		// return error
		return nil, err
	}

	// return result
	return &result, nil
}

// GetUserByID implements storage.UserI
func (u *userRepo) GetUserByID(ctx context.Context, id int) (*models.User, error) {

	// response result
	var result models.User

	// query
	query := `
		SELECT
			` + fieldsOfUser + `
		FROM users
		WHERE
			id = $1
		LIMIT 1`

	// execute query and scan result
	err := u.db.QueryRowContext(ctx, query, id).Scan(
		&result.ID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
		&result.FirstName,
		&result.LastName,
		&result.UserName,
		&result.Email,
		&result.Password,
	)

	// check error
	if err != nil {

		// log error
		u.log.Info().Msg("Method GetUserByID: error while getting user by id: " + err.Error())

		// return error
		return nil, err
	}

	// return result
	return &result, nil
}

// GetUserByUserName implements storage.UserI
func (u *userRepo) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {

	// response result
	var result models.User

	// query
	query := `
		SELECT
			` + fieldsOfUser + `
		FROM
			users
		WHERE
			user_name = $1
		LIMIT 1`

	// execute query and scan result
	err := u.db.QueryRowContext(ctx, query,
		userName,
	).Scan(
		&result.ID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
		&result.FirstName,
		&result.LastName,
		&result.UserName,
		&result.Email,
		&result.Password,
	)

	// check error
	if err != nil {

		// log error
		u.log.Info().Msg("Method GetUserByUserName: error while getting user by user name: " + err.Error())

		// return error
		return nil, err
	}

	// return result
	return &result, nil
}

// GetUsers implements storage.UserI
func (u *userRepo) GetUsers(ctx context.Context, meta *models.Meta) (*models.GetAllUsersResponse, error) {

	// response result
	var result models.GetAllUsersResponse

	// query
	query := `
		SELECT
			` + fieldsOfUser + `
		FROM
			users
		WHERE
			deleted_at = 0
		ORDER BY
			id DESC
		LIMIT $1
		OFFSET $2`

	// calculate limit and offset
	limit, offset := meta.GetLimitAndOffset()

	// execute query and scan result
	rows, err := u.db.QueryContext(ctx, query, limit, offset)

	// check error
	if err != nil {

		// log error
		u.log.Info().Msg("Method GetUsers: error while getting users: " + err.Error())

		// return error
		return nil, err
	}

	// close rows
	defer rows.Close()

	// loop rows
	for rows.Next() {

		// response result
		var user models.User

		// scan result
		err = rows.Scan(
			&user.ID,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Email,
			&user.Password,
		)

		// check error
		if err != nil {

			// log error
			u.log.Info().Msg("Method: Getusers Error: " + err.Error())

			// return error
			return nil, err
		}
		// append result
		result.Data = append(result.Data, &user)
	}

	// calculate total
	query = `
		SELECT
			COUNT(1)
		FROM
			users
		WHERE
			deleted_at = 0`

	// execute query and scan result
	err = u.db.QueryRowContext(ctx, query).Scan(
		&result.Meta.TotalData,
	)

	// check error
	if err != nil {

		// log error
		u.log.Info().Msg("Method GetUsers: error while getting total users: " + err.Error())

		// return error
		return nil, err
	}

	// calculate meta
	result.Meta = meta.SetTotalData(result.Meta.TotalData)

	// return result
	return &result, nil

}

// UpdateUser implements storage.UserI
func (u *userRepo) UpdateUser(ctx context.Context, use *models.User) (*models.User, error) {

	// response result
	var result models.User

	// query
	query := `
		UPDATE
			users
		SET
			first_name = $1,
			last_name = $2,
			user_name = $3,
			email = $4,
			password = $5,
			updated_at = $6
		WHERE
			id = $7 AND deleted_at = 0
		RETURNING ` + fieldsOfUser

	// execute query and scan result
	err := u.db.QueryRowContext(ctx, query,
		use.FirstName,
		use.LastName,
		use.UserName,
		use.Email,
		use.Password,
		time.Now().Unix(),
		use.ID,
	).Scan(
		&result.ID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
		&result.FirstName,
		&result.LastName,
		&result.UserName,
		&result.Email,
		&result.Password,
	)

	// check error
	if err != nil {

		// log error
		u.log.Info().Msg("Method UpdateUser: error while updating user: " + err.Error())

		// return error
		return nil, err
	}

	// return result
	return &result, nil
}

func (p *userRepo) RemoveFromDB(ctx context.Context, id int) error {

	// query
	query := `
		DELETE FROM
			users
		WHERE
			id = $1`

	// execute query
	_, err := p.db.ExecContext(ctx, query, id)

	// check error
	if err != nil {

		// log error
		p.log.Info().Msg("Method: RemoveFromDB Error: " + err.Error())

		// return error
		return err
	}

	// if no error, return nil
	return nil
}
