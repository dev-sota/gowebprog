# Go Web Programming with MySQL

```
$ go run data.go server.go
```

```
# Create
curl -i -X POST -H "Content-Type: application/json"  -d '{"content":"My first post","author":"Sau Sheong"}' http://127.0.0.1:8080/post/

# Read
curl -i -X GET http://127.0.0.1:8080/post/1

# Update
curl -i -X PUT -H "Content-Type: application/json"  -d '{"content":"Updated post","author":"Sau Sheong"}' http://127.0.0.1:8080/post/1

# Delete
curl -i -X DELETE http://127.0.0.1:8080/post/1
```
