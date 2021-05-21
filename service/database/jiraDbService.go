package service

import (
	"github.com/nansuri/gp-server/config"

	model "github.com/nansuri/gp-server/model"
	logger "github.com/sirupsen/logrus"
)

func DbQueryTicketInfo(fieldName string) []model.FieldItem {

	var fieldItems []model.FieldItem
	var fieldItem model.FieldItem

	fieldItem.ID = 0

	db := config.Connect()

	result, err := db.Query("SELECT value,extend_info FROM jira_field_data WHERE field_name=?", fieldName)
	if err != nil {
		logger.WithFields(logger.Fields{"Error": err.Error()}).Info("DbQueryTicketInfo")
	}
	defer db.Close()

	for result.Next() {
		err := result.Scan(&fieldItem.Value, &fieldItem.Description)
		if err != nil {
			logger.WithFields(logger.Fields{"Error": err.Error()}).Info("DbQueryTicketInfo")
		}
		fieldItem.ID++
		fieldItems = append(fieldItems, fieldItem)
	}

	return fieldItems
}
