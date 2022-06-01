package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-maintain-vehicle-association-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToVehicleInvoice(raw []byte, l *logger.Logger) ([]VehicleInvoice, error) {
	pm := &responses.VehicleInvoice{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to VehicleInvoice. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	vehicleInvoice := make([]VehicleInvoice, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		vehicleInvoice = append(vehicleInvoice, VehicleInvoice{
			VehicleID:                data.VehicleID,
			Number:                   data.Number,
			ItemNumber:               data.ItemNumber,
			ID:                       data.ID,
			CategoryCode:             data.CategoryCode,
			CategoryName:             data.CategoryName,
			Date:                     data.Date,
			NetAmount:                data.NetAmount,
			Currency:                 data.Currency,
			TypeCode:                 data.TypeCode,
			TypeName:                 data.TypeName,
			PaymentStatusCode:        data.PaymentStatusCode,
			PaymentStatusName:        data.PaymentStatusName,
			PaymentDate:              data.PaymentDate,
			CounterPartyID:           data.CounterPartyID,
			CounterPartyName:         data.CounterPartyName,
			CancelledInvoice:         data.CancelledInvoice,
			CancellationInvoice:      data.CancellationInvoice,
			CancelledInvoiceNumber:   data.CancelledInvoiceNumber,
			SupplierInvoiceReference: data.SupplierInvoiceReference,
			DocumentTypeCode:         data.DocumentTypeCode,
			DocumentTypeName:         data.DocumentTypeName,
			MaterialNumber:           data.MaterialNumber,
			MaterialDescription:      data.MaterialDescription,
			Quantity:                 data.Quantity,
			UomCode:                  data.UomCode,
			UomName:                  data.UomName,
			GrossAmount:              data.GrossAmount,
			TaxCode:                  data.TaxCode,
			TaxName:                  data.TaxName,
			TaxAmount:                data.TaxAmount,
			BillToPartyID:            data.BillToPartyID,
			BillToPartyName:          data.BillToPartyName,
			PostingDate:              data.PostingDate,
			AdditionalText:           data.AdditionalText,
			SourceSystemID:           data.SourceSystemID,
			CreatedAt:                data.CreatedAt,
			ChangedAt:                data.ChangedAt,
		})
	}

	return vehicleInvoice, nil
}

func ConvertToVehicleProcurement(raw []byte, l *logger.Logger) ([]VehicleProcurement, error) {
	pm := &responses.VehicleProcurement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to VehicleProcurement. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	vehicleProcurement := make([]VehicleProcurement, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		vehicleProcurement = append(vehicleProcurement, VehicleProcurement{
			VehicleID:                    data.VehicleID,
			DocumentNumber:               data.DocumentNumber,
			ItemNumber:                   data.ItemNumber,
			ID:                           data.ID,
			SupplierID:                   data.SupplierID,
			SupplierName:                 data.SupplierName,
			CostCenterNumber:             data.CostCenterNumber,
			CostCenterDescription:        data.CostCenterDescription,
			MileageLimitFrequencyCode:    data.MileageLimitFrequencyCode,
			MileageLimitFrequencyName:    data.MileageLimitFrequencyName,
			Date:                         data.Date,
			Amount:                       data.Amount,
			Currency:                     data.Currency,
			PlantID:                      data.PlantID,
			PlantName:                    data.PlantName,
			TypeCode:                     data.TypeCode,
			TypeName:                     data.TypeName,
			StatusCode:                   data.StatusCode,
			StatusName:                   data.StatusName,
			AssetNumber:                  data.AssetNumber,
			InternalOrderNumber:          data.InternalOrderNumber,
			SourceSystemID:               data.SourceSystemID,
			OrganizationID:               data.OrganizationID,
			OrganizationName:             data.OrganizationName,
			GroupID:                      data.GroupID,
			GroupName:                    data.GroupName,
			CompanyCode:                  data.CompanyCode,
			CompanyName:                  data.CompanyName,
			AdditionalText:               data.AdditionalText,
			LeasingContractID:            data.LeasingContractID,
			ContractStartDate:            data.ContractStartDate,
			ContractEndDate:              data.ContractEndDate,
			PaymentFrequencyCode:         data.PaymentFrequencyCode,
			PaymentFrequencyName:         data.PaymentFrequencyName,
			ServiceAmount:                data.ServiceAmount,
			ServiceRateFrequencyCode:     data.ServiceRateFrequencyCode,
			ServiceRateFrequencyName:     data.ServiceRateFrequencyName,
			EarlyTerminationChargeAmount: data.EarlyTerminationChargeAmount,
			MileageLimit:                 data.MileageLimit,
			MileageLimitUomCode:          data.MileageLimitUomCode,
			MileageLimitUomName:          data.MileageLimitUomName,
			DocumentTypeCode:             data.DocumentTypeCode,
			DocumentTypeName:             data.DocumentTypeName,
			DeliveryDate:                 data.DeliveryDate,
			TaxCode:                      data.TaxCode,
			TaxName:                      data.TaxName,
			TaxAmount:                    data.TaxAmount,
			CreatedAt:                    data.CreatedAt,
			ChangedAt:                    data.ChangedAt,
		})
	}

	return vehicleProcurement, nil
}
