package dpfm_api_input_reader

import (
	"data-platform-api-supply-chain-relationship-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *GeneralSDC) ConvertToSupplyChainRelationshipGeneral() *requests.SupplyChainRelationshipGeneral {
	data := sdc.SupplyChainRelationshipGeneral
	return &requests.SupplyChainRelationshipGeneral{
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
	}
}

func (sdc *DeliveryRelationSDC) ConvertToSupplyChainRelationshipDeliveryRelation() *requests.SupplyChainRelationshipDeliveryRelation {
	data := sdc.SupplyChainRelationshipDeliveryRelation
	return &requests.SupplyChainRelationshipDeliveryRelation{
		SupplyChainRelationshipID:         data.SupplyChainRelationshipID,
		SupplyChainRelationshipDeliveryID: data.SupplyChainRelationshipDeliveryID,
		Buyer:                             data.Buyer,
		Seller:                            data.Seller,
		DeliverToParty:                    data.DeliverToParty,
		DeliverFromParty:                  data.DeliverFromParty,
	}
}

func (sdc *BillingRelationSDC) ConvertToSupplyChainRelationshipBillingRelation() *requests.SupplyChainRelationshipBillingRelation {
	data := sdc.SupplyChainRelationshipBillingRelation
	return &requests.SupplyChainRelationshipBillingRelation{
		SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
		SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
		Buyer:                            data.Buyer,
		Seller:                           data.Seller,
		BillToParty:                      data.BillToParty,
		BillFromParty:                    data.BillFromParty,
	}
}

func (sdc *PaymentRelationSDC) ConvertToSupplyChainRelationshipPaymentRelation() *requests.SupplyChainRelationshipPaymentRelation {
	data := sdc.SupplyChainRelationshipPaymentRelation
	return &requests.SupplyChainRelationshipPaymentRelation{
		SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
		SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
		Buyer:                            data.Buyer,
		Seller:                           data.Seller,
		BillToParty:                      data.BillToParty,
		BillFromParty:                    data.BillFromParty,
		Payer:                            data.Payer,
		Payee:                            data.Payee,
	}
}

func (sdc *DeliveryPlantRelationSDC) ConvertToSupplyChainRelationshipDeliveryPlantRelation() *requests.SupplyChainRelationshipDeliveryPlantRelation {
	data := sdc.SupplyChainRelationshipDeliveryPlantRelation
	return &requests.SupplyChainRelationshipDeliveryPlantRelation{
		SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
		SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
		Buyer:                                  data.Buyer,
		Seller:                                 data.Seller,
		DeliverToParty:                         data.DeliverToParty,
		DeliverFromParty:                       data.DeliverFromParty,
		DeliverToPlant:                         data.DeliverToPlant,
		DeliverFromPlant:                       data.DeliverFromPlant,
	}
}

func (sdc *ProductionPlantRelationSDC) ConvertToSupplyChainRelationshipProductionPlantRelation() *requests.SupplyChainRelationshipProductionPlantRelation {
	data := sdc.SupplyChainRelationshipProductionPlantRelation
	return &requests.SupplyChainRelationshipProductionPlantRelation{
		SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
		SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
		Buyer:                                    data.Buyer,
		Seller:                                   data.Seller,
		ProductionPlantBusinessPartner:           data.ProductionPlantBusinessPartner,
		ProductionPlant:                          data.ProductionPlant,
	}
}
