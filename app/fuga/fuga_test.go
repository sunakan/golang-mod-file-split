package fuga

// fugaパッケージからhogeパッケージをimportする時
// module名/
import (
	"golang-mod-file-split/hoge"
	"testing"
)

func TestUser2(t *testing.T) {
	user := hoge.NewUser("suna2")
	if user.Name != "suna2" {
		t.Errorf("add() = %v, want %v", user.Name, "suna2")
	}
	t.Log("DONE")
}
