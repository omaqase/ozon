package user

import "github.com/google/uuid"

type Entity struct {
	ID       uuid.UUID `db:"id"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	Salt     string    `db:"salt"`
}
