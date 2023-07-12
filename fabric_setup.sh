
cp ~/mmchain/smartcontract.go ~/mmchain/fabric-samples/asset-transfer-basic/chaincode-go/chaincode/.
cd ~/mmchain/fabric-samples/test-network
./network.sh up createChannel -c mychannel -ca 
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go/ -ccl go
cd ~/mmchain/fabric-samples/asset-transfer-basic/rest-api-go
go mod download
go run main.go


