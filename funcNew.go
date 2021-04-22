package main

import (
	"github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/kemper.com.br/util"
)

func New() (referenceInicialized *SQLiteUser, err error) {
	referenceInicialized = &SQLiteUser{}
	err = referenceInicialized.Connect(constants.KSQLiteConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = referenceInicialized.Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
