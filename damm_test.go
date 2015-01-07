package damm

import "testing"

func TestDamm(t *testing.T) {
	s := "572"
	c, err := CheckDigit(s)
	if err != nil {
		t.Fatalf("CheckDigit error %s", err)
	}
	if c != "4" {
		t.Errorf("check digit %s; want %s", c, "4")
	}
	s1 := s + c
	t.Logf("s1 %s s %s c %s", s1, s, c)
	ok := Validate(s1)
	if !ok {
		t.Errorf("validate %s failed", s1)
	}
	s2 := "5274"
	ok = Validate(s2)
	if ok {
		t.Errorf("validate %s returned true", s2)
	}
}

func TestDamm2(t *testing.T) {
	s := "12345678901"
	c, err := CheckDigit(s)
	if err != nil {
		t.Fatalf("CheckDigit error %s", err)
	}
	t.Logf("check digit for %s: %s", s, c)
	s2 := "123156789018"
	ok := Validate(s2)
	if ok {
		t.Errorf("validate %s returned true", s2)
	}
}
