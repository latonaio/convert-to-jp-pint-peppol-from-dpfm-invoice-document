package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header        Header          `json:"Header"`
	Item          []Item          `json:"Item"`
	Address       Address         `json:"Address"`
	HeaderPartner []HeaderPartner `json:"HeaderPartner"`
	ItemPartner   []ItemPartner   `json:"ItemPartner"`
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
