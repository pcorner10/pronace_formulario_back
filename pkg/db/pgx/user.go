package db

import "pronaces_back/pkg/domain"

// TODO
func (r *psxStore) LoginUser(email, password string) (*domain.ReponseLogin, error) {
	return nil, nil
}

// TODO
func (r *psxStore) RegisterUser(user *domain.User) error {
	return nil
}

func (r *psxStore) CreateUser(user domain.User) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id`
	_, err = tx.Exec(query, user.UserName, user.Email, user.Password, user.Role)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *psxStore) GetUserByEmail(email string) (*domain.User, error) {

	// implementar la lógica de obtener un usuario por ID con Psx
	user := &domain.User{}

	err := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *psxStore) UpdateUser(user domain.User) error {
	// implementar la lógica de actualizar un usuario con Psx
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE users SET username = $1, email = $2, password = $3, role = $4 WHERE id = $5`
	_, err = tx.Exec(query, user.UserName, user.Email, user.Password, user.Role, user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *psxStore) DeleteUser(userID uint64) error {

	// implementar la lógica de eliminar un usuario con Psx
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM users WHERE id = $1`
	_, err = tx.Exec(query, userID)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
