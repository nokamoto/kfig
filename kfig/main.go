package main

import "fmt"
import "flag"
import "gopkg.in/yaml.v2"
import "io/ioutil"
import "os"
import "net/http"

func main() {
	filename := flag.String("yaml", "default.yaml", "a yaml configration file")

	api := flag.String("admin", "http://localhost:8001", "a kong admin api")

	flag.Parse()

	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s not found\n", *filename)
		os.Exit(1)
	}

	config := Config{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Fprintf(os.Stderr, "%s is not valid: %s\n", *filename, err)
		os.Exit(1)
	}

	fmt.Printf("%s/status\n", *api)

	if res, err := http.Get(fmt.Sprintf("%s/status", *api)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else if res.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "%d\n", res.StatusCode)
		os.Exit(1)
	} else {
		fmt.Println("ok")
	}

	config.callConsumers(*api)
}
