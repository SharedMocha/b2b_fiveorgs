/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	_ "encoding/pem"
	"strings"
	"encoding/pem"
	"crypto/x509"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type B2B_3B11 struct {
	OrderNo  	    	string 				`json:"orderno"`
	LineNo  	    	string 				`json:"lineno"`
	PID  	    	    string 				`json:"pid"`
	Quantity  	    	string 				`json:"quantity"`

}
type B2B_3B3 struct {
	OrderConfirm  	    string 				`json:"orderconfirm"`
	ETA  	    	    string 				`json:"eta"`
	POD  	    	    string 				`json:"pod"`
}
type B2B_3B13 struct {
	ShipNotice  	    string 				`json:"shipnotice"`
	FELocation  	    string 				`json:"felocation"`
	FEName  	        string 				`json:"fename"`

}

type RMA struct {
	RMANO 			string 				`json:"rmano"`
	B2B_3B11 		B2B_3B11 			`json:"3b11"`
	B2B_3B3 		B2B_3B3 			`json:"3b3"`
	B2B_3B13 		B2B_3B13 			`json:"3b13"`
}

const indexName = `RMANO`

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "create" {
		// Make payment of X units from A to B
		return t.create(stub, args)
	} else if function == "update3b11" {
		// Deletes an entity from its state
		return t.update3B11(stub, args)
	} else if function == "update3b13" {
		// Deletes an entity from its state
		return t.update3B13(stub, args)
	} else if function == "update3b3OrderConfirm" {
		// Deletes an entity from its state
		return t.update3B3OrderConfirm(stub, args)
	}  else if function == "update3b3ETA" {
		// Deletes an entity from its state
		return t.update3B3ETA(stub, args)
	}  else if function == "update3b3OrderPOD" {
		// Deletes an entity from its state
		return t.update3B3POD(stub, args)
	}  else if function == "updateFElocation" {
		// Deletes an entity from its state
		return t.update3B3POD(stub, args)
	}  else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"move\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) create(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	rmano := args[0]

	value, err := json.Marshal(RMA{RMANO:rmano})
	if err != nil {
		return shim.Error("cannot marshal")
	}
	key, err := stub.CreateCompositeKey(indexName, []string{rmano})
	if err != nil {
		return shim.Error("cannot create composite key")
	}
	err = stub.PutState(key, []byte(value))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
