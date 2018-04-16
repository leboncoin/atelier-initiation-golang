# Goroutines et channels

Nous allons voir comment les goroutines permettent de paralléliser les exécutions et comment les channels permettent d'échanger des données.

L'exercice consiste à créer le client d'un service HTTP qui fournit le nombre de vues sur une annonce sur un site en ligne.
Ce service est assez lent, ce qui ne pose pas de problème pour une seule annonce.

Lorsque nous voudrons obtenir les informations pour des milliers d'annonces dans le but d'obtenir la somme totale des vues, nous allons rencontrer des problèmes de performance. Il faudra donc paralléliser les appels.

## Préparation

Nous fournissons un faux serveur pour le nombre de vues. Pour un même identifiant d'annonce, il retournera toujours le même nombre.

Dans le répertoire `views-server`, vous trouverez l'implémentation de ce serveur. Il suffit de le compiler et de l'exécuter.

```
$ go build
$ ./views-server
this returns the number of views for a given ad id
try:
  curl http://localhost:9091/views/2aae6c35c94fcfb415dbe95f408b9ce91ee846ed

2018/04/05 22:01:32 serving on :9091
```

L'exécution fournit un exemple d'appel. Il s'agit d'un simple GET avec un identifiant d'annonce spécifié à la fin de l'URL (n'importe quelle string marchera très bien).

Vous pouvez laissez le serveur tourner.

## Premier appel

Codez un appel GET à ce serveur. Une fonction simple pour ce faire est `func Get(url string) (*http.Response, error)`, dans le package [net/http](https://golang.org/pkg/net/http/). La réponse est la première valeur de retour de la fonction `Get`.

Le body est une variable fournie par l'objet `http.Response` qui implémente l'interface `io.Reader`. Comme pour les `InputStream` en Java, il faut le lire pour accéder à son contenu. Pour cela, la fonction [ioutil.ReadAll](https://golang.org/pkg/io/ioutil/#ReadAll) sera utile.

Enfin, convertissez la string retournée dans le body en un entier avec [strconv.Atoi](https://golang.org/pkg/strconv/#Atoi).

## Appels en série

Ajoutez une boucle `for` pour faire 10 appels les uns après les autres et faire la somme des résulats.

Pour créer les ids des annonces, vous pouvez simplement les écrire manuellement dans un tableau et les lire dans la boucle, ou les générer à partir de la valeur courante de l'itération.
Nous recommandons en tout cas de prendre des valeurs prévisibles, non aléatoires, afin d'obtenir les mêmes résultats à chaque exécution. Cela vous aidera à vérifier que le code se comporte comme prévu.

Si tout se passe comme prévu, l'exécution de ce code devrait prendre de longues secondes :-)

## Appels en parallèle

Dans votre boucle `for`, englobez l'appel au serveur HTTP dans une goroutine. Il faudra également passer en paramètre l'identifiant afin d'éviter que la variable soit partagée entre les goroutines.

Cela devrait ressembler à cela :

```
	go func(idInGoroutine string) {
		// le code
	}(idFromForLoop)
```

Un problème que vous allez rencontrer à l'exécution est que la fonction `main` va s'achever plus vite que les goroutines n'aient le temps de s'exécuter complètement. Comme pour la JVM, le programme va donc sortir sans que tout soit fini.

Nous verrons une solution élégante à l'étape suivante. Pour le moment, nous allons simplement attendre une durée arbitraire :

```
	time.Sleep(5 * time.Second)
```

L'exécution de ce code devrait être bien plus rapide (pas beaucoup plus que 5 secondes, donc). Tentez également d'augmenter le nombre d'appels simultanés à 100 ou 1000.

## Aggrégation des résultats

Créez un channel de type `int`, puis dans la goroutine, insérez le nombre de vues dans le channel.

L'insertion ressemble à cela :
```
	monChannel <- maVariable
```

A la suite de la boucle `for` qui lance les goroutines, ajoutez une nouvelle boucle `for` avec le même nombre d'itérations, qui lit les données dans le channel et en fait la somme.

La lecture ressemble à cela :
```
	maVariable := <-monChannel
```

La durée d'exécution du code devrait maintenant être la même que celle de la plus lente des requêtes HTTP, soit environ 2 secondes si vous avez assez d'annonces.
Au delà des performances, vous avez surtout gagné la garantie d'obtenir toutes les valeurs.

## Exercice bonus : Gestion d'un timeout

Nous allons gérer un timeout, au cas où certaines requêtes prendraient trop de temps. Pour cela, nous allons utiliser la clause `select` (qui écoute sur plusieurs channels) et la fonction `time.After` (qui envoie une valeur sur un channel si une durée s'est écoulée).

Dans la deuxième clause `for` de l'exercice précédent, ajoutez un `select` qui va écouter d'une part sur le channel des nombres de vue et d'autre part sur un `timeoutChan`. Faites échouer le code si le channel `timeoutChan` fournit une valeur.

### Quelques explications

La clause `select` écoute sur plusieurs channels simultanément et rend la main dès qu'un des channels fournit une valeur. Nous l'utilisons souvent à l'intérieur d'une boucle `for` infinie qui processe des valeurs jusqu'à ce qu'une erreur arrive, par exemple.

```
for {
	select {
	case msg := <-messageChan:
		log.Infof("message received: %v", msg)
	case err := <-errorChan:
		log.Infof("something went wrong :-(: %v", err)
	}
}
```

La fonction `time.After` est très pratique pour gérer les timeouts. Elle se combine souvent avec le `select` pour exprimer que l'on veut attendre une réponse, mais pas plus qu'une certaine durée.

Nous l'utilisons beaucoup dans nos tests.
```
select {
case m := <-partitionConsumer.Messages():
	assert.Equal(t, "john", m.Name)
	break
case <-time.After(1 * time.Second):
	assert.Fail(t, "did not get the message after waiting 1 sec")
}
```

Nous l'avons utilisée aussi pour gérer des timeouts sur les appels HTTP dans le code de production. Néanmoins, les `context` en Go semblent un meilleur paradigme pour cela et nous les utilisons de plus en plus souvent pour ce besoin.
