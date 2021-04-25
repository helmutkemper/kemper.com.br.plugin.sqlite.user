package main

import (
	"github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/kemper.com.br/util"
)

func (e *SQLiteUser) New() (referenceInicialized interface{}, err error) {
	referenceInicialized = &SQLiteUser{}
	err = referenceInicialized.(*SQLiteUser).Connect(constants.KSQLiteConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = referenceInicialized.(*SQLiteUser).Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
