# Number Guessing Game #
Sample solution for the [number guessing game](https://roadmap.sh/projects/number-guessing-game) challenge from [roadmap.sh](https://roadmap.sh/). Number guessing is a game where the computer randomly selects a number and the user has to guess it. The user will be given a limited number of chances to guess the number. If the user guesses the number correctly, the game will end, and the user will win. Otherwise, the game will continue until the user runs out of chances.. The application run from the command line interface (CLI).

## How run the application?
1. Download [Go language](https://go.dev), setup system env variable & check installation ``go version``.
2. Download **ZIP** file or use ``git clone https://github.com/mrizkisaputra/backend-projects.git``.
3. Open terminal, navigate to directory project **number-guessing-game** ``cd mrizkisaputra-backend-projects/number-guessing-game``.
4. Run ``go install``

## Example
```shell
# running the installation result
number-guessing-game


Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have a chances base on the difficulty level
to guess the correct number.
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
---------------------------------------
Enter your choice: 2
Great! You have selected the Medium difficulty level.
Let's start the game!

Enter your guess: 25
Incorrect! The number is less than 50.

Enter your guess: 25
Incorrect! The number is greater than 25.

Enter your guess: 35
Incorrect! The number is less than 35.

Enter your guess: 30
Congratulations! You guessed the correct number in 4 attempts.
the time you need is 0.7seconds.


you want play again (y/n):


```