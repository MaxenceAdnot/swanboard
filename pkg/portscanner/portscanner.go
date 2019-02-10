package portscanner

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type PortScanner struct {
	Ip   string
	Lock *semaphore.Weighted
}

func Ulimit() int64 {
	return 500
}

func ScanPort(ip string, port int, timeout time.Duration) int {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		}
		return 0
	}

	conn.Close()
	return port
}

func (ps *PortScanner) Start(timeout time.Duration) []int {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	openPorts := []int{}

	portRange := []int{80, 8080, 8000, 3000, 3030}

	for _, port := range portRange {
		ps.Lock.Acquire(context.TODO(), 1)
		wg.Add(1)
		go func(port int) {
			defer ps.Lock.Release(1)
			defer wg.Done()
			scanned := ScanPort(ps.Ip, port, timeout)
			if scanned > 0 {
				openPorts = append(openPorts, port)
			}
		}(port)
	}

	wg.Wait()
	return openPorts
}
