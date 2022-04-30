package setup

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LogrusSetup() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "02.01.2006 15:04:05"
	logrus.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
}

func LoadEnv() error {
	return godotenv.Load()
}

// LoadEnv loads env vars from .env
// const projectDirName = "scraper_trendyol"

// func LoadEnv() error {
// 	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
// 	cwd, _ := os.Getwd()
// 	rootPath := re.Find([]byte(cwd))

// 	err := godotenv.Load(string(rootPath) + `/.env`)

// 	return err
// 	if err != nil {
// 		logging.Error().WithFields(log.Fields{
// 			"cause": err,
// 			"cwd":   cwd,
// 		}).Fatal("Problem loading .env file")

// 		os.Exit(-1)
// 	}
// }
