package responses

type VehicleInvoice struct {
	Value []struct {
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
	} `json:"value"`
}
