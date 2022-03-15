package helper

import "testing"

func TestHello(t *testing.T) {

	result := Hello("nopal")
	if result != "Hello nopal" {
		//eror
		t.Fail()
	}
	fmt.println("Berhasil")

}

func TestHelloNopal(t *testing.T) {

	result := Hello("bukan nopal")
	if result != "salah nopal" {
		//eror
		t.FailNow()
	}
	fmt.println("done")
}
