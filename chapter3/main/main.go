package main

import "github.com/jmoiron/sqlx"

func main() {
	// Register(&Handler{
	// 	Db: sqlx.DB{},
	// })
	RegisterWithClosure(&Handler{
		Db: sqlx.DB{},
	})
}
