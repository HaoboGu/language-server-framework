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

type MyServer struct {
    // Base server is an empty server
    server.LanguageServer
}

func NewMyServer(port int, wd string, config server.Config) *MyServer {
    s := MyServer{
        LanguageServerBase: lsframework.NewBaseServer(port, wd, config)
    }
}

func (MyServer *s) Initialize(params InitializeParams) {
    // Do initialization
}

func (MyServer *s) Completion(param CompletionParams) {
    // Your completion code
}

func Main() {
    config := lsframework.DefaultConfig
    // Your server will listen to this port via tcp
    port := 12345 
    // Your server's working directory
    wd := "/"
    server := NewMyServer(port, wd, config)

    // Start server, waiting for client's connection
    // Once the connection is established, the server will keep listening the requests and notifications
    if err:= server.Start(); err != nil {
        log.Error("The server crashed")
    }
    log.Info("The server is shut down")
    os.Exit(0)
}
```

