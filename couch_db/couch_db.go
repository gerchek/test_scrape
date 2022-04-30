package couch_db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"scraper_trendyol/models"
	"scraper_trendyol/models/bagisto"
	"scraper_trendyol/pkg/logging"
	"scraper_trendyol/uploader"

	"strings"

	"github.com/Jeffail/gabs/v2"
	_ "github.com/go-kivik/couchdb/v4" // The CouchDB driver
	kk "github.com/go-kivik/kivik/v4"

	// "github.com/google/go-cmp/cmp"
	// "github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sirupsen/logrus"
)

//CouchDBClient inserts to couch_db and calls upload function
type CouchDBClient struct {
	// couch database
	cdb *kk.DB

	// uploader
	uploader uploader.Uploader
}

// NewCouchDB create new instance of CouchDB
func NewCouchDB(database string) (CouchDBClient, error) {

	couchDbClient := CouchDBClient{}

	dataSource := os.Getenv("couch_db_source")

	client, err := kk.New("couch", dataSource)
	if err != nil {
		logging.Error(err)
		return couchDbClient, err
	}

	connStatus, _ := client.Ping(context.TODO())
	if !connStatus {

		/*cmd := exec.Command("sudo service couchdb restart")

		err := cmd.Run()
		if err != nil {
			logrus.Error("cmd run error", err)
			connStatusErr := errors.New("can not connect to couch_db")
			logging.Error(connStatusErr)
			return couchDbClient, connStatusErr
		} else {
			logrus.Infoln("couch_db stared successfully")
		}*/

		connStatusErr := errors.New("can not connect to couch_db")
		logging.Error(connStatusErr)
		return couchDbClient, connStatusErr
	}

	couchDbClient.cdb = client.DB(database)

	couchDbClient.uploader, err = uploader.NewUploader()
	if err != nil {
		return couchDbClient, err
	}

	return couchDbClient, nil
}

//InsertData receives json docs and inserts it to CouchDB
func (cc *CouchDBClient) InsertProductData(doc *gabs.Container) error {

	// get product data from json
	docId, _ := doc.Search("product_group_id").Data().(string)

	// Find doc
	row := cc.cdb.Get(context.TODO(), docId)

	// if document is not added before
	if row.Body == nil {

		// Product not added before, insert product
		err := cc.insertProductDocument(docId, doc, "create")
		return err

	} else {
		logrus.Infoln("document already inserted: ", docId)

		/* // product saved in couchDB
		var product bagisto.BagistoProductModel

		// document product
		var docProduct bagisto.BagistoProductModel
		err := json.Unmarshal(doc.Bytes(), &docProduct)
		if err != nil {
			logging.Error(err)
			return err
		}

		// scan couchDb product
		err = row.ScanDoc(&product)
		if err != nil {
			logging.Error(err)
			return err
		}

		// check product changed
		isEqual := cmp.Equal(docProduct, product, cmpopts.IgnoreFields(docProduct, "Categories"))

		// if is not equal, product changed and remove old product and insert new one
		if !isEqual {
			// remove doc by revision
			_, err = cc.cdb.Delete(context.TODO(), docId, row.Rev)
			if err != nil {
				logging.Error(err)
				return err
			}

			cc.insertProductDocument(docId, doc)
			return err
		} */
	}

	return nil
}

//UpdateProductData receives json docs and inserts it to CouchDB
func (cc *CouchDBClient) UpdateProductData(doc *gabs.Container) error {

	// get product data from json
	docId, _ := doc.Search("product_group_id").Data().(string)
	err := cc.insertProductDocument(docId, doc, "update")

	if err != nil {
		return err
	}

	return nil
}

func (cc *CouchDBClient) insertProductDocument(docId string, doc *gabs.Container, operation string) error {

	// document product
	var docProduct bagisto.BagistoProductModel
	err := json.Unmarshal(doc.Bytes(), &docProduct)
	if err != nil {
		logging.Error(err)
		return err
	}

	err = cc.uploader.UploadToRemoteServer(docProduct, operation)

	if err != nil {
		return err
	}

	fmt.Println("Data inserted to Bagisto server")

	if operation == "create" {
		// insert to couch_db
		rev, err := cc.cdb.Put(context.TODO(), docId, doc)
		if err != nil {
			logging.Error(err)
			return err
		}

		fmt.Println("data inserted with revision \n", rev)
	} else {
		row := cc.cdb.Get(context.TODO(), docId)

		doc.Set(row.Rev, "_rev")

		rev, err := cc.cdb.Put(context.TODO(), docId, doc)
		if err != nil {
			logging.Error(err)
			return err
		}

		fmt.Println("data updated with revision \n", rev)
	}

	return nil
}

//InsertCategory receives json docs and inserts it to CouchDB
func (cc *CouchDBClient) InsertCategoryData(doc *gabs.Container) error {

	// get product data from json
	docId, _ := doc.Search("id").Data().(string)

	// Product not added before, insert product
	_, err := cc.cdb.Put(context.TODO(), docId, doc)
	if err != nil {
		logging.Error(err)
		return err
	}

	logrus.Println("doc with id", docId, "inserted successfully")

	return nil
}

func (cc *CouchDBClient) GetCategoryWeight(slug string) string {

	category := models.Category{}

	query := map[string]interface{}{
		"selector": map[string]interface{}{
			"slug": strings.ToLower(slug),
		},
	}

	rows, err := cc.cdb.Find(context.TODO(), query)

	if err != nil {
		logging.Error(err)
		return "0.5"
	}

	for rows.Next() {

		if err := rows.ScanDoc(&category); err != nil {
			logging.Error(err)
			return "0.5"
		}

	}

	if (category == models.Category{}) {
		err := fmt.Errorf("category not found")
		logging.Error(err.Error())

		return "0.5"
	}

	return category.Weight
}

func (cc *CouchDBClient) IsProductExists(productGroupId string) bool {

	row := cc.cdb.Get(context.TODO(), productGroupId)

	return row.Body != nil

}

// // =================================================================================================start

func (cc *CouchDBClient) GetExcelDoc() {

	row := cc.cdb.Get(context.TODO(), "ty_categories")
	fmt.Println(row)
}
