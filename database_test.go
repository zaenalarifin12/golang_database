package golang_database

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func Test(t *testing.T) {

}

func TestConnection(t *testing.T) {
	//db, err := sql.Open("mysql", "root:rahasia@tcp(localhost:3306)/golanllllg_database")
	//if err != nil {
	//	panic(err.Error())
	//}
	db := GetConnection();
	defer db.Close()
}