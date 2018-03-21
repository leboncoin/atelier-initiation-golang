# Premier programme Go

Selon les conventions Go, les projets se placent dans `$GOPATH/src/nomduprojet`.

Dans un répertoire `$GOPATH/src/hello`, créez un fichier `helloworld.go` avec le contenu suivant :

```
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```


```
$ cd $GOPATH/src/hello
$ go build
$ ./hello
hello, world
```

## Quelques explications

`go build` compile le package du répertoire courant. Le binaire résultant prend le nom du répertoire (et non le nom du fichier `.go`, ni le nom du package). Il y a un seul binaire par package et non un binaire par fichier source.

`main` est un nom de package spécial. Il indique à Go que le binaire doit être exécutable (pas une simple librairie). On ne peut pas avoir de package sans nom.

Le package `fmt` fournit des fonctions d'intéraction avec la console. C'est un peu l'équivalent de `java.lang.System`, mais il faut l'importer explicitement.

`fmt.Printf` est une fonction très utilisée qui permet d'afficher quelque chose, avec le système de formattage classique du C. Noter le `\n` à la fin.

En Go, il est plus fréquent qu'en Java d'utiliser le formattage pour l'affichage, en partie parce qu'il n'y a pas de méthode par défaut `.toString()`.

Alors qu'en Java, on va souvent faire des choses comme cela :
```
System.out.println("user: " + user);
```
En go, vous verrez souvent ceci :
```
fmt.Println("user: %v\n", user)
```

`%v` utilise un format par défaut. Voir [la doc du package fmt](https://golang.org/pkg/fmt/) pour les détails.
