var chai = require('chai');
var chaiHttp = require('chai-http');
var mongoose = require("mongoose");

var server = require('../server');
var TodoModel = require("../api/models/todo.model");
var TodoController = require("../api/controllers/todo.controller")

var should = chai.should();
chai.use(chaiHttp);

describe('Todos', () => {

  TodoModel.find({}).remove().exec()

  beforeEach((done) => {
    let newTodo = new TodoModel({title: 'Test Todo'});
    newTodo.save()
    done()
  });
  afterEach((done) => {
    TodoModel.find({}).remove().exec()
    done();
  });

  it('should list ALL Todos on /todos GET', (done) => {
    chai.request(server)
      .get('/todos')
      .end((err, res) => {
        res.should.have.status(200);
        res.should.be.json;
        res.body.should.be.a('array');
        res.body[0].should.have.property('_id');
        res.body[0].should.have.property('title');
        res.body[0].should.have.property('status');
        res.body[0].should.have.property('created_date');
        res.body[0].title.should.equal('Test Todo');
        res.body[0].status.should.equal('In progress');
        done();
      });
  });

  it('should list a SINGLE Todo on /todo/<id> GET', (done) =>{
    let newTodo = new TodoModel({title: 'New Todo'});
    newTodo.save((err, data) => {
      chai.request(server)
        .get('/todo/'+data._id)
        .end((err, res) =>{
          res.should.have.status(200);
          res.should.be.json;
          res.body.should.be.a('object');
          res.body.should.have.property('_id');
          res.body.should.have.property('title');
          res.body.should.have.property('status');
          res.body.should.have.property('created_date');
          res.body.title.should.equal('New Todo');
          res.body.status.should.equal('In progress');
          res.body._id.should.equal(data.id);
          done();
        });
    });
  });

  it('should add a SINGLE Todo on /todo POST', (done) =>{
    chai.request(server)
      .post('/todo')
      .send({'title': 'POST Todo'})
      .end((err, res) => {
        res.should.have.status(200);
        res.should.be.json;
        res.body.should.be.a('object');
        res.body.should.have.property('object');
        res.body.should.have.property('message');
        res.body.message.should.equal('Todo successfully created');
        res.body.object.should.be.a('object');
        res.body.object.should.have.property('title');
        res.body.object.should.have.property('status');
        res.body.object.should.have.property('_id');
        res.body.object.should.have.property('created_date');
        res.body.object.title.should.equal('POST Todo');
        res.body.object.status.should.equal('In progress');
        done();
      });
  });

  it('should update a SINGLE todo on /todo/<id> PUT', (done) => {
    chai.request(server)
      .get('/todos')
      .end((err, res) => {
        chai.request(server)
          .put('/todo/'+res.body[0]._id)
          .send({'title': 'Updated'})
          .end((error, response) => {
            response.should.have.status(200);
            response.should.be.json;
            response.body.should.be.a('object');
            response.body.should.have.property('object');
            response.body.should.have.property('message');
            response.body.message.should.equal('Todo successfully updated');
            response.body.object.should.be.a('object');
            response.body.object.should.have.property('title');
            response.body.object.should.have.property('_id');
            response.body.object.should.have.property('status');
            response.body.object.should.have.property('created_date');
            response.body.object.title.should.equal('Updated');
            done();
        });
      });
  });

  it('should delete a SINGLE todo on /todo/<id> DELETE', (done) => {
    chai.request(server)
      .get('/todos')
      .end((err, res) => {
        chai.request(server)
          .delete('/todo/'+res.body[0]._id)
          .end((error, response) => {
            response.should.have.status(200);
            response.should.be.json;
            response.body.should.be.a('object');
            response.body.should.have.property('message');
            response.body.message.should.equal('Todo successfully deleted');
            done();
        });
      });
  });

  it('should change status of a SINGLE todo on /finish/<id> PUT', (done) => {
    chai.request(server)
      .get('/todos')
      .end((err, res) => {
        chai.request(server)
          .put('/finish/'+res.body[0]._id)
          .end((error, response) => {
            response.should.have.status(200);
            response.should.be.json;
            response.body.should.be.a('object');
            response.body.should.have.property('object');
            response.body.should.have.property('message');
            response.body.message.should.equal('Todo is finished');
            response.body.object.should.be.a('object');
            response.body.object.should.have.property('title');
            response.body.object.should.have.property('_id');
            response.body.object.should.have.property('status');
            response.body.object.should.have.property('created_date');
            response.body.object.status.should.equal('Done');
            done();
        });
      });
  });

});