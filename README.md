<p align="center">
	<img width="560" height="100" src="https://github.com/EnesCakir/balance/blob/master/logo.png">
	<br> <br>
    A package for validating the balance of parentheses
</p>

## Getting started
Install `balance`:
```shell
$ go get github.com/enescakir/balance
```

Add `balance` to your imports to start using
```go
import "github.com/enescakir/balance"
```


## Usage

#### `Check(str string) (valid bool, err error)`
It checks given string for parentheses balance for `{}`, `()`, `[]` pairs

```go
valid, err := balance.Check("{()[]}(())")
// Returns: valid => true, err => nil

valid, err := balance.Check("([)]")
// Returns: valid => false, err => MismatchError

valid, err := balance.Check("[[]")
// Returns: valid => false, err => UnclosedParenthesesError

valid, err := balance.Check("({a})")
// Returns: valid => false, err => UnknownCharacterError
```

#### `CheckCustom(str string, opens string, closes string) (bool, err)`
It checks given string for parentheses balance for custom pairs.

`opens` and `closes` strings should have pair elements in same order.

Given pair elements have to be unique. `CheckCustom` function doesn't work properly without unique elements.

```go
valid, err := balance.CheckCustom("<<>><>", "<", ">")
// Returns: valid => true, err => nil

valid, err := balance.CheckCustom(")))()(((", ")", "(")
// Returns: valid => true, err => nil

valid, err := balance.CheckCustom("<><><>", "<<", ">")
// Returns: valid => false, err => CustomPairError
```

## HTTP Server
It's a simple HTTP server that shows example usage for `balance` package.

It checks parenthesis balance of given and save request history to MySQL database for calculating some metrics.

### Usage
```shell
# Enter to server directory
$ cd server

# Copy config.example.json to config.json
$ cp config.example.json config.json

# Put database credentials into config.json

# Download goland MySQL driver
$ go get -u github.com/go-sql-driver/mysql

# Run server 
$ go run main.go

# Visit `http://localhost:8080/`
```

#### Docker
You can use balance server with Docker

*Notice:* `Dockerfile` is at project root, not the server root. Because this repository isn't public right now. 
When we put it to the server root, it can't download `github.com/enescakir/balance` package.

```shell
###
# Running balance server with MySQL server
$ docker-compose -f docker-compose.yml up balance

# Visit `http://localhost:8080/`

###
# Running balance server with another MySQL server
# Build docker image
$ docker build -t balance .

# Run docker image with your MySQL database credentials
$ docker run -it -p 8080:8080 \
    -e DATABASE_HOST=127.0.0.1 \
    -e DATABASE_PORT=3306 \
    -e DATABASE_NAME=balance \
    -e DATABASE_USER=root \
    -e DATABASE_PASSWORD=secret \
    balance
```

### Endpoints     
**POST** /isbalanced

Checks the parentheses balance of given expression 
```json
// Example request body:
{
    "expr": "{[()]}"
}
```
```json
// Response for "{[()]}":
{
    "valid": true
}

// Response for "[())]":
{
    "valid": false,
    "error": "Mismatch at index: 3"
}
```

**GET** /logs

Returns the collection of logs at given date range.

It accepts `start` and `end` as query parameter. They are optional. 
*Example:* `/logs?start=2019-05-11+12:17:00`

```json
// Example response body
[
  {
    "id": 148,
    "query": "[()]",
    "status": 1,
    "response_time": 230647,
    "created_at": "2019-05-12T12:58:52Z"
  }, {
    "id": 147,
    "query": "[(()]",
    "status": 2,
    "response_time": 89170,
    "created_at": "2019-05-12T12:57:39Z"
   }
]
```

**GET** /logs/status

Returns the collection of status:count pairs at given date range.

It accepts `start` and `end` as query parameter. They are optional. 
*Example:* `/logs/status?start=2019-05-11+12:17:00`

```json
// Example response body
[
    {
        "status": 1,
        "count": 6
    },
    {
        "status": 2,
        "count": 2
    }
]
```

**GET** /logs/histogram

Returns the collection of label:responseTime bins at given date range.
Labels are milliseconds.

It accepts `start` and `end` as query parameter. They are optional. 
*Example:* `/logs/histogram?start=2019-05-11+12:17:00`

```json
// Example response body
[
    {
        "label": "0-10",
        "count": 3
    },
    {
        "label": "10-20",
        "count": 4
    },
    {
        "label": "20-30",
        "count": 2
    },
    {
        "label": "30-40",
        "count": 3
    },
    {
        "label": "40-50",
        "count": 5
    }
]
```

## Contributing

`balance` is an open source project run by `Enes Çakır`, and contributions are welcome! Check out the [Issues](https://github.com/enescakir/balance/issues) page to see if your idea for a contribution has already been mentioned, and feel free to raise an issue or submit a pull request.

## License
Copyright (c) 2019 Enes Çakır. All rights reserved. Use of this source code is
governed by a MIT license that can be found in the LICENSE file.
