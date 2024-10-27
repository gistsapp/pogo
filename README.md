# POGO

> A SQL utility library for Golang that will help you scale your codebase.

## Why using Pogo ?

Golang philosophy is basically "let's write simple stuff efficiently". For most people, it means writing everything by themselves. However, at Gists, we feel like writing everything from scratch is wrong when it comes to Web services. Go devs shouldn't have to write and maintain a code base that involves rewriting the same SQL boilerplate code over and over again. That's why we created POGO.

We manage that by just providing a wrapper around `scan()` while leveraging customizable structure tags to make the mapping between your query parameters and your data structure easy.
POGO is way more simpler than an ORM (for now its basically a one function library) while at the same time more scalable than writing your code using `pq`.

## How to use it ?

```bash
$ go get github.com/gistsapp/pogo
```

Then write a datastructure :

```go
type User struct {
  id string `pogo:"user_id"`
  name string `pogo:"name"`
}
```

Finally query your structure like so :

```go
type User struct {
	UserId string `pogo:"user_id"`
	Name   string `pogo:"name"`
}


func createDatabase() *pogo.Database {
	return pogo.NewDatabase("postgres", "postgres", "0.0.0.0", "5432", "pogo")
}

func GetUsers() ([]User, error) {
	db := createDatabase()
	users := make([]User, 0)
	err := pogo.SuperQuery(db, "SELECT :fields FROM users ORDER BY user_id", &users) // :fields will be replaced by the fields of the User struct
  return users, err
}
```
