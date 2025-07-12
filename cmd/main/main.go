package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/pauloavelar/go-dukpt-playground/dukpt"
	"github.com/pauloavelar/go-dukpt-playground/track2"
)

var (
	testBDK = mustDecodeHex("0123456789ABCDEFFEDCBA9876543210") // 16 bytes
	testKSN = mustDecodeHex("12345678901234567890")             // 10 bytes
)

func main() {
	t2 := track2.Data{
		PAN:               "1234123412341234",
		ExpYear:           2025,
		ExpMonth:          12,
		ServiceCode:       "601",
		DiscretionaryData: "123123",
	}

	formatted, err := t2.FormatISO()
	if err != nil {
		panic(err)
	}

	encT2, err := dukpt.EncryptTrack2(formatted, testBDK, testKSN)
	if err != nil {
		panic(err)
	}

	fmt.Println("Track2    :", strings.ToUpper(hex.EncodeToString(formatted)))
	fmt.Println("Track2 ENC:", strings.ToUpper(hex.EncodeToString(encT2)))
}

func mustDecodeHex(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(fmt.Sprintf("failed to decode hex string %s: %v", s, err))
	}

	return data
}
