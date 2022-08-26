package libs

import (
	"testing"
)

//InitMySql use a function "once". Connection is generated once and if you try to run both tests at same time
//error test doesn't work 'cause instance is already exists in first test
func TestInitMySql(t *testing.T) {
	t.Run("InitMySqlSuccess", func(t *testing.T) {
		got := InitMySql()
		if got == nil {
			t.Errorf("Got %v - Want %v", got, "*gorm.DB")
		}
	})
	// t.Run("InitMySqlError", func(t *testing.T) {
	// 	user = "error"
	// 	password = "1235465461"
	// 	got := InitMySql()
	// 	if got != nil {
	// 		t.Errorf("Got %v - Want %v", got, "*gorm.DB")
	// 	}
	// })
}
