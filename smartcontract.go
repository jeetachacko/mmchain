package chaincode

import (
	"encoding/json"
	"fmt"
	"time"
        "strconv"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Record struct {
	DocType      string `json:"docType"`
	Id           string `json:"Id"`
	RecordingId  int `json:"recordingId"`
	UserId       string `json:"userId"`
	ContractType int    `json:"contractType"`
	Created  string `json:"created"`
}


func (s *SmartContract) Create_contract(ctx contractapi.TransactionContextInterface,  recordingId int, userId int, contractType int) error {

	current_time := time.Now().Local()
	current_date := current_time.Format("02-01-2006")
	record := &Record{
		DocType:        "record",
		RecordingId:    recordingId,
		UserId:         strconv.Itoa(userId),
		ContractType:   contractType,
		Created:    current_date,
	}
	recordJSON, err := json.Marshal(record)
	if err != nil {
		return err
	}
	id := strconv.Itoa(recordingId) + " " + strconv.Itoa(userId)

	return ctx.GetStub().PutState(id, recordJSON)
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*Record, error) {
	var records []*Record
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var record Record
		err = json.Unmarshal(queryResult.Value, &record)
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
		fmt.Printf("%v",records)
	}

	return records, nil
}

func (t *SmartContract) Get_contracts(ctx contractapi.TransactionContextInterface, userId int) ([]*Record, error) {
	queryString := fmt.Sprintf(`{"selector":{"docType":"record","userId":"%s"}}`, strconv.Itoa(userId))
	return getQueryResultForQueryString(ctx, queryString)
}

func getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*Record, error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	return constructQueryResponseFromIterator(resultsIterator)
}

func (s *SmartContract) Get_allcontracts(ctx contractapi.TransactionContextInterface) ([]*Record, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []*Record
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var record Record
		err = json.Unmarshal(queryResponse.Value, &record)
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
	}

	return records, nil
}

