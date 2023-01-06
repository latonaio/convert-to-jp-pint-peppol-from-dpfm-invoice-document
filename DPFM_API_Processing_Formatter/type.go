package dpfm_api_processing_formatter

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

type SDC struct {
	MappingHeader        *MappingHeader          `json:"MappingHeader"`
	MappingItem          *[]MappingItem          `json:"MappingItem"`
	MappingHeaderPartner *[]MappingHeaderPartner `json:"MappingHeaderPartner"`
	MappingAddress       *MappingAddress         `json:"MappingAddress"`
	MappingItemPartner   *[]MappingItemPartner   `json:"MappingItemPartner"`
}
type Header struct {
	InvoiceNumber                   string          `json:"InvoiceNumber"`
	InvoiceIssueDate                *string         `json:"InvoiceIssueDate"`
	InvoiceIssueTime                *string         `json:"InvoiceIssueTime"`
	InvoiceTypeCode                 *string         `json:"InvoiceTypeCode"`
	InvoiceCurrencyCode             *string         `json:"InvoiceCurrencyCode"`
	TaxAccountingCurrency           *string         `json:"TaxAccountingCurrency"`
	TAXPointDate                    *string         `json:"TAXPointDate"`
	TAXPointDateCode                *string         `json:"TAXPointDateCode"`
	PaymentDueDate                  *string         `json:"PaymentDueDate"`
	BuyerReference                  *string         `json:"BuyerReference"`
	ProjectReference                *string         `json:"ProjectReference"`
	ContractReference               *string         `json:"ContractReference"`
	PurchaseOrderReference          *string         `json:"PurchaseOrderReference"`
	SalesOrderReference             *string         `json:"SalesOrderReference"`
	ReceivingAdviceReference        *string         `json:"ReceivingAdviceReference"`
	DespatchAdviceReference         *string         `json:"DespatchAdviceReference"`
	TenderOrLotReference            *string         `json:"TenderOrLotReference"`
	PaymentTerms                    *string         `json:"PaymentTerms"`
	TermsAmount                     *string         `json:"TermsAmount"`
	TermsInstallmentDueDate         *string         `json:"TermsInstallmentDueDate"`
	InvoiceNoteSubjectCode          *string         `json:"InvoiceNoteSubjectCode"`
	InvoiceNote                     *string         `json:"InvoiceNote"`
	BusinessProcessType             *string         `json:"BusinessProcessType"`
	PrecedingInvoiceReference       *string         `json:"PrecedingInvoiceReference"`
	PrecedingInvoiceIssueDate       *string         `json:"PrecedingInvoiceIssueDate"`
	ActualDeliveryDate              *string         `json:"ActualDeliveryDate"`
	InvoicingPeriodStartDate        *string         `json:"InvoicingPeriodStartDate"`
	InvoicingPeriodEndDate          *string         `json:"InvoicingPeriodEndDate"`
	PaymentInstructionsID           *string         `json:"PaymentInstructionsID"`
	PaymentMeansTypeCode            *string         `json:"PaymentMeansTypeCode"`
	PaymentCardPrimaryAccountNumber *string         `json:"PaymentCardPrimaryAccountNumber"`
	PaymentCardHolderName           *string         `json:"PaymentCardHolderName"`
	PaidAmount                      *string         `json:"PaidAmount"`
	PaymentType                     *string         `json:"PaymentType"`
	SumOfInvoiceLineNetAmount       *string         `json:"SumOfInvoiceLineNetAmount"`
	InvoiceTotalAmountWithoutTAX    *string         `json:"InvoiceTotalAmountWithoutTAX"`
	InvoiceTotalTAXAmount           *string         `json:"InvoiceTotalTAXAmount"`
	InvoiceTotalAmountWithTAX       *string         `json:"InvoiceTotalAmountWithTAX"`
	AmountDueForPayment             *string         `json:"AmountDueForPayment"`
	ExternalDocumentLocation        *string         `json:"ExternalDocumentLocation"`
	AttachedDocument                *string         `json:"AttachedDocument"`
	Item                            []Item          `json:"Item"`
	HeaderPartner                   []HeaderPartner `json:"HeaderPartner"`
	Address                         []Address       `json:"Address"`
}

