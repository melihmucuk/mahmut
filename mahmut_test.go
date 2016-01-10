package mahmut

import "testing"

func TestEncDec(t *testing.T) {
	secret := "your secret key!"
	expected := "mahmut sıkı bağla da düşmeyelim yeğenim"

	bagla, err := Bagla(secret, expected)
	if err != nil {
		t.Errorf("Test failed, error: %s", err.Error())
	}

	coz, err := Coz(secret, bagla)
	if err != nil {
		t.Errorf("Test failed, error: %s", err.Error())
	}

	if coz != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, coz)
	}
}
