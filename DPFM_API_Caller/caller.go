package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-supply-chain-relationship-exconf-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-supply-chain-relationship-exconf-rmq-kube/DPFM_API_Output_Formatter"
	"encoding/json"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type ExistenceConf struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewExistenceConf(ctx context.Context, db *database.Mysql, l *logger.Logger) *ExistenceConf {
	return &ExistenceConf{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (e *ExistenceConf) Conf(
	msg rabbitmq.RabbitmqMessage,
	accepter []string,
) interface{} {
	var ret interface{}
	ret = map[string]interface{}{
		"ExistenceConf": false,
	}
	input := make(map[string]interface{})
	err := json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		return ret
	}

	ok := false

	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				_, ok = input["SupplyChainRelationshipGeneral"]
				if ok {
					input := &dpfm_api_input_reader.GeneralSDC{}
					err = json.Unmarshal(msg.Raw(), input)
					ret = e.confSupplyChainRelationshipGeneral(input)
				}
			}()
		case "DeliveryRelation":
			func() {
				_, ok = input["SupplyChainRelationshipDeliveryRelation"]
				if ok {
					input := &dpfm_api_input_reader.DeliveryRelationSDC{}
					err = json.Unmarshal(msg.Raw(), input)
					ret = e.confSupplyChainRelationshipDeliveryRelation(input)
				}
			}()
		case "DeliveryRelationBySRCID":
			func() {
				_, ok = input["SupplyChainRelationshipDeliveryRelation"]
				if ok {
					input := &dpfm_api_input_reader.DeliveryRelationSDCBySRCID{}
					err = json.Unmarshal(msg.Raw(), input)
					ret = e.confSupplyChainRelationshipDeliveryRelationBySRCID(input)
				}
			}()
		case "DeliveryPlant":
			func() {
				_, ok = input["SupplyChainRelationshipDeliveryPlantRelation"]
				if ok {
					input := &dpfm_api_input_reader.DeliveryPlantRelationSDC{}
					err = json.Unmarshal(msg.Raw(), input)
					ret = e.confSupplyChainRelationshipDeliveryPlantRelation(input)
				}
			}()
		case "DeliveryPlantBySRCID":
			func() {
				_, ok = input["SupplyChainRelationshipDeliveryPlantRelation"]
				if ok {
					input := &dpfm_api_input_reader.DeliveryPlantRelationSDC{}
					err = json.Unmarshal(msg.Raw(), input)
					ret = e.confSupplyChainRelationshipDeliveryPlantRelationByySRCID(input)
				}
			}()
		case "Billing":
			func() {
				_, ok = input["SupplyChainRelationshipBillingRelation"]
				if ok {
					input := &dpfm_api_input_reader.BillingRelationSDC{}
					err = json.Unmarshal(msg.Raw(), input)
					ret = e.confSupplyChainRelationshipBillingRelation(input)
				}
			}()
		case "BillingBySRCID":
			func() {
				_, ok = input["SupplyChainRelationshipBillingRelation"]
				if ok {
					input := &dpfm_api_input_reader.BillingRelationSDC{}
					err = json.Unmarshal(msg.Raw(), input)
					ret = e.confSupplyChainRelationshipBillingRelation(input)
				}
			}()
		case "Payment":
			func() {
				_, ok = input["SupplyChainRelationshipPaymentRelation"]
				if ok {
					input := &dpfm_api_input_reader.PaymentRelationSDC{}
					err = json.Unmarshal(msg.Raw(), input)
					ret = e.confSupplyChainRelationshipPaymentRelation(input)
				}
			}()
		case "PaymentBySRCID":
			func() {
				_, ok = input["SupplyChainRelationshipPaymentRelation"]
				if ok {
					input := &dpfm_api_input_reader.PaymentRelationSDC{}
					err = json.Unmarshal(msg.Raw(), input)
					ret = e.confSupplyChainRelationshipPaymentRelationBySRCID(input)
				}
			}()
		case "Transaction":
			func() {
				_, ok = input["SupplyChainRelationshipTransaction"]
				if ok {
					input := &dpfm_api_input_reader.TransactionRelationSDC{}
					err = json.Unmarshal(msg.Raw(), input)
					ret = e.confSupplyChainRelationshipTransaction(input)
				}
			}()
		case "TransactionBySRCID":
			func() {
				_, ok = input["SupplyChainRelationshipTransaction"]
				if ok {
					input := &dpfm_api_input_reader.TransactionRelationSDC{}
					err = json.Unmarshal(msg.Raw(), input)
					ret = e.confSupplyChainRelationshipTransactionBySRCID(input)
				}
			}()

		case "ProductionPlant":
			func() {
				_, ok = input["SupplyChainRelationshipProductionPlantRelation"]
				if ok {
					input := &dpfm_api_input_reader.ProductionPlantRelationSDC{}
					err = json.Unmarshal(msg.Raw(), input)
					ret = e.confSupplyChainRelationshipProductionPlantRelation(input)
				}
			}()
		default:
		}
	}

	//err = xerrors.Errorf("can not get exconf check target")

	goto endProcess

