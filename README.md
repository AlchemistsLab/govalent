# Govalent

govalent is a go client library for Covalent Rest APIs

## Installation

```sh
go get github.com/zaebee/govalent
```

### Usage without a Client

If you are dealing with one account. There is no need to create a new client. you can simply call `govalent.$resource$()`

```go
import (
	"github.com/zaebee/govalent"
)

// Setup
govalent.APIKey = ""
info, err := govalent.ClassA().GetHistoricalPortfolio("56", "0xb1b3f0e569a19E407cEb7bFAEA3486F0D9d2488B")
if err != nil {
	fmt.Printf("err = %v", err)
	return
}
fmt.Printf("%v", info)
```

### Usage with a Client

If you are dealing with multiple accounts. You can create a new `govalent.Client` by the following

```go
import (
    "github.com/zaebee/govalent"
)
covalentClient := govalent.Client{}
covalentClient.Init("YOUR_API_KEY")
info, err := covalentClient.ClassA.GetHistoricalPortfolio("56", "0xb1b3f0e569a19E407cEb7bFAEA3486F0D9d2488B")
if err != nil {
	fmt.Printf("err = %v", err)
	return
}
fmt.Printf("%v", info)
```

### Class A endpoints

#### Get Token Balances 

```go
import (
	"github.com/zaebee/govalent"
	"github.com/zaebee/govalent/class_a"
)

balanceParams := class_a.BalanceParams{
	Nft: true,
}

p, err := govalent.ClassA().GetTokenBalances("56", "0xb1b3f0e569a19E407cEb7bFAEA3486F0D9d2488B", balanceParams)
if err != nil {
	fmt.Printf("err = %v", err)
	return
}
fmt.Printf("%v", p)
```

#### Get Historical Portfolio

```go
import (
    "github.com/zaebee/govalent"
    "github.com/zaebee/govalent/class_a"
)

p, err := govalent.ClassA().GetHistoricalPortfolio("56", "0xb1b3f0e569a19E407cEb7bFAEA3486F0D9d2488B")
if err != nil {
	fmt.Printf("err = %v", err)
	return
}
fmt.Printf("%v", p)
```

#### Get ERC20 token transfers

```go
import (
    "github.com/zaebee/govalent"
)
// TODO
```

### Class B endpoints

TODO
