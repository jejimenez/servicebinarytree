package main

import(
	"flag"
	//"fmt"
	"log"
	"net/http"

    "servicebinarytree/pkg/storage/inmemory"
    sampledata "servicebinarytree/pkg/models/sample-data"
    "servicebinarytree/pkg/models"
    "servicebinarytree/pkg/http/rest"
)

func main() {
    
    withData := flag.Bool("withData", false, "start with daple data")
    flag.Parse()

    var binarytrees map[string]*models.BinaryTree
    if *withData {
        binarytrees = sampledata.Binarytrees
    }


    repo := inmemory.InitBinarytreeRepository(binarytrees)

    s := rest.New(repo)

    log.Fatal(http.ListenAndServe(":8080", s.Handler()))

}