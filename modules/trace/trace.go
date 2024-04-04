package trace

import (
	"fmt"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
)

var (
	machineIDOnce sync.Once
	machineID     string
	counter       uint32
)

func getMachineID() string {
	machineIDOnce.Do(func() {
		addrs, err := net.InterfaceAddrs()
		if err == nil {
			for _, addr := range addrs {
				if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						machineID = ipNet.IP.String()
						break
					}
				}
			}
		}

		if machineID == "" {
			machineID = "127.0.0.1"
		}
	})

	return machineID
}

func GenerateTraceID() string {
	machineID := getMachineID()
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	processID := os.Getpid()
	id := atomic.AddUint32(&counter, 1)
	id = id % 1000
	traceID := fmt.Sprintf("%08x%05d%013d%04d", machineIDToHex(machineID), processID, timestamp, id)

	return traceID
}

func machineIDToHex(machineID string) uint32 {
	ip := net.ParseIP(machineID)
	ipUint32 := uint32(ip[12])<<24 | uint32(ip[13])<<16 | uint32(ip[14])<<8 | uint32(ip[15])
	return ipUint32
}

func GetTraceID(ctx *context.Context) string {
	traceID, ok := ctx.GetUserValue(TraceIDKey).(string)
	if !ok {
		return ""
	}
	return traceID
}

const (
	TraceIDKey = "traceID"
)
