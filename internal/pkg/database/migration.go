package database

import (
	"construction-organization-system/internal/pkg/log"
	"database/sql"
	"errors"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Migrate(db *sql.DB, scriptsDir string) error {
	scripts, err := os.ReadDir(scriptsDir)
	if err != nil {
		log.Logger.WithError(err).Errorln("Error on reading directory: ", scriptsDir)
		return err
	}

	sortMigrationScripts(scripts)

	for _, script := range scripts {
		bytes, err := os.ReadFile(scriptsDir + "/" + script.Name())
		if err != nil {
			log.Logger.WithError(err).Errorln("Error on reading script:", script.Name())
			return err
		}

		query := string(bytes)
		result, err := db.Exec(query)
		if err != nil {
			log.Logger.WithError(err).Errorln("Error executing db query from", script.Name())
			return err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Logger.WithError(err).Errorln("Error on getting affected rows")
			return err
		}

		log.Logger.Infoln("Script", script.Name(), "executed successfully. Rows affected: ", rowsAffected)
	}

	return nil
}

func sortMigrationScripts(scripts []os.DirEntry) {
	sort.Slice(scripts, func(i, j int) bool {
		num1, _ := getMigrationNumber(scripts[i].Name())
		num2, _ := getMigrationNumber(scripts[j].Name())

		return num1 < num2
	})
}

func getMigrationNumber(scriptName string) (int, error) {
	scriptNumberStr, _, _ := strings.Cut(scriptName, "_")
	scriptNumber, err := strconv.Atoi(scriptNumberStr)
	if err != nil {
		return -1, errors.New("no migration number in script name: " + scriptName)
	}

	return scriptNumber, nil
}
