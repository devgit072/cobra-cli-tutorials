Cobra-cli is undisputed tool to create powerfull modern day cli interfaces in Golang.     
All major companies like Kubertenetes, Docker uses Cobra cli tool to create cli.

Here are the steps to create cli using Cobra
1) go get -u github.com/spf13/cobra/cobra (Will create the cobra executables in $GOPATH/bin)
2) Typical golang organisation of Cobra based cli will look like:
```
  ▾ cliName/
    ▾ cmd/
        add.go
        yourCommandImpl.go
        commands.go
        root.go
      main.go
```
3) 