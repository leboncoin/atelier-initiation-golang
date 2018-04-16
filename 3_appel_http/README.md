# Appel HTTP

Ecrivons un programme qui appelle le service codé lors de l'[étape précédente](../2_serveur_http).

Dans le package [net/http](https://golang.org/pkg/net/http/), vous trouverez les fonctions associées aux méthodes HTTP (GET, POST, DELETE, etc.).

## Simple appel

Codez un appel à `http://localhost:8080/test/blah`.

La fonction à utiliser est `func Get(url string) (*http.Response, error)`. Vous pouvez ignorer la réponse pour le moment (pour ignorer les valeurs de retour, on peut utiliser le placeholder `_` pour chaque valeur de retour).

A noter que, pour être exhaustif, il est utile en production de vérifier que l'erreur retournée n'est pas `nil`.

## Lecture de la réponse

Changez l'appel pour contacter `http://localhost:8080/ping`.

La réponse est la première valeur de retour de la function `Get`.

Le body est une variable fournie par l'objet Response qui implémente l'interface `io.Reader`. Comme pour les `InputStream` en Java, il faut le lire pour accéder à son contenu.

Pour cela, la fonction [ioutil.ReadAll](https://golang.org/pkg/io/ioutil/#ReadAll) sera utile.

## Envoi d'un body dans la requête

Nous allons faire un appel POST à `http://localhost:8080/bye`, avec un body.

Il existe une fonction `Post(url string, contentType string, body io.Reader) (*Response, error)`. Elle prend un `io.Reader` en paramètre qu'il va falloir peupler avec une string.

La solution pour cela est `func NewReader(s string) *bytesReader`, dans le package `strings`.

## Exercice bonus : http.NewRequest

Chez leboncoin, à la place des fonctions `Get`, `Post`, etc., nous utilisons plutôt `func NewRequest(method, url string, body io.Reader) (*Request, error)`.

C'est en partie parce qu'il n'y a pas de fonctions pour toutes les méthodes HTTP (il manque DELETE, OPTIONS, PUT...) et en partie parce que cela nous permet de factoriser plus de code.

Nous vous invitons à réécrire le code précédent en remplaçant `Get` et `Post` par `http.NewRequest`.
