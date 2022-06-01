package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey  string `json:"connection_key"`
	Result         bool   `json:"result"`
	RedisKey       string `json:"redis_key"`
	Filepath       string `json:"filepath"`
	VehicleInvoice struct {
		VehicleID                string `json:"vehicleID"`
		Number                   string `json:"number"`
		ItemNumber               string `json:"itemNumber"`
		ID                       string `json:"id"`
		CategoryCode             string `json:"categoryCode"`
		CategoryName             string `json:"categoryName"`
		Date                     string `json:"date"`
		NetAmount                string `json:"netAmount"`
		Currency                 string `json:"currency"`
		TypeCode                 string `json:"typeCode"`
		TypeName                 string `json:"typeName"`
		PaymentStatusCode        string `json:"paymentStatusCode"`
		PaymentStatusName        string `json:"paymentStatusName"`
		PaymentDate              string `json:"paymentDate"`
		CounterPartyID           string `json:"counterPartyID"`
		CounterPartyName         string `json:"counterPartyName"`
		CancelledInvoice         bool   `json:"cancelledInvoice"`
		CancellationInvoice      bool   `json:"cancellationInvoice"`
		CancelledInvoiceNumber   string `json:"cancelledInvoiceNumber"`
		SupplierInvoiceReference string `json:"supplierInvoiceReference"`
		DocumentTypeCode         string `json:"documentTypeCode"`
		DocumentTypeName         string `json:"documentTypeName"`
		MaterialNumber           string `json:"materialNumber"`
		MaterialDescription      string `json:"materialDescription"`
		Quantity                 string `json:"quantity"`
		UomCode                  string `json:"uomCode"`
		UomName                  string `json:"uomName"`
		GrossAmount              string `json:"grossAmount"`
		TaxCode                  string `json:"taxCode"`
		TaxName                  string `json:"taxName"`
		TaxAmount                string `json:"taxAmount"`
		BillToPartyID            string `json:"billToPartyID"`
		BillToPartyName          string `json:"billToPartyName"`
		PostingDate              string `json:"postingDate"`
		AdditionalText           string `json:"additionalText"`
		SourceSystemID           string `json:"sourceSystemID"`
		CreatedAt                string `json:"createdAt"`
		ChangedAt                string `json:"changedAt"`
		VehicleProcurement       struct {
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
		} `json:"VehicleProcurement"`
	} `json:"VehicleInvoice"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	PromotionCode string   `json:"vehicle_association_code"`
	Deleted       bool     `json:"deleted"`
}
