package models

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"time"
)

type User struct {
	Id        int
	Name      string
	Email     string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func AllUsers(ctx context.Context) ([]*User, error) {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return nil, errors.New("Models: could not get database connection to pool")
	}
	rows, err := db.Query("SELECT * FROM classmate")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		fmt.Println(user)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

/*
CREATE TABLE "classmate" (
	id bigserial primary key,
	name varchar(50) NOT NULL,
	email varchar(50) NOT NULL,
	created_at timestamp DEFAULT current_timestamp,
	updated_at timestamp DEFAULT current_timestamp
);

INSERT INTO "classmate" (name,email) VALUES ('rick','plumbus@fleeb.com'), ('morty', 'dumbus@fleeb.com');

*/
