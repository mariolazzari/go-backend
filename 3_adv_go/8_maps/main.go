package main

func main() {
	m := make(map[string]int)
	// m["a"] = 1
	// a, exists := m["a"]

	if _, ok := m["a"]; ok {
		// ...
	}

	delete(m, "a")
	clear(m)
}
