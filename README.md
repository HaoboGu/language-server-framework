# language-server-framework
A language server framework written in Go which provides an easy way making language servers

# Design(WIP)
## Install
```shell
go get -u github.com/haobogu/lsframework
```
## Usage example
Example(go style pesudo code)
```go
import "github.com/haobogu/lsframework/server"

// Create your own processor
type MyProcessor struct {}

// Implement server.Processor interface
...
...

// Pass your processor to server host
s := server.NewBaseServer(port, ".", server.Config{}, MyProcessor{})

// Start your server
if err := s.Start(); err != nil {
    log.Error("The server crashed")
}
```

