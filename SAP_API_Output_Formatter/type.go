package sap_api_output_formatter

type VMSForVehicleLocation struct {
	ConnectionKey             string `json:"connection_key"`
	Result                    bool   `json:"result"`
	RedisKey                  string `json:"redis_key"`
	Filepath                  string `json:"filepath"`
	APISchema                 string `json:"api_schema"`
	VMSForVehicleLocationCode string `json:"vms_for_vehicle_location_code"`
	Deleted                   bool   `json:"deleted"`
}

type VehicleLocation struct {
	VMSVehicleLocation      string `json:"VMSVehicleLocation"`
	VMSVehicleLocation_Text string `json:"VMSVehicleLocation_Text"`
	AddressNumber           string `json:"AddressNumber"`
	BusinessPartnerName1    string `json:"BusinessPartnerName1"`
	CityName                string `json:"CityName"`
	PostalCode              string `json:"PostalCode"`
	HouseNumber             string `json:"HouseNumber"`
	StreetName              string `json:"StreetName"`
	Country                 string `json:"Country"`
	Region                  string `json:"Region"`
	LogicalSystem           string `json:"LogicalSystem"`
}
