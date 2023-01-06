package convert_complementer

import (
	dpfm_api_input_reader "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Processing_Formatter"
)

// Mapping Headerの処理
func (c *ConvertComplementer) ComplementMappingHeader(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*dpfm_api_processing_formatter.MappingHeader, error) {
	res, err := psdc.ConvertToMappingHeader(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}
