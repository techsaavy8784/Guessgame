# Game Overview
In this game, a player, referred to as the "game creator," initiates a game by setting a secret number, a reward, and a participation fee. Other players, known as
"guessers," attempt to guess this number within a specified range and timeframe.
The guessers who come closest to the secret number win a share of the total reward.

## Assumptions

Only the game creator have access to the secret number.

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### How to create a game

```
gamed tx game create-game [secretNumber] [reward] [entryFee] [duration] --from [creator]
```
### How to submit a guess

```
gamed tx game submit-guess [gameId] [guess] --from [guesser]
```

### How to end a game

```
gamed tx game end-game [gameId] --from [creator]
```

### How to query the current game status

```
gamed q game list-game
```

```
gamed q bank balances $(gamed keys show [player] -a)
```