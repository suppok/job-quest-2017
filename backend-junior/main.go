package main

func main() {
	a := App{}
	// You need to set your Username and Password here
	a.Initialize("suppok", "1234", "todo_list_api")

	a.Run(":8080")
}
