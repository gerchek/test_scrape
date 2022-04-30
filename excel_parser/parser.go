package excel_parser

import (
	"fmt"
	"os"
	"scraper_trendyol/couch_db"
	"scraper_trendyol/pkg/logging"

	"github.com/Jeffail/gabs/v2"
	"github.com/xuri/excelize/v2"
)

type ExcelParser struct {
	// couch database client
	cc couch_db.CouchDBClient

	jsonMap map[int]string
}

func NewExcelParser() (ExcelParser, error) {

	ep := ExcelParser{}

	cdb, err := couch_db.NewCouchDB(os.Getenv("db_ty_categories"))
	if err != nil {
		return ep, err
	}

	ep.cc = cdb

	ep.jsonMap = map[int]string{
		0: "id",
		1: "order",
		2: "parent_id",
		3: "name",
		4: "sarga_id",
		5: "slug",
		6: "weight",
		7: "createdAt",
		8: "updatedAt",
	}

	return ep, nil
}

func (ep *ExcelParser) ParseExcelAndInsert() error {

	f, err := excelize.OpenFile("excel_parser/ty_categories.xlsx")
	if err != nil {
		logging.Error(err)
		return err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the ty_categories.
	rows, err := f.GetRows("ty_categories")
	if err != nil {
		logging.Error(err)
		return err
	}
	for _, row := range rows {
		json := gabs.New()

		for i, colCell := range row {
			// fmt.Print(colCell, "\t")
			json.Set(colCell, ep.jsonMap[i])

		}
		// fmt.Print(json.String())

		ep.cc.InsertCategoryData(json)

		// fmt.Println()
	}

	return nil
}

func (ep *ExcelParser) GetCategoryWeight(categorySlug string) string {
	return ep.cc.GetCategoryWeight(categorySlug)
}

// // ---------------------------------------------------------------------------------------------------------crud

func (ep *ExcelParser) GetExcelData() {

	ep.cc.GetExcelDoc()
}
