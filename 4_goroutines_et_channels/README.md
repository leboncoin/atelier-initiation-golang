# Goroutines et channels

Nous allons voir comment les goroutines permettent de paralléliser les exécutions et comment les channels permettent d'échanger des données.

L'exercice consiste à la création d'un client qui appelle un service HTTP fournissant le nombre de vues sur une annonce sur un site en ligne.
Ce service est assez lent, ce qui ne pose pas de problème pour une seule annonce.

Lorsque nous voudrons obtenir les informations pour des milliers d'annonces dans le but d'obtenir la somme totale des vues, il faudra paralléliser les appels et faire la somme des vues de chaque annonce.

## Préparation

Nous fournissons un faux serveur pour le nombre de vues. Pour un même identifiant d'annonce, il retournera toujours le même nombre de vues.

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

Vous pouvez partir de [la solution de l'exercice précédent](../3_appel_http/solutions) pour faire un premier appel à ce serveur.
Convertissez la string retournée dans le body en un entier (cf. https://golang.org/pkg/strconv/#Atoi).

## Appels en série

Ajoutez une boucle `for` pour faire 10 appels les uns après les autres et faire la somme des résulats.

Pour créer les ids des annonces, vous pouvez simplement les écrire manuellement dans un tableau et les lire dans la boucle, ou les générer à partir de la valeur courante de l'itération. Nous recommandons en tout cas de prendre des valeurs prévisibles, non aléatoires, afin d'obtenir les mêmes résultats à chaque exécution.

Si tout se passe comme prévu, l'exécution de ce code devrait prendre de longues secondes :-)

## Appels en parallèle

Dans votre boucle `for`, englobez l'appel au serveur HTTP dans une goroutine. Il faudra également passer en paramètre l'identifiant afin d'éviter que la variable soit partagée entre les goroutines.

Cela devrait ressembler à cela :

```
	go func(idInGoroutine string) {
		// le code
	}(idFromForLoop)
```

Un problème que vous allez rencontrer à l'exécution est que la fonction `main` va sortir plus vite que les goroutines n'aient le temps de s'achever.
Nous verrons une solution élégante au paragraphe suivant. Pour le moment, nous allons simplement attendre une durée arbitraire :

```
	time.Sleep(5 * time.Second)
```

L'exécution de ce code devrait être bien plus rapide. Tentez également d'augmenter le nombre d'appels simultanés à 100 ou 1000.

## Aggrégation des résultats

Créez un channel de type `int`, puis dans la goroutine, insérez le nombre de vues dans le channel.

L'insertion ressemble à cela :
```
	monChannel <- maVariable
```

A la suite de la boucle `for` qui lance les goroutines, ajoutez une boucle `for` avec le même nombre d'itérations, qui lit les données dans le channel et fait la somme.

La lecture ressemble à cela :
```
	maVariable := <-monChannel
```

La durée d'exécution du code devrait maintenant être la même que celle de la plus lente des requêtes HTTP, soit environ 2 secondes si vous avez un grand nombre d'annonces.
Mais surtout, vous avez la garantie d'obtenir toutes les valeurs.

## Gestion d'un timeout

En exercice bonus, nous allons gérer un timeout, au cas où certaines requêtes prendraient trop de temps.

Pour cela, nous allons utiliser la clause `select` qui permet d'écouter sur plusieurs channels simulatément.

Egalement, nous allons tirer parti de la fonction `time.After` qui crée un channel qui recevra une valeur lorsqu'une certaine durée s'est écoulée.
```
	timeoutChan := time.After(3 * time.Second)
```

Dans la deuxième clause `for` du paragraphe précédent, ajoutez un `select` qui va écouter d'une part sur le channel des nombres de vue et d'autre part sur le `timeoutChan` (et échouer si ce channel fournit une valeur).
