package db

import (
	"context"

	"github.com/Faaizz/simple_http_chatapp/types"
)

var dba types.DBAdapter

// set database adapter
func SetDatabaseAdapter(dbaInit types.DBAdapter) {
	dba = dbaInit
}

// CheckExists sets TableName on DatabaseAdapter and checks if table tableName exists
func CheckExists(tableName string) error {
	dba.SetTableName(tableName)
	return dba.CheckExists(context.TODO())
}

// PutConn adds an entry into the table with a connectionId, username pair
func PutConn(pcIn types.User) error {
	return dba.PutConn(context.TODO(), pcIn)
}

func Delete(data map[string]string) error {
	return nil
}

func GetUserConnId(username string) (string, error) {
	return "connectionId", nil
}
