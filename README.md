Project context : 📽️

Part 1 ⬜️

  - For Hangman,you have 10 attempts to complete the game.

  - First, the program choose randomly a word in the file.

  - After that the programm reveal n random letters in the word, where n is the len(word) / 2 - 1

  - The program read the standard input to suggest a letter.

  - If the letter is not present, it print an error message and the number of attempts decreases (10->9->…0)

  - If the letter is present, it reveal all the letters corresponding in the word.

  - The program continues until the word is either found, or the numbers of attempts is 0.

Part 2 🟦

  - You have 10 positions corresponding to the 10 attempts needed to complete the game.

  - You must be need to parse the file at 10 positions and display the appropriate position of "José" as the count of attemps decreases.

Technology uses : 💻

  - Pour ce projet nous sommes sur le langage de programmation Golang.


Son système d'installation : 🏗️

    - Prendre le lien de la création de notre répository sur Gitea puis de le clone dans notre fichier VS-code.  

    - Créer 3 fonctions = Verticale , Horizontale, Diagonale.

Comment le lancer : 🏁

    - Grâce à GIT et ses commandes il nous suffis juste de taper la commande : "go run ("nom du notre package")".

    - Et grâce à ces commandes nous pouvons donc ajouter notre travail avec la commande : "git add . ", puis l'initialiser avec la commande : "git commit -m "notre_texte" et enfin l'envoyer avec la commande : "git push".