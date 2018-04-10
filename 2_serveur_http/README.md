# Serveur HTTP


Dans un répertoire `$GOPATH/src/service`, créez un fichier `service.go`.

Pour écrire ce microservice, vous aurez besoin du package `net/http` de la bibliothèque standard du langage (cf. https://golang.org/pkg/net/http/).

Ce package met a disposition deux fonctions importantes, `HandleFunc` et `ListenAndServe`.

#### HandleFunc

```HandleFunc(pattern string, handler func(ResponseWriter, *Request))```

Cette fonction enregistre le `handler` passé en paramètre sur un pattern de path donné.

Le `handler` est un pointeur sur une fonction ayant la signature suivante : `func(ResponseWriter, *Request)`.

Dans cette fonction, il y a deux paramètres :

   - Un `ResponseWriter` qui permet de construire une réponse HTTP ; il contient une fonction pour écrire un tableau de bytes (pour convertir une `string` en tableau de bytes, utilisez `[]byte(myString)`)
   - Un pointeur sur une `Request` qui représente la requête HTTP reçue, y compris le path de l'appel et le body

#### ListenAndServe

```ListenAndServe(addr string, handler Handler)```

Cette fonction démarre un serveur HTTP qui écoute sur une interface et un port passés en paramètre. Le paramètre `addr` sera de la forme `:8080`.

Attention : cette fonction ne rend pas la main ; l'exécution est donc bloquée au niveau de cet appel.

### Affichage du path

Le microservice devra commencer par exposer une route `/test/` et afficher dans la console le path de la requête appelante.

```
$ curl localhost:8080/test/blah
```

Pas de réponse visible du côté de la console appelante, mais la console du serveur devrait afficher ceci :
```
/test/blah
```


### Pong

Le microservice exposera une route `/ping` sur laquelle il renverra simplement le texte `pong`.

```
$ curl localhost:8080/ping
pong
```

### Appel avec paramètre de path

Le microservice exposera ensuite une route `/hello/<nom>` sur laquelle il renverra `Hello, \<nom>!`.

```
$ curl localhost:8080/hello/world
Hello, world!
```

### Appel avec body

Le microservice exposera finalement une route `/bye` sur laquelle il renverra `Bye, \<nom>!`.

Le \<nom> viendra du body de la requete HTTP

```
$ curl -X POST -d 'Chris' localhost:8080/bye
Bye, Chris!
```
