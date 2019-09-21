package main

import(
    "net/http"
    "encoding/json"
    "flag"
    "log"
    //"os"
    "io/ioutil"
    //"fmt"

    "servicebinarytree/pkg/storage/inmemory"
    sampledata "servicebinarytree/pkg/models/sample-data"
    "servicebinarytree/pkg/models"
    "servicebinarytree/pkg/http/rest"
)

type Configuration struct {
    Server struct {
        Address string
    }
}

var conf Configuration


// Initialize the server values from config file

func init() {

    file, _ := ioutil.ReadFile("conf/config.json")
 
    err := json.Unmarshal([]byte(file), &conf)

    if err != nil {
        log.Fatal("could not Unmarshal config file: ", err)
    }
    
    if conf.Server.Address == ""{
        log.Fatal("could not find the Address in config")   
    }

}

func main() {
    
    withData := flag.Bool("withData", false, "start with daple data")
    flag.Parse()

    var binarytrees map[string]*models.BinaryTree
    if *withData {
        binarytrees = sampledata.Binarytrees
    }


    repo := inmemory.InitBinarytreeRepository(binarytrees)

    s := rest.New(repo)

    log.Printf("Serving at %s",conf.Server.Address)
    log.Fatal(http.ListenAndServe(conf.Server.Address, s.Handler()))

}