# Installer Go

Voir les instructions pour la dernière version sur https://golang.org/doc/install (sous macOS, `brew install go` est la solution la plus simple).

Si vous possédez déjà Go sur votre machine, inutile de le mettre à jour pour cet atelier.

Notez que vous devriez avoir une variable PATH à jour, de façon à pouvoir lancer la commande `go` dans n'importe quel répertoire.

```
$ go version
go version go1.9 linux/amd64
```

Le répertoire racine de vos projets devrait être `$HOME/go`, chemin par défaut pour les projets Go. Si vous souhaitez un chemin particulier, positionnez une variable `GOPATH` (recommandé).

```
$ export GOPATH=$HOME/monchemin/go
```
