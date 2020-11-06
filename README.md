# FL-chain-lab-dev

## install

```bash=
mkdir -p $GOPATH/src/github.com

cd $GOPATH/src/github.com

git clone https://github.com/Intelligent-Systems-Lab/FL-chain-lab-dev.git

cd FL-chain-lab-dev/app

# install dependency
go get ./...

go install

cd ../abci-client

go run main.go
```
