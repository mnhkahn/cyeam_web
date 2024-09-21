package controllers

import (
	"os"
	"testing"

	ddlmaker "github.com/mnhkahn/ddl-maker"
	"github.com/stretchr/testify/assert"
)

func TestDDL(t *testing.T) {
	data := `{"a":1}`

	conf := ddlmaker.Config{
		DB: ddlmaker.DBConfig{
			Driver:  "mysql",
			Engine:  "InnoDB",
			Charset: "utf8mb4",
		},
		OutFilePath: os.TempDir(),
	}

	dm, err := ddlmaker.New(conf)
	assert.Nil(t, err)
	res, err := dm.GenerateJSON(data)
	assert.Nil(t, err)
	t.Log(string(res))
}
