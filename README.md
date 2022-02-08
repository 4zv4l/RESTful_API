# RESTful_API
A basic RESTful API made in Go
# What is RESTful API?
RESTful API is a set of rules for designing a web service.
Allowing you to interact with the data in a uniform way.
# Basic API call
```
curl -X GET http://localhost:8080/language/1
```
# Response
```
{
  "id": 1,
  "name": "Go",
  "description": "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."
}
```