package responses

type VehicleLocation struct {
	D struct {
		Results []struct {
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
		} `json:"results"`
	} `json:"d"`
}
