package main

func main() {
	run_containers := []string{"red", "redslv", "mycat", "mycatslv", "porn", "not"}
	for i, x := range run_containers {
		println(i, x)
	}
}
