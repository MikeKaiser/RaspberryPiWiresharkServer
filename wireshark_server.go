// compile with "go build wireshark_server.go"

package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func execute(command string) {
	out, err := exec.Command("/bin/sh", "-c", command).CombinedOutput()
	if out != nil {
		fmt.Printf("out: %s\n", out)
	}
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
}

func executeNonBlocking(command string) {
	err := exec.Command("/bin/sh", "-c", command).Run()
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=/start>Start-Capture</a>\n<a href=/stop>Stop-Capture</a>\n<a href=/shutdown>Shutdown RPi</a>\n<a href=/chmod>Adjust-Permissions</a>")
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	executeNonBlocking("sudo tshark -i eth0 -w capture_" + time.Now().Format("20060102_150405") + ".pcap -F pcap </dev/null &>/dev/null &")
	fmt.Fprintf(w, "<a href=/stop>Stop-Capture</a>")
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
	execute("sudo pkill tshark")
	execute("sudo chmod 777 *.pcap")
	fmt.Fprintf(w, "<a href=/start>Start-Capture</a>\n<a href=/chmod>Adjust-Permissions</a>")
}

func chmodHandler(w http.ResponseWriter, r *http.Request) {
	execute("sudo chmod 777 *.pcap")
	mainHandler(w, r)
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	executeNonBlocking("sudo shutdown -h now &")
	fmt.Fprintf(w, "Shutting Down")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/stop", stopHandler)
	http.HandleFunc("/shutdown", shutdownHandler)
	http.HandleFunc("/chmod", chmodHandler)
	log.Fatal(http.ListenAndServe(":1988", nil))
}
