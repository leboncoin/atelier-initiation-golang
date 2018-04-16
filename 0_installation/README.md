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

## Quel IDE ?

Pour cet atelier, n"importe quel éditeur texte fera très bien l'affaire.

Au boncoin, nous aimons bien IntelliJ (version payante), avec son plugin Go (Golang est une version d'IntelliJ qui intégre directement le plugin). On y retrouve certaines fonctionnalités de la version Java d'IntelliJ, comme le refactoring, mais il y a encore beaucoup de choses à améliorer. C'est probablement l'IDE le plus avancé en ce moment.

Parmi les outils gratuits, Visual Studio Code, avec son plugin Go est vraiment de bonne qualité.

Si vous possédez déjà Sublime Text, vous pourrez également utiliser son plugin Go, néanmoins relativement limité.

La plupart des éditeurs populaires ont leur propre intégration avec Go.
