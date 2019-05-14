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

## How It Works
The balance checking algorithm uses stack at the core. The stack is simple First in Last out data structure.

The algorithm iterates over the given string. If it encounters an opening character, pushes the character to stack. On the other hand, if it gets a closing character, pops the top element from the stack and checks them for pairing. Eventually, if they aren't matching pairs, it raises `MismatchError`. If the encountered character is not a member of either the opening or closing sets, it raises `UnknownCharacterError`. End of the iteration, the stack should be empty. If not, there is an unclosed parenthesis. So it raises `UnclosedParenthesesError`

## HTTP Server
It's a simple HTTP server that shows example usage for `balance` package.

It checks parenthesis balance of given string and save request history to memory or MySQL database for calculating some metrics.

Visit `server` directory for [detailed documentation](https://github.com/EnesCakir/balance/tree/master/server).

![Dashboard](https://github.com/EnesCakir/balance/blob/master/dashboard.png)

## Testing
You can use `go test` method for unit tests.

**Test coverage:** 100%

## Contributing
`balance` is an open source project run by `Enes Çakır`, and contributions are welcome! Check out the [Issues](https://github.com/enescakir/balance/issues) page to see if your idea for a contribution has already been mentioned, and feel free to raise an issue or submit a pull request.

## License
Copyright (c) 2019 Enes Çakır. All rights reserved. Use of this source code is
governed by a MIT license that can be found in the LICENSE file.
