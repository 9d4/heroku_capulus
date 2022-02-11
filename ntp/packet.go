package ntp

import (
	"encoding/binary"
	"log"
	"net"
	"time"

	"github.com/traperwaze/heroku_capulus/config"
)

type packet struct {
	Settings       uint8  // leap yr indicator, ver number, and mode
	Stratum        uint8  // stratum of local clock
	Poll           int8   // poll exponent
	Precision      int8   // precision exponent
	RootDelay      uint32 // root delay
	RootDispersion uint32 // root dispersion
	ReferenceID    uint32 // reference id
	RefTimeSec     uint32 // reference timestamp sec
	RefTimeFrac    uint32 // reference timestamp fractional
	OrigTimeSec    uint32 // origin time secs
	OrigTimeFrac   uint32 // origin time fractional
	RxTimeSec      uint32 // receive time secs
	RxTimeFrac     uint32 // receive time frac
	TxTimeSec      uint32 // transmit time secs
	TxTimeFrac     uint32 // transmit time frac
}

var host string = config.Config.NtpServer

var conn *net.Conn

func init() {
	connectNTP()
}

func connectNTP() {
	_conn, err := net.Dial("udp", host)
	if err != nil {
		log.Fatal("failed to connect ntp:", err)
	}

	if err := _conn.SetDeadline(time.Now().Add(time.Duration(15) * time.Second)); err != nil {
		log.Fatal("failed to set deadline:", err)
	}

	conn = &_conn
}

func requestForTime() packet {
	// 00 011 011 (0x1B)
	// 00 : leap year indicator
	// 011 : version 3
	// 011 : client mode

	req := &packet{Settings: 0x1B}

	if err := binary.Write(*conn, binary.BigEndian, req); err != nil {
		log.Fatal("failed to send request:", err)
	}

	res := &packet{}
	if err := binary.Read(*conn, binary.BigEndian, res); err != nil {
		log.Fatal("failed to read response:", err)
	}

	return *res
}

func parseTime(rsp packet) time.Time {
	const ntpEpochOffset = 2208988800 // 70 years of second

	secs := float64(rsp.TxTimeSec) - ntpEpochOffset
	nanos := (int64(rsp.TxTimeFrac) * 1e9) >> 32

	return time.Unix(int64(secs), nanos)
}

func GetTime() time.Time {
	rsp := requestForTime()
	return parseTime(rsp)
}
