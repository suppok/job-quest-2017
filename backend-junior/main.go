package main

func main() {
	a := App{}
	a.Initialize("suppok", "1234", "todo_list_api")

	a.Run(":8080")
}
