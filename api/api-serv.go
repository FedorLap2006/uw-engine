package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	wbf "github.com/FedorLap2006/WebFrame/Go"
	"github.com/gorilla/mux"
)

func readConf(path string, v *interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	file, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	defer f.Close()
	json.Unmarshal(file, &v)

	return nil
}

func main() {
	conf_file := ""
	if len(os.Args) == 2 {
		conf_file = os.Args[2]
	} else {
		log.Fatalln("unknown arguments: Usage - api-serv <ConfFile> or api-serv <Port>")
	}

	r := mux.NewRouter()
	r.HandleFunc("/h", wbf.HandleHTTP(func(c *wbf.Context) {
		c.WriteStringIO("starter kit of server...")
	}))
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
