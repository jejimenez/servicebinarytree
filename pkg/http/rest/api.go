package rest

import (
    "net/http"
    "encoding/json"
    "fmt"
    "strconv"

    servicebinarytree "servicebinarytree/pkg"

    "github.com/gorilla/mux"
    "servicebinarytree/pkg/models"
    btservice "servicebinarytree/pkg/service"
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
    r.HandleFunc("/lowestancestor/{treename:[a-zA-Z0-9_]+}/{value1:[0-9_]+}/{value2:[0-9_]+}", a.LowestAncestor).Methods(http.MethodGet)

    a.handler = r
    return a
}


// Create binary tree hanlder 
func (a *api) CreateBinarytree(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")
    decoder := json.NewDecoder(r.Body)

    var g *models.BinaryTree
    
    err := decoder.Decode(&g)
    //fmt.Println(g)
    //fmt.Println(r.Body)

    if err != nil || g == nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode("Error unmarshalling request body")
        return
    }


    if err = a.repositorybt.CheckIfExists(g.Name); err != nil {
        w.WriteHeader(http.StatusPreconditionFailed)         
        json.NewEncoder(w).Encode(fmt.Sprintf("Binarytree %s already exists.", g.Name))
        return 
    }

    err = a.repositorybt.CreateBinarytree(g)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)         
        json.NewEncoder(w).Encode("Error when trying to store the binary tree.")
        return
    }

    //json.NewEncoder(w).Encode(err)
    w.WriteHeader(http.StatusCreated)
}

// Look for the lowest ancestor of given binarytree and two element nodes
func (a *api) LowestAncestor(w http.ResponseWriter, r *http.Request){
    
    w.Header().Set("Content-Type", "application/json")

    //fmt.Println(r)
    params := mux.Vars(r)
    //params := r.URL.Query()

    if len(params) < 1 {
        fmt.Println("Url Param 'key' is missing")
        return
    }
    if _, ok := params["treename"]; !ok {
        w.WriteHeader(http.StatusUnprocessableEntity)         
        json.NewEncoder(w).Encode("Param treename is mandatory.")
        return
    }
    if _, ok := params["value1"]; !ok {
        w.WriteHeader(http.StatusUnprocessableEntity)         
        json.NewEncoder(w).Encode("Param value1 is mandatory.")
        return
    }
    if _, ok := params["value2"]; !ok {
        w.WriteHeader(http.StatusUnprocessableEntity)         
        json.NewEncoder(w).Encode("Param value2 is mandatory.")
        return
    }
    if _, err := strconv.Atoi(params["value1"]); err != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)         
        json.NewEncoder(w).Encode("Value1 is not a number.")
        return
    }
    if _, err := strconv.Atoi(params["value2"]); err != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)         
        json.NewEncoder(w).Encode("Value2 is not a number.")
        return
    }
    if params["value1"] == params["value2"] {
        w.WriteHeader(http.StatusUnprocessableEntity)         
        json.NewEncoder(w).Encode("Value1 is equal to value2.")
        return
    }

    nm  := params["treename"]
    v1,_:= strconv.Atoi(params["value1"])
    v2,_:= strconv.Atoi(params["value2"])

    bt,_  := a.repositorybt.GetBinarytreeByName(nm)
    ca, err := btservice.GetLowestCommonAncestor(bt, v1, v2)

    if err != nil {
        w.WriteHeader(http.StatusNotFound)         
        json.NewEncoder(w).Encode("Not found.")
        return
    }
    

    //s, _ := json.MarshalIndent(bt, "", "\t")
    //fmt.Println(string(s))

    //fmt.Println(ca)
    //fmt.Println(err)
    lar := models.LowestAncestorResp{Treename:nm,Value1:v1,Value2:v2,Ancestor:ca}


    w.WriteHeader(http.StatusFound)
    json.NewEncoder(w).Encode(lar)


    //fmt.Println(val)
    //if params.treenam 

    //fmt.Println(params)
    //fmt.Println(params.treen)

    //StatusUnprocessableEntity
}

