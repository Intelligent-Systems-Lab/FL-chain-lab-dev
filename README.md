# FL-chain-lab-dev

## install

```bash=
mkdir -p $GOPATH/src/github.com

cd $GOPATH/src/github.com

git clone https://github.com/Intelligent-Systems-Lab/FL-chain-lab-dev.git

cd abci-client

go mod init github.com/me/example
go build -o kv-buildin  main.go app.go

./kv-buildin
```
