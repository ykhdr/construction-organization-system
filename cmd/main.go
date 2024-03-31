package main

import (
	"construction-organization-system/internal/pkg/database"
	"construction-organization-system/internal/pkg/log"
	"construction-organization-system/pkg/config"
)

func main() {
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		log.Logger.WithError(err).Errorln("Db config isn't init")
		return
	}

	log.Logger.Infoln("Connect to database...")
	newDB, err := database.NewDB(dbConfig)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on database connection")
		return
	}

	log.Logger.Infoln("Successful database connection")

	log.Logger.Infoln("Migrate database...")
	err = database.Migrate(newDB, "db/migrations")
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on migrating")
		return
	}

	log.Logger.Infoln("Successful database migrate")
}
