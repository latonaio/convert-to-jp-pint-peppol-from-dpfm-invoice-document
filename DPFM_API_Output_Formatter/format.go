package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Processing_Formatter"
	"encoding/json"
)

func ConvertToHeader(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*Header, error) {
	mappingHeader := psdc.MappingHeader

	header := Header{}

	data, err := json.Marshal(mappingHeader)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &header)
	if err != nil {
		return nil, err
	}

	return &header, nil
}

func ConvertToItem(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]Item, error) {
	mappingItem := psdc.MappingItem

	item := []Item{}
	inputItem := sdc.InvoiceDocumentHeader.InvoiceDocumentItem
	inputData, err := json.Marshal(inputItem)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inputData, &item)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(mappingItem)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
