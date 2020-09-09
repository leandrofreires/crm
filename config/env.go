package config

import (
	"os"
)

type Environment struct {
	databaseURL string
}

var Env Environment

func init() {
	Env.databaseURL = os.Getenv("DATABASE_URL")
}
func (e *Environment) GetDatabaseURL() string {
	return e.databaseURL
}
