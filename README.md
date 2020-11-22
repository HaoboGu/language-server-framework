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
type MyServer struct {
    LanguageServerBase
}

func NewMyServer() *MyServer {
    config := lsframework.DefaultConfig
    server := MyServer{
        LanguageServerBase: lsframework.NewBaseServer(os.Stdin, os.Stdout, config)
    }
}

func (MyServer *s) Initialize(params InitializeParams) {
    // Do initialization
}

func (MyServer *s) Completion(param CompletionParams) {
    // Functional code for completion
}

func Main() {
    server := NewMyServer()
    // Start server, waiting for client's connection
    // Once the connection is established, the server will keep listening the requests and notifications
    if err:= server.Start(); err != nil {
        Logger.Error("The server crashed")
    }
    Logger.Info("The server is shut down")
    exit(0)
}

```