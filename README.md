# balance ([{}])
    A package for validating the balance of parentheses

### Getting started
Install `balance`:
```shell
    go get github.com/enescakir/balance
```

Add `balance` to your imports to start using
```go
    import "github.com/enescakir/balance"
```


### Usage

```go

package main

import "github.com/enescakir/balance"

func main() {
	// Check given string for parentheses balance
	valid, err := balance.Check("{()[]}(())")
	// Returns: valid => true, err => nil

	valid, err := balance.Check("([)]")
	// Returns: valid => false, err => MismatchError

	valid, err := balance.Check("[[]")
	// Returns: valid => false, err => UnclosedParenthesesError

	valid, err := balance.Check("({a})")
	// Returns: valid => false, err => UnknownCharacterError
}

```

### HTTP Server

### Contributing

`balance` is an open source project run by `Enes Çakır`, and contributions are welcome! Check out the [Issues](https://github.com/enescakir/balance/issues) page to see if your idea for a contribution has already been mentioned, and feel free to raise an issue or submit a pull request.

### License
Copyright (c) 2019 Enes Çakır. All rights reserved. Use of this source code is
governed by a MIT license that can be found in the LICENSE file.
