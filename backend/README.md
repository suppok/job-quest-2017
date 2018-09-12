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

__________
### Prerequisites
* Node.js
* MongoDB
### Install
* clone project
* ```npm install```
### Run
* Run MongoDB
    ##### If you want to run server
    * ```npm start```
    ##### If you want to run test
    * Before you run, you have to comment 3 lines of code in server.js that I comment in source code already. Because it's about caching
    * ```npm test```
### API
* ### GET
    * #### /todos 
    * get all todos
    * example: [http://localhost:3000/todos](http://localhost:3000/todos)
    * #### /todo/:todo_id
    * get single todo by id
    * example: [http://localhost:3000/1q2w3e4r](http://localhost:3000/1q2w3e4r)
    * It will get todo which id is '1q2w3e4r'
* ### POST
    * #### /todo
    * create new todo
    * send information by body of request
    * example: [http://localhost:3000/todo](http://localhost:3000/todos) ```(body = {'title': 'New Todo'})```
    * It will create a new todo that title is 'New Todo'
* ### PUT
    * #### /todo/:todo_id
    * update single todo
    * send information by body of request
    * example: [http://localhost:3000/todo/1q2w3e4r](http://localhost:3000/todos) ```(body = {'title': 'New Todo - Updated'})```
    * It will update title of a todo which id is '1q2w3e4r' from old title to 'New Todo - Updated'
    * #### /finish/:todo_id
    * update status of a single todo
    * example: [http://localhost:3000/finish/1q2w3e4r](http://localhost:3000/todos)
    * It will status of a todo which id is '1q2w3e4r' from 'In progress' to 'Done'
* ### DELETE
    * #### /todo/:todo_id
    * delete single todo
    * example: [http://localhost:3000/todo/1q2w3e4r](http://localhost:3000/todos)
    * It will delete a todo which id is '1q2w3e4r'
