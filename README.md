# sqlite-go-connect

Simple go app to load the SQLite file (`movies.db`) included in this repo.

## Installation

1. Install go (tested with v1.11) and set $GOPATH
2. Install dep (tested with v0.5.0)
3. Run `dep ensure` to install dependencies

## Running

1. Run `go run main.go` to run the app
2. Make sure you see output like this
```
------------------------------------
Title: Sherlock Jr.
Director: Buster Keaton
Year: 1924
Runtime: 45
------------------------------------
Title: The General
Director: Clyde Bruckman, Buster Keaton
Year: 1926
Runtime: 67
------------------------------------
Title: The Kid
Director: Charles Chaplin
Year: 1921
Runtime: 68
------------------------------------
...
```