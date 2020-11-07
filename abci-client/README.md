# abci-test

This is an exercise in [Creating a built-in application in Go](https://docs.tendermint.com/master/tutorials/go-built-in.html)

In this exercise, we will build a build-in application. Before start, we need to understand abci and learn golang, because they only provide abci-lib for golang in a build-in application.

### What is different between build-in application and non build-in application?

* `build-in application` : build-in application combines a tendermint node when we execute an application.

ex:
```bash=
$ tendermint init # this create configuration file in $HOME/.tendermint
$ ./build-in-app  # this app lanch a node and application with configuration file in $HOME/.tendermint/config/config.toml
OR
$ ./build-in-app --config <Path to config.toml file>
```

* `non build-in application` : non build-in application need to launch tendermint node and application separately.

ex:
```bash=
$ tendermint init # this create configuration file in $HOME/.tendermint
$ tendermint node --proxy_app=kvstore

$ ./non-build-in-app  # this app lanch a node and application with configuration file in $HOME/.tendermint/config/config.toml
OR
$ ./non-build-in-app --config <Path to config.toml file>
```

### Not familiar with go?

[my note](https://hackmd.io/@tonyguo/Hkm0hFWdv)

### What is abci?

abci protocol was implemented with GRPC, connect application (abci-application) and tendermint core (consensus)

[Application Architecture Guide](https://docs.tendermint.com/master/app-dev/app-architecture.html)

## Hand on 

In the host
```bash=
# launch docker
wget -O - https://raw.githubusercontent.com/Intelligent-Systems-Lab/FL-chain-lab-dev/main/docker/run.sh | bash
```

In the container
```bash=
mkdir -p $GOPATH/src/github.com

cd $GOPATH/src/github.com

git clone https://github.com/Intelligent-Systems-Lab/FL-chain-lab-dev.git

cd FL-chain-lab-dev/abci-client

go mod init github.com/me/example
go build -o kv-buildin  main.go app.go

./kv-buildin
```
## Types of abci

Still working ☕ ☕

## Wireshark monitor

Still working ☕ ☕