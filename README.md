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

#### Endpoints
- **GET** /dashboard

Shows request history and some metrics
     
- **POST** /isbalanced
    

**Request:**
```json
// Example request body:
{
    "expr": "{[()]}"
}
```
    
## Contributing

`balance` is an open source project run by `Enes Çakır`, and contributions are welcome! Check out the [Issues](https://github.com/enescakir/balance/issues) page to see if your idea for a contribution has already been mentioned, and feel free to raise an issue or submit a pull request.

## License
Copyright (c) 2019 Enes Çakır. All rights reserved. Use of this source code is
governed by a MIT license that can be found in the LICENSE file.
