package main

import "fmt"

type User struct {
	Id    int
	Name  string
	Token string
}

var usersFromBD = []User{
	{1, "Peter", "123456789"},
	{2, "Ann", "923456789"},
	{2, "Kate", "8923456789"},
}
var users Cache[int, *User]
var tokens Cache[string, int]

// обобщенная структура кэш
type Cache[K comparable, V any] struct {
	m map[K]V
}

func (c *Cache[K, V]) Init() {
	c.m = make(map[K]V)
}

func (c Cache[K, V]) Set(key K, value V) {
	c.m[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	k, ok := c.m[key]
	return k, ok
}
func main() {
	//init caches
	users.Init()
	tokens.Init()
	for _, user := range usersFromBD {
		users.Set(user.Id, &user)
		tokens.Set(user.Token, user.Id)
	}

	token := "123456789"
	//find
	userId, ok := tokens.Get(token)
	if ok {
		user, ok := users.Get(userId)
		if ok {
			fmt.Println("User name is:", user.Name)
		}

	}
	fmt.Println("User not Found")
}
