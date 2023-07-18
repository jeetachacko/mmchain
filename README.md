# MMChain 

A blockchain-based music usage tracking system that ensures a private and secure approach to enable provenance tracking in the music industry. This set of scripts helps to quickly setup a Hyperledger Fabric network on a single node, deploy the smart contract and start a REST server with endpoints for invoking the smart contract functions. The scripts, rest server, and smart contract implementations are based on [fabric-samples](https://github.com/hyperledger/fabric-samples/tree/c04253d55407e5fe7217d4931738fe7273b4a8a5)

## Steps for Setup

Clone the repository to home directory (~)
```shell
git clone https://github.com/jeetachacko/mmchain.git
```
```shell
cd mmchain
```

Install Prerequisites (the scripts are interactive and need user inputs like yes, enter to continue and setting a new/old password):  
```shell
./prerequisites1.sh
```
```shell
cd mmchain
```
```shell
./prerequisites2.sh
```

Fabric Setup: 
```shell
source ~/.bashrc
```
```shell
source ~/.profile
```
```shell
./fabric_setup.sh
```
Open a new terminal to execute the API requests. The value of the requests can be edited in the script as desired.
```shell
cd mmchain
```
```shell
./initialize_ledger.sh
```
```shell
./api_requests.sh
```

To delete and restart the Fabric blockchain network
```shell
./fabric_delete.sh
```
```shell
./fabric_setup.sh
```