endProcess:
	if err != nil {
		e.l.Error(err)
	}
	return ret
}

func (e *ExistenceConf) confSupplyChainRelationshipGeneral(input *dpfm_api_input_reader.GeneralSDC) *dpfm_api_output_formatter.SupplyChainRelationshipGeneral {
	exconf := dpfm_api_output_formatter.SupplyChainRelationshipGeneral{
		ExistenceConf: false,
	}
	if input.SupplyChainRelationshipGeneral.SupplyChainRelationshipID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipGeneral.Buyer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipGeneral.Seller == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.SupplyChainRelationshipGeneral{
		SupplyChainRelationshipID: *input.SupplyChainRelationshipGeneral.SupplyChainRelationshipID,
		Buyer:                     *input.SupplyChainRelationshipGeneral.Buyer,
		Seller:                    *input.SupplyChainRelationshipGeneral.Seller,
		ExistenceConf:             false,
	}

	rows, err := e.db.Query(
		`SELECT SupplyChainRelationshipID, Buyer, Seller
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_general_data 
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?, ?, ?);`, exconf.SupplyChainRelationshipID, exconf.Buyer, exconf.Seller,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}
func (e *ExistenceConf) confSupplyChainRelationshipDeliveryRelation(input *dpfm_api_input_reader.DeliveryRelationSDC) *dpfm_api_output_formatter.SupplyChainRelationshipDeliveryRelation {
	exconf := dpfm_api_output_formatter.SupplyChainRelationshipDeliveryRelation{
		ExistenceConf: false,
	}
	if input.SupplyChainRelationshipDeliveryRelation.SupplyChainRelationshipID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryRelation.SupplyChainRelationshipDeliveryID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryRelation.Buyer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryRelation.Seller == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryRelation.DeliverToParty == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryRelation.DeliverFromParty == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.SupplyChainRelationshipDeliveryRelation{
		SupplyChainRelationshipID:         *input.SupplyChainRelationshipDeliveryRelation.SupplyChainRelationshipID,
		SupplyChainRelationshipDeliveryID: *input.SupplyChainRelationshipDeliveryRelation.SupplyChainRelationshipDeliveryID,
		Buyer:                             *input.SupplyChainRelationshipDeliveryRelation.Buyer,
		Seller:                            *input.SupplyChainRelationshipDeliveryRelation.Seller,
		DeliverToParty:                    *input.SupplyChainRelationshipDeliveryRelation.DeliverToParty,
		DeliverFromParty:                  *input.SupplyChainRelationshipDeliveryRelation.DeliverFromParty,
		ExistenceConf:                     false,
	}

	rows, err := e.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, Buyer, Seller, DeliverToParty, DeliverFromParty
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_src_delivery_relation_data 
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, Buyer, Seller, DeliverToParty, DeliverFromParty) = (?, ?, ?, ?, ?, ?);`, exconf.SupplyChainRelationshipID, exconf.SupplyChainRelationshipDeliveryID, exconf.Buyer, exconf.Seller, exconf.DeliverToParty, exconf.DeliverFromParty,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confSupplyChainRelationshipDeliveryRelationBySRCID(
	input *dpfm_api_input_reader.DeliveryRelationSDCBySRCID,
) *dpfm_api_output_formatter.SupplyChainRelationshipDeliveryRelation {
	exconf := dpfm_api_output_formatter.SupplyChainRelationshipDeliveryRelation{
		ExistenceConf: false,
	}
	if input.SupplyChainRelationshipDeliveryRelation.SupplyChainRelationshipID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryRelation.Buyer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryRelation.Seller == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.SupplyChainRelationshipDeliveryRelation{
		SupplyChainRelationshipID: *input.SupplyChainRelationshipDeliveryRelation.SupplyChainRelationshipID,
		Buyer:                     *input.SupplyChainRelationshipDeliveryRelation.Buyer,
		Seller:                    *input.SupplyChainRelationshipDeliveryRelation.Seller,
		ExistenceConf:             false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_scr_delivery_relation_data 
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?, ?, ?);`, exconf.SupplyChainRelationshipID, exconf.Buyer, exconf.Seller,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confSupplyChainRelationshipBillingRelation(input *dpfm_api_input_reader.BillingRelationSDC) *dpfm_api_output_formatter.SupplyChainRelationshipBillingRelation {
	exconf := dpfm_api_output_formatter.SupplyChainRelationshipBillingRelation{
		ExistenceConf: false,
	}
	if input.SupplyChainRelationshipBillingRelation.SupplyChainRelationshipID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipBillingRelation.SupplyChainRelationshipBillingID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipBillingRelation.Buyer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipBillingRelation.Seller == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipBillingRelation.BillToParty == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipBillingRelation.BillFromParty == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.SupplyChainRelationshipBillingRelation{
		SupplyChainRelationshipID:        *input.SupplyChainRelationshipBillingRelation.SupplyChainRelationshipID,
		SupplyChainRelationshipBillingID: *input.SupplyChainRelationshipBillingRelation.SupplyChainRelationshipBillingID,
		Buyer:                            *input.SupplyChainRelationshipBillingRelation.Buyer,
		Seller:                           *input.SupplyChainRelationshipBillingRelation.Seller,
		BillToParty:                      *input.SupplyChainRelationshipBillingRelation.BillToParty,
		BillFromParty:                    *input.SupplyChainRelationshipBillingRelation.BillFromParty,
		ExistenceConf:                    false,
	}

	rows, err := e.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipBillingID, Buyer, Seller, BillToParty, BillFromParty
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_src_billing_relation_data 
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipBillingID, Buyer, Seller, BillToParty, BillFromParty) = (?, ?, ?, ?, ?, ?);`, exconf.SupplyChainRelationshipID, exconf.SupplyChainRelationshipBillingID, exconf.Buyer, exconf.Seller, exconf.BillToParty, exconf.BillFromParty,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confSupplyChainRelationshipBillingRelationBySRCID(input *dpfm_api_input_reader.BillingRelationSDC) *dpfm_api_output_formatter.SupplyChainRelationshipBillingRelation {
	exconf := dpfm_api_output_formatter.SupplyChainRelationshipBillingRelation{
		ExistenceConf: false,
	}
	if input.SupplyChainRelationshipBillingRelation.SupplyChainRelationshipID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipBillingRelation.Buyer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipBillingRelation.Seller == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.SupplyChainRelationshipBillingRelation{
		SupplyChainRelationshipID: *input.SupplyChainRelationshipBillingRelation.SupplyChainRelationshipID,
		Buyer:                     *input.SupplyChainRelationshipBillingRelation.Buyer,
		Seller:                    *input.SupplyChainRelationshipBillingRelation.Seller,
		ExistenceConf:             false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_src_billing_relation_data 
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?, ?, ?);`, exconf.SupplyChainRelationshipID, exconf.Buyer, exconf.Seller,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confSupplyChainRelationshipPaymentRelation(input *dpfm_api_input_reader.PaymentRelationSDC) *dpfm_api_output_formatter.SupplyChainRelationshipPaymentRelation {
	exconf := dpfm_api_output_formatter.SupplyChainRelationshipPaymentRelation{
		ExistenceConf: false,
	}
	if input.SupplyChainRelationshipPaymentRelation.SupplyChainRelationshipID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipPaymentRelation.SupplyChainRelationshipBillingID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipPaymentRelation.SupplyChainRelationshipPaymentID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipPaymentRelation.Buyer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipPaymentRelation.Seller == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipPaymentRelation.BillToParty == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipPaymentRelation.BillFromParty == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipPaymentRelation.Payer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipPaymentRelation.Payee == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.SupplyChainRelationshipPaymentRelation{
		SupplyChainRelationshipID:        *input.SupplyChainRelationshipPaymentRelation.SupplyChainRelationshipID,
		SupplyChainRelationshipBillingID: *input.SupplyChainRelationshipPaymentRelation.SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID: *input.SupplyChainRelationshipPaymentRelation.SupplyChainRelationshipPaymentID,
		Buyer:                            *input.SupplyChainRelationshipPaymentRelation.Buyer,
		Seller:                           *input.SupplyChainRelationshipPaymentRelation.Seller,
		BillToParty:                      *input.SupplyChainRelationshipPaymentRelation.BillToParty,
		BillFromParty:                    *input.SupplyChainRelationshipPaymentRelation.BillFromParty,
		Payer:                            *input.SupplyChainRelationshipPaymentRelation.Payer,
		Payee:                            *input.SupplyChainRelationshipPaymentRelation.Payee,
		ExistenceConf:                    false,
	}

	rows, err := e.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipBillingID, SupplyChainRelationshipPaymentID,Buyer, Seller, BillToParty, BillFromParty, Payer, Payee
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_src_payment_relation_data 
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipBillingID, SupplyChainRelationshipPaymentID,Buyer, Seller, BillToParty, BillFromParty, Payer, Payee) = (?, ?, ?, ?, ?, ?, ?, ?, ?);`, exconf.SupplyChainRelationshipID, exconf.SupplyChainRelationshipBillingID, exconf.SupplyChainRelationshipPaymentID, exconf.Buyer, exconf.Seller, exconf.BillToParty, exconf.BillFromParty, exconf.Payer, exconf.Payee,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confSupplyChainRelationshipPaymentRelationBySRCID(input *dpfm_api_input_reader.PaymentRelationSDC) *dpfm_api_output_formatter.SupplyChainRelationshipPaymentRelation {
	exconf := dpfm_api_output_formatter.SupplyChainRelationshipPaymentRelation{
		ExistenceConf: false,
	}
	if input.SupplyChainRelationshipPaymentRelation.SupplyChainRelationshipID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipPaymentRelation.Buyer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipPaymentRelation.Seller == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.SupplyChainRelationshipPaymentRelation{
		SupplyChainRelationshipID: *input.SupplyChainRelationshipPaymentRelation.SupplyChainRelationshipID,
		Buyer:                     *input.SupplyChainRelationshipPaymentRelation.Buyer,
		Seller:                    *input.SupplyChainRelationshipPaymentRelation.Seller,
		ExistenceConf:             false,
	}

	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_src_payment_relation_data 
		WHERE (SupplyChainRelationshipID ,Buyer, Seller) = (?, ?, ?);`, exconf.SupplyChainRelationshipID, exconf.Buyer, exconf.Seller,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}
func (e *ExistenceConf) confSupplyChainRelationshipDeliveryPlantRelation(input *dpfm_api_input_reader.DeliveryPlantRelationSDC) *dpfm_api_output_formatter.SupplyChainRelationshipDeliveryPlantRelation {
	exconf := dpfm_api_output_formatter.SupplyChainRelationshipDeliveryPlantRelation{
		ExistenceConf: false,
	}
	if input.SupplyChainRelationshipDeliveryPlantRelation.SupplyChainRelationshipID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryPlantRelation.SupplyChainRelationshipDeliveryID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryPlantRelation.SupplyChainRelationshipDeliveryPlantID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryPlantRelation.Buyer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryPlantRelation.Seller == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryPlantRelation.DeliverToParty == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryPlantRelation.DeliverFromParty == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryPlantRelation.DeliverToPlant == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryPlantRelation.DeliverFromPlant == nil {
		return &exconf
	}

	exconf = dpfm_api_output_formatter.SupplyChainRelationshipDeliveryPlantRelation{
		SupplyChainRelationshipID:              *input.SupplyChainRelationshipDeliveryPlantRelation.SupplyChainRelationshipID,
		SupplyChainRelationshipDeliveryID:      *input.SupplyChainRelationshipDeliveryPlantRelation.SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID: *input.SupplyChainRelationshipDeliveryPlantRelation.SupplyChainRelationshipDeliveryPlantID,
		Buyer:                                  *input.SupplyChainRelationshipDeliveryPlantRelation.Buyer,
		Seller:                                 *input.SupplyChainRelationshipDeliveryPlantRelation.Seller,
		DeliverToParty:                         *input.SupplyChainRelationshipDeliveryPlantRelation.DeliverToParty,
		DeliverFromParty:                       *input.SupplyChainRelationshipDeliveryPlantRelation.DeliverFromParty,
		DeliverToPlant:                         *input.SupplyChainRelationshipDeliveryPlantRelation.DeliverToPlant,
		DeliverFromPlant:                       *input.SupplyChainRelationshipDeliveryPlantRelation.DeliverFromPlant,
		ExistenceConf:                          false,
	}

	rows, err := e.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, SupplyChainRelationshipDeliveryPlantID, Buyer, Seller, DeliverToParty, DeliverFromParty, DeliverToPlant, DeliverFromPlant
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_src_delivery_plant_rel_data 
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, SupplyChainRelationshipDeliveryPlantID, Buyer, Seller, DeliverToParty, DeliverFromParty, DeliverToPlant, DeliverFromPlant) = (?, ?, ?, ?, ?, ?, ?, ?, ?);`, exconf.SupplyChainRelationshipID, exconf.SupplyChainRelationshipDeliveryID, exconf.SupplyChainRelationshipDeliveryPlantID, exconf.Buyer, exconf.Seller, exconf.DeliverToParty, exconf.DeliverFromParty, exconf.DeliverToPlant, exconf.DeliverFromPlant,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confSupplyChainRelationshipDeliveryPlantRelationByySRCID(input *dpfm_api_input_reader.DeliveryPlantRelationSDC) *dpfm_api_output_formatter.SupplyChainRelationshipDeliveryPlantRelation {
	exconf := dpfm_api_output_formatter.SupplyChainRelationshipDeliveryPlantRelation{
		ExistenceConf: false,
	}
	if input.SupplyChainRelationshipDeliveryPlantRelation.SupplyChainRelationshipID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryPlantRelation.Buyer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipDeliveryPlantRelation.Seller == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.SupplyChainRelationshipDeliveryPlantRelation{
		SupplyChainRelationshipID: *input.SupplyChainRelationshipDeliveryPlantRelation.SupplyChainRelationshipID,
		Buyer:                     *input.SupplyChainRelationshipDeliveryPlantRelation.Buyer,
		Seller:                    *input.SupplyChainRelationshipDeliveryPlantRelation.Seller,
		ExistenceConf:             false,
	}
	rows, err := e.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_src_delivery_plant_rel_data 
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?, ?, ?);`, exconf.SupplyChainRelationshipID, exconf.Buyer, exconf.Seller,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confSupplyChainRelationshipProductionPlantRelation(input *dpfm_api_input_reader.ProductionPlantRelationSDC) *dpfm_api_output_formatter.SupplyChainRelationshipProductionPlantRelation {
	exconf := dpfm_api_output_formatter.SupplyChainRelationshipProductionPlantRelation{
		ExistenceConf: false,
	}
	if input.SupplyChainRelationshipProductionPlantRelation.SupplyChainRelationshipID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipProductionPlantRelation.SupplyChainRelationshipProductionPlantID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipProductionPlantRelation.Buyer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipProductionPlantRelation.Seller == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipProductionPlantRelation.ProductionPlantBusinessPartner == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipProductionPlantRelation.ProductionPlant == nil {
		return &exconf
	}

	exconf = dpfm_api_output_formatter.SupplyChainRelationshipProductionPlantRelation{
		SupplyChainRelationshipID:                *input.SupplyChainRelationshipProductionPlantRelation.SupplyChainRelationshipID,
		SupplyChainRelationshipProductionPlantID: *input.SupplyChainRelationshipProductionPlantRelation.SupplyChainRelationshipProductionPlantID,
		Buyer:                                    *input.SupplyChainRelationshipProductionPlantRelation.Buyer,
		Seller:                                   *input.SupplyChainRelationshipProductionPlantRelation.Seller,
		ProductionPlantBusinessPartner:           *input.SupplyChainRelationshipProductionPlantRelation.ProductionPlantBusinessPartner,
		ProductionPlant:                          *input.SupplyChainRelationshipProductionPlantRelation.ProductionPlant,
		ExistenceConf:                            false,
	}

	rows, err := e.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipProductionPlantID, Buyer, Seller, ProductionPlantBusinessPartner, ProductionPlant
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_src_prod_plant_relation_data 
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipProductionPlantID, Buyer, Seller, ProductionPlantBusinessPartner, ProductionPlant) = (?, ?, ?, ?, ?, ?);`, exconf.SupplyChainRelationshipID, exconf.SupplyChainRelationshipProductionPlantID, exconf.Buyer, exconf.Seller, exconf.ProductionPlantBusinessPartner, exconf.ProductionPlant,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confSupplyChainRelationshipTransactionBySRCID(input *dpfm_api_input_reader.TransactionRelationSDC) *dpfm_api_output_formatter.SupplyChainRelationshipTransaction {
	exconf := dpfm_api_output_formatter.SupplyChainRelationshipTransaction{
		ExistenceConf: false,
	}
	if input.SupplyChainRelationshipTransaction.SupplyChainRelationshipID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipTransaction.Buyer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipTransaction.Seller == nil {
		return &exconf
	}

	exconf = dpfm_api_output_formatter.SupplyChainRelationshipTransaction{
		SupplyChainRelationshipID: *input.SupplyChainRelationshipTransaction.SupplyChainRelationshipID,
		Buyer:                     *input.SupplyChainRelationshipTransaction.Buyer,
		Seller:                    *input.SupplyChainRelationshipTransaction.Seller,
		ExistenceConf:             false,
	}

	rows, err := e.db.Query(
		`SELECT SupplyChainRelationshipID, Buyer, Seller
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_src_transaction_data 
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?, ?, ?);`, exconf.SupplyChainRelationshipID, exconf.Buyer, exconf.Seller,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) confSupplyChainRelationshipTransaction(input *dpfm_api_input_reader.TransactionRelationSDC) *dpfm_api_output_formatter.SupplyChainRelationshipTransaction {
	exconf := dpfm_api_output_formatter.SupplyChainRelationshipTransaction{
		ExistenceConf: false,
	}
	if input.SupplyChainRelationshipTransaction.SupplyChainRelationshipID == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipTransaction.Buyer == nil {
		return &exconf
	}
	if input.SupplyChainRelationshipTransaction.Seller == nil {
		return &exconf
	}

	exconf = dpfm_api_output_formatter.SupplyChainRelationshipTransaction{
		SupplyChainRelationshipID: *input.SupplyChainRelationshipTransaction.SupplyChainRelationshipID,
		Buyer:                     *input.SupplyChainRelationshipTransaction.Buyer,
		Seller:                    *input.SupplyChainRelationshipTransaction.Seller,
		ExistenceConf:             false,
	}

	rows, err := e.db.Query(
		`SELECT SupplyChainRelationshipID, Buyer, Seller
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_src_transaction_data 
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?, ?, ?);`, exconf.SupplyChainRelationshipID, exconf.Buyer, exconf.Seller,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}
	defer rows.Close()

	exconf.ExistenceConf = rows.Next()
	return &exconf
}
