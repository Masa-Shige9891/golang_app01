package main

func main() {

	// db before the server
	dbUrl := "postgres://postgres:postgres@localhost:5430/restgolang"
	db, err := getDB(dbUrl)
	if err != nil {
		panic("Failed to connect to the DB")
	}

	addr := ":8080"
	s := NewAPIServer(addr, db)
	s.Run()
}
