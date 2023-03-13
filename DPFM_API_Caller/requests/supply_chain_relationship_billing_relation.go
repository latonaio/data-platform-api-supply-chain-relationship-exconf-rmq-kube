package requests

type SupplyChainRelationshipBillingRelation struct {
	SupplyChainRelationshipID        *int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int `json:"SupplyChainRelationshipBillingID"`
	Buyer                            *int `json:"Buyer"`
	Seller                           *int `json:"Seller"`
	BillToParty                      *int `json:"BillToParty"`
	BillFromParty                    *int `json:"BillFromParty"`
}
