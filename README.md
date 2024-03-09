## go-webapp
## crash course 42Bangkok -- Go, Postgres, Docker, and Nginx
#basic command
- build
```bash
go build main.go
```

- execute
```bash
./main.go
```

- build and execute
```bash
go run main.go 
```


- install library
```bash
go mod init github/xx/xx

```

# Tasks 1
1. Create new handler
2. Make this handler return text:
    "Hellooo from <your nickname>!"
3. Connect this handler to url/about

# Tasks 2
1. Change the IserPage handler
2. It should find user in users map and show page with that user


# Status Code
404 no handler