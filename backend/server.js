let express = require('express'),
  app = express(),
  port = process.env.PORT || 3000,
  mongoose = require('mongoose'),
  bodyParser = require('body-parser'),
  morgan = require('morgan'),
  apicache = require('apicache'),
  cache = apicache.middleware
  
// mongoose instance connection url connection
mongoose.Promise = global.Promise
mongoose.connect('mongodb://localhost:27017/tododb')

let Todo = require('./api/models/todo.model')

app.use(bodyParser.urlencoded({ extended: true }))
app.use(bodyParser.json())

//For logging all request
app.use(morgan('combined'))

/**
 * For caching rest API; limit 1 request per second
 * Before run test, you have to comment these 3 lines below
 */
app.use(cache('1 seconds'))
app.get('/will-be-cached', (req, res) => {
  res.json({ success: true })
})

let routes = require('./api/routes/todo.routes') //importing route
routes(app); //register the route

let server = require('http').Server(app)

server.listen(port, () => {
	console.log('Todo RESTful API server started on: ' + port, new Date())
})

module.exports = server

