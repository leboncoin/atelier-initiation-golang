# Installer Go

Voir les instructions pour la dernière version : https://golang.org/doc/install

Si vous possédez déjà Go sur votre machine, inutile de le mettre à jour pour cet atelier.

Notez que vous devriez avoir une variable PATH à jour, de façon à pouvoir lancer la commande `go` dans n'importe quel répertoire.

```
$ go version
go version go1.9 linux/amd64
```

Le répertoire racine de vos projets devrait être `$HOME/go`. Dans le cas contraire, positionnez une variable `GOPATH` (recommandé).

```
$ export GOPATH=$HOME/devoxx/go
```
