package main

import (
	"encoding/binary"
	"fmt"
)

func main()  {
	data := []byte{0x12, 0x34, 0x56, 0x78}
	fmt.Printf("%b\n", data)
	//for i, b := range data {
	//	data[i] = (b&0x0F)<<4 | (b&0xF0)>>4
	//}
	data = []byte{0x56, 0x78,0x12, 0x34}
	fmt.Printf("%b\n", data)
	fmt.Printf("%X\n", data)

	//povID := uuid.FromStringOrNil("c28c2e75-9a19-4dc1-9497-48c855c5b154")
	//newPov := base64.StdEncoding.EncodeToString(povID.Bytes())
	//fmt.Println(newPov)
	//log.Info().Hex("pov", povID.Bytes()).Str("pkg", string(hex.Dump(povID.Bytes()))).Msg("sss")
	//
	//d := make(map[string]map[string]int)
	//c := make(map[string]int)
	//c["a"] = 10
	//d["a"] = c
	//
	//
	//fmt.Println(d["a"])
	//
	//pasePov, _ := base64.StdEncoding.DecodeString("+SVJuwrtTYqc+VbIwOSu1A==")
	//fmt.Println(uuid.FromBytesOrNil(pasePov).String())
	//
	//num := 1
	//
	//if num == 1 {
	//	fmt.Println("num value")
	//}
	//
	//data = []byte{0x00,0x08}
	//fmt.Println(int(binary.BigEndian.Uint16(data)))
fmt.Println(binary.BigEndian.Uint32([]byte{0x00, 0x00, 0x01, 0x84}))
}