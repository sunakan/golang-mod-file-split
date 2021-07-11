package hoge

import (
	"testing"
)

func TestUser(t *testing.T) {
	user := NewUser("suna")
	if user.Name != "suna" {
		t.Errorf("add() = %v, want %v", user.Name, "suna")
	}
	t.Log("DONE")
}
