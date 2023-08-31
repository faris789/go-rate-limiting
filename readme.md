To setup the project
```
go mod init github.com/faris789
go mod tidy
```

To run the program

```
go run main.go
```

To Check the Rate-limitng try hitting the API localhost URL in the terminal using curl
```
$ for i in {1..30}; do curl http://localhost:8081/sqs/metadata; done
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Successful","body":"Hi! You've reached the API. How can I help you?"}
{"status":"Request Failed","body":"The API is at full capacity, try again later."}
{"status":"Request Failed","body":"The API is at full capacity, try again later."}
{"status":"Request Failed","body":"The API is at full capacity, try again later."}
{"status":"Request Failed","body":"The API is at full capacity, try again later."}
{"status":"Request Failed","body":"The API is at full capacity, try again later."}
```

The below command gets only the status code 
```
for i in {1..30}; do curl -s -o /dev/null -I -w "%{http_code}" http://localhost:8081/sqs/metadata; echo ; done
```
