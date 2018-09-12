'use strict';
var mongoose = require('mongoose');
var Schema = mongoose.Schema;

var TodoSchema = new Schema({
  title: {
    type: String,
  },
  status: {
    type: String,
    enum: ['In progress', 'Done'],
    default: 'In progress'
  },
  created_date: {
    type: Date,
    default: Date.now
  }
  
});

module.exports = mongoose.model('Todo', TodoSchema);