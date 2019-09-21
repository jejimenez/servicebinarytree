# Prueba Técnica Masivian
##Prueba Técnica Desarrollador Backend

API REST que permite crear un arbol binario y buscar el ancestro común de dos nodos.
La solución se desarrolló en Go pensada como un microservicio que responde a requerimientos de alta concurrencia.

El api rest tiene 2 endpoints:

- /tree 
- /lowestancestor

###/tree
Este end point recibe la información POST del arbol que se va a crear. Se debe nombrar para identificarlo de los demás arboles. El json queda de la siguiente forma:

```json
{
         "name":"treename",
         "root":{
            "value":67,
            "left":{
               "value":39,
               "left":{
                  "value":28,
                  "left":null,
                  "right":{
                     "value":29,
                     "left":null,
                     "right":null
                  }
               },
               "right":{
                  "value":44,
                  "left":null,
                  "right":null
               }
            },
            "right":{
               "value":76,
               "left":{
                  "value":74,
                  "left":null,
                  "right":null
               },
               "right":{
                  "value":85,
                  "left":{
                     "value":83,
                     "left":null,
                     "right":null
                  },
                  "right":{
                     "value":87,
                     "left":null,
                     "right":null
                  }
               }
            }
         }
      }

```
###/lowestancestor
####/lowestancestor/{treename:[a-zA-Z0-9_]+}/{value1:[0-9_]+}/{value2:[0-9_]+}
Éste endpoint recibe el nombre del arbol, el valor 1 y el valor 2 y devuelve el valor del ancestro común más cercano. Ej:

http://localhost:8080/lowestancestor/treename/29/44

Devuelve :
```json
{
    "treename": "treename",
    "value1": 29,
    "value2": 44,
    "ancestor": 39
}

```
##Ejecutar el Servidor
Para desplegar y ejecutar el servidor se puede a través de Docker:

    git clone https://github.com/jejimenez/servicebinarytree.git
    cd servicebinarytree/
    make docker
    make run

Si se tiene instalado go se puede desplegar con los siguientes comandos

    git clone https://github.com/jejimenez/servicebinarytree.git
    cd servicebinarytree/
    go build cmd/binarytree/main.go 
    ./main

Las pruebas se pueden ejecutar de la siguiente manera después de haber clonado el repositorio y en la carpeta servicebinarytree

    make test
    make unittest
