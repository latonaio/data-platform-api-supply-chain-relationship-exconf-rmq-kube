package dpfm_api_input_reader

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
	BusinessPartner struct {
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
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type GeneralSDC struct {
	ConnectionKey                  string `json:"connection_key"`
	Result                         bool   `json:"result"`
	RedisKey                       string `json:"redis_key"`
	Filepath                       string `json:"filepath"`
	APIStatusCode                  int    `json:"api_status_code"`
	RuntimeSessionID               string `json:"runtime_session_id"`
	BusinessPartner                *int   `json:"business_partner"`
	ServiceLabel                   string `json:"service_label"`
	SupplyChainRelationshipGeneral struct {
		SupplyChainRelationshipID *int `json:"SupplyChainRelationshipID"`
		Buyer                     *int `json:"Buyer"`
		Seller                    *int `json:"Seller"`
	} `json:"SupplyChainRelationshipGeneral"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type DeliveryRelationSDC struct {
	ConnectionKey                           string `json:"connection_key"`
	Result                                  bool   `json:"result"`
	RedisKey                                string `json:"redis_key"`
	Filepath                                string `json:"filepath"`
	APIStatusCode                           int    `json:"api_status_code"`
	RuntimeSessionID                        string `json:"runtime_session_id"`
	BusinessPartner                         *int   `json:"business_partner"`
	ServiceLabel                            string `json:"service_label"`
	SupplyChainRelationshipDeliveryRelation struct {
		SupplyChainRelationshipID         *int `json:"SupplyChainRelationshipID"`
		SupplyChainRelationshipDeliveryID *int `json:"SupplyChainRelationshipDeliveryID"`
		Buyer                             *int `json:"Buyer"`
		Seller                            *int `json:"Seller"`
		DeliverToParty                    *int `json:"DeliverToParty"`
		DeliverFromParty                  *int `json:"DeliverFromParty"`
	} `json:"SupplyChainRelationshipDeliveryRelation"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type BillingRelationSDC struct {
	ConnectionKey                          string `json:"connection_key"`
	Result                                 bool   `json:"result"`
	RedisKey                               string `json:"redis_key"`
	Filepath                               string `json:"filepath"`
	APIStatusCode                          int    `json:"api_status_code"`
	RuntimeSessionID                       string `json:"runtime_session_id"`
	BusinessPartner                        *int   `json:"business_partner"`
	ServiceLabel                           string `json:"service_label"`
	SupplyChainRelationshipBillingRelation struct {
		SupplyChainRelationshipID        *int `json:"SupplyChainRelationshipID"`
		SupplyChainRelationshipBillingID *int `json:"SupplyChainRelationshipBillingID"`
		Buyer                            *int `json:"Buyer"`
		Seller                           *int `json:"Seller"`
		BillToParty                      *int `json:"BillToParty"`
		BillFromParty                    *int `json:"BillFromParty"`
	} `json:"SupplyChainRelationshipBillingRelation"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type PaymentRelationSDC struct {
	ConnectionKey                          string `json:"connection_key"`
	Result                                 bool   `json:"result"`
	RedisKey                               string `json:"redis_key"`
	Filepath                               string `json:"filepath"`
	APIStatusCode                          int    `json:"api_status_code"`
	RuntimeSessionID                       string `json:"runtime_session_id"`
	BusinessPartner                        *int   `json:"business_partner"`
	ServiceLabel                           string `json:"service_label"`
	SupplyChainRelationshipPaymentRelation struct {
		SupplyChainRelationshipID        *int `json:"SupplyChainRelationshipID"`
		SupplyChainRelationshipBillingID *int `json:"SupplyChainRelationshipBillingID"`
		SupplyChainRelationshipPaymentID *int `json:"SupplyChainRelationshipPaymentID"`
		Buyer                            *int `json:"Buyer"`
		Seller                           *int `json:"Seller"`
		BillToParty                      *int `json:"BillToParty"`
		BillFromParty                    *int `json:"BillFromParty"`
		Payer                            *int `json:"Payer"`
		Payee                            *int `json:"Payee"`
	} `json:"SupplyChainRelationshipPaymentRelation"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type DeliveryPlantRelationSDC struct {
	ConnectionKey                                string `json:"connection_key"`
	Result                                       bool   `json:"result"`
	RedisKey                                     string `json:"redis_key"`
	Filepath                                     string `json:"filepath"`
	APIStatusCode                                int    `json:"api_status_code"`
	RuntimeSessionID                             string `json:"runtime_session_id"`
	BusinessPartner                              *int   `json:"business_partner"`
	ServiceLabel                                 string `json:"service_label"`
	SupplyChainRelationshipDeliveryPlantRelation struct {
		SupplyChainRelationshipID              *int    `json:"SupplyChainRelationshipID"`
		SupplyChainRelationshipDeliveryID      *int    `json:"SupplyChainRelationshipDeliveryID"`
		SupplyChainRelationshipDeliveryPlantID *int    `json:"SupplyChainRelationshipDeliveryPlantID"`
		Buyer                                  *int    `json:"Buyer"`
		Seller                                 *int    `json:"Seller"`
		DeliverToParty                         *int    `json:"DeliverToParty"`
		DeliverFromParty                       *int    `json:"DeliverFromParty"`
		DeliverToPlant                         *string `json:"DeliverToPlant"`
		DeliverFromPlant                       *string `json:"DeliverFromPlant"`
	} `json:"SupplyChainRelationshipDeliveryPlantRelation"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type ProductionPlantRelationSDC struct {
	ConnectionKey                                  string `json:"connection_key"`
	Result                                         bool   `json:"result"`
	RedisKey                                       string `json:"redis_key"`
	Filepath                                       string `json:"filepath"`
	APIStatusCode                                  int    `json:"api_status_code"`
	RuntimeSessionID                               string `json:"runtime_session_id"`
	BusinessPartner                                *int   `json:"business_partner"`
	ServiceLabel                                   string `json:"service_label"`
	SupplyChainRelationshipProductionPlantRelation struct {
		SupplyChainRelationshipID                *int    `json:"SupplyChainRelationshipID"`
		SupplyChainRelationshipProductionPlantID *int    `json:"SupplyChainRelationshipProductionPlantID"`
		Buyer                                    *int    `json:"Buyer"`
		Seller                                   *int    `json:"Seller"`
		ProductionPlantBusinessPartner           *int    `json:"ProductionPlantBusinessPartner"`
		ProductionPlant                          *string `json:"ProductionPlant"`
	} `json:"SupplyChainRelationshipProductionPlantRelation"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	Deleted   bool     `json:"deleted"`
}

type TransactionRelationSDC struct {
	ConnectionKey                      string `json:"connection_key"`
	Result                             bool   `json:"result"`
	RedisKey                           string `json:"redis_key"`
	Filepath                           string `json:"filepath"`
	APIStatusCode                      int    `json:"api_status_code"`
	RuntimeSessionID                   string `json:"runtime_session_id"`
	BusinessPartner                    *int   `json:"business_partner"`
	ServiceLabel                       string `json:"service_label"`
	SupplyChainRelationshipTransaction struct {
		SupplyChainRelationshipID *int `json:"SupplyChainRelationshipID"`
		Buyer                     *int `json:"Buyer"`
		Seller                    *int `json:"Seller"`
	} `json:"SupplyChainRelationshipTransaction"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
}
