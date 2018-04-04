# Serveur HTTP


Dans un répertoire `$GOPATH/src/service`, créez un fichier `service.go`:


Pour ecrire ce microservice vous aurez besoin du package `net/http` de la bibliotheque standard du langage.

Ce package met a disposition deux fonctions: 

#### HandleFunc(pattern string, handler func(ResponseWriter, *Request))

Enregistre le **handler** passé en paramètre sur un pattern donné.

La variable **handler** est un pointeur sur fonction ayant la signature suivante: `func(ResponseWriter, *Request)`

Le handler est donc une fonction qui prend deux paramètres:

   - Un `ResponseWriter` qui permet de construire une réponse HTTP
   - Un pointeur sur une `Request` qui représente la requête HTTP reçue

#### ListenAndServe(addr string, handler Handler)

Démarre un serveur HTTP qui écoute sur une interface et un port passés en paramètre

### Partie 1

Le microservice devra commencer par exposer une première route `/ping` sur laquelle il renverra simplement **pong**

```
$ curl localhost:8080/ping
pong
```

### Partie 2

Le microservice exposera ensuite une route `/hello/<nom>` sur laquelle il renverra **Hello, \<nom>!**

```
$ curl localhost:8080/hello/world
Hello, world!
```

### Partie 3

Le microservice exposera finalement une route `/bye` sur laquelle il renverra **Bye, \<nom>!**

Le \<nom> viendra du body de la requete HTTP

```
$ curl -X POST -d 'Chris' localhost:8080/bye
Bye, Chris!
```
