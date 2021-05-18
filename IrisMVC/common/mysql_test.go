package common

import "testing"

func TestNewMysqlConn(t *testing.T) {
	_, err := NewMysqlConn()
	if err != nil {
		t.Error(err)
	}
}
