# Premier programme Go

Selon les conventions Go, les projets se placent dans `$GOPATH/src/nomduprojet`.

Dans votre répertoire, créez un fichier `helloworld.go` avec le contenu suivant :

```
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```


```
$ cd $GOPATH/src/hello
$ touch helloworld.go
```

## Quelques explications

`main` est un nom de package spécial. Il indique à Go que le binaire doit être exécutable (pas une simple librairie).

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
