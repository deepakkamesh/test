package main

import (
	"fmt"
	"time"

	"github.com/tarm/serial"
)

func main() {
	fmt.Printf("hello serial1\n")
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 4800, ReadTimeout: time.Second * 1}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Printf("Err opening serial")
		return
	}
	// Setup
	/*buf := make([]byte, 10)
	for {
		n, err := s.Read(buf)
		if err != nil {
			fmt.Printf("prob %s", err)
		}
		fmt.Printf("Read s: %x\n", buf[:n])
		if string(buf[:n]) == "Z" {
			WriteSerial(s, []byte{0xc3})
		}
			if string(buf[:n]) == "" {
				break

			}
	 }*/

	// Write some
	//b := []byte{0x4, 0x2A, 0x0, 0x6, 0x23, 0x0}
	WriteSerial(s, []byte{0x4, 0x2A})
	ReadSerial(s)
	WriteSerial(s, []byte{0x00})
	ReadSerial(s)
	WriteSerial(s, []byte{0x86, 0x23})
	ReadSerial(s)
	WriteSerial(s, []byte{0x00})
	ReadSerial(s)

	for {
	}
}
func ReadSerial(s *serial.Port) {
	buf := make([]byte, 128)
	n, err := s.Read(buf)

	if err != nil {
		fmt.Printf("read err %s", err)
	}
	fmt.Printf("Read: %x \n", buf[:n])
}

func WriteSerial(s *serial.Port, b []byte) {

	_, err := s.Write(b)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("Wrote: %x\n", b)
}
