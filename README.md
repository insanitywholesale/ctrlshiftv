# ctrlshiftv

paste microservice written in Go

## example use

run with `go run main.go` and in a different terminal send data to it

with curl:
```bash
curl --header "Content-Type: application/json" -d '{"content": "mypaste"}' http://localhost:8000/
```

which should yeild a response like below:
```json
{"code":"297ZHGHMg","content":"mypaste","created_at":1597777998}
```

and then view the contents with:
```bash
curl http://localhost:8000/297ZHGHMg
```

which should return:
```bash
mypaste
```
