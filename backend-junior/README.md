### backend

#### TODO

* Write a To-do API in language of your choice having following endpoints

  * GET / Get all todos
  * POST / Create new todo
  * DELETE /:id Delete specific todo item
  * POST /:id/toggle Toggle the state of specific todo item (Todo / Done)

* Async/Await
* ES6, 7 syntax
* API Testing
* HTTP request logger
* Caching REST API

### Bonus

* express-validator (to add validations for limit length of characters in todo)
* CSRF (for cross-site fraud prevention)

___
### Prerequisites
* Golang
* Mysql
### Install
* clone project
* ```go install```
### Run
##### If you want to run server
* ```go run main.go model.go app.go```
##### If you want to run test
* ```go test -v```
### API
* ### GET
    * #### /todos 
    * get all todos
    * example: [http://localhost:8080/todos](http://localhost:8080/todos)
    * #### /todo/{id:[0-9]+}
    * get single todo by id
    * example: [http://localhost:8080/1234](http://localhost:8080/1234)
    * It will get todo which id is '1234'
* ### POST
    * #### /todo
    * create new todo
    * send information by body of request
    * example: [http://localhost:8080/todo](http://localhost:8080/todo) ```(body = {'title': 'New Todo'})```
    * It will create a new todo that title is 'New Todo'
* ### PUT
    * #### /todo/{id:[0-9]+}
    * update single todo
    * send information by body of request
    * example: [http://localhost:8080/todo/1234](http://localhost:8080/todo) ```(body = {'title': 'New Todo - Updated'})```
    * It will update title of a todo which id is '1234' from old title to 'New Todo - Updated'
* ### DELETE
    * #### /todo/{id:[0-9]+}
    * delete single todo
    * example: [http://localhost:8080/todo/1234](http://localhost:8080/todo)
    * It will delete a todo which id is '1234'
