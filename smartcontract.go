package chaincode

import (
	"encoding/json"
	"fmt"
	"time"
        "strconv"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Record struct {
	ContractType int    `json:"contractType"`
	CreatedDate  string `json:"createdDate"`
	RecordingId  int `json:"recordingId"`
	UserId       string `json:"userId"`
}


func (s *SmartContract) Create_contract(ctx contractapi.TransactionContextInterface,  recordingId int, userId int, contractType int) error {
	exists, err := s.RecordExists(ctx, strconv.Itoa(userId))
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the record %s already exists", strconv.Itoa(userId))
	}

	current_time := time.Now().Local()
	current_date := current_time.Format("02-01-2006")
	record := Record{
		RecordingId:    recordingId,
		UserId:         strconv.Itoa(userId),
		ContractType:   contractType,
		CreatedDate:    current_date,
	}
	recordJSON, err := json.Marshal(record)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(strconv.Itoa(userId), recordJSON)
}

func (s *SmartContract) Get_contracts(ctx contractapi.TransactionContextInterface, userId int) (*Record, error) {
	recordJSON, err := ctx.GetStub().GetState(strconv.Itoa(userId))
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if recordJSON == nil {
		return nil, fmt.Errorf("the record %s does not exist", strconv.Itoa(userId))
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


