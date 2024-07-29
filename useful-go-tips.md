# Useful Go Tips

## Command Line Go Commands
```bash

# Go module commands
## Used to work with modules
go mod

## Initializes a new module in current directory
go mod init

## Install the current directory
go install .

# Random Ones

## Add to .bashrc to add GOPATH
export PATH=$PATH:$(go env GOPATH)/bin
## Downloading default go library if GOPROXY is set (like at work)
GOPROXY=https://proxy.golang.org,direct go install [package-link]

git remote set-url "origin" git@github.com-personal:jemajet/web-dev-with-go.git
```


## Web Development Go Tips
