package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-maintain-vehicle-association-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetMaintainVehicleAssociation(vehicleID, supplierID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "VehicleInvoice":
			func() {
				c.VehicleInvoice(vehicleID)
				wg.Done()
			}()
		case "VehicleProcurement":
			func() {
				c.VehicleProcurement(supplierID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) VehicleInvoice(vehicleID string) {
	vehicleInvoiceData, err := c.callMaintainVehicleAssociationSrvAPIRequirementVehicleInvoice("invoice", vehicleID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(vehicleInvoiceData)
}

func (c *SAPAPICaller) callMaintainVehicleAssociationSrvAPIRequirementVehicleInvoice(api, vehicleID string) ([]sap_api_output_formatter.VehicleInvoice, error) {
	url := strings.Join([]string{c.baseURL, "vehicleInvoice", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithVehicleInvoice(req, vehicleID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToVehicleInvoice(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) VehicleProcurement(supplierID string) {
	vehicleProcurementData, err := c.callMaintainVehicleAssociationSrvAPIRequirementVehicleProcurement("procurement", supplierID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(vehicleProcurementData)
}

func (c *SAPAPICaller) callMaintainVehicleAssociationSrvAPIRequirementVehicleProcurement(api, supplierID string) ([]sap_api_output_formatter.VehicleProcurement, error) {
	url := strings.Join([]string{c.baseURL, "vehicleProcurement", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithVehicleProcurement(req, supplierID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToVehicleProcurement(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithVehicleInvoice(req *http.Request, vehicleID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("VehicleID eq '%s'", vehicleID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithVehicleProcurement(req *http.Request, supplierID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("SupplierID eq '%s'", supplierID))
	req.URL.RawQuery = params.Encode()
}
