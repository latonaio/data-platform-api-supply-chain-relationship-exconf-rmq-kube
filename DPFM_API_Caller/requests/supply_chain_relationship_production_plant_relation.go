package requests

type SupplyChainRelationshipProductionPlantRelation struct {
	SupplyChainRelationshipID                *int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID *int    `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    *int    `json:"Buyer"`
	Seller                                   *int    `json:"Seller"`
	ProductionPlantBusinessPartner           *int    `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          *string `json:"ProductionPlant"`
}
