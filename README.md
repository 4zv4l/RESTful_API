# RESTful_API
A basic RESTful API made in Go
# What is RESTful API?
RESTful API is a set of rules for designing a web service.
Allowing you to interact with the data in a uniform way.
# Basic API calls
**CREATE**
```
curl -X POST -H "Content-Type: application/json" -d '{"id": "3","Title": "C","desc": "A cool language","Content": "stuff about C",}' http://localhost:8080/language
```
**READ**
```
curl -X GET http://localhost:8080/language/1
```
```
{
  "id": 1,
  "Title": "Go",
  "desc": "A nice language !",
  "Content": "Some text about Go",
}
```
**UPDATE**
```
curl -X PUT -H "Content-Type: application/json" -d '{"id": "1","Title": "Go","desc": "A SUPER language","Content": "Changed text",}' http://localhost:8080/language/1
```
**DELETE**
```
curl -X DELETE http://localhost:8080/language/1
```
**GET all**
```
curl -X GET http://localhost:8080/languages
```
```
[
  {
    "id": 1,
    "Title": "Go",
    "desc": "A nice language !",
    "Content": "Some text about Go",
  },
  {
    "id": 2,
    "Title": "Rust",
    "desc": "a cool language",
    "Content": "stuff about Rust",
  }
]
```