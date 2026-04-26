package handlers

import "sync"

type User struct { //Data type for users. Password stored as hash, not as string
	Username string `json:"username"`
	Password string `json:"password"`
}

type Album struct { //Data type for albums
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var (
	users  = []User{} //User slice
	albums = []Album{ //Albums slice. Preinitialized for sake of demonstartion
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Time Out", Artist: "Dave Brubeck", Price: 37.99},
		{ID: "3", Title: "Flying Beagle", Artist: "Himiko Kikuchi", Price: 69.99},
	}

	revokedTokens = make(map[string]bool) //Tokens that have been revoked and no longer work

	signKey = []byte("PleaseGod12345Ilikh") //Sign key to generate tokens

	mu sync.Mutex //Mutex
)
