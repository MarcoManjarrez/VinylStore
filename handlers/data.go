package handlers

import "sync"

// User represents the credentials for a user
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Album represents a vinyl record in the store
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// In-memory data stores
var (
	users = []User{}
	// Pre-populate with data from the instructions
	albums = []Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Time Out", Artist: "Dave Brubeck", Price: 37.99},
		{ID: "3", Title: "Flying Beagle", Artist: "Himiko Kikuchi", Price: 69.99},
	}

	// Map to keep track of revoked tokens (Logout implementation)
	revokedTokens = make(map[string]bool)

	// Secret key for JWT signing
	signKey = []byte("PleaseGod12345Ilikh")

	// Mutex to ensure thread safety when modifying slices and maps simultaneously
	mu sync.Mutex
)
