package tests

import (
	"testing"

	"github.com/gistsapp/pogo/pogo"
	_ "github.com/lib/pq"
)

func TestSelectUsers(t *testing.T) {
	t.Run("Get all fields", func(t *testing.T) {
		db := CreateDatabase()

		ClearDatabase(db)
		SeedDatabase(db)
		type User struct {
			UserId string `pogo:"user_id"`
			Name   string `pogo:"name"`
		}
		users := make([]User, 0)
		err := pogo.SuperQuery(db, "SELECT user_id, name FROM users ORDER BY user_id", &users)
		if err != nil {
			t.Fatalf("Failed to run super query: %v", err)
		}

		expected_names := []string{"test1", "test2", "test3"}

		for i, user := range users {
			if user.Name != expected_names[i] {
				t.Fatalf("wrong name, expected %v but got %v", expected_names[i], user.Name)
			}
		}

	})

	t.Run("Get all fields with :field", func(t *testing.T) {

		db := CreateDatabase()
		ClearDatabase(db)
		SeedDatabase(db)
		type User struct {
			UserId string `pogo:"user_id"`
			Name   string `pogo:"name"`
		}
		users := make([]User, 0)
		err := pogo.SuperQuery(db, "SELECT :fields FROM users ORDER BY user_id", &users)
		if err != nil {
			t.Fatalf("Failed to run super query: %v", err)
		}

		expected_names := []string{"test1", "test2", "test3"}

		for i, user := range users {
			if user.Name != expected_names[i] {
				t.Fatalf("wrong name, expected %v but got %v", expected_names[i], user.Name)
			}
		}

	})
}
