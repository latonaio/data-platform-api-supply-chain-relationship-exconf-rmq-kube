package dpfm_api_output_formatter

type MetaData struct {
	ConnectionKey                                  string                                          `json:"connection_key"`
	Result                                         bool                                            `json:"result"`
	RedisKey                                       string                                          `json:"redis_key"`
	Filepath                                       string                                          `json:"filepath"`
	APIStatusCode                                  int                                             `json:"api_status_code"`
	RuntimeSessionID                               string                                          `json:"runtime_session_id"`
	BusinessPartnerID                              *int                                            `json:"business_partner"`
	ServiceLabel                                   string                                          `json:"service_label"`
	SupplyChainRelationshipGeneral                 *SupplyChainRelationshipGeneral                 `json:"SupplyChainRelationshipGeneral,omitempty"`
	SupplyChainRelationshipBillingRelation         *SupplyChainRelationshipBillingRelation         `json:"SupplyChainRelationshipBillingRelation,omitempty"`
	SupplyChainRelationshipPaymentRelation         *SupplyChainRelationshipPaymentRelation         `json:"SupplyChainRelationshipPaymentRelation,omitempty"`
	SupplyChainRelationshipDeliveryRelation        *SupplyChainRelationshipDeliveryRelation        `json:"SupplyChainRelationshipDeliveryRelation,omitempty"`
	SupplyChainRelationshipDeliveryPlantRelation   *SupplyChainRelationshipDeliveryPlantRelation   `json:"SupplyChainRelationshipDeliveryPlantRelation,omitempty"`
	SupplyChainRelationshipProductionPlantRelation *SupplyChainRelationshipProductionPlantRelation `json:"SupplyChainRelationshipProductionPlantRelation,omitempty"`
	SupplyChainRelationshipTransaction             *SupplyChainRelationshipTransaction             `json:"SupplyChainRelationshipTransaction,omitempty"`
	APISchema                                      string                                          `json:"api_schema"`
	Accepter                                       []string                                        `json:"accepter"`
	Deleted                                        bool                                            `json:"deleted"`
}

type SupplyChainRelationshipGeneral struct {
	SupplyChainRelationshipID int  `json:"SupplyChainRelationshipID"`
	Buyer                     int  `json:"Buyer"`
	Seller                    int  `json:"Seller"`
	ExistenceConf             bool `json:"ExistenceConf"`
}

type SupplyChainRelationshipDeliveryRelation struct {
	SupplyChainRelationshipID         int  `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID int  `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             int  `json:"Buyer"`
	Seller                            int  `json:"Seller"`
	DeliverToParty                    int  `json:"DeliverToParty"`
	DeliverFromParty                  int  `json:"DeliverFromParty"`
	ExistenceConf                     bool `json:"ExistenceConf"`
}

type SupplyChainRelationshipBillingRelation struct {
	SupplyChainRelationshipID        int  `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int  `json:"SupplyChainRelationshipBillingID"`
	Buyer                            int  `json:"Buyer"`
	Seller                           int  `json:"Seller"`
	BillToParty                      int  `json:"BillToParty"`
	BillFromParty                    int  `json:"BillFromParty"`
	ExistenceConf                    bool `json:"ExistenceConf"`
}

type SupplyChainRelationshipPaymentRelation struct {
	SupplyChainRelationshipID        int  `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int  `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID int  `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int  `json:"Buyer"`
	Seller                           int  `json:"Seller"`
	BillToParty                      int  `json:"BillToParty"`
	BillFromParty                    int  `json:"BillFromParty"`
	Payer                            int  `json:"Payer"`
	Payee                            int  `json:"Payee"`
	ExistenceConf                    bool `json:"ExistenceConf"`
}

type SupplyChainRelationshipDeliveryPlantRelation struct {
	SupplyChainRelationshipID              int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int    `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int    `json:"Buyer"`
	Seller                                 int    `json:"Seller"`
	DeliverToParty                         int    `json:"DeliverToParty"`
	DeliverFromParty                       int    `json:"DeliverFromParty"`
	DeliverToPlant                         string `json:"DeliverToPlant"`
	DeliverFromPlant                       string `json:"DeliverFromPlant"`
	ExistenceConf                          bool   `json:"ExistenceConf"`
}

type SupplyChainRelationshipProductionPlantRelation struct {
	SupplyChainRelationshipID                int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID int    `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int    `json:"Buyer"`
	Seller                                   int    `json:"Seller"`
	ProductionPlantBusinessPartner           int    `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string `json:"ProductionPlant"`
	ExistenceConf                            bool   `json:"ExistenceConf"`
}

type SupplyChainRelationshipTransaction struct {
	SupplyChainRelationshipID int  `json:"SupplyChainRelationshipID"`
	Buyer                     int  `json:"Buyer"`
	Seller                    int  `json:"Seller"`
	ExistenceConf             bool `json:"ExistenceConf"`
}
