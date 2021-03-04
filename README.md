# Graceful Shutdown In Go
This is the Graceful Shutdown sample code in Go, feel free to learn.
## How to run
```go run main.go```
## How to test
1. Run the service
2. Open the other terminal, request to the API using this command line
```curl -X GET http://localhost:8081/api/v1/test-graceful-shutdown```
3. Back to the service terminal Stop the service
```Ctrl + C```
4. Waiting the response from api

### You can access this blog post for details.
https://anhlamweb.com/bai-viet-75/tim-hieu-ve-graceful-shutdown-graceful-shutdown-trong-golang.html
