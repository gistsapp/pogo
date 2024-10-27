package tests

import (
	"testing"

	"github.com/gistsapp/pogo/pogo"
)

func TestInsertUsers(t *testing.T) {
	t.Run("Insert a user", func(t *testing.T) {
		db := CreateDatabase()
		type User struct {
			UserId string `pogo:"user_id"`
			Name   string `pogo:"name"`
		}
		user := make([]User, 0)
		err := pogo.SuperQuery(db, "INSERT INTO users(name) VALUES($1) RETURNING user_id, name", &user, "test")
		if err != nil {
			t.Fatalf("Failed to insert user: %v", err)
		}

		expected_names := []string{"test"}

		for i, user := range user {
			if user.Name != expected_names[i] {
				t.Fatalf("wrong name, expected %v but got %v", expected_names[i], user.Name)
			}
		}
	})

}
