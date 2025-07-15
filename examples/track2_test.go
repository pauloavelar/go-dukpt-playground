package examples

import (
	"encoding/hex"
	"fmt"

	"github.com/pauloavelar/go-dukpt-playground/dukpt"
)

func ExampleEncryptTrack2() {
	enc, err := dukpt.EncryptTrack2([]byte("TODO"), []byte{}, []byte{})

	fmt.Println(hex.EncodeToString(enc), err)
}
