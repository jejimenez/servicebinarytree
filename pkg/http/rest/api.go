package rest

import (
    "net/http"
    "encoding/json"
    "fmt"

    servicebinarytree "servicebinarytree/pkg"

    "github.com/gorilla/mux"
    "servicebinarytree/pkg/models"
)

func (a *api) Handler() http.Handler {
    return a.handler
}

type api struct {
    handler     http.Handler
    repositorybt servicebinarytree.BinarytreeRepository
}

type Rest interface {
    Handler() http.Handler
    CreateBinarytree(w http.ResponseWriter, r *http.Request)
    LowestAncestor(w http.ResponseWriter, r *http.Request)
}

// Config the end points 
func New(repobt servicebinarytree.BinarytreeRepository) Rest {
    a := &api{repositorybt: repobt}

    r := mux.NewRouter()
    r.HandleFunc("/tree", a.CreateBinarytree).Methods(http.MethodPost)
    r.HandleFunc("/lowestancestor", a.LowestAncestor).Methods(http.MethodGet)

    a.handler = r
    return a
}


// Create binary tree hanlder 
func (a *api) CreateBinarytree(w http.ResponseWriter, r *http.Request) {

    decoder := json.NewDecoder(r.Body)

    var g *models.BinaryTree
    w.Header().Set("Content-Type", "application/json")
    
    err := decoder.Decode(&g)
    //fmt.Println(g)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode("Error unmarshalling request body")
        return
    }

    if err = a.repositorybt.CheckIfExists(g.Name); err != nil {
        w.WriteHeader(http.StatusPreconditionFailed) // We use not found for simplicity
        json.NewEncoder(w).Encode(fmt.Sprintf("Binarytree %s already exists", g.Name))
        return 
    }

    err = a.repositorybt.CreateBinarytree(g)

    if err != nil {
        w.WriteHeader(http.StatusPreconditionFailed) // We use not found for simplicity
        json.NewEncoder(w).Encode("Error when trying to store the binary tree. ")
        return
    }

    //json.NewEncoder(w).Encode(err)
    w.WriteHeader(http.StatusCreated)

}

// Look for the lowest ancestor of given binarytree and two element nodes
func (a *api) LowestAncestor(w http.ResponseWriter, r *http.Request){
    
}

