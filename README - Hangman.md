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

  - For this project we are using the Golang programming language.


Installation system  : 🏗️

    - Take the link to our Gitea repository and clone it in our VS-code file.  .  

    - Create 3 functions = Vertical, Horizontal, Diagonal.

How to launch : 🏁

    - Thanks to GIT and its commands, all we need to do is type the command: "go run ("name of our package")".

    - And thanks to these commands, we can add our work with the command: "git add . "then initialize it with the command: "git commit -m "notre_texte" and finally send it with the command: "git push". 