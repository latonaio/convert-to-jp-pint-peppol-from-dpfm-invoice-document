package dpfm_api_input_reader

import (
	"convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.InvoiceDocumentHeader {
	data := sdc.InvoiceDocumentHeader
	return &requests.InvoiceDocumentHeader{
		InvoiceDocument:                   data.InvoiceDocument,
		CreationDate:                      data.CreationDate,
		LastChangeDate:                    data.LastChangeDate,
		SupplyChainRelationshipID:         data.SupplyChainRelationshipID,
		SupplyChainRelationshipBillingID:  data.SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID:  data.SupplyChainRelationshipPaymentID,
		BillToParty:                       data.BillToParty,
		BillFromParty:                     data.BillFromParty,
		BillToCountry:                     data.BillToCountry,
		BillFromCountry:                   data.BillFromCountry,
		Payer:                             data.Payer,
		Payee:                             data.Payee,
		InvoiceDocumentDate:               data.InvoiceDocumentDate,
		InvoiceDocumentTime:               data.InvoiceDocumentTime,
		InvoicePeriodStartDate:            data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:              data.InvoicePeriodEndDate,
		AccountingPostingDate:             data.AccountingPostingDate,
		IsExportImport:                    data.IsExportImport,
		HeaderBillingIsConfirmed:          data.HeaderBillingIsConfirmed,
		HeaderBillingConfStatus:           data.HeaderBillingConfStatus,
		TotalNetAmount:                    data.TotalNetAmount,
		TotalTaxAmount:                    data.TotalTaxAmount,
		TotalGrossAmount:                  data.TotalGrossAmount,
		TransactionCurrency:               data.TransactionCurrency,
		Incoterms:                         data.Incoterms,
		PaymentTerms:                      data.PaymentTerms,
		DueCalculationBaseDate:            data.DueCalculationBaseDate,
		PaymentDueDate:                    data.PaymentDueDate,
		NetPaymentDays:                    data.NetPaymentDays,
		PaymentMethod:                     data.PaymentMethod,
		ExternalReferenceDocument:         data.ExternalReferenceDocument,
		DocumentHeaderText:                data.DocumentHeaderText,
		HeaderPaymentBlockStatus:          data.HeaderPaymentBlockStatus,
		HeaderPaymentRequisitionIsCreated: data.HeaderPaymentRequisitionIsCreated,
		InvoiceDocumentIsCancelled:        data.InvoiceDocumentIsCancelled,
		CancelledInvoiceDocument:          data.CancelledInvoiceDocument,
	}
}

