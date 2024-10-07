package service

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func BinToDex(b string) string {
	dotI := strings.Index(b, ".")
	resI := 0.0
	i := dotI
	for _, bb := range b {
		if bb != 46 {
			a := (float64(bb) - 48.0) * math.Pow(2.0, float64(i-1))
			resI += a
			// fmt.Println(bb-48, a, i-1, resI)
			i--
		}
	}
	return strconv.FormatFloat(resI, 'f', 6, 64)
}

func DexToBin(pi float64) {
	buf := new(bytes.Buffer)
	// var pi float64 = 120.5
	err := binary.Write(buf, binary.BigEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("%b", buf.Bytes())
	fmt.Println("")
}
