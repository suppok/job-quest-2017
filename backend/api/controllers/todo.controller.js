'use strict';

var mongoose = require('mongoose'),
  TodoModel = mongoose.model('Todo');

module.exports = {

  getAllTodos: async (req, res) => {
    let result = await TodoModel.find({})
    return res.json(result)
  },

  getTodo: async (req, res) => {
    let result = await TodoModel.findById(req.params.todo_id)
    return res.json(result)
  },

  createNewTodo: async (req, res) => {
    let newTodo = new TodoModel(req.body);
    await newTodo.save()
    let result = {
      object: newTodo,
      message: 'Todo successfully created'
    }
    return res.json(result)
  },

  updateTodo: async (req, res) => {
    let todoResponse = await TodoModel.findById(req.params.todo_id)
    todoResponse.title = req.body.title
    await todoResponse.save()
    let result = {
      object: todoResponse,
      message: 'Todo successfully updated'
    }
    return res.json(result)
  },
  
  deleteTodo: async (req, res) => {
    let todoResponse = await TodoModel.remove({_id: req.params.todo_id})
    return res.json({ message: 'Todo successfully deleted' })
  },

  finishTodo: async (req, res) => {
    let todoResponse = await TodoModel.findOne({_id: req.params.todo_id})
    todoResponse.status = 'Done'
    await todoResponse.save()
    let result = {
      object: todoResponse,
      message: 'Todo is finished'
    }
    return res.json(result)
  }

}