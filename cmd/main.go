package main

import (
	"construction-organization-system/internal/config"
	"construction-organization-system/internal/database"
	"construction-organization-system/internal/log"
)

func main() {
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		log.Logger.WithError(err).Errorln("Db config isn't init")
		return
	}

	log.Logger.Infoln("Connect to database...")
	db, err := database.NewDB(dbConfig)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on db connection")
		return
	}

	log.Logger.Infoln("Successful database connection")

	log.Logger.Infoln("Migrate database...")
	err = database.Migrate(db, "deploy/migrations")
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on migrating")
		return
	}

	log.Logger.Infoln("Successful database migrate")
}
