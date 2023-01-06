package convert_complementer

import (
	"context"
	dpfm_api_input_reader "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "convert-to-jp-pint-peppol-from-dpfm-invoice-document/DPFM_API_Processing_Formatter"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ConvertComplementer struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewConvertComplementer(ctx context.Context, db *database.Mysql, l *logger.Logger) *ConvertComplementer {
	return &ConvertComplementer{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (c *ConvertComplementer) CreateSdc(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(10)

	// Header
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		psdc.MappingHeader, e = c.ComplementMappingHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	// Item
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		psdc.MappingItem, e = c.ComplementMappingItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	c.l.Info(psdc)
	osdc, err = c.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
