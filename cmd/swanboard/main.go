package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/maxenceadnot/swanboard/pkg/portscanner"
	"golang.org/x/sync/semaphore"
)

func main() {

	http.Handle("/", http.FileServer(rice.MustFindBox("../../web/dist").HTTPBox()))

	http.HandleFunc("/api/getusers", func(w http.ResponseWriter, r *http.Request) {
		userInfo := []map[string]string{}
		out, err := exec.Command("ipsec", "leases").Output()
		if err != nil {
			log.Fatal(err)
			return
		}
		arrayStr := strings.Split(string(out), "\n")
		for _, line := range arrayStr[1:] {
			userLine := strings.Fields(line)
			if len(userLine) > 0 {
				portsOpen := []int{}
				if userLine[1] == "online" {
					ps := &portscanner.PortScanner{
						Ip:   userLine[0],
						Lock: semaphore.NewWeighted(portscanner.Ulimit()),
					}
					portsOpen = ps.Start(1000 * time.Millisecond)
				}
				entry := map[string]string{"ip": userLine[0], "status": userLine[1], "username": strings.Replace(userLine[2], "'", "", -1), "ports": strings.Trim(strings.Join(strings.Fields(fmt.Sprint(portsOpen)), ","), "[]")}
				userInfo = append(userInfo, entry)
			}
		}
		jsonRes, err := json.MarshalIndent(userInfo, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, string(jsonRes))
	})

	http.ListenAndServe(":80", nil)
}
