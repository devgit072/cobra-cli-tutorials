<h4>What is CobraCLI tool</h4>
Cobra-cli is undisputed tool to create powerfull modern day cli interfaces in Golang.     
All major companies like Kubernetes, Docker, CockroachDB uses Cobra cli tool to create cli.

<h6>Here are the steps to create cli using Cobra</h6>
1) go get -u github.com/spf13/cobra/cobra (Will create the cobra executables in $GOPATH/bin)  

2) Typical golang organisation of Cobra based cli will look like:

  

    ▾ cliName/
        ▾ cmd/
            add.go
            yourCommandImpl.go
            commands.go
            root.go
          main.go

3) `cobra init` will create the basic go file and directory organisation as mentioned above.    
    Use `-l none` to not generate default Apache license.
4) Inside root.go, mention all the usage and description.
5) Add the command that you want to add with `cobra add my_command.go`

I have created two command : math-fun(say.go) and add(add.go). Code are self documented with ample comments.