type Item struct {
	InvoiceNumber                          string        `json:"InvoiceNumber"`
	InvoiceLineIdentifier                  string        `json:"InvoiceLineIdentifier"`
	InvoiceLineNote                        *string       `json:"InvoiceLineNote"`
	ItemStandardIdentifier                 *string       `json:"ItemStandardIdentifier"`
	ItemClassificationIdentifier           *string       `json:"ItemClassificationIdentifier"`
	InvoiceLineDocumentIdentifier          *string       `json:"InvoiceLineDocumentIdentifier"`
	DocumentTypeCode                       *string       `json:"DocumentTypeCode"`
	InvoiceLineObjectIdentifier            *string       `json:"InvoiceLineObjectIdentifier"`
	InvoicedQuantity                       *string       `json:"InvoicedQuantity"`
	InvoicedQuantityUnitOfMeasureCode      *string       `json:"InvoicedQuantityUnitOfMeasureCode"`
	InvoiceLineNetAmountn                  *string       `json:"InvoiceLineNetAmountn"`
	ReferencedPurchaseOrderLineReference   *string       `json:"ReferencedPurchaseOrderLineReference"`
	OrderReference                         *string       `json:"OrderReference"`
	InvoiceLinePeriodStartDate             *string       `json:"InvoiceLinePeriodStartDate"`
	InvoiceLinePeriodEndDate               *string       `json:"InvoiceLinePeriodEndDate"`
	InvoiceLineChargeAmount                *string       `json:"InvoiceLineChargeAmount"`
	ItemNetPrice                           *string       `json:"ItemNetPrice"`
	ItemPriceBaseQuantity                  *string       `json:"ItemPriceBaseQuantity"`
	ItemPriceBaseQuantityUnitOfMeasureCode *string       `json:"ItemPriceBaseQuantityUnitOfMeasureCode"`
	ActualDeliveryDate                     *string       `json:"ActualDeliveryDate"`
	InvoicedItemTAXCategoryCode            *string       `json:"InvoicedItemTAXCategoryCode"`
	ItemName                               *string       `json:"ItemName"`
	ItemDescription                        *string       `json:"ItemDescription"`
	TAXCategoryCode                        *string       `json:"TAXCategoryCode"`
	TAXCategoryRate                        *string       `json:"TAXCategoryRate"`
	ItemCountryOfOrigin                    *string       `json:"ItemCountryOfOrigin"`
	ItemPartner                            []ItemPartner `json:"ItemPartner"`
}

type HeaderPartner struct {
	InvoiceNumber   string  `json:"InvoiceNumber"`
	PartnerFunction string  `json:"PartnerFunction"`
	PartnerID       string  `json:"PartnerID"`
	PartnerName     *string `json:"PartnerName"`
	AddressID       *int    `json:"AddressID"`
}

type ItemPartner struct {
	InvoiceNumber         string  `json:"InvoiceNumber"`
	InvoiceLineIdentifier string  `json:"InvoiceLineIdentifier"`
	PartnerFunction       string  `json:"PartnerFunction"`
	PartnerID             string  `json:"PartnerID"`
	PartnerName           *string `json:"PartnerName"`
	AddressID             *int    `json:"AddressID"`
}

type Address struct {
	InvoiceNumber string  `json:"InvoiceNumber"`
	AddressID     int     `json:"AddressID"`
	PostalCode    *string `json:"PostalCode"`
	LocalRegion   *string `json:"LocalRegion"`
	Country       *string `json:"Country"`
	District      *string `json:"District"`
	StreetName    *string `json:"StreetName"`
	CityName      *string `json:"CityName"`
	Building      *string `json:"Building"`
	Floor         *int    `json:"Floor"`
	Room          *int    `json:"Room"`
}

// 項目マッピング変換
type MappingHeader struct {
	InvoiceIssueDate             *string `json:"InvoiceIssueDate"`
	InvoiceIssueTime             *string `json:"InvoiceIssueTime"`
	TAXPointDate                 *string `json:"TAXPointDate"`
	PaymentDueDate               *string `json:"PaymentDueDate"`
	BuyerReference               *string `json:"BuyerReference"`
	PaymentTerms                 *string `json:"PaymentTerms"`
	TermsAmount                  *string `json:"TermsAmount"`
	InvoiceNote                  *string `json:"InvoiceNote"`
	PrecedingInvoiceReference    *string `json:"PrecedingInvoiceReference"`
	ActualDeliveryDate           *string `json:"ActualDeliveryDate"`
	InvoicingPeriodStartDate     *string `json:"InvoicingPeriodStartDate"`
	InvoicingPeriodEndDate       *string `json:"InvoicingPeriodEndDate"`
	PaymentMeansTypeCode         *string `json:"PaymentMeansTypeCode"`
	PaidAmount                   *string `json:"PaidAmount"`
	SumOfInvoiceLineNetAmount    *string `json:"SumOfInvoiceLineNetAmount"`
	InvoiceTotalAmountWithoutTAX *string `json:"InvoiceTotalAmountWithoutTAX"`
	InvoiceTotalTAXAmount        *string `json:"InvoiceTotalTAXAmount"`
	InvoiceTotalAmountWithTAX    *string `json:"InvoiceTotalAmountWithTAX"`
	ExternalDocumentLocation     *string `json:"ExternalDocumentLocation"`
}

