package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Input_Reader"
	"strconv"
)

// Header
// データ連携基盤 JP PINT Peppol Electronic Invoice Headerとデータ連携基盤 Invoice Documentとの項目マッピング
func (psdc *SDC) ConvertToMappingHeader(sdc *dpfm_api_input_reader.SDC) (*MappingHeader, error) {
	var res MappingHeader
	data := sdc.InvoiceDocumentHeader
	dataItem := sdc.InvoiceDocumentHeader.InvoiceDocumentItem

	for _, dataItem := range dataItem {
		totalGrossAmount := strconv.FormatFloat(float64(*data.TotalGrossAmount), 'f', -1, 32)
		grossAmount := strconv.FormatFloat(float64(*dataItem.GrossAmount), 'f', -1, 32)
		netAmount := strconv.FormatFloat(float64(*dataItem.NetAmount), 'f', -1, 32)
		totalNetAmount := strconv.FormatFloat(float64(*data.TotalNetAmount), 'f', -1, 32)
		totalTaxAmount := strconv.FormatFloat(float64(*data.TotalTaxAmount), 'f', -1, 32)

		res = append(res, MappingHeader{
			InvoiceIssueDate:             data.InvoiceDocumentDate,
			InvoiceIssueTime:             data.InvoiceDocumentTime,
			TAXPointDate:                 data.AccountingPostingDate,
			PaymentDueDate:               data.PaymentDueDate,
			BuyerReference:               data.DueCalculationBaseDate,
			PaymentTerms:                 data.PaymentTerms,
			TermsAmount:                  &totalGrossAmount,
			InvoiceNote:                  data.DocumentHeaderText,
			PrecedingInvoiceReference:    dataItem.ExternalReferenceDocument,
			ActualDeliveryDate:           dataItem.ActualGoodsIssueDate,
			InvoicingPeriodStartDate:     data.InvoicePeriodStartDate,
			InvoicingPeriodEndDate:       data.InvoicePeriodEndDate,
			PaymentMeansTypeCode:         data.PaymentMethod,
			PaidAmount:                   &grossAmount,
			SumOfInvoiceLineNetAmount:    &netAmount,
			InvoiceTotalAmountWithoutTAX: &totalNetAmount,
			InvoiceTotalTAXAmount:        &totalTaxAmount,
			InvoiceTotalAmountWithTAX:    &totalGrossAmount,
			//ExternalDocumentLocation:     dataPartner.ExternalDocumentID,                            Invoice Document Partner を使う
		})
	}

	return &res, nil
}

// Item
// データ連携基盤 JP PINT Peppol Electronic Invoice Itemとデータ連携基盤 Invoice Documentとの項目マッピング
func (psdc *SDC) ConvertToMappingItem(sdc *dpfm_api_input_reader.SDC) (*[]MappingItem, error) {
	var res []MappingItem
	dataItem := sdc.InvoiceDocumentHeader.InvoiceDocumentItem

	for _, dataItem := range dataItem {
		invoiceQuantity, err := ParseStringPtr(data.PriceDetnExchangeRate)
		if err != nil {
			return nil, err
		}

		res = append(res, MappingItem{
			InvoiceLineIdentifier:                dataItem.InvoiceDocumentItem,
			InvoiceLineNote:                      dataItem.InvoiceDocumentItemText,
			ItemStandardIdentifier:               dataItem.ProductStandardID,
			InvoicedQuantity:                     dataItem.InvoiceQuantity,
			InvoicedQuantityUnitOfMeasureCode:    dataItem.InvoiceQuantityUnit,
			InvoiceLineNetAmount:                 dataItem.NetAmount,
			ReferencedPurchaseOrderLineReference: dataItem.OrderItem,
			OrderReference:                       dataItem.OrderID,
			InvoiceLinePeriodStartDate:           dataItem.InvoicePeriodStartDate,
			InvoiceLinePeriodEndDate:             dataItem.InvoicePeriodEndDate,
			InvoiceLineChargeAmount:              dataItem.GrossAmount,
			ItemNetPrice:                         dataItem.NetAmount,
			// ItemPriceBaseQuantityUnitOfMeasureCode: dataItem.ItemPriceBaseQuantityUnitOfMeasureCode,
			// ActualDeliveryDate:                     dataItem.ActualDeliveryDate,
			// TAXCategoryCode:                        dataItem.TAXCategoryCode,
			// ItemCountryOfOrigin:                    dataItem.ItemCountryOfOrigin,
		})
	}

	return &res, nil
}
