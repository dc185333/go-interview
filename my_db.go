package main

// myDB simulates a simple key-value store that
// stores string values at string keys.
type myDB struct {
	store map[string]string
}

// Get retrieves the value associated with the
// passed key.
func (db myDB) Get(key string) (string, bool) {
	val, ok := db.store[key]
	return val, ok
}

// Set adds a new string value to the key-value
// store at the passed key.
func (db myDB) Set(key, value string) {
	db.store[key] = value
}

// Load seeds the key-value store with initial
// key-value pairs {a: valA, b: valB}.
func (db myDB) Load() {
	db.store = map[string]string{
		"a": "valA",
		"b": "valB",
	}
}
