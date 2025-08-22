package tools

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"os"
	"sync/atomic"
	"time"
)

var (
	localIP     net.IP
	processID   uint16
	sequenceNum uint32
)

func init() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIP = ipnet.IP.To4()
				break
			}
		}
	}
	if localIP == nil {
		panic("can't get local IP address")
	}
	processID = uint16(os.Getpid())
}
func NewId() string {
	ts := uint32(time.Now().Unix())
	seq := atomic.AddUint32(&sequenceNum, 1)
	idBytes := make([]byte, 16)
	idBytes[0] = byte(ts >> 24)
	idBytes[1] = byte(ts >> 16)
	idBytes[2] = byte(ts >> 8)
	idBytes[3] = byte(ts)
	copy(idBytes[4:8], localIP.To4())
	idBytes[8] = byte(processID >> 8)
	idBytes[9] = byte(processID)
	idBytes[10] = byte(seq >> 24)
	idBytes[11] = byte(seq >> 16)
	idBytes[12] = byte(seq >> 8)
	idBytes[13] = byte(seq)
	idBytes[14] = 0
	idBytes[15] = 0
	return hex.EncodeToString(idBytes)
}

func Rand4Num() string {
	intn := rand.Intn(9999)
	if intn < 1000 {
		intn = intn + 1000
	}
	return fmt.Sprintf("%d", intn)
}

func Unq(prefix string) string {
	milli := time.Now().UnixMilli()
	intn := rand.Intn(999999)
	return fmt.Sprintf("%s%d%d", prefix, milli, intn)
}
