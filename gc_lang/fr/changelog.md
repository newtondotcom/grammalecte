# CHANGELOG (FR)

## Grammalecte 2.3 (15 décembre 2025)

- Nouvelles règles de contrôle.
- Suppression de faux positifs.
- Quelques bugs supprimés.
- [Thésaurus] Mise à jour (Merci à Algoo) + régularisation de nombreuses entrées.
- [Firefox/Thunderbird] Ajout du thésaurus de LibreOffice dans la WebExtension.
- [CLI] Thésaurus consultable depuis la CLI.
- [Conjugueur] Corrections concernant les participes passés.
- [Dictionnaire] Mise à jour (v7.7), et contrôle strict sur les entrées (de nombreuses erreurs d’étiquetage corrigées).
- [Graphspell] Amélioration du moteur de suggestions (plus rapide, plus précis).
- [Graphspell] Nouvelles suggestions ad hoc.
- [LibreOffice] Formateur de texte: peut dorénavant agir sur une sélection de texte (Merci à Jean-Marc Zambon).
- [LibreOffice] Ajustement de la couleur de l’UI pour être lisible pour ceux qui utilisent un thème sombre.
- [LibreOffice] Correction: normalisation des URL de l’import/export des fichiers.
- [LibreOffice] Bug concernant la fermeture et la relance de l’éditeur lexical.
- [Serveur] Mise à jour vers Bottle.py v0.13.4.
- [Build] Meilleur contrôle des entrées du dictionnaire Hunspell.
- [Build] Suppression de la dépendance à distutils (pour conformité avec les dernières versions de Python).


## Grammalecte 2.2 (mars 2023)

- Mise à jour pour satisfaire les prérequis de Firefox.


## décembre 2022

Perte de la base de données de Dicollecte, qui servait de forum à Grammalecte et de site de conception du dictionnaire pour Hunspell.

Cause: L’hébergeur de Grammalecte a été racheté par une grosse compagnie qui a mis fin aux services d’hébergement que nous utilisions.
Aucun des messages envoyés par cette compagnie ne m’étant parvenu, je n’ai pu procéder à la sauvegarde de la base de données avant l’extinction définitive des services d’hébergement.
La dernière sauvegarde de la base de données en ma possession était trop vieille pour relancer cette partie du site web.

Cette perte n’a pas impacté le code de Grammalecte directement, mais ont été perdues les discussions sur le forum ainsi que les dernières propositions d’ajout de nouveaux mots pour le dictionnaire Hunspell.


## Grammalecte 2.1.2 (08 mars 2021)

- [LibreOffice] Correction des bugs d’importation/exportation du dictionnaire personnel.
- Quelques nouvelles règles.

La mise à jour ne concerne que LibreOffice.


## Grammalecte 2.1.1 (22 février 2021)

- Quelques nouvelles règles.
- Éradication de nombreux faux positifs.
- [Graphspell] amélioration du moteur de suggestion orthographique.

Cette version est essentiellement une version de consolidation pour supprimer quelques faux positifs rédhibitoires sur la version précédente.
Afin de prévenir l’apparition de nouveaux faux positifs, Grammalecte effectue dorénavant environ 350 000 tests de contrôle.


## Grammalecte 2.1.0 (10 février 2021)

- Nombreuses nouvelles règles.
- Moins de faux positifs.
- Suggestions grammaticales : tout le code de suggestion a été revu, corrigé et testé, ce qui permis de supprimer bon nombre de micro-bugs, de filtrer dans certains cas nombre de suggestions absurdes, d’affiner les suggestions dans d’autres cas.
- Suggestions orthographiques : extension de l’échantillon de graphies proches repérées.
- [Writer] Formateur de texte: éditeur de transformations personnalisées.
- [Thunderbird] La signature ne sera dorénavant plus vérifiée et ne devrait donc plus être altérée.
- [Thunderbird] Réduction des altérations non souhaitables.


## Grammalecte 2.0 (05 décembre 2020)

### Nouveautés

- [Dictionnaire] Modification du format interne du dictionnaire de Grammalecte. À savoir : Grammalecte 2 peut lire les dictionnaires de Grammalecte 1, mais l’inverse n’est pas vrai.
- [Dictionnaire] Moteur de suggestion orthographique accéléré.
- [Writer et Python] Modification du format de fichier du dictionnaire (format binaire vers JSON).
- [Firefox et Thunderbird] Lexicographe : analyseur grammatical et lexical.
- Nouvelles règles de contrôle.
- Beaucoup d’anciennes règles ont été consolidées.
- Moins de faux positifs.
- Correction de nombreux bugs mineurs, comme des messages mal écrits ou mal configurés.
- Procédures de test consolidées.
- Nettoyage du vieux code.

### Les dictionnaires

Le format interne des dictionnaires a été modifié suite à des optimisations faites par un contributeur (IllusionPerdu) qui cherchait à améliorer la vitesse du dictionnaire.
Jusqu’alors le dictionnaire interne était constitué d’une très longue chaîne d’octets représentant les nodes et les arcs d’un graphe de mots, qu’on parcourait en calculant la valeur des nombres à partir des octets constituant les nodes et les adresses vers d’autres nodes. Par exemple, si l’adresse d’un node était représentée par 3 octets, on convertissait ces trois octets en un entier qui donnait la position du node dans la chaîne d’octets du dictionnaire. Autrement dit, c’était un format de données très compact, et c’était d’ailleurs le but recherché, en 2012, quand j’ai créé ce format.
Accélérer le dictionnaire consiste tout simplement à précalculer toutes les données du graphe à partir de la chaîne d’octets, et à stocker ça dans une liste d’entier. Bien sûr, ça va plus vite, mais ça consomme beaucoup plus de mémoire.
Le problème, c’est que si ce précalcul passait inaperçu sur Firefox et Thunderbird, il ralentissait énormément le démarrage de Writer.
C’est pourquoi j’ai décidé de modifier le format de données du dictionnaire et de stocker dorénavant les données dans une liste d’entiers, au lieu d’une chaîne d’octets.

Tout est transparent pour les utilisateurs. Ceux qui ont créé des dictionnaires personnels n’ont rien à faire, Grammalecte 2.0 fera la conversion automatiquement. Mais ceux qui exportent et importent leurs dictionnaires personnels doivent savoir que Grammalecte 2.x reconnaîtra les dictionnaires faits avec Grammalecte 1.x, mais Grammalecte 1.x ne pourra pas lire les dictionnaires conçus à partir de Grammalecte 2.x.

Autre détail interne, qui n’a aucune conséquence pour les utilisateurs, le format de fichiers des dictionnaires pour Python et Writer est désormais le JSON, en lieu et place d’un format binaire ad hoc. Ça simplifie la maintenance de n’avoir qu’un seul format de fichiers pour Python et JavaScript. Le désavantage, c’est que la taille des dictionnaires pour Python et Writer a quasiment triplé, passant en moyenne de 1,5 Mo à presque 4 Mo.

### Le lexicographe

En version bêta, voire alpha, uniquement sur Firefox et Thunderbird pour l’instant, le lexicographe propose dorénavant deux types d’analyse : une lexicale (c’est-à-dire en indiquant ce que sont les mots sans tenir compte du contexte), une grammaticale (en tenant compte du contexte pour dire de quelle nature sont les mots).
La seconde analyse est encore inachevée et très expérimentale, et elle le sera certainement encore pour un bon bout de temps, attendu que c’est plus complexe que ce à quoi je m’attendais et que ce que je prévoyais de faire a des incidences sur le correcteur grammatical lui-même. Pour aboutir, il faudra refondre nombre de règles de fonctionnement. Ça s’améliorera donc progressivement avec les versions suivantes.
Par la suite, cette nouvelle fonction devrait me permettre d’écrire des règles de contrôle plus précises.

### Règles de contrôle de Grammalecte

En un an et demi, beaucoup de nouvelles règles et d’améliorations ont été apportées. Quelques chiffres :

Grammalecte 2.0 (décembre 2020) :

- Analyse des paragraphes : 170 règles par expression régulière.
- Analyse des phrases : 10 134 règles par description de tokens réparties en 10 graphes différents, et 21 règles par expression régulière.

Grammalecte 1.0 (avril 2019) :

- Analyse des paragraphes : 165 règles par expression régulière.
- Analyse des phrases : 5 377 règles par description de tokens réparties en 9 graphes différents, et 38 règles par expression régulière.

D’une manière générale, Grammalecte 2 est la continuation directe de Grammalecte 1. Il n’y a pas de gros changements internes, seulement de nouvelles fonctionnalités dans les mécanismes d’analyse et d’innombrables micro-améliorations diverses.



## Grammalecte 1.12.0 (19 août 2020)

- Moins de faux positifs.
- Nouvelles règles de contrôle.
- [Thunderbird] correction du bug qui déstructurait certains mails en HTML.
- [Firefox] Les corrections grammaticales sont à présent répercutées dans les iframes aussi.

En ce qui concerne Twitter, j’ai désactivé la rétroaction sur la zone de texte, attendu qu’avec le nouveau fonctionnement de Twitter, écrire dans cette zone s’avère compliqué. Il faut donc copier la correction dans la zone soi-même.


## Grammalecte 1.11.0 (23 juillet 2020)

- Moins de faux positifs.
- Nouvelles règles de contrôle.
- [Thunderbird] MailExtension compatible avec Thunderbird 78+. *

La MailExtension pour Thunderbird est très semblable à la WebExtension pour Firefox et Chrome. Elle repose sur le même code et fonctionne grosso modo de la même manière, sauf en partie pour l’interface graphique, qui est plus limitante pour Thunderbird.

La grosse différence réside dans le fait que le panneau de composition de Thunderbird ne tolère visiblement pas l’insertion de nodes de type “input”, c’est-à-dire de widgets pour recevoir des données (champs de texte, cases à cocher, etc.). J’ignore encore si c’est un bug, un oubli ou une restriction délibérée, car je n’ai pas encore reçue de réponse de la part des développeurs.

Par ailleurs, tous les événements se déclenchant lorsque vous faites quelque chose, comme écrire du texte, sont apparemment inopérants. Conséquemment, Grammalecte ne peut pas savoir si vous écrivez ou non quelque part, et c’est pourquoi l’auto-rafraîchissement n’est pas possible en l’occurrence.

À part ça, tout est à peu près semblable à l’extension pour Firefox.


## Grammalecte 1.10.0 (31 mai 2020)

- Moins de faux positifs.
- Nouvelles règles de contrôle.


## Grammalecte 1.9.1 et 1.9.2 (04 mai 2020)

(Pour Firefox/Chrome et Thunderbird seulement.)

- Moins de faux positifs.
- Quelques nouvelles règles.
- [Firefox] Correction d’un bug pouvant faire planter certaines applis en JavaScript.


## Grammalecte 1.9.0 (21 avril 2020)

- Moins de faux positifs.
- Nouvelles règles de contrôle.
- [Réglages] Option chimie activée par défaut (sauf pour Writer).
- [LibreOffice] Moteur grammatical capable de reconnaître les diacritiques combinants.
- [Firefox] Le bouton tournoyant à présent au-dessus de la page au lieu d’être inclus à l’intérieur de celle-ci, ce qui devrait éviter de nombreux effets de bord (décalage ou disparation) constatés sur certains sites auparavant.
- [Firefox] API web : Ajustements et correction de bugs (notamment pour Chrome).
- [Graphspell] Meilleure suggestion orthographique pour la chimie. Exemple : C5H10N2O3 -> C₅H₁₀N₂O₃.
- [Build] bugs corrigés pour construire les extensions pour Firefox et Thunderbird.
- [Build] Multi-processus pour accélérer la construction des graphes.
- [Build] Construction des graphes : clarification du code et syntaxe des règles plus stricte.


## Grammalecte 1.8.1 (28 mars 2020)

- [Firefox] Correction du bug qui faisait disparaître certaines parties du texte en rencontrant certaines erreurs (rares, sauf si vous utilisiez l’option OCR).
- [Firefox] Suppression d’une fonctionnalité qui faisait planter Twitter en appuyant sur Backspace : ce n’est pas Grammalecte qui plantait, mais bien Twitter qui n’appréciait visiblement pas ce que Grammalecte pouvait faire. C’est le problème avec les zones de texte non standard, impossible de savoir ce qui se passe vraiment. Aucune erreur rapportée, juste un plantage impossible à analyser de notre côté.


## Grammalecte 1.8.0 (27 mars 2020)

- Moins de faux positifs.
- Nouvelles règles de contrôle.
- Bugs corrigés.
- [CLI] Mode interactif pour correction des fichiers.
- [LibreOffice] Boîtes de dialogue non modales pour le conjugueur et l’éditeur lexical.
- [LibreOffice] Meilleur compte des mots par le recenseur (mots élidés et mots composés).
- [Firefox] Apparence du surlignement des erreurs revue.
- [Firefox] API web pour les sites web.


## Grammalecte 1.7.0 (03 février 2020)

- Moins de faux positifs.
- Nouvelles règles de contrôle.
- Bugs corrigés.
- [Thunderbird] Sélection du dictionnaire orthographique.
- [LibreOffice] Recenseur de mots : export des données.
- [Dictionnaires] Refonte de nombreux drapeaux (le masculin est désormais considéré comme le lemme des mots de genre variable).


## Grammalecte 1.6.0 (27 novembre 2019)

- Moins de faux positifs.
- Nouvelles règles de contrôle.
- [Firefox] Choix du dictionnaire orthographique (Classique, Réforme ou Toutes variantes).
- [Dictionnaires] Nouveaux mots grammaticaux : iel et iels, à la fois comme pronom personnel sujet et objet. Ces pronoms ont été intégrés aux règles de contrôle grammatical.


## Grammalecte 1.5.0 (30 septembre 2019)

- Moins de faux positifs.
- Nouvelles règles de contrôle.
- [Graphspell] Meilleures suggestions.
- [LibreOffice] Sélecteur de fichiers pour l’éditeur lexical.
- [LibreOffice] Recenseur de mots : bouton pour naviguer dans le document d’une occurrence d’un mot à l’autre, fenêtre non modale.


## Grammalecte 1.4.0 (06 septembre 2019)

- Moins de faux positifs.
- Nouvelles règles de contrôle.
- [Thunderbird] Portage de l’extension pour Thunderbird 68.
- [LibreOffice 6.4+] UI pour la métabarre en mode Ruban.

L’adaptation pour Thunderbird 68 n’est pas parfaitement identique, attendu que certaines choses étaient dysfonctionnelles, comme les barres de progression, ainsi que divers détails. Contrairement à ce que je pensais, je n’ai pas porté l’extension vers MailExtension, le nouveau système ne semblant pas adapté à l’ancien. On est clairement entre deux mondes, et ce qui vient est encore en cours de construction. Il s’agit donc d’une adaptation de l’ancien système, et ça reste encore l’ancien système.


## Grammalecte 1.3.2 (21 août 2019)

- Moins de faux positifs.
- [Bug] Pas de sortie d’erreurs en mode JSON.
- [Server] Fonctionnel en mode WSGI.

Seules les versions Python (LibreOffice, CLI et Serveur) ont été mises à jour.


## Grammalecte 1.3.1 (21 août 2019)

(Version boguée et supprimée.)


## Grammalecte 1.3.0 (17 août 2019)

- Nouvelles règles de contrôle.
- Correction de faux positifs.
- [Firefox] Autorafraîchissement : correction du bug qui décalait le curseur après coup.
- [Firefox] Option pour mémoriser l’autorafraîchissement.
- [Firefox] Correction du bug qui gelait dans certains cas le correcteur grammatical.
- [Firefox] Bouton pour redémarrer le serveur grammatical en cas de panne ou gel de ce dernier.
- [Serveur] Le serveur renvoie dorénavant le content-type “application/json; charset=UTF-8” quand il renvoie du JSON.


## Grammalecte 1.2.1 (16 juillet 2019)

- Quelques nouvelles règles de contrôle.
- Quelques faux positifs corrigés.
- [LibreOffice] correction d’un bug de l’éditeur lexical.
- [LibreOffice] console Python accessible depuis le menu À propos.

Cette version mineure uniquement publiée pour corriger un bug sur LibreOffice n’est pas proposée pour Firefox et Chrome.


## Grammalecte 1.2 (04 juillet 2019)

- Nouvelles règles de contrôle.
- Moins de faux positifs.
- Améliorations du moteur de suggestion.
- [Firefox] Correction du bug de l’option Restructuration du formateur de texte.
- [LibreOffice 6.3+] Options graphiques : couleurs et forme du soulignement.
- [Serveur] Multiprocessus.


## Grammalecte 1.1.1 (22 mai 2019)

- Quelques nouvelles règles de contrôle (surtout pour l’option OCR).
- Moins de faux positifs.
- [Firefox] Option pour l’auto-rafraîchissement de la correction grammaticale.
- [Firefox] Correction de la copie dans le presse-papiers.
- [Firefox] Correction de bugs d’interface.


## Grammalecte 1.1 (15 mai 2019)

- Nouvelles règles de contrôle.
- Moins de faux positifs.
- [Graphspell] Moteur de suggestions amélioré : peut désormais suggérer des mots avec apostrophes (lorsquil -> lorsqu’il).
- [Firefox] Interface refaite :
· Disparition du menu sous les boutons inclus dans les pages.
· Menu inclus dans le panneau de contrôle.
· Panneau multifonctions.
· Formateur de texte disponible pour tous les types de texte contrôlés.
· Ajout du conjugueur dans le panneau de contrôle.


## Grammalecte 1.0.2 (02 mai 2019)

- Quelques règles supplémentaires.
- Moins de faux positifs.
- [Éditeur lexical] Correction du bug à l’enregistrement constaté dans certaines circonstances.


## Grammalecte 1.0.1 (17 avril 2019)

- Quelques règles supplémentaires.
- Moins de faux positifs (notamment avec l’option OCR).
- [Firefox] Interface : auto-rafraîchissement de la correction grammaticale (3 secondes après la dernière frappe).


## Grammalecte 1.0 (13 avril 2019)

Depuis la version précédente, le moteur interne de Grammalecte a été complètement réécrit. Au lieu de chercher des erreurs avec une longue suite d’expressions régulières, le moteur sépare dorénavant les phrases en tokens (c’est-à-dire en mots et divers éléments textuels) puis parcourt plusieurs graphes de tokens à la recherche de motifs d’erreurs.

Cette nouvelle méthode n’a pas complètement remplacé l’ancienne qui reste toujours pertinente dans quelques cas de figure. La combinaison des deux solutions apporte une grande souplesse d’utilisation.

L’énorme avantage qu’apporte la nouvelle méthode, c’est de pouvoir gérer des cas bien trop complexes pour l’ancienne (uniquement possibles au prix d’un coût de traitement énorme). Il est bien moins difficile (quoique pas facile pour autant) de connaître le contexte et de traiter un nombre croissant de cas particuliers. Pour un coût la plupart du temps négligeable. Les graphes permettent en effet de traiter un nombre considérable de règles sans gréver notablement les performances d’analyse.

Quoique traitant plus de deux fois plus de règles qu’autrefois, Grammalecte va environ deux fois plus vite pour traiter un même texte qu’avec la version précédente (pour le cas Python/LibreOffice). Avec JavaScript (Firefox/Thunderbird/Chrome), l’amélioration n’est pas si notable, mais ça va environ 20 % plus vite qu’avec Python. Par ailleurs, l’ajout de nouvelles règles ne change presque rien aux performances actuelles, si tant est qu’on prenne garde à les écrire de manière à minimiser le coût de traitement.

Grammalecte 0.6.5 :

- Analyse des paragraphes : 159 règles par expression régulière.
- Analyse des phrases : 1934 règles par expression régulière.

Grammalecte 1.0 :

- Analyse des paragraphes : 165 règles par expression régulière.
- Analyse des phrases : 5377 règles par description de tokens réparties en 9 graphes différents, et 38 règles par expression régulière.

Bien que la manière de compter les règles soit quelque peu différente (certaines règles antérieures ont été fusionnées, beaucoup d’autres ont été décomposées), la nouvelle version comporte de très nombreuses nouvelles règles de contrôle. La réécriture quasi totale des anciennes règles m’a permis de resserrer les mailles du filet dans de nombreux cas de figure. Les nouvelles règles concernent à peu près tous les domaines, mais plus notablement la détection de certaines confusions et différentes formes de conjugaison. J’espère que cela sera visible à l’usage.

En ce qui concerne les faux positifs, beaucoup ont été éradiqués. Mais les nouvelles règles engendrent elles aussi inévitablement des possibilités de faire de nouveaux faux positifs. Selon le principe suivi par Grammalecte depuis le commencement, tout est fait pour les supprimer autant que possible et vous ne devriez pas en rencontrer plus qu’auparavant, et j’espère même moins.

Autres nouveautés et améliorations :

- Écriture des règles grandement facilitée (ça ne concerne que le développement, mais cette facilité rend plus aisée l’amélioration de l’existant).
- Nouvelle option pour la normalisation des graphies épicènes (avec points médians).
- [LibreOffice 6.2+] Les erreurs sont dorénavant affichées avec des couleurs différentes.
- [LibreOffice] Révision de l’interface du conjugueur, du formateur de texte et de l’éditeur lexical.
- [Graphspell] Amélioration du moteur de suggestion orthographique. Le correcteur orthographique, sait à présent séparer les mots qui ont été fusionnés (desordres -> des ordres).
- [Firefox 63+] Grâce au ShadowDOM, les menus et panneaux de Grammalecte inclus dans une page ne sont plus altérés par les règles d’affichage de la page web (fonctionnalité apportée par un contributeur externe).

Cette nouvelle version de Grammalecte clôt les travaux engagés par la campagne de financement. Mais cela ne signifie aucunement que Grammalecte ne sera plus amélioré. Le développement va se poursuivre.

Nouvelles fonctionnalités envisagées :

- Une interface plus fluide pour Firefox (et Chrome).
- Étendre les dictionnaires existants avec d’autres fournis par un serveur de dictionnaires communautaires. Cette fonctionnalité devait à l’origine être prête pour Grammalecte 1.0, mais ayant pris beaucoup de retard pour diverses raisons, j’ai préféré publier la nouvelle version sans plus attendre.
- Il est probable qu’il faudra une nouvelle fois refaire l’extension pour Thunderbird, attendu que les développeurs semblent vouloir réécrire l’interface de programmation (API), plus ou moins conformément au modèle des WebExtensions pour Firefox, avec des spécificités propres à Thunderbird…
- Bien sûr, améliorer encore le cœur du correcteur grammatical lui-même.



## Grammalecte 0.6.5 (18 avril 2018)

- Moins de faux positifs.
- Divers bugs corrigés.
- [CLI] option pour charger le dictionnaire personnalisé.
- [Thunderbird] Mise à jour de compatibilité pour Thunderbird 60.
- [Formateur de texte] Titres de civilité.
- [Graphspell] Vitesse du moteur de suggestions.


## Grammalecte 0.6.3 et 0.6.4 (avril 2018)

- Nouvelles règles de contrôle.
- Moins de faux positifs.
- Divers bugs corrigés.
- [Graphspell] Accélération de la suggestion orthographique.
- [Graphspell] Suggestion orthographique améliorée.
- [Graphspell] La suggestion orthographique tient mieux compte des problèmes de casse.
- [LibreOffice] Éditeur lexical.
- [LibreOffice] Graphspell remplace Hunspell (options pour garder Hunspell).
- [Thunderbird] Éditeur lexical.
- [Firefox] Meilleure protection des panneaux contre l’héritage CSS des pages.
- [Firefox] Les corrections sont dorénavant aussi répercutées dans les nodes HTML éditables.


## Grammalecte 0.6.2 (04 mars 2018)

- Nouvelles règles de contrôle.
- Moins de faux positifs.
- Divers bugs corrigés.
- Dictionnaires mis à jour avec les demandes des contributeurs.
- Graphspell, le correcteur orthographique de Grammalecte, est désormais un paquet indépendant.
- [Graphspell] Quelques améliorations du moteur de suggestion orthographique.
- [Graphspell] Prise en compte des ligatures (ﬃ, ﬄ, ﬀ, ﬅ, ﬁ, ﬂ, ﬆ) et du caractère ‹ſ› qui remplace le ‹s› dans les anciens livres, et normalisation Unicode NFC.
- [LibreOffice] Interface du conjugueur revue.
- [LibreOffice] Nouvel outil : Recenseur de mots.
- [Firefox] Préversion de l’éditeur lexical (Unique accès via Ctrl+Maj+8).
- Beaucoup d’améliorations dans la mécanique interne.


### Le recenseur de mots

Lorsque vous demandez un comptage par lemme chaque lemme potentiel d’un mot verra son nombre d’occurrences incrémenté de 1.
Et il y a en français de nombreux mots qui sont les déclinaisons de lemmes différents.

Ce qui signifie, par exemple, que :

- chaque occurrence du mot “suis” augmente le nombre pour “être” et “suivre”,
- chaque occurrence du mot “œuvre” augmente le nombre pour “œuvre” et “œuvrer”,
- chaque occurrence du mot “lui” augmente le nombre pour “lui” et “luire”,
- chaque occurrence du mot “mais” augmente le nombre pour “mai” et “mais”,
- chaque occurrence du mot “saurai” augmente le nombre pour “savoir” et “saurer”.

Vous pouvez trier différemment la liste des mots en cliquant sur l’entête d’une colonne.


### L’éditeur lexical

Dernièrement, j’ai beaucoup travaillé sur l’éditeur lexical, mais à cause du retard pris en décembre et janvier, celui-ci n’est pas encore au point. Seuls Firefox et Chrome en intègrent une préversion. Je n’ai pas intégré l’éditeur dans LibreOffice. Celui de Thunberbird n’a pas encore été conçu.

Merci de bêta-tester.

C’est une fonctionnalité non officielle susceptible de changements majeurs.


### Droits demandés sur Firefox/Chrome

Les nouveaux droits exigés par Grammalecte 0.6.12 pour Firefox/Chrome ont provoqué quelque émoi. En effet, l’extension demande à pouvoir télécharger des fichiers et accéder et modifier l’historique des téléchargements…
Sauf que, en réalité, elle ne fait rien de tout ça.

Les droits sont accordés par paquets. Grammalecte demande les droits sur l’API appelée “downloads” qui permet de télécharger des fichiers et de lire et modifier l’historique du navigateur. Mais Grammalecte ne s’en sert pas pour ça.

Il n’existe actuellement aucune interface pour enregistrer un fichier sur le disque dur. Une question de sécurité, paraît-il. Sauf que l’interface “downloads” permet d’enregistrer sur le disque dur ce qui est téléchargé sur Internet… Allez comprendre.

Grammalecte ruse pour enregistrer un fichier qui n’a pas été téléchargé sur Internet. L’extension crée un fichier en mémoire (un blob), puis crée une URL vers ce fichier, puis demande à l’interface “downloads” d’enregistrer ce fichier sur le disque.

En résumé, Grammalecte télécharge depuis la mémoire de l’ordinateur pour pouvoir enregistrer sur le disque dur. Voilà tout. Mais l’extension ne vous espionne pas et ne change rien à votre insu.

Ceci est nécessaire pour exporter certaines données si vous le demandez.

Le code pour les experts :

    let xBlob = new Blob([ JSON.stringify(this.oIBDAWG.getJSON()) ], {type: 'application/json'});
    let sURL = URL.createObjectURL(xBlob);
    browser.downloads.download({ filename: "fr.personal.json", url: sURL, saveAs: true });



## Grammalecte 0.6.1 (fin novembre 2017)

- Nouvelles règles de contrôle.
- Moins de faux positifs.
- Divers bugs corrigés.
- Amélioration du moteur de suggestion du correcteur orthographique.
- [Firefox] Détection de locutions dans le lexicographe.
- [Firefox] Bugs d’interface corrigés.


## Grammalecte 0.6 (31 octobre 2017)

La nouvelle version 0.6 de Grammalecte est la continuation des versions 0.5.x qui ont vu apparaître beaucoup de nouvelles fonctionnalités.

La principale nouveauté de la version 0.6, c’est la refonte complète de l’extension pour Firefox, désormais compatible avec Firefox 57+. Elle utilise la nouvelle API appelée WebExtension, appelée être normalisée dans les années à venir pour être compatible avec tous les navigateurs. browserext.github.io…

En conséquence, l’extension pour Firefox a été adaptée pour Chrome.

L’autre modification importante de la version 0.6, c’est la suppression de la dette technique concernant le fonctionnement interne du correcteur grammatical. Celui-ci se base désormais bien plus sur le désambiguïsateur que sur le processeur de texte pour éviter les innombrables problèmes de faux positifs.

Le dernier point majeur de la version 0.6, c’est l’apparition d’un moteur de suggestion orthographique. C’est une affaire complexe. Le moteur de suggestion fonctionne sur une base de substitutions phonétiques, des interversions de lettres, et d’autres modifications systématiques. C’est encore une fonctionnalité en bêta, mais qui est fonctionnelle avec la majeure partie des erreurs orthographiques rencontrées. À noter que ce nouveau moteur de suggestion pourra faire l’objet d’adaptations simples pour faire des suggestions au cas par cas si besoin.
Pour l’instant, seuls les extensions pour Firefox et Chrome utilise ce nouveau moteur de suggestion.

Et, bien sûr, cette nouvelle version apporte, comme toujours :
- de nouvelles règles de contrôle,
- moins de faux positifs,
- des corrections de bugs,
- un meilleur contrôle des suggestions grammaticales.

La route est désormais ouverte pour des modifications plus profondes du correcteur grammatical.


## Grammalecte 0.5.18 (26 juillet 2017)

- Moins de faux positifs.
- Nouvelles règles de contrôle.
- [Firefox] Refonte du code de genèse du HTML (hypothétique faille de sécurité bouchée), code plus propre.
- [Firefox] Panneau grammatical : nouvelle interface, mode édition toujours actif, texte (théoriquement) incassable…
- [Firefox] Lexicographe : prise en compte de la ponctuation.
- [Firefox] Les espaces insécables sont remplacées par des espaces insécables fines dans les suggestions du correcteur.


## Grammalecte 0.5.17 (09 juin 2017)

- Moins de faux positifs.
- Quelques nouvelles règles.
- [Sources] Révision du processus de build.


## Grammalecte 0.5.16 (26 avril 2017)

- Moins de faux positifs.
- [LibreOffice] Ajustements de l’agencement du conjugueur pour LO 5.3.
- [Firefox] Normalisation Unicode NFC.
- [Thunderbird] Nouvelle extension :  https://addons.mozilla.org/fr/thunderbird/addon/grammalecte-fr-thunderbird/
- [Sources] Révision du processus de build.
- [Sources] Intégration du processus de build du lexique et des dictionnaires pour Hunspell.


## Grammalecte 0.5.15 (31 janvier 2017)

- Le correcteur cesse dorénavant d’empiler les erreurs les unes sur les autres (dans 99% des cas) lorsqu’une erreur est détectée par différentes règles. Les règles de contrôle peuvent à présent être priorisées.
- Moins de faux positifs.
- Meilleure désambiguïsation.
- Nouvelles règles.
- Option OCR améliorée.
- Extension des tests.
- [LibreOffice/Linux] Correction du bug du lexicographe.
- [Linux] Make.py opérationnel.


## Grammalecte 0.5.14 (07 décembre 2016)

- Correction du bug avec « a base de » et « a cause de ».
- Corrige le bug faisant disparaître du dictionnaire orthographique les unités de mesure préfixée au singulier.
- Moins de faux positifs.

C’est une petite version de maintenance.


## Grammalecte 0.5.13 (26 novembre 2016)

Mise à jour corrective. J’ai ajouté beaucoup de tests et déniché bon nombre de bugs. À présent, toutes les règles de contrôle sont testées au moins une fois. Plus d’une fois pour beaucoup.
J’ai aussi revu et clarifié le processus de construction des extensions.


## Grammalecte 0.5.12.1 (05 novembre 2016)

Mise à jour corrective. J’ai corrigé des bugs problématiques concernant le moteur de suggestions (le problème concernait plus Firefox que Writer, mais mettez quand même à jour).


## Grammalecte 0.5.12 (02 novembre 2016)

- Moins de faux positifs.
- Quelques améliorations diverses et corrections de bugs.
- Réorganisation de nombreuses règles.
- Amélioration significative de moteur de suggestion et du soulignement des mots possiblement erronés. Par exemple, si vous écrivez “la mec”, le correcteur vous propose à présent deux suggestions (le -> la et mec -> Mecque).


## Grammalecte 0.5.11 (08 octobre 2016)

(La version 0.5.10 a sauté à cause d’un cafouillage lors des mises à jour sur addons.mozilla.org.)

- Moins de faux positifs.
- Meilleures suggestions.
- Quelques nouvelles règles.

## Grammalecte 0.5.9 (1ᵉʳ septembre 2016)

- Moins de faux positifs.
- Meilleures suggestions.
- Quelques nouvelles règles.
- [Firefox] Affichage corrigé sur les thèmes de bureau sombres.
- [Firefox] Correction du bug des étiquettes dans le lexicographe.
- [Writer] Correction de la position de la fenêtre des options.


## Grammalecte 0.5.8 (17 août 2016)

- Beaucoup moins de faux positifs.
- Amélioration du moteur de suggestion.
- Correction de bugs du moteur de suggestion.
- [Firefox] L’extension fonctionne à présent en mode navigation privée.
- [Firefox] Modifications des options par défaut (concernant les espaces insécables sur les nombres et les unités de mesure).
- [Writer] Ajout des locales africaines : fr_BF (Burkina Faso) fr_CI (Côte d’Ivoire) fr_SN (Sénégal) fr_ML (Mali) fr_NE (Niger) fr_TG (Togo) fr_BJ (Bénin).
- [Writer] Centrage des fenêtres de dialogue avec LibreOffice 5.2.

Cette version est surtout une amélioration de l’existant : j’ai corrigé une quantité importante de faux positifs grâce à des tests avec Firefox, j’ai renforcé les tests unitaires, corrigé des tas de petits détails, etc.

Comme d’habitude, il faut attendre que Mozilla valide la version pour Firefox pour qu’elle soit automatiquement téléchargée.


## Grammalecte 0.5.6 (26 mai 2016)

- Beaucoup moins de faux positifs.
- Corrige quelques bugs.
- Quelques améliorations sur la suggestion de certains types d’erreurs.


## Grammalecte 0.5.5 (16 mai 2016)

La nouvelle version de Grammalecte apporte un lot considérable de changements, la plupart internes, donc peu visibles pour l’utilisateur de LibreOffice.