func (t *SimpleChaincode) update3B11(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	/*
	cert, _ := stub.GetCreator()
	org := getOrganization(cert)

	if org != "cisco" {
		return pb.Response{Message:"only cisco is authorized to update via 3B11", Status:401}
	}
	*/

	rmano := args[0]
	orderno := args[1]
	lineno := args[2]
	pid := args[3]
	quantity := args[4]



	rma, err := getRMA(stub, rmano)
	if err != nil {
		return shim.Error("cannot get rma")
	}

	/*
	if part == "engine" {
		car.Engine.Maker = maker
	} else {
		return pb.Response{Message:"cannot set anything but engine", Status:400}
	}
	*/
	rma.B2B_3B11.OrderNo = orderno
	rma.B2B_3B11.LineNo = lineno
	rma.B2B_3B11.PID = pid
	rma.B2B_3B11.Quantity = quantity




	err = putRMA(stub, rma)
	if err != nil {
		return shim.Error("cannot put rma")
	}

	return shim.Success(nil)
}
func (t *SimpleChaincode) update3B13(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	/*
	cert, _ := stub.GetCreator()
	org := getOrganization(cert)

	if org != "cisco" {
		return pb.Response{Message:"only cisco is authorized to update via 3B11", Status:401}
	}
	*/

	rmano := args[0]
	status := args[1]
	//felocation :=  args[2]

	rma, err := getRMA(stub, rmano)
	if err != nil {
		return shim.Error("cannot get rma")
	}

	/*
	if part == "engine" {
		car.Engine.Maker = maker
	} else {
		return pb.Response{Message:"cannot set anything but engine", Status:400}
	}
	*/
	rma.B2B_3B13.ShipNotice = status
	//rma.B2B_3B13.FELocation = felocation


	err = putRMA(stub, rma)
	if err != nil {
		return shim.Error("cannot put rma")
	}

	return shim.Success(nil)
}
func (t *SimpleChaincode) updateFElocation(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	/*
	cert, _ := stub.GetCreator()
	org := getOrganization(cert)

	if org != "cisco" {
		return pb.Response{Message:"only cisco is authorized to update via 3B11", Status:401}
	}
	*/

	rmano := args[0]
	fename := args[1]
	felocation :=  args[2]

	rma, err := getRMA(stub, rmano)
	if err != nil {
		return shim.Error("cannot get rma")
	}

	/*
	if part == "engine" {
		car.Engine.Maker = maker
	} else {
		return pb.Response{Message:"cannot set anything but engine", Status:400}
	}
	*/
	rma.B2B_3B13.FELocation = felocation
	rma.B2B_3B13.FEName = fename



	err = putRMA(stub, rma)
	if err != nil {
		return shim.Error("cannot put rma")
	}

	return shim.Success(nil)
}
func (t *SimpleChaincode) update3B3OrderConfirm(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	/*
	cert, _ := stub.GetCreator()
	org := getOrganization(cert)

	if org != "cisco" {
		return pb.Response{Message:"only cisco is authorized to update via 3B11", Status:401}
	}
	*/

	rmano := args[0]
	orderconfirmstatus := args[1]


	rma, err := getRMA(stub, rmano)
	if err != nil {
		return shim.Error("cannot get rma")
	}

	/*
	if part == "engine" {
		car.Engine.Maker = maker
	} else {
		return pb.Response{Message:"cannot set anything but engine", Status:400}
	}
	*/
	rma.B2B_3B3.OrderConfirm = orderconfirmstatus

	err = putRMA(stub, rma)
	if err != nil {
		return shim.Error("cannot put rma")
	}

	return shim.Success(nil)
}
func (t *SimpleChaincode) update3B3ETA(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	/*
	cert, _ := stub.GetCreator()
	org := getOrganization(cert)

	if org != "cisco" {
		return pb.Response{Message:"only cisco is authorized to update via 3B11", Status:401}
	}
	*/

	rmano := args[0]
	eta := args[1]


	rma, err := getRMA(stub, rmano)
	if err != nil {
		return shim.Error("cannot get rma")
	}

	/*
	if part == "engine" {
		car.Engine.Maker = maker
	} else {
		return pb.Response{Message:"cannot set anything but engine", Status:400}
	}
	*/
	rma.B2B_3B3.ETA = eta

	err = putRMA(stub, rma)
	if err != nil {
		return shim.Error("cannot put rma")
	}

	return shim.Success(nil)
}
func (t *SimpleChaincode) update3B3POD(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	/*
	cert, _ := stub.GetCreator()
	org := getOrganization(cert)

	if org != "cisco" {
		return pb.Response{Message:"only cisco is authorized to update via 3B11", Status:401}
	}
	*/

	rmano := args[0]
	pod := args[1]


	rma, err := getRMA(stub, rmano)
	if err != nil {
		return shim.Error("cannot get rma")
	}

	/*
	if part == "engine" {
		car.Engine.Maker = maker
	} else {
		return pb.Response{Message:"cannot set anything but engine", Status:400}
	}
	*/
	rma.B2B_3B3.POD = pod

	err = putRMA(stub, rma)
	if err != nil {
		return shim.Error("cannot put rma")
	}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	it, err := stub.GetStateByPartialCompositeKey(indexName, []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer it.Close()

	arr := []RMA{}
	for it.HasNext() {
		next, err := it.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		var rma RMA
		err = json.Unmarshal(next.Value, &rma)
		if err != nil {
			return shim.Error(err.Error())
		}

		arr = append(arr, rma)
	}

	ret, err := json.Marshal(arr)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(ret)
}
func getRMA(stub shim.ChaincodeStubInterface, rmano string) (RMA, error) {
	key, err := stub.CreateCompositeKey(indexName, []string{rmano})
	if err != nil {
		return RMA{}, err
	}

	value, err := stub.GetState(key)
	if err != nil {
		return RMA{}, err
	}

	var rma RMA
	err = json.Unmarshal(value, &rma)
	if err != nil {
		return RMA{}, err
	}

	return rma, nil
}
func putRMA(stub shim.ChaincodeStubInterface, rma RMA) error {
	key, err := stub.CreateCompositeKey(indexName, []string{rma.RMANO})
	if err != nil {
		return err
	}

	value, err := json.Marshal(rma)
	if err != nil {
		return err
	}

	err = stub.PutState(key, []byte(value))
	if err != nil {
		return err
	}

	return nil
}

func getOrganization(certificate []byte) string {
	data := certificate[strings.Index(string(certificate), "-----") : strings.LastIndex(string(certificate), "-----")+5]
	block, _ := pem.Decode([]byte(data))
	cert, _ := x509.ParseCertificate(block.Bytes)
	organization := cert.Issuer.Organization[0]
	//logger.Info("getOrganization: " + organization)

	ret := strings.Split(organization, ".")[0]

	return ret
}


func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
