package responses

type VehicleProcurement struct {
	Value []struct {
		VehicleID                    string `json:"vehicleID"`
		DocumentNumber               string `json:"documentNumber"`
		ItemNumber                   string `json:"itemNumber"`
		ID                           string `json:"id"`
		SupplierID                   string `json:"supplierID"`
		SupplierName                 string `json:"supplierName"`
		CostCenterNumber             string `json:"costCenterNumber"`
		CostCenterDescription        string `json:"costCenterDescription"`
		MileageLimitFrequencyCode    string `json:"mileageLimitFrequencyCode"`
		MileageLimitFrequencyName    string `json:"mileageLimitFrequencyName"`
		Date                         string `json:"date"`
		Amount                       string `json:"amount"`
		Currency                     string `json:"currency"`
		PlantID                      string `json:"plantID"`
		PlantName                    string `json:"plantName"`
		TypeCode                     string `json:"typeCode"`
		TypeName                     string `json:"typeName"`
		StatusCode                   string `json:"statusCode"`
		StatusName                   string `json:"statusName"`
		AssetNumber                  string `json:"assetNumber"`
		InternalOrderNumber          string `json:"internalOrderNumber"`
		SourceSystemID               string `json:"sourceSystemID"`
		OrganizationID               string `json:"organizationID"`
		OrganizationName             string `json:"organizationName"`
		GroupID                      string `json:"groupID"`
		GroupName                    string `json:"groupName"`
		CompanyCode                  string `json:"companyCode"`
		CompanyName                  string `json:"companyName"`
		AdditionalText               string `json:"additionalText"`
		LeasingContractID            string `json:"leasingContractID"`
		ContractStartDate            string `json:"contractStartDate"`
		ContractEndDate              string `json:"contractEndDate"`
		PaymentFrequencyCode         string `json:"paymentFrequencyCode"`
		PaymentFrequencyName         string `json:"paymentFrequencyName"`
		ServiceAmount                string `json:"serviceAmount"`
		ServiceRateFrequencyCode     string `json:"serviceRateFrequencyCode"`
		ServiceRateFrequencyName     string `json:"serviceRateFrequencyName"`
		EarlyTerminationChargeAmount string `json:"earlyTerminationChargeAmount"`
		MileageLimit                 string `json:"mileageLimit"`
		MileageLimitUomCode          string `json:"mileageLimitUomCode"`
		MileageLimitUomName          string `json:"mileageLimitUomName"`
		DocumentTypeCode             string `json:"documentTypeCode"`
		DocumentTypeName             string `json:"documentTypeName"`
		DeliveryDate                 string `json:"deliveryDate"`
		TaxCode                      string `json:"taxCode"`
		TaxName                      string `json:"taxName"`
		TaxAmount                    string `json:"taxAmount"`
		CreatedAt                    string `json:"createdAt"`
		ChangedAt                    string `json:"changedAt"`
	} `json:"value"`
}
