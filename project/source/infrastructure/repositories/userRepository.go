package repositories

import (
	"context"
	"fmt"
	"projectGoLang/source/domain/entity"
)

func (r *Repository) CreateUser(user entity.User) error {
	q := `
		INSERT INTO user_password(username,password) values($1,$2)
	`

	_, err := r.client.Query(context.TODO(), q, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindUserID(user entity.User) entity.User {
	q := `
		SELECT user_id, username from user_password where username = $1
	`
	err := r.client.QueryRow(context.TODO(), q, user.Username).Scan(&user.UserID, &user.Username)
	if err != nil {
		fmt.Errorf("Creating user impossible to DB", err)
		return user
	}
	return user
}

func (r *Repository) FindUserByID(user_id int) (entity.User, error) {
	q := `
		SELECT user_id, username from user_password where user_id = $1
	`
	user := entity.User{}
	err := r.client.QueryRow(context.TODO(), q, user_id).Scan(&user.UserID, &user.Username)
	if err != nil {
		fmt.Errorf("Creating user impossible to DB", err)
		return user, err
	}
	return user, nil
}

func (r *Repository) FindUserPassword(user entity.User) entity.User {
	q := `
		SELECT password from user_password where username = $1
	`
	err := r.client.QueryRow(context.TODO(), q, user.Username).Scan(&user.Password)
	if err != nil {
		fmt.Errorf("Creating user impossible to DB", err)
		return user
	}
	return user
}
