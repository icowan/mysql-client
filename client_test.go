/**
 * @Time : 2020/8/24 10:45 AM
 * @Author : solacowa@gmail.com
 * @File : client_test
 * @Software: GoLand
 */

package mysqlclient

import (
	"fmt"
	"testing"
)

func TestNewMysql(t *testing.T) {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local&timeout=20m&collation=utf8mb4_unicode_ci",
		"root",
		"password",
		"127.0.0.1",
		3306,
		"dbname")

	db, err := NewMysql(dbUrl, true)
	if err != nil {
		t.Error(err)
	}

	db.Select("SUM")
}
