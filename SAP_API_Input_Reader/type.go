package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		Dosument       string `json:"document_no"`
		Plant          string `json:"deliver_to"`
		OrderQuantity  string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		NetPriceAmount string `json:"price"`
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
	APISchema                string   `json:"api_schema"`
	Accepter                 []string `json:"accepter"`
	MaterialCode             string   `json:"material_code"`
	Supplier                 string   `json:"plant/supplier"`
	Stock                    string   `json:"stock"`
	SalesOrderType           string   `json:"document_type"`
	MaintenanceItemNo        string   `json:"document_no"`
	ScheduleLineDeliveryDate string   `json:"planned_date"`
	ValidatedDate            string   `json:"validated_date"`
	Deleted                  bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	VehicleLocation struct {
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
	} `json:"VehicleLocation"`
	APISchema                 string   `json:"api_schema"`
	Accepter                  []string `json:"accepter"`
	VMSForVehicleLocationCode string   `json:"vms_for_vehicle_location_code"`
	Deleted                   bool     `json:"deleted"`
}