type MappingItem struct {
	InvoiceNumber                          string  `json:"InvoiceNumber"`
	InvoiceLineIdentifier                  string  `json:"InvoiceLineIdentifier"`
	InvoiceLineNote                        *string `json:"InvoiceLineNote"`
	ItemStandardIdentifier                 *string `json:"ItemStandardIdentifier"`
	InvoicedQuantity                       *string `json:"InvoicedQuantity"`
	InvoicedQuantityUnitOfMeasureCode      *string `json:"InvoicedQuantityUnitOfMeasureCode"`
	InvoiceLineNetAmountn                  *string `json:"InvoiceLineNetAmountn"`
	ReferencedPurchaseOrderLineReference   *string `json:"ReferencedPurchaseOrderLineReference"`
	OrderReference                         *string `json:"OrderReference"`
	InvoiceLinePeriodStartDate             *string `json:"InvoiceLinePeriodStartDate"`
	InvoiceLinePeriodEndDate               *string `json:"InvoiceLinePeriodEndDate"`
	InvoiceLineChargeAmount                *string `json:"InvoiceLineChargeAmount"`
	ItemNetPrice                           *string `json:"ItemNetPrice"`
	ItemPriceBaseQuantity                  *string `json:"ItemPriceBaseQuantity"`
	ItemPriceBaseQuantityUnitOfMeasureCode *string `json:"ItemPriceBaseQuantityUnitOfMeasureCode"`
	ActualDeliveryDate                     *string `json:"ActualDeliveryDate"`
	TAXCategoryCode                        *string `json:"TAXCategoryCode"`
	ItemCountryOfOrigin                    *string `json:"ItemCountryOfOrigin"`
}

type MappingHeaderPartner struct {
	InvoiceNumber   string  `json:"InvoiceNumber"`
	PartnerFunction string  `json:"PartnerFunction"`
	PartnerID       string  `json:"PartnerID"`
	PartnerName     *string `json:"PartnerName"`
	AddressID       *int    `json:"AddressID"`
}

type MappingItemPartner struct {
	InvoiceNumber   string  `json:"InvoiceNumber"`
	PartnerFunction string  `json:"PartnerFunction"`
	PartnerID       string  `json:"PartnerID"`
	PartnerName     *string `json:"PartnerName"`
	AddressID       *int    `json:"AddressID"`
}

type MappingAddress struct {
	InvoiceNumber string  `json:"InvoiceNumber"`
	AddressID     int     `json:"AddressID"`
	PostalCode    *string `json:"PostalCode"`
	LocalRegion   *string `json:"LocalRegion"`
	Country       *string `json:"Country"`
	District      *string `json:"District"`
	StreetName    *string `json:"StreetName"`
	CityName      *string `json:"CityName"`
	Building      *string `json:"Building"`
	Floor         *int    `json:"Floor"`
	Room          *int    `json:"Room"`
}

// 番号変換
type CodeConversionKey struct {
	SystemConvertTo   string `json:"SystemConvertTo"`
	SystemConvertFrom string `json:"SystemConvertFrom"`
	LabelConvertTo    string `json:"LabelConvertTo"`
	LabelConvertFrom  string `json:"LabelConvertFrom"`
	CodeConvertFrom   string `json:"CodeConvertFrom"`
	BusinessPartner   int    `json:"BusinessPartner"`
}

type CodeConversionQueryGets struct {
	CodeConversionID  int    `json:"CodeConversionID"`
	SystemConvertTo   string `json:"SystemConvertTo"`
	SystemConvertFrom string `json:"SystemConvertFrom"`
	LabelConvertTo    string `json:"LabelConvertTo"`
	LabelConvertFrom  string `json:"LabelConvertFrom"`
	CodeConvertFrom   string `json:"CodeConvertFrom"`
	CodeConvertTo     string `json:"CodeConvertTo"`
	BusinessPartner   int    `json:"BusinessPartner"`
}

type CodeConversionHeader struct {
	InvoiceDocument       int     `json:"InvoiceDocument"`
	InvoiceNumber         string  `json:"InvoiceNumber"`
	InvoiceCurrencyCode   *string `json:"InvoiceCurrencyCode"`
	TaxAccountingCurrency *string `json:"TaxAccountingCurrency"`
}

type CodeConversionItem struct {
	InvoiceLineIdentifier string `json:"InvoiceLineIdentifier"`
}

type ConversionData struct {
	InvoiceDocument int    `json:"InvoiceDocument"`
	InvoiceNumber   string `json:"InvoiceNumber"`
}

//個別処理
