package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/qredo/apitool/defs"
	"github.com/qredo/apitool/webui"

	"github.com/pkg/errors"
)

var (
	flagURL       = flag.String("url", "", "URL Path to be used")
	flagMethod    = flag.String("method", "", "HTTP method to be used")
	flagAPIKey    = flag.String("api-key", "", "API key")
	flagAPISecret = flag.String("secret", "", "API Secret")
	flagTimestamp = flag.String("timestamp", "", "Timestamp to be used")
	flagPort      = flag.String("port", "4569", "Port to listen to for the web ui")
)

func printHelpAndExit() {
	fmt.Printf("Usage: %s [options] sign|send|ui\n\nOptions:\n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
	os.Exit(1)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func setupRequest(req *defs.Request) error {
	if *flagTimestamp != "" {
		req.Timestamp = strings.TrimSpace(*flagTimestamp)
	} else {
		req.Timestamp = fmt.Sprintf("%v", time.Now().Unix())
	}

	req.ApiKey = strings.TrimSpace(*flagAPIKey)
	req.Method = *flagMethod
	req.URL = *flagURL

	if err := req.Validate(); err != nil {
		return errors.Wrap(err, "validation")
	}

	if *flagAPISecret == "" {
		return errors.New("no secret")
	}

	switch *flagMethod {
	case "POST", "PUT", "PATCH":
		req.Body = getBody()
	}

	secret, err := base64.URLEncoding.DecodeString(*flagAPISecret)
	if err != nil {
		return errors.Wrap(err, "invalid secret")
	}

	req.Sign(secret)

	return nil
}

func getBody() []byte {

	terminatorMessage := "hit Ctrl+D (twice if body has no trailing new line) to end"
	if runtime.GOOS == "windows" {
		terminatorMessage = "hit Ctrl+Z followed by <enter> to end"
	}

	fmt.Printf("body (%s):\n", terminatorMessage)
	reader := bufio.NewReader(os.Stdin)
	readBuffer := make([]byte, 10024)
	var bodyBuffer []byte
	for {
		n, err := reader.Read(readBuffer)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if n == 0 {
			continue
		}
		if n > 1 {
			if readBuffer[n-2] == '\r' && readBuffer[n-1] == '\n' {
				readBuffer[n-2] = '\n'
				n--
			}
		}
		bodyBuffer = append(bodyBuffer, readBuffer[:n]...)
	}

	if len(bodyBuffer) != 0 {
		return bodyBuffer
	}

	return nil
}

func main() {
	argc := len(os.Args)
	if argc < 2 {
		printHelpAndExit()
	}
	if err := flag.CommandLine.Parse(os.Args[1 : argc-1]); err != nil {
		printHelpAndExit()
	}

	req := &defs.Request{}

	switch os.Args[argc-1] {
	case "sign":
		checkErr(setupRequest(req))
		fmt.Printf("qredo-api-sign: %s\n", req.Signature)
		fmt.Printf("qredo-api-key: %s\n", req.ApiKey)
		fmt.Printf("qredo-api-ts: %s\n", req.Timestamp)
	case "send":
		checkErr(setupRequest(req))
		resp, err := req.Send()
		checkErr(err)
		fmt.Printf("\n%s\n", resp)
	case "ui":
		webui.Serve(*flagPort)
	default:
		fmt.Println("no command")
	}
}
