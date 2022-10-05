# Go Chat with bot

# How to run
1) Run backend
```bash
cd back && PORT=9090 go run cmd/main.go
```

2) Run bot
```bash
cd bot && BACKEND_URL=localhost:9090 go run cmd/main.go
```

3) Run frontend
```bash
cd front && PORT=8080 BACKEND_URL=localhost:9090 go run cmd/main.go
```

4) Open a web browser and navigate to `localhost:8080?name=<your_name>`