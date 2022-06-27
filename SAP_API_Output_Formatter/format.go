package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-vms-for-vehicle-location-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToVehicleLocation(raw []byte, l *logger.Logger) ([]VehicleLocation, error) {
	pm := &responses.VehicleLocation{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to VehicleLocation. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	vehicleLocation := make([]VehicleLocation, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		vehicleLocation = append(vehicleLocation, VehicleLocation{
			VMSVehicleLocation:      data.VMSVehicleLocation,
			VMSVehicleLocation_Text: data.VMSVehicleLocation_Text,
			AddressNumber:           data.AddressNumber,
			BusinessPartnerName1:    data.BusinessPartnerName1,
			CityName:                data.CityName,
			PostalCode:              data.PostalCode,
			HouseNumber:             data.HouseNumber,
			StreetName:              data.StreetName,
			Country:                 data.Country,
			Region:                  data.Region,
			LogicalSystem:           data.LogicalSystem,
		})
	}

	return vehicleLocation, nil
}