- Le moteur interne de Grammalecte est dorénavant indépendant. Plus besoin de Hunspell ni de LibreOffice.
- La syntaxe d’écriture des règles de contrôle a pas mal changé (la gestion de la délimitation des mots et de la casse est facilitée).
- Il y a un désambiguïsateur (il est encore peu utilisé, mais ça va aller en s’améliorant).
- D’un moteur multi-passes, on est revenu à un moteur bi-passes (une pour le paragraphe, une pour les phrases), qui permet de gérer le flot des règles avec plus de finesse. Chaque règle de contrôle est dorénavant multi-actions, elle peut générer plusieurs erreurs, modifier le texte et faire de la désambiguïsation.
- Il y a dorénavant des tests unitaires (pas complets, mais ça va aller en s’améliorant).
- Le correcteur a été porté en JavaScript pour Firefox, avec la syntaxe ES6 (non portable, car il utilise des spécificités à Firefox).
- Le moteur de suggestion a été amélioré. (Du travaille => travail).
- De nouvelles règles ont été ajoutées.
- Pas mal de bugs corrigés.
- Moins de faux positifs.
- Le correcteur est utilisable en ligne de commande (cli.py).


## Grammalecte 0.4.10.7 (27 février 2016)

- Correction du bug d’affichage des options pour les langues autres que le français et l’anglais.


## Grammalecte 0.4.10.6 (19 février 2016)

(Pour LibreOffice uniquement.)

Rien de nouveau pour cette version. Il s’agit seulement d’un changement de numérotation nécessaire pour LO 5.1, qui intègre une version des dictionnaires plus récentes. Du coup, LO 5.1 vous propose une une mise à jour qui va supprimer Grammalecte de votre profil.


## Grammalecte 0.4.10.5 (05 novembre 2015)

- Correction d’un bug du moteur de suggestion.
- Correction d’un plantage dans certains cas particuliers.


## Grammalecte 0.4.10.4

- Correction d’un bug du formateur de texte (ils t’ont -> ils-t-ont)


## Grammalecte 0.4.10.3 (14 septembre 2015)

- Correction d’un bug qui faisait planter le logiciel.


## Grammalecte 0.4.10.2 (14 septembre 2015)

- Mise à jour du dictionnaire.
- Redimensionnement du conjugueur pour LO 5.


## Grammalecte 0.4.10.1 (25 mai 2015)

- Corrections de quelques faux positifs gênants.


## Grammalecte 0.4.10 (16 mai 2015)

- Moins de faux positifs (comme d’habitude).
- Amélioration des règles du préprocesseur de texte.
- Nouvelles règles de contrôle sur des confusions élémentaires.

Le code a été un peu revu et réorganisé. La gestion de la casse a été revue. Mais c’est loin d’être fini.


## Grammalecte 0.4.9 (21 avril 2015)

- Refonte de la boîte de dialogue des options.
- Quelques améliorations et corrections de bugs.
- Formateur de texte : contournements des bugs d’expressions régulières de LibreOffice 4.4


### À propos des bugs concernant les expressions régulières

Le formateur de texte repose principalement sur le moteur d’expressions régulières de LibreOffice qui utilise lui-même une librairie appelée ICU. Celle-ci semble pas mal buguée et ses bugs semblent varier avec les versions. Du coup, il est devenu prudent de simplifier les expressions régulières pour contourner les problèmes. Il est difficile de savoir ce qui se passe.

Même si je pense avoir contourné tous les problèmes, je vous recommande donc toujours de lancer le formateur de texte AVANT de relire et non après.


### Avenir de Grammalecte

Ça fait assez longtemps que Grammalecte s’éloigne peu à peu de Lightproof. Mais à présent, quelle que soit l’issue de la campagne de financement participatif, Grammalecte va prendre le large.


## Grammalecte 0.4.8 (27 mars 2015)

- Beaucoup d’améliorations sur les règles de contrôle grammaticales.
- Moins de faux positifs.
- Correction de divers bugs.
- Formateur de texte : correction d’un bug sur les ligatures œ Œ.
- Formateur de texte : nouvelle option sur les ligatures typographiques.
- Formateur de texte : option pour remplacer les puces par des tirets cadratins.
- Boîte de dialogue “À propos”.
- Modification du champ “Auteur” : passage de Basic à Python et refonte de l’interface.
- Performances : moins de latence au démarrage de Writer. Grammalecte ne charge plus tous les outils dès le démarrage, mais attend qu’on les réclame.
- Performances : des optimisations de code.
- Performances : beaucoup moins de latence sur le menu contextuel des erreurs dans les gros paragraphes (plus de 3000 signes).

Et un nouveau logo.


## Grammalecte 0.4.7.1 (1ᵉʳ janvier 2015)

- Réduction de la taille du conjugueur pour passer sur les petits écrans (1366 × 768)


## Grammalecte 0.4.7 (31 décembre 2014)

- Moins de faux positifs.
- Quelques bugs corrigés.
- Quelques améliorations du correcteur.
- Détection des ligatures typographiques.
- Quelques améliorations du formateur de texte.


## Grammalecte 0.4.6 (07 novembre 2014)

- Moins de faux positifs.
- Quelques bugs corrigés.
- Quelques améliorations mineures.

Rien de révolutionnaire. De la consolidation.


## Grammalecte 0.4.5 (22 septembre 2014)

- Moins de faux positifs.
- Quelques bugs corrigés.

Cette version sort plus tôt que prévu à cause de la mise à jour des dictionnaires dans les versions de LibreOffice à paraître. Si vous mettez à jour LibreOffice avec une précédente version de Grammalecte, LibreOffice va considérer que l’extension incluse est plus récente et désinstaller Grammalecte.


## Grammalecte 0.4.4.1 (28 août 2014)

- Un faux positif gênant supprimé.
- Correction du plantage du formateur de texte si l’extension est installée pour tous les utilisateurs.

Note : la sauvegarde des options du formateur de texte n’est effective que si l’extension est installée pour l’utilisateur. (Du moins, sur Windows ; sur Linux, je ne sais pas.)


## Grammalecte 0.4.4 (26 août 2014)

- Beaucoup moins de faux positifs.
- Formateur de texte : mémorisation des options appliquées + retour du bouton Fermer.
- Amélioration du préprocesseur de texte.
- Bug d’OpenOffice contourné (ça ne devrait plus planter).


## Grammalecte 0.4.3 (15 juillet 2014)

- Quelques faux positifs en moins.
- Correction des erreurs dans l’interface des options en anglais.
- Amélioration des options de restructuration du formateur de texte.
- Modification du champ “Auteur”.


## Grammalecte 0.4.2 (28 juin 2014)

- Corrections de faux positifs.
- Correction du bug qui décalait d’un caractère à droite les erreurs suivant les phrases se finissant par un sigle de deux lettres avec points.
- Formateur de texte : options de restructuration.

En ce qui concerne le formateur de texte, l’option “Fusion des paragraphes contigus” fusionne tous les paragraphes qui ne sont pas séparés par un paragraphe vide.

Par exemple, le texte suivant :

    Titre 1

    Lorem ipsum dolor sit amet, consectetur adipiscing elit.
    Nullam auctor diam quam, ut scelerisque lectus laoreet
    vulputate. Nulla facilisi. Sed nec mi neque. Etiam urna lectus,
    sollicitudin tincidunt euismod eu, commodo quis enim.
    Mauris in risus nec nisl vestibulum sollicitudin. Vivamus
    feugiat magna et enim mattis eleifend. Morbi iaculis sem
    at suscipit cursus. Aenean a libero aliquet, tempus velit et,
    ultrices nisi. Sed vestibulum, lacus at pharetra rhoncus, est
    felis laoreet risus, vitae malesuada elit elit a mauris. Nulla a
    rhoncus mi, eu tempor purus. Donec eu scelerisque nisi.
    Mauris id mauris libero.

    Titre 2

    Ut rhoncus eu ipsum vitae interdum. Nam volutpat semper
    placerat. Sed et suscipit justo. Proin ultricies pellentesque
    ornare. Class aptent taciti sociosqu ad litora torquent per
    conubia nostra, per inceptos himenaeos. Duis mattis leo
    ipsum, a feugiat sapien imperdiet non. Ut rhoncus et
    massa vitae fermentum. Duis non dictum ipsum, nec
    condimentum sapien. Pellentesque pharetra tortor
    blandit dictum sollicitudin.


