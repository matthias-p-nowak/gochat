package db

import (
	"context"
	"gochat/db/ent"
	"gochat/utils"

	"gopkg.in/ini.v1"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var (
	DbStorage *ent.Client // The database client
)

func Initialize(config *ini.File) {
	utils.Log("Initializing storage")
	dbSection := config.Section("database")
	dbDriver := dbSection.Key("driver").String()
	dbDatabase := dbSection.Key("database").String()
	dbOptions := dbSection.Key("dbOptions").String()
	var err error
	DbStorage, err = ent.Open(dbDriver, dbDatabase+dbOptions)
	if err != nil {
		utils.Fatal(err)
	}
	ctx := context.Background()

	// Run the auto migration tool.
	if err := DbStorage.Schema.Create(ctx); err != nil {
		utils.Fatal(err)
	}
}

func SyncStorage() {
	utils.Log("Syncing storage")
}
