package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"net/http"
	"net/url"
)

const (

	// BANNER is what printed for help/info output
	BANNER = `
  ____  _     ____  ____  _____ _____    _____ _      ____
 /  __\/ \ /\/  __\/  __\/  __//__ __\  /  __// \  /|/   _\
 |  \/|| | |||  \/||  \/||  \    / \    |  \  | |\ |||  /
 |  __/| \_/||  __/|  __/|  /_   | |    |  /_ | | \|||  \__
 \_/   \____/\_/   \_/   \____\  \_/    \____\\_/  \|\____/

 Puppet External Node Classifier
 Version: %s

`

	// VERSION is the binary version.
	VERSION = "v0.0.2"
)

var (
	node    string
	version bool
	host    string
	port    int
)

func init() {
	// parse flags
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&version, "v", false, "print version and exist (shorthand)")
	flag.StringVar(&host, "host", "localhost", "puppet dashboard URL")
	flag.IntVar(&port, "port", 3000, "port number")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if version {
		fmt.Printf("%s\n", VERSION)
		os.Exit(0)
	}

	if flag.NArg() >= 1 {
		node = flag.Arg(0)
	} else {
		usageAndExit("", 0)
	}
}

func main() {

	dashboardURL := fmt.Sprintf("http://%s:%d", host, port)
	u, err := url.Parse(dashboardURL + "/nodes/" + node)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Accept", "text/yaml")
	resp, err := client.Do(req)

	if nil != err {
		fmt.Println("Connection error:", err.Error())
		os.Exit(-1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if nil != err {
		fmt.Println("Error reading body: ", err.Error())
	}

	fmt.Printf("%v", string(body))
}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}

func redirectPolicyFunc(r *http.Request, rr []*http.Request) error {
	return errors.New("disable")
}
