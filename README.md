# Battleship API

## Project overview

This is an API designed to allow two players to play a game of battleship over the internet. The game takes place on a 10x10 grid with each space numbered from 1-100, starting with 1 at the top left and ending with 100 at the bottom right.

There are 5 ship types:

- Carrier (Length of 5 spaces)
- Battleship (Length of 4 spaces)
- Cruiser (Length of 3 spaces)
- Submarine (Length of 3 spaces)
- Destroyer (Length of 2 spaces)

## Installation and setup

If you don't have Go already installed you can either run

```
$ brew install go
```

Or you can go [here](https://go.dev/doc/install) and follow the intructions to install it.

Once Go is installed we can get the project setup

```
$ git clone https://github.com/Chrixs/battleship
$ cd {yourprojectsdirectory}/battleship

$ go mod tidy
$ go build && go run server.go
```

This will start a server at http://localhost:1323

Tests can be run from the top level directory with

```
$ go test ./...
```

## API Usage

All request data should be Json formated

### Get players

`GET http://localhost:1323/players` displays all players and their related data.

### Deploy Ship

A deploy request would look something like this<br>
`PUT http://localhost:1323/player/{id}/deploy`

```
{
    "shipType": "Battleship",
    "coordinate": 87,
    "isVertical": false
}
```

Or with `curl`

```
curl --header "Content-Type: application/json" --request PUT --data '{"shipType":"Carrier","coordinate":1,"isVertical":false}' http://localhost:1323/player/1/deploy
```

which should return the deployed ship object.

```
{
    "status": 200,
    "success": true,
    "message": "Battleship sucessfully deployed",
    "data": {
        "type": "Battleship",
        "length": 4,
        "health": 4,
        "coordinates": [
            87,
            88,
            89,
            90
        ]
    }
}
```

This will place a ship at the coordinate specified plus cover the grid squares to the right of the coordinate if `isVertical` is `false` or vertically downwards if `true`, to the length of the ship type. Ships cannot exceed the 10x10 grid nor overlap eachother.

### Fire

A fire request would look like this. <br>
`PUT http://localhost:1323/player/{id}/fire`

```
{
    "coordinate": 50
}
```

Or with `curl`

```
curl --header "Content-Type: application/json" --request PUT --data '{"coordinate":50}' http://localhost:1323/player/1/fire
```

Where the player ID is the attacking player.
On a miss or a hit you should expect something like this.

```
{
    "status": 200,
    "success": true,
    "data": {
        "Status": "hit"
    }
}
```

If all enemy player ships are destroyed after a successful shot the response will indicate your win.

```
{
    "status": 200,
    "success": true,
    "data": {
        "status": "sunk",
        "shipType": "Destroyer",
        "winner": true
    }
}
```

### Reset

The game state can be reset with <br>
`GET http://localhost:1323/reset`