deviendra :

    Titre 1
    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam auctor diam quam, ut scelerisque lectus laoreet vulputate. Nulla facilisi. Sed nec mi neque. Etiam urna lectus, sollicitudin tincidunt euismod eu, commodo quis enim. Mauris in risus nec nisl vestibulum sollicitudin. Vivamus feugiat magna et enim mattis eleifend. Morbi iaculis sem at suscipit cursus. Aenean a libero aliquet, tempus velit et, ultrices nisi. Sed vestibulum, lacus at pharetra rhoncus, est felis laoreet risus, vitae malesuada elit elit a mauris. Nulla a rhoncus mi, eu tempor purus. Donec eu scelerisque nisi. Mauris id mauris libero.
    Titre 2
    Ut rhoncus eu ipsum vitae interdum. Nam volutpat semper placerat. Sed et suscipit justo. Proin ultricies pellentesque ornare. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Duis mattis leo ipsum, a feugiat sapien imperdiet non. Ut rhoncus et massa vitae fermentum. Duis non dictum ipsum, nec condimentum sapien. Pellentesque pharetra tortor blandit dictum sollicitudin.



Rappel : dans LibreOffice/OpenOffice, les paragraphes sont limités à 65535 caractères. Passer outre cette limite provoque un plantage de Writer. Cette limitation est levée avec LibreOffice 4.3 qui sortira prochainement.


## Grammalecte 0.4.1 (13 juin 2014)

- Correction de divers faux positifs.
- Vérification des dates (idée de Dominiko).


## Grammalecte 0.4.0.1 (22 mai 2014)

- Corrections de quelques bugs.


## Grammalecte 0.4 (20 mai 2014)

Ça a bien failli être une version 0.3.10, mais avec toutes les nouveautés apportées et certaines difficultés dont Grammalecte triomphe enfin, je me suis dit que ça ne ferait pas de mal de monter le numéro de version secondaire, d’autant plus que nombre de mécanismes internes ont été modifiés et que je pars sur une base consolidée.


### Fonctionnalités nouvelles

- Mécanisme de suggestion.
- Petite réorganisation des options.
- Corrections de divers faux positifs.
- Renforcement des règles de contrôle.
- Recherche de répétitions dans la phrase et dans le paragraphe (options à activer).
- Quelques bugs corrigés.

### Moteur interne

- Règles de désambiguïsation dans le préprocesseur de texte.
- Contrôle des options revu : l’option est vérifiée avant la mise en œuvre d’une règle.
- Écriture simplifiée des options dans les règles.
- Optimisation et clarification de code.
- Chaque règle a désormais un identifiant unique (auparavant, les règles avec un même motif regex partageaient le même identifiant).


### À propos du mécanisme de suggestion

Dorénavant, au lieu d’avoir juste un message vous disant ce qui est erroné, Grammalecte fournit à présent la forme correcte. La suggestion, c’est quand même une affaire délicate, et il y a encore une marge d’amélioration pour tous les cas difficiles, mais ça fonctionnera la plupart du temps.

Exemples :

- Il fais beau. -> suggestion : fait.
- Il est parties -> suggestion : parti.
- Des hommes grande -> suggestion : grands.
- Martine est intelligent -> suggestion : intelligente.
- Tu te prenait pour un cador -> suggestion : prenais

Globalement, pour tout ce qui concerne les conjugaisons, ça devrait être fiable, attendu que ça repose sur le conjugueur et que les verbes sont bien normés. En revanche, pour tout ce qui est singulier/pluriel/féminin/masculin, les noms ou les adjectifs, c’est plus difficile et bien moins normé que les verbes. Du coup, le moteur tâtonne à la recherche de formes fléchies potentiellement correctes. Par ailleurs, tous les mots ne sont pas générés par un lemme commun. Par exemple, vieil, vieille(s) et vieux n’étant pas liés dans le dictionnaire, le moteur de suggestion ne peut rien faire. Il faudrait normaliser tout le dictionnaire (encore un travail de longue haleine) pour pouvoir fournir à coup sûr des réponses. Ceci dit, normalement, la plupart du temps, ça trouvera ce que ça cherche, le dictionnaire étant quand même globalement bien conçu, hormis pour les cas particuliers à la marge.

Autre problème qui nécessiterait de revoir le fonctionnement du correcteur : le moteur essaie parfois de suggérer sur le mauvais mot. Exemple : si vous écrivez « des grands maisons », Grammalecte essaie de trouver la forme masculine de maison. N’y parvenant pas, il signale l’erreur mais ne suggère rien.

Autre bizarrerie : si vous écrivez « elles sont aveugle », Grammalecte vous suggérera aveugles et aveuglées. Est-ce un problème ? Pas sûr.



## Grammalecte 0.3.9.1 (19 avril 2014)

Cette nuit, j’ai rêvé qu’il y avait un bug dans mon code à un endroit précis. Vérification faite, c’était bien le cas. Donc, nouvelle version. J’ai aussi vu que le Conjugueur plantait parfois avec AOO.

- Correction d’un bug du préprocesseur de texte.
- Correction du plantage du conjugueur sur AOO.


## Grammalecte 0.3.9 (18 avril 2014)

- moins de faux positifs,
- amélioration du préprocesseur de texte,
- correction du bug concernant la détection des lettres isolées,
- corrections de divers bugs,
- raccourci clavier pour accéder au menu,
- + version rétrocompatible avec AOO 4.

### Concernant la version pour AOO 4

AOO 4 contient la version 2.7 de Python qui est un peu mieux que Python 2.6, et c’est pourquoi je me suis attelé à écrire un mini-script de conversion. Ceci dit, Python 2.7 reste obsolète et contient des bugs qu’il est difficile d’ignorer (les mêmes que ceux de la version 2.6). Ce n’est pas dramatique, mais cela occasionne des effets curieux.
Par exemple, si vous écrivez « Était-elle fatiguer ? », AOO 4 ne verra rien.
Si vous écrivez « Mais était-elle fatiguer ? », AOO 4 verra l’erreur.
En fait, dès qu’une phrase (ou une expression régulière) commence par un signe avec diacritique, celui-ci semble ignoré, ce qui provoque parfois des erreurs dont certaines ont été contournées dans le code de Grammalecte, mais il est impossible de tout gérer de la sorte.

Par ailleurs, la conversion pour AOO 4 n’est peut-être pas parfaite. Ne vous étonnez donc pas si certaines erreurs détectées avec LibreOffice ne le sont pas avec OpenOffice. Je m’efforcerai de vérifier avec le temps.

Ceci dit, cette version reste quand même meilleure que toutes les précédentes. (On a vécu avec ça pendant longtemps avec Python 2.6 inclus dans LibO 3 et AOO 3, et on a fait avec.)

Je ne sais pas combien de temps la rétrocompatibilité sera maintenue. Cela sera ainsi tant que Grammalecte n’inclura pas une version binaire indexable de ses propres dictionnaires. Ce jour-là, c’en sera définitivement fini des versions rétrocompatibles avec Python 2.7.


## Grammalecte 0.3.8.3 (05 mars 2014)

- Suppression de faux positifs.


## Grammalecte 0.3.8.2 (03 mars 2014)

- Suppression de faux positifs.

Navré pour ces versions à répétition, mais il y avait encore un bug rédhibitoire.


## Grammalecte 0.3.8.1 (02 mars 2014)

- Corrige un bug du formateur de texte.


## Grammalecte 0.3.8 (1ᵉʳ mars 2014)

- Moins de faux positifs (comme toujours).
- Beaucoup de bugs corrigés.
- Nouvelles règles de contrôle.
- Préprocesseur de texte consolidé.
- Le menu global est désormais visible dans les documents maîtres et dans l’éditeur HTML, ainsi que dans Impress.
- Quelques améliorations dans le formateur de texte.

Dans cette version, un gros travail de consolidation a été fait, une attention particulière a été portée sur les règles existantes, dont beaucoup ont été améliorées, et de nouvelles règles de contrôle ont été ajoutées. Vous devriez donc observer que le correcteur détecte plus d’erreurs qu’autrefois.


## Grammalecte 0.3.7 (18 janvier 2014)

Prérequis : Python 3.3

Nouveautés :

- Beaucoup moins de faux positifs.
- Menu global pour accéder aux outils et options (les utilisateurs ignoraient souvent leur existence).
- Quelques nouvelles options pour le formateur de texte.
- Nouvel outil : le conjugueur (qui peut conjuguer tous les verbes du dictionnaire, environ 8000).
- Mise à jour du dictionnaire 5.0 (il y avait encore des problèmes avec certaines graphies élidées).
- Quelques bugs corrigés.

### À propos du conjugueur

Vous pouvez activer directement certaines options en tapant “ne pas” ou “se” avant le verbe ou “?” après celui-ci.

En sus du menu Grammalecte, vous pouvez accéder au conjugueur via le menu contextuel de Writer, en cliquant sur un verbe, sous une forme conjuguée ou non. Par exemple, sur le mot “suis”, il vous proposera de conjuguer “être” et “suivre”.

Cet outil est encore en version bêta, parce qu’il faudrait revoir l’étiquetage des 8000 verbes pour permettre la conjugaison de ceux-ci pronominalement et dans les temps composés (ce sont les participes passés qui posent problème, attendu que les accords sont soumis à des règles complexes). Il faudrait par exemple compléter l’étiquetage pour signaler si le verbe s’utilise avec l’auxiliaire être ou avoir. Du coup, pour la plupart des verbes, j’ai délibérément désactivé certaines options pour éviter d’éventuelles erreurs. Ne vous attendez pas à ce que ce soit corrigé prochainement, ça va demander beaucoup de temps, et ce n’est pas prioritaire à mes yeux.

Par ailleurs, certaines particularités ne sont pas gérées, par exemple le fait qu’il peut exister plusieurs graphies par conjugaison, comme je peux, je puis, ou les variantes graphiques introduites par la réforme de l’orthographe. Mon script prend la première graphie qu’il trouve et ignore le reste. Les variantes graphiques sont oubliées, la réforme est ignorée, sauf si vous demandez à l’infinitif une graphie réformée, comme connaitre (sans l’accent circonflexe).


## Grammalecte 0.3.6.2 (10 décembre 2013)

- Dictionnaires corrigés inclus.


## Grammalecte 0.3.6.1 (10 décembre 2013)

- Corrige quelques bugs,
- Version rétrocompatible pour OpenOffice et LibreOffice 3.x (conçue à l’arrache et la truelle).

Maintenir deux versions est une solution temporaire. Quand j’en aurai marre de ces histoires de rétrocompatibilité avec les anciennes versions de Python, tout ça passera à la trappe. Python 3 est sorti en 2008. Il est plus que temps de se débarrasser Python 2.x.


## Grammalecte 0.3.6 (08 décembre 2013)

- Beaucoup d’améliorations sur le préprocesseur de texte, le champ d’erreurs que le correcteur peut détecter s’est pas mal accru.
- Réduction des faux positifs.
- Intégration du dictionnaire 5.0 qui a été optimisé pour simplifier le fonctionnement du correcteur grammatical, ce qui se traduit par moins de consommation de mémoire et une vitesse accrue.
- Réparation du lexicographe pour OpenOffice.


## Grammalecte 0.3.5 (24 août 2013)

- Moins de faux positifs.
- Correction de bugs concernant le lexicographe.
- Nouvelle numérotation de l’extension afin d’éviter les mises à jour intempestives avec l’extension dictionnaire incluse dans LO.

Finalement, pas de changement global de numérotation. En revanche, l’extension Grammalecte (considérée comme un pack) sera dorénavant numérotée en fonction des dictionnaires inclus.


## Grammalecte 0.3.4 (13 août 2013)

- Corrige un bug important du formateur de texte.
- Quelques faux positifs supprimés.


## Grammalecte 0.3.3 (04 août 2013)

Attention, à compter de cette version, Grammalecte inclut les dictionnaires orthographiques.

Il y aura donc dorénavant deux extensions avec le même identifiant :
- Grammalecte avec les dictionnaires,
- les dictionnaires sans Grammalecte.

L’installation de l’une écrasera l’autre.

L’identifiant unique retenu est celui de l’extension multi-dictionnaires incluse dans LibreOffice 4. L’installation de Grammalecte prendra donc sa place avec les dictionnaires à jour.

Toute autre version des extensions dictionnaires devra être désinstallée (LibreOffice 3.x et OpenOffice). Les anciennes versions de Grammalecte devront aussi être désinstallées, attendu que la nouvelle version ne s’installera pas à leur place.

Modifications :
- Moins de faux positifs.
- Préprocesseur de texte amélioré.
- Formateur de texte : espaces insécables fines, points médians sur certaines unités.

C’est plutôt une version de consolidation, aucun bouleversement majeur.


## Grammalecte 0.3.2 (23 juin 2013)

- Formateur de texte amélioré pour préserver le format existant.
- Quelques faux positifs en moins.


## Grammalecte 0.3.1 (18 juin 2013)

- Moins de faux positifs.
- Correction de bugs concernant LibreOffice 3.x et OpenOffice.


## Grammalecte 0.3 (17 juin 2013)

### Prérequis

Cette version nécessite l’installation de la dernière extension multi-dictionnaires, la version 4.11, car l’étiquetage grammatical a changé. Sans cette extension, le correcteur grammatical sera dans certains cas incapable de comprendre ce que sont certains mots. Pareillement, évitez d’installer la nouvelle extension multi-dictionnaires avec les anciennes versions de Grammalecte.


### Modifications

- Nouvelles règles de contrôle.
- Moteur multi-passes et réorganisation des règles de contrôle.
- Moins faux positifs.
- Bugs divers corrigés.
- Le lexicographe : informe de la nature grammaticale des mots dans le menu contextuel.
- Le formateur de texte : formatage automatique du texte (apostrophes, espaces, etc.), accessible depuis le menu Outils.


### Moteur multi-passes et préprocesseur de texte

La nouveauté majeure, la plus discrète, concerne le moteur interne du correcteur. En version 0.1, Grammalecte, à l’instar de LanguageTool, n’effectuait qu’une seule passe sur le texte. Avec la version 0.2, le moteur en faisait deux (une paragraphe par paragraphe, une phrase par phrase) et utilisait un préprocesseur pour simplifier le texte entre les deux passes.

Avec la version 0.3, le correcteur peut effectuer un nombre de passes théorique illimité, avec pour chacune d’elles un préprocesseur qui simplifie le texte et ôte ce qui n’est plus utile à la vérification. À présent, pour le français, cinq passes sont faites, et d’autres viendront peut-être à l’avenir selon les besoins.

Voici grosso modo comment ça se passe :

* Passe 0, paragraphe par paragraphe :
    - Préprocesseur : suppression des URL.
    - Vérification des signes graphiques : apostrophes, ponctuation, espaces, guillemets, etc.

* Passe 1, phrase par phrase :
    - Préprocesseur : suppression des signes graphiques encombrants (guillemets, certains tirets).
    - Contrôle de cohérence entre certaines catégories de mots.

* Passe 2 :
    - Préprocesseur : suppression des adverbes, locutions adverbiales, expressions usuelles, etc.
    - Contrôle des groupes nominaux.

* Passe 3 & 4 :
    - Préprocesseur : suppression de certains pronoms et adverbes entourant les verbes.
    - Contrôle des conjugaisons.

Si le moteur multi-passes est au point, en revanche, la réorganisation des règles de contrôle n’en est qu’à son commencement. Le potentiel du nouveau moteur se déploiera au fur et à mesure des versions ultérieures. Comme d’habitude, je préfère opérer une montée progressive des ajouts et des modifications pour éviter une explosion des faux positifs.


### Le lexicographe

Cet outil est très simple à utiliser, il est purement informatif. Sur un clic droit, le menu contextuel indique quelle est la nature grammaticale du mot sur lequel se trouve le curseur : nom, adjectif, adverbe, verbe, etc. Cette fonction est un peu expérimentale, mais fonctionne sans poser de problème particulier. Si vous constatez une erreur dans l’étiquetage, il faut proposer une modification du dictionnaire, car cet outil ne fait que transcrire de manière lisible les étiquettes inscrites dedans.

L’étiquetage du dictionnaire n’est pas toujours cohérent, observerez-vous peut-être. Par exemple, pour un nom masculin ou féminin, il peut y avoir une entrée étiquetée épicène, ou bien deux entrées, l’une disant que le nom est masculin, l’autre qu’il est féminin.

Requis : LibreOffice 4+, OpenOffice 3.4+. Ça ne fonctionne pas avec OpenOffice.org et LibreOffice 3.x à cause d’un bug de l’API corrigé dernièrement. Même avec ces versions, vous pouvez tout de même installer l’extension. Le correcteur grammatical fonctionnera, mais cette fonction sera inopérante.

Note : cet outil permettra à l’avenir d’afficher d’autres informations lexicales, sémantiques ou étymologiques. C’est déjà possible en fait, mais ça complique la tâche du correcteur grammatical, car ces informations sont disposées au beau milieu des étiquettes grammaticales (c’est dû à la structure du dictionnaire, et ce n’est pas modifiable). Pour bénéficier de ces informations, il faudrait ou bien tolérer une perte de performance du correcteur grammatical, ou bien recomposer tout le lexique dans un automate à états finis binaire indexable. La 2e solution a ma préférence, c’est le remède à beaucoup de problèmes, mais ce n’est pas pour tout de suite.


### Le formateur de texte

Cet outil, là encore assez simple, sert à automatiser le processus de correction des erreurs les plus communes, notamment typographiques. Si vous en avez marre de corriger une par une les apostrophes droites, les espaces surnuméraires, les guillemets et toutes sortes d’erreurs récurrentes, cet outil est fait pour vous. Il est préférable de l’essayer en premier lieu sur des textes assez courts, afin de vérifier qu’il ne corrompt rien.

Note : l’outil touche aux espaces, mais pas aux tabulations.


## Grammalecte 0.2.8 (28 mai 2013)

