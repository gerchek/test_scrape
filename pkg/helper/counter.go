package helper

import "github.com/sirupsen/logrus"

// product total count per category
var ProductTotalCount = 1000

var TotalProductLimit int = 10000
var InsertedProductCount int = 0

func IsProductLimitReached() bool {
	logrus.Infoln("TotalProductLimit: ", TotalProductLimit)
	logrus.Infoln("InsertedProductCount: ", InsertedProductCount)
	return TotalProductLimit <= InsertedProductCount
}
