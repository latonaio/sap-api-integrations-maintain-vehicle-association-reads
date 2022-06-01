package main

import (
	sap_api_caller "sap-api-integrations-maintain-vehicle-association-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-maintain-vehicle-association-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Maintain_Vehicle_Association_Vehicle_Invoice_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/sapdvh/odata/v4/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"VehicleInvoice", "VehicleProcurement",
		}
	}

	caller.AsyncGetMaintainVehicleAssociation(
		inoutSDC.VehicleInvoice.VehicleID,
		inoutSDC.VehicleInvoice.VehicleProcurement.SupplierID,
		accepter,
	)
}
