package main

func main() {

	// For Loop => Similar to foreach in PHP

	users := []string{"John", "Jane", "Doe"}
	

	for key, user := range users {
		println("Key:", key, "User:", user)
	}

}
