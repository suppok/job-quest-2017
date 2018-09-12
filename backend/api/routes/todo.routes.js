'use strict';

module.exports = (app) => {
  let todo = require('../controllers/todo.controller')

  app.route('/todos')
    .get(todo.getAllTodos)

  app.route('/todo/:todo_id')
    .get(todo.getTodo)
    .put(todo.updateTodo)
    .delete(todo.deleteTodo)

  app.route('/todo')
    .post(todo.createNewTodo)

  app.route('/finish/:todo_id')
    .put(todo.finishTodo)
}