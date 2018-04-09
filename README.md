# Atelier Initiation au langage Go

Développeur·se C ou Java ? Initiez-vous au Go avec nous, dans un atelier qui vous fera écrire votre premier service HTTP.

Nous ferons un survol des spécificités du langage, en insistant particulièrement sur les différences avec Java. Puis nous accompagnerons les participant·e·s dans l'implémentation d'un service ReST simple.

Ce repository contient les instructions pour les exercices. Le plus simple est de suivre ces instructions en ligne, mais vous pouvez aussi cloner le repository.

La première étape à suivre est l'[installation du compilateur Go](0_installation).

Ensuite, parcourez chaque chapitre dans l'ordre.

# Slides

Vous pourrez retrouver les slides présentés pendant l'atelier [ici](https://docs.google.com/presentation/d/1stST9AzzvQzZF2F2xrDz-GEIEEwWlqSrgBBLCvLOgks/edit?usp=sharing).

# Repository Git

La dernière version de ce repository est disponible sur https://github.com/leboncoin/atelier-initiation-golang.

# Comment créer une clef USB avec le contenu nécessaire pour faire l'atelier sans connexion internet

1. créer un répertoire `atelier-go` sur la clef : `cd /media/{user}/{nom-clef-usb} && mkdir atelier-go`
2. se placer dans le répertoire en question : `cd /media/{user}/{nom-clef-usb}/atelier-go/`
3. y clone le repo : `git clone https://github.com/leboncoin/atelier-initiation-golang.git`
4. créer un répertoire pour les binaires d'installation de Go : `mkdir /media/{user}/{nom-clef-usb}/atelier-go/install`
5. y partir de https://golang.org/dl/, y télécharger les binaires `darwin-amd64.pkg`, `linux-386.tar.gz`, `linux-amd64.tar.gz`, `windows-386.msi`, `windows-amd64.msi`.

Le total fait environ 550MB.
