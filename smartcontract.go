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
	RecordingId  int `json:"recordingId"`
	UserId       string `json:"userId"`
	ContractType int    `json:"contractType"`
	Created  string `json:"created"`
}


func (s *SmartContract) Create_contract(ctx contractapi.TransactionContextInterface,  recordingId int, userId int, contractType int) error {
	exists, err := s.RecordExists(ctx, strconv.Itoa(recordingId))
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the record %s already exists", strconv.Itoa(recordingId))
	}

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

	return ctx.GetStub().PutState(strconv.Itoa(recordingId), recordJSON)
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) ([]*Record, error) {
	var records []*Record
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var record Record
		//var r map[string]json.RawMessage
		err = json.Unmarshal(queryResult.Value, &record)
		//err = json.Unmarshal([]byte(queryResult.Value), &r)
		if err != nil {
			return nil, err
		}
		//delete(r, "docType")
		records = append(records, &record)
		//records = append(records, &r)
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


func (s *SmartContract) Get_contract_byRecordingId(ctx contractapi.TransactionContextInterface, recordingId int) (*Record, error) {
	recordJSON, err := ctx.GetStub().GetState(strconv.Itoa(recordingId))
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if recordJSON == nil {
		return nil, fmt.Errorf("the record %s does not exist", strconv.Itoa(recordingId))
	}

	var record Record
	err = json.Unmarshal(recordJSON, &record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}


func (s *SmartContract) Get_allcontracts(ctx contractapi.TransactionContextInterface) ([]*Record, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all records in the chaincode namespace.
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

func (s *SmartContract) RecordExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
        recordJSON, err := ctx.GetStub().GetState(id)
        if err != nil {
                return false, fmt.Errorf("failed to read from world state: %v", err)
        }

        return recordJSON != nil, nil
}