func (sdc *SDC) ConvertToItem(num int) *requests.InvoiceDocumentItem {
	dataHeader := sdc.InvoiceDocumentHeader
	data := sdc.InvoiceDocumentHeader.InvoiceDocumentItem[num]

	return &requests.InvoiceDocumentItem{
		InvoiceDocument:                         dataHeader.InvoiceDocument,
		InvoiceDocumentItem:                     data.InvoiceDocumentItem,
		InvoiceDocumentItemCategory:             data.InvoiceDocumentItemCategory,
		SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
		SupplyChainRelationshipDeliveryID:       data.SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID:  data.SupplyChainRelationshipDeliveryPlantID,
		InvoiceDocumentItemText:                 data.InvoiceDocumentItemText,
		InvoiceDocumentItemTextByBuyer:          data.InvoiceDocumentItemTextByBuyer,
		InvoiceDocumentItemTextBySeller:         data.InvoiceDocumentItemTextBySeller,
		Product:                                 data.Product,
		ProductGroup:                            data.ProductGroup,
		ProductStandardID:                       data.ProductStandardID,
		CreationDate:                            data.CreationDate,
		CreationTime:                            data.CreationTime,
		ItemBillingIsConfirmed:                  data.ItemBillingIsConfirmed,
		ItemBillingConfStatus:                   data.ItemBillingConfStatus,
		Buyer:                                   data.Buyer,
		Seller:                                  data.Seller,
		DeliverToParty:                          data.DeliverToParty,
		DeliverFromParty:                        data.DeliverFromParty,
		DeliverToPlant:                          data.DeliverToPlant,
		DeliverToPlantStorageLocation:           data.DeliverToPlantStorageLocation,
		DeliverFromPlant:                        data.DeliverFromPlant,
		DeliverFromPlantStorageLocation:         data.DeliverFromPlantStorageLocation,
		ProductionPlantBusinessPartner:          data.ProductionPlantBusinessPartner,
		ProductionPlant:                         data.ProductionPlant,
		ProductionPlantStorageLocation:          data.ProductionPlantStorageLocation,
		ServicesRenderedDate:                    data.ServicesRenderedDate,
		InvoiceQuantity:                         data.InvoiceQuantity,
		InvoiceQuantityUnit:                     data.InvoiceQuantityUnit,
		InvoiceQuantityInBaseUnit:               data.InvoiceQuantityInBaseUnit,
		BaseUnit:                                data.BaseUnit,
		ActualGoodsIssueDate:                    data.ActualGoodsIssueDate,
		ActualGoodsIssueTime:                    data.ActualGoodsIssueTime,
		ActualGoodsReceiptDate:                  data.ActualGoodsReceiptDate,
		ActualGoodsReceiptTime:                  data.ActualGoodsReceiptTime,
		ItemGrossWeight:                         data.ItemGrossWeight,
		ItemNetWeight:                           data.ItemNetWeight,
		ItemWeightUnit:                          data.ItemWeightUnit,
		NetAmount:                               data.NetAmount,
		TaxAmount:                               data.TaxAmount,
		GrossAmount:                             data.GrossAmount,
		GoodsIssueOrReceiptSlipNumber:           data.GoodsIssueOrReceiptSlipNumber,
		TransactionCurrency:                     data.TransactionCurrency,
		PricingDate:                             data.PricingDate,
		TransactionTaxClassification:            data.TransactionTaxClassification,
		ProductTaxClassificationBillToCountry:   data.ProductTaxClassificationBillToCountry,
		ProductTaxClassificationBillFromCountry: data.ProductTaxClassificationBillFromCountry,
		DefinedTaxClassification:                data.DefinedTaxClassification,
		Project:                                 data.Project,
		OrderID:                                 data.OrderID,
		OrderItem:                               data.OrderItem,
		OrderType:                               data.OrderType,
		ContractType:                            data.ContractType,
		OrderVaridityStartDate:                  data.OrderVaridityStartDate,
		OrderVaridityEndDate:                    data.OrderVaridityEndDate,
		InvoicePeriodStartDate:                  data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:                    data.InvoicePeriodEndDate,
		DeliveryDocument:                        data.DeliveryDocument,
		DeliveryDocumentItem:                    data.DeliveryDocumentItem,
		OriginDocument:                          data.OriginDocument,
		OriginDocumentItem:                      data.OriginDocumentItem,
		ReferenceDocument:                       data.ReferenceDocument,
		ReferenceDocumentItem:                   data.ReferenceDocumentItem,
		ReferenceDocumentType:                   data.ReferenceDocumentType,
		ExternalReferenceDocument:               data.ExternalReferenceDocument,
		ExternalReferenceDocumentItem:           data.ExternalReferenceDocumentItem,
		TaxCode:                                 data.TaxCode,
		TaxRate:                                 data.TaxRate,
		CountryOfOrigin:                         data.CountryOfOrigin,
		CountryOfOriginLanguage:                 data.CountryOfOriginLanguage,
		ItemPaymentRequisitionIsCreated:         data.ItemPaymentRequisitionIsCreated,
		ItemPaymentBlockStatus:                  data.ItemPaymentBlockStatus,
	}
}

func (sdc *SDC) ConvertToPartner(num int) *requests.InvoiceDocumentPartner {
	dataHeader := sdc.InvoiceDocumentHeader
	data := sdc.InvoiceDocumentHeader.InvoiceDocumentPartner[num]

	return &requests.InvoiceDocumentPartner{
		InvoiceDocument:         dataHeader.InvoiceDocument,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
	}
}

func (sdc *SDC) ConvertToItemPricingElement(hpNum, hppNum int) *requests.InvoiceDocumentItemPricingElement {
	dataHeader := sdc.InvoiceDocumentHeader
	dataItem := sdc.InvoiceDocumentHeader.InvoiceDocumentItem[hpNum]
	data := dataItem.InvoiceDocumentItemPricingElement[hppNum]

	return &requests.InvoiceDocumentItemPricingElement{
		InvoiceDocument:            dataHeader.InvoiceDocument,
		InvoiceDocumentItem:        dataItem.InvoiceDocumentItem,
		PricingProcedureCounter:    data.PricingProcedureCounter,
		ConditionRecord:            data.ConditionRecord,
		ConditionSequentialNumber:  data.ConditionSequentialNumber,
		ConditionType:              data.ConditionType,
		PricingDate:                data.PricingDate,
		ConditionRateValue:         data.ConditionRateValue,
		ConditionCurrency:          data.ConditionCurrency,
		ConditionQuantity:          data.ConditionQuantity,
		ConditionQuantityUnit:      data.ConditionQuantityUnit,
		TaxCode:                    data.TaxCode,
		ConditionAmount:            data.ConditionAmount,
		TransactionCurrency:        data.TransactionCurrency,
		ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
	}
}

func (sdc *SDC) ConvertToAddress(num int) *requests.InvoiceDocumentAddress {
	dataHeader := sdc.InvoiceDocumentHeader
	data := sdc.InvoiceDocumentHeader.InvoiceDocumentAddress[num]

	return &requests.InvoiceDocumentAddress{
		InvoiceDocument: dataHeader.InvoiceDocument,
		AddressID:       data.AddressID,
		PostalCode:      data.PostalCode,
		LocalRegion:     data.LocalRegion,
		Country:         data.Country,
		District:        data.District,
		StreetName:      data.StreetName,
		CityName:        data.CityName,
		Building:        data.Building,
		Floor:           data.Floor,
		Room:            data.Room,
	}
}
