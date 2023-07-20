package main

import (
	//"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"
	"unsafe"

	serial "github.com/goburrow/serial"
	_ "github.com/lib/pq"
)

func main() {
	//var U1 float32
	port, err := serial.Open(
		&serial.Config{
			Address:  "COM5",
			BaudRate: 9600,
			DataBits: 8,
			StopBits: 1,
			Parity:   "N",
			Timeout:  1 * time.Second,
		})
	if err != nil {
		log.Fatal("Comport open fail")
	}
	defer port.Close()

	wbuf := []byte{03, 03, 70, 02, 00, 04, 71, 19} //這裡要改驗證 7E E9 寫一個可以算crc的程式
	//data := make([]byte, 1024)
	dataTmp := make([]byte, 1024)

	_, err = port.Write(wbuf)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1000 * time.Millisecond) //等待回傳所需的時間1000ms

	dataLength := 0

	expectedDataLength := 9 // 假設要讀取 14 個字節的資料

	for dataLength < expectedDataLength {
		n := 1
		n, err = port.Read(dataTmp[dataLength:]) //讀資料回來，使用 dataTmp[dataLength:] 確保從目前位置開始存放資料

		if err != nil {
			log.Fatal("Error reading data:", err)
		}

		dataLength += n
	}
	fmt.Println(dataTmp)
	// encodedString := hex.EncodeToString(data[:dataLength])
	// fmt.Println("Encoded Hex String: ", encodedString)

	// U1 = hexToDec(encodedString[6:13]) //順序不知道對不對
	// fmt.Println(U1)
}

func hexToDec(hexString string) float32 {
	s := hexString[4:8] + hexString[0:4]
	fmt.Println(s)
	n, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		panic(err)
	}
	n2 := uint32(n)
	f := *(*float32)(unsafe.Pointer(&n2))
	return f
}
