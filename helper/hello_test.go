package helper

import "testing"

func TestHello(t *testing.T) {

	result := Hello("nopal")
	if result != "nopal" {
		//eror
		panic("Hasil e salah cok")
	}
}
