# POGO

> Stop writing 20 lines of golang to just `SELECT` your data by getting rid of the `scan()` function.

## Why using Pogo ?

Golang philosophy is basically "let's write simple stuff efficiently". For most people, it means writing everything by themself but at Gists, we feel like that this take is wrong. Go devs shouldn't have to write and maintain a code base that involves 20 lines functions for just a classic `SELECT` or `INSERT` sql queries.

We manage that by just providing a wrapper around `scan()` and customizable structure tags to making easy the mapping between your query parameters and your data structure. Therefore POGO is way more simple than an ORM but at the same time more scalable than writing your code using `pq`.

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
func (u *User) Save() error {
  return pogo.Save("INSERT INTO users(user_id, name) VALUES (:user_id, :values)")
}



func (u *User) Get() error {
  return pogo.Save("SELECT user_id, name")
}
```
