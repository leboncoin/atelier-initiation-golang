# Goroutines et channels

Nous allons voir comment les goroutines permettent de paralléliser les exécutions et comment les channels permettent d'échanger des données.

L'exercice consiste à la création d'un client qui appelle un service HTTP qui fournit le nombre de vues sur une annonce.
Ce service est assez lent, ce qui ne pose pas de problème pour une seule annonce.

Lorsque nous voudrons obtenir les informations pour des milliers d'annonces dans le but d'obtenir la somme totale des vues, il faudra paralléliser les appels et aggréger les nombres obtenus.

## Préparation

Nous fournissons un faux serveur pour le nombre de vues. Pour un même identifiant d'annonce, il retournera toujours le même nombre de vues.

Dans le répertoire `views-server`, vous trouverez son implémentation. Il suffit de le compiler et de l'exécuter.

L'exécution vous fournira un exemple d'appel. Il s'agit d'un simple GET avec un identifiant d'annonce à la fin de l'URL (n'importe quelle string marchera très bien).

Laissez le serveur tourner.

## Premier appel

Partez de la solution de l'exercice précédent pour faire un premier appel à ce serveur.
Pensez à convertir la string retournée dans le body en un entier (cf. https://golang.org/pkg/strconv/#Atoi).

## Appels multiples non parallélisés

Ajoutez une boucle `for` pour faire 10 appels les uns après les autres et faire la somme des résulats.

Pour créer les ids des annonces, vous pouvez simplement les écrire manuellement dans un tableau et les lire dans la boucle, ou les générer à partir de la valeur courante de l'itération. Nous recommandons en tout cas de prendre des valeurs prévisibles, non aléatoires, afin d'obtenir les mêmes résultats à chaque itération.

Si tout se passe comme prévu, l'exécution de ce code devrait prendre de longues secondes :-)

## Appels multiples en parallèle

Commentez temporairement le calcul de la somme des résultats.

Dans votre boucle `for`, englobez l'appel au serveur HTTP dans une goroutine. Cela devrait ressembler à cela :

```
	go func() {
		// le code
	}()
```

L'exécution de ce code devrait être bien plus rapide. Tentez également d'augmenter le nombre d'appels simultanés à 100 ou 1000.

## Aggrégation des résultats

Créez un channel de type `int`, puis dans la goroutine, insérez le nombre de vues dans le channel.

L'insertion ressemble à cela :
```
	monChannel <- maVariable
```

A la suite la la boucle `for` qui lance les goroutines, ajoutez une boucle `for` avec le même nombre d'itérations, qui lit les données dans le channel et fait la somme.

La lecture ressemble à cela :
```
	maVariable := <-monChannel
```