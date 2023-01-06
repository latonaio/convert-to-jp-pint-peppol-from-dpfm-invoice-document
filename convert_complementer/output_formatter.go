package convert_complementer

import (
	dpfm_api_input_reader "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Processing_Formatter"
)

func (c *ConvertComplementer) SetValue(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) (*dpfm_api_output_formatter.SDC, error) {
	var header *dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var err error

	header, err = dpfm_api_output_formatter.ConvertToHeader(*sdc, *psdc)
	if err != nil {
		return nil, err
	}

	item, err = dpfm_api_output_formatter.ConvertToItem(*sdc, *psdc)
	if err != nil {
		return nil, err
	}

	osdc.Message = dpfm_api_output_formatter.Message{
		Header: *header,
		Item:   *item,
	}

	return osdc, nil
}
