package test1

import "testing"

type ucTest struct {
	in, out string
}

var ucTests = []ucTest{
	ucTest{"abc", "ABC"},
	ucTest{"cvo-az", "CVO-AZ"},
	ucTest{"Antwerp", "ANTWERP"},
}

func TestUC(t *testing.T) {
	for _, ut := range ucTests {
		uc := UpperCase(ut.in)
		if uc != ut.out {
			t.Errorf("UpperCase(%s) = %s, must be %s", ut.in, uc, ut.out)
		}
	}
}

func BenchmarkUpperCase(b *testing.B) {
	for _, ut := range ucTests {
		uc := UpperCase(ut.in)
		if uc != ut.out {
			b.Errorf("UpperCase(%s) = %s, must be %s", ut.in, uc, ut.out)
		}
	}
}

// test格式
func TestXxx(t *testing.T) {

}
