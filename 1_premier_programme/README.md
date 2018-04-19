# Premier programme Go

## Préparation

Selon les conventions Go, les projets se placent dans `$GOPATH/src/nomduprojet`.

Dans un répertoire `$GOPATH/src/hello`, créez un fichier `helloworld.go` avec le contenu suivant :

```
package main

func main() {
}
```

Pour l'exécuter :

```
$ cd $GOPATH/src/hello
$ go build
$ ./hello
```

Le programme ne retourne rien pour le moment.

### Quelques explications

`go build` compile le package du répertoire courant. Le binaire résultant prend le nom du répertoire (et non pas le nom du fichier `.go`, ni le nom du package). Il y a un seul binaire par package et non un binaire par fichier source.

`main` est un nom de package spécial. Il indique à Go que le binaire doit être exécutable (et non une librairie qui serait plus tard importée dans un autre package). On ne peut pas avoir de package sans nom.

## Affichage d'`Hello, World`

Vous allez avoir besoin d'importer le package `fmt`.

```
import (
	"fmt"
)
```

`fmt` contient plusieurs fonctions utiles pour formatter des données et les afficher en ligne de commande.

La fonction la plus simple à utiliser est `fmt.Println`. Utilisez plutôt `fmt.Printf` qui est plus puissante.

Implémentez l'affichage de `Hello, World`. Vous devriez obtenir ce qui suit :
```
$ go build
$ ./hello
Hello, World
```

Il est aussi possible de demander à Go d'exécuter le code sans créer de binaire :
```
$ go run helloworld.go
Hello, World
```

Maintenant que vous avez écrit plusieurs lignes, il est probable que le code ne soit pas formatté selon les standards Go. `go fmt` va vous aider ici.

```
$ go fmt
helloworld.go
```

`go fmt` affiche la liste des fichiers reformattés, ou rien si aucun fichier n'a été modifié.

En général, les IDE s'intégrent facilement avec `go fmt`. Le respect des conventions de formattage est très répandu dans la communauté Go.

### Quelques explications

Le package `fmt` fournit des fonctions d'intéraction avec la console. C'est un peu l'équivalent de `java.lang.System`, mais il faut l'importer explicitement, alors que `java.lang` est importé implicitement en Java.
En Go, tous les packages utilisés doivent être explicitement importés.

`fmt.Printf` est une fonction très utilisée qui permet d'afficher des données, en utilisant le système de formattage classique du C. Ne pas oublier le `\n` à la fin.

En Go, il est plus fréquent qu'en Java d'utiliser le formattage pour afficher des données, en partie parce qu'il n'y a pas de méthode par défaut comme `.toString()`.
Il existe d'ailleurs de nombreuses fonctions qui possèdent une variante de même nom, suffixé par un `f` : `fmt.Scan`/`fmt.Scanf`, `fmt.Fprint`/`fmt.Fprintf`, `assert.Error`/`assert.Errorf` (package `testify/assert`), etc.

Alors qu'en Java, on va souvent faire des choses comme cela :
```
System.out.println("user: " + user);
```
En Go, la convention, très répandue, est d'utiliser ceci :
```
fmt.Printf("user: %v\n", user)
```

`%v` utilise un format par défaut. Voir [la doc du package fmt](https://golang.org/pkg/fmt/) pour les détails.

## Les tests

Nous allons modifier le code pour saluer une personne (affichage de "Hello, Jane"), dont le prénom sera passé en paramètre. Commençons par écrire un test.

Dans le même répertoire, créez un fichier `helloworld_test.go` avec le contenu suivant :

```
package main

import (
	"testing"
)

func Test_1_plus_1_should_be_2(t *testing.T) {
	result := 1 + 2 // code under test; will fail at first
	if result != 2 {
		t.Errorf("expected '%v', got '%v'", 2, result)
	}
}
```

Lancez les tests :

```
$ go test
--- FAIL: Test_1_plus_1_should_be_2 (0.00s)
	helloworld_test.go:10: expected '2', got '3'
FAIL
exit status 1
FAIL	hello	0.001s
```

Corrigez le code sous test. Vous devriez le voir passer :
```
$ go test
PASS
ok  	hello	0.001s
```

Modifiez maintenant le test pour qu'il vérifie qu'une fonction `greet` dans le même package qui prend un prénom en paramètre. Cette fonction retourne un texte commençant par `Hello, ` et se terminant par le prénom.

Dans un premier temps, écrivez uniquement le test, sans le code correspondant.

Executer `go test` échouera à la compilation, puisque la fonction n'existe pas.

```
$ go test
# hello
./helloworld_test.go:8:12: undefined: greet
FAIL	hello [build failed]
```

Créez cette fonction (retournant tout d'abord une chaîne vide), puis relancez le test.

```
$ go test
--- FAIL: Test_greet_should_return_the_name_prefixed_with_a_greeting (0.00s)
	helloworld_test.go:10: expected 'Hello, Jane', got ''
FAIL
exit status 1
FAIL	hello	0.001s
```

Enfin, implémentez le corps de la fonction. Vous devriez obtenir un message de succès de la part de `go test` :

```
$ go test
PASS
ok  	hello	0.001s
```

### Quelques explications

La convention en Go est de placer les tests unitaires dans le même package et dans le même répertoire que le code sous test.

Les fichiers contenant les tests *doivent* avoir un nom qui termine par `_test.go`. Les fonctions qui s'y trouvent *doivent* avoir un nom qui commence par `Test`.

`go test` est fourni avec l'installation standard de Go, ce qui en fait le standard pour lancer les tests. Néanmoins, il y a peu de support pour industrialiser les tests et les rendre plus lisibles.

Il existe donc des packages open-source pour écrire des assertions, mocks, etc. Le plus populaire semble être [testify](https://github.com/stretchr/testify) que nous utilisons beaucoup chez leboncoin.

### Exercice bonus : Les tests avec Testify

Réécrivez le test en utilisant `testify/assert`.

Au préalable, récupérez le package avec la commande

```
go get github.com/stretchr/testify
```

Le test ressemblera alors à ceci :

```
	assert.Equal(t, "Hello, Jane", greet("Jane"))
```
