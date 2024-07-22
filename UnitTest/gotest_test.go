package UnitTest

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("除法函式測試沒透過") // 如果不是如預期的那麼就報錯
	} else {
		t.Log("第一個測試通過了") //記錄一些你期望記錄的資訊
	}
}

func Test_Division_2(t *testing.T) {
	t.Error("就是不透過")
}