- Toujours moins de faux positifs.
- Amélioration du préprocesseur de texte.
- Quelques bugs corrigés.
- Mise en place du système de notification de mise à jour prévu par LibreOffice/OpenOffice.


## Grammalecte 0.2.7 (1ᵉʳ avril 2013)

- Quelques bugs corrigés.
- Moins de faux positifs.
- Quelques améliorations diverses.


## Grammalecte 0.2.6 (11 mars 2013)

- Compatibilité avec LibreOffice 4.
- Quelques bugs corrigés.
- Quelques améliorations diverses.


## Grammalecte 0.2.5 (24 janvier 2013)

- Beaucoup, beaucoup moins de faux positifs.
- Un préprocesseur de texte encore amélioré.
- Quelques bugs corrigés.
- Un peu de nettoyage de code.

Cette version est compatible avec Python 3. Mais elle est vraiment compatible avec LibreOffice 4. J’ai mis longtemps à comprendre qu’un des problèmes concernant les espaces entre les phrases ne venait pas de mon code ou de Python 3, mais d’une modification non documentée de l’API de LibreOffice 4. Cette modification est peut-être un bug, j’attends toujours la réponse des dévs sur ce point. Si ce n’est pas un bug, je serai forcé de revoir en partie et peut-être en profondeur le code de Grammalecte, notamment le système à passages multiples que j’avais commencé à mettre en place.


## Grammalecte 0.2.4 (17 novembre 2012)

- Beaucoup de faux positifs en moins.
- Beaucoup d’améliorations dans le préprocesseur de texte.
- De nouvelles règles qui étendent la recherche des erreurs.
- Des bugs corrigés dans les messages rapportés.
- Un contournement d’un bug concernant Hunspell : sur les gros documents, au commencement, Hunspell est tellement sollicité pour vérifier l’orthographe qu’il ne donne pas toujours les informations demandées par le correcteur grammatical ; du coup, j’ai rendu le Grammalecte très insistant pour avoir les données dont il a besoin.
- Une option spécifique pour l’apostrophe typographique qui semble gêner pas mal de monde.


## Grammalecte 0.2.3 (07 octobre 2012)

- Nombreux faux positifs supprimés.


## Grammalecte 0.2.2 (02 octobre 2012)

- Suppression d’un faux positif rédhibitoire.


## Grammalecte 0.2.1 (30 septembre 2012)

- Quelques règles supplémentaires.
- Améliorations diverses.
- Corrections de bugs divers.
- Suppression de quelques faux positifs.
- Beaucoup d’améliorations sur le préprocesseur de texte.


## Grammalecte 0.2 (03 septembre 2012)

Les nouveautés se situent surtout sous le capot.


### Soulignement restreint à la zone d’erreur

Auparavant, Grammalecte soulignait presque toujours toute la zone qu’il avait dû examiner pour repérer une erreur. Par exemple, si vous écriviez « Le petit chaperon rouges », il soulignait « Le petit chaperon rouges » au lieu seulement de mettre en exergue « rouges ». C’était une contrainte technique. Il n’était possible de faire autrement qu’au prix d’une lourde chute des performances et d’une complexité plutôt rédhibitoire.
À présent, restreindre la zone signalée au mot erroné est aisé et sans conséquences sur les performances.

### Le préprocesseur de texte

Grammalecte, contrairement à LanguageTool et à Lightproof, analyse les textes que LibreOffice lui envoie en deux passes : paragraphe par paragraphe, puis phrase par phrase. Cet avantage permet d’introduire un nouvel outil, le préprocesseur de texte, dont le rôle est de modifier (en interne, pour le propre usage du correcteur uniquement) la phrase qu’il va analyser ; ces modifications consistent notamment en une simplification du texte, afin de le dépouiller de tout ce qui est gênant pour débusquer les erreurs.

Grammalecte 0.2 fonctionne ainsi :
- analyse du paragraphe (sans modification du texte) >> remontée des erreurs à LibreOffice,
- transformation du texte par le préprocesseur de texte,
- analyse phrase par phrase du texte modifié >> remontée des erreurs à LibreOffice.

L’épure du texte consiste notamment à ôter les adverbes, les locutions adverbiales, les expressions usuelles, les caractères indésirables qui ponctuent régulièrement les écrits.

Cette épure a plusieurs conséquences :

- elle simplifie drastiquement la désambiguïsation du texte (un souci permanent), et permet même souvent de se passer de faire de la désambiguïsation,
- elle réduit conséquemment le nombre de faux positifs,
- elle simplifie la recherche des erreurs et soulage le programmeur des maux de tête récurrents que la correction grammaticale automatique peut provoquer,
- elle permet la vérification, autrefois presque impossible, de certains accords distants.

Ce n’est pas la panacée, mais ça débroussaille déjà pas mal et ça permet vraiment d’y voir plus clair.

Il y a environ 140 règles dans le préprocesseur, qui opèrent un nombre plus grand encore de modifications diverses.


## Grammalecte 0.1.2 (15 janvier 2012)

- correction de bug critique.


## Grammalecte 0.1.1 (15 janvier 2012)

- Réduction drastique des faux positifs.
- Vérification des mots composés à trait d’union (désactivée par défaut).
- Meilleur contrôle des verbes à l’infinitif.
- De nouvelles règles (je n’ai toujours pas l’habitude de tenir un journal précis des modifications).
- Bugs mineurs corrigés.

Il y a 580 règles de contrôle.

Le contrôle des mots composés à traits d’union est désactivé par défaut, parce que c’est un nid de faux positifs potentiels. Ce contrôle est utile, car le correcteur orthographique, lorsqu’il ne reconnaît pas un mot composé, le scinde en deux et vérifie chaque partie séparément : si les deux parties existent dans le dictionnaire, il ne se passe rien. Avec la nouvelle option, le correcteur grammatical vérifie toutes les associations de mots comportant un trait d’union. Si le mot composé n’existe pas, le correcteur le signale, examine si la graphie soudée existe et la propose le cas échéant.

Prérequis : Dictionnaire français 4.4.1.


## Grammalecte 0.1 (03 octobre 2011)

- Beaucoup moins de faux positifs.
- Quelques règles supplémentaires.
- Divers bugs mineurs corrigés.
- Légère réorganisation des options (+ descriptif).
- Désactivation globale des règles de contrôle.

Prérequis : Dictionnaire français 4.3.


## Grammalecte 0.0.12 (04 août 2011)

- Les options sont dorénavant mémorisées sous Windows et Mac.
- La mécanique interne a été revue (analyse par paragraphe par paragraphe ou par phrase par phrase, au choix).
- Optimisation du code.
- Des règles supplémentaires.
- Des faux positifs supprimés.
- Des bugs concernant certaines règles ont été corrigés.


## Grammalecte 0.0.11 (04 juillet 2011)

Améliorations :

- le correcteur fonctionne dorénavant sur toutes les locales françaises (Belgique, France, Canada, Suisse, etc.),
- des règles supplémentaires,
- des faux positifs supprimés.

Pour rappel :

- Grammalecte est incompatible avec LanguageTool (qu’il faut désinstaller) ;
- cette extension a besoin de l’un des derniers dictionnaires français (4.2) pour fonctionner correctement ;
- écrite en Python, cette extension a besoin que la passerelle Python-UNO soit installée (c’est le cas par défaut sous Windows) ;
- depuis LibreOffice 3.4, il faut aussi que le « fournisseur de script pour Python » soit installé.


## Grammalecte 0.0.10 (10 avril 2011)

- Correction de nombreux faux positifs.


## Grammalecte 0.0.9 (28 mars 2011)

- Nouvelles règles.
- Correction de faux positifs.

Il est indispensable d'installer le nouveau dictionnaire 4.1 pour un fonctionnement normal.

Dorénavant, l’analyse se fait phrase par phrase au lieu de paragraphe par paragraphe. En effet, lorsque ceux-ci étaient longs (plusieurs pages), le correcteur ramait horriblement, au point de suspendre OpenOffice.org. Par sécurité, une phrase de plus de 1500 caractères ne sera pas analysée.


## Grammalecte 0.0.8 (06 février 2011)

- Améliorations diverses.


## Grammalecte 0.0.7 (04 février 2011)

- Correction d’un bug à propos des options.


## Grammalecte 0.0.x (janvier 2011)

Préversions du correcteur grammatical.

Grammalecte est un correcteur qui préférera rester discret s’il n’est pas possible de savoir avec une assez grande certitude qu’il s’agit d’une erreur. La ligne de conduite est « le moins de faux positifs possible ».

Il faut installer la dernière version des dictionnaires (v4+), dont l’étiquetage grammatical est indispensable au bon fonctionnement de cette extension.

Grammalecte considère que le bon usage consiste à utiliser les apostrophes typographiques. Activez donc le remplacement des apostrophes droites lors de la frappe, dans Menu Outils > Autocorrections > Options linguistiques.
