package convert_complementer

import (
	dpfm_api_input_reader "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Processing_Formatter"
)

// Mapping Itemの処理
func (c *ConvertComplementer) ComplementMappingItem(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*dpfm_api_processing_formatter.MappingItem, error) {
	res, err := psdc.ConvertToMappingItem(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}
