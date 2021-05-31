# Govalent

govalent is a go client library for Covalent Rest APIs

## Installation

```sh
go get github.com/AlchemistsLab/govalent
```

### Usage without a Client

If you are dealing with one account. There is no need to create a new client. you can simply call `govalent.$resource$()`

```go
import (
	"fmt"
	"github.com/AlchemistsLab/govalent"
)

// Setup
govalent.APIKey = ""
info, err := govalent.ClassA().HistoricalPortfolio("56", "0xb1b3f0e569a19E407cEb7bFAEA3486F0D9d2488B")
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
    "fmt"
    "github.com/AlchemistsLab/govalent"
)
client := govalent.Client{}
client.Init("YOUR_API_KEY")
info, err := client.ClassA.HistoricalPortfolio("56", "0xb1b3f0e569a19E407cEb7bFAEA3486F0D9d2488B")
if err != nil {
	fmt.Printf("err = %v", err)
	return
}
fmt.Printf("%v", info)
```

### Class A endpoints

#### Get Chains

```go
import (
	"fmt"
	"github.com/AlchemistsLab/govalent"
)

func main() {
	govalent.APIKey = ""
	chains, err := govalent.ClassA().Chains()
	if err != nil {
		fmt.Printf("err = %v", err)
		return
	}
	fmt.Println("Get All chains")
	for _, chain := range chains.Items {
		fmt.Printf("Chain[%+v/%v]\n", chain.ID, chain.Name)
	}
	fmt.Println("Get All chain statuses")
	chainStatus, err := govalent.ClassA().ChainsStatus()
	if err != nil {
		fmt.Printf("err = %v", err)
		return
	}
	for _, chain := range chainStatus.Items {
		fmt.Printf("Chain[%v/%v], BlockHeight: %v\n", chain.ID, chain.Name, chain.SyncedBlockHeight)
	}
}
```

#### Get Token Balances 

```go
import (
    "fmt"
    "github.com/AlchemistsLab/govalent"
    "github.com/AlchemistsLab/govalent/class_a"
)

balanceParams := class_a.BalanceParams{
	Nft: true,
}

p, err := govalent.ClassA().TokenBalances("56", "0xb1b3f0e569a19E407cEb7bFAEA3486F0D9d2488B", balanceParams)
if err != nil {
	fmt.Printf("err = %v", err)
	return
}
fmt.Printf("%v", p)
```

#### Get Historical Portfolio

```go
import (
    "fmt"
    "github.com/AlchemistsLab/govalent"
)

p, err := govalent.ClassA().HistoricalPortfolio("56", "0xb1b3f0e569a19E407cEb7bFAEA3486F0D9d2488B")
if err != nil {
	fmt.Printf("err = %v", err)
	return
}
fmt.Printf("%v", p)
```

#### Get Transactions

```go
import (
    "fmt"
    "github.com/AlchemistsLab/govalent"
)

p, err := govalent.ClassA().Transactions("56", "0xb1b3f0e569a19E407cEb7bFAEA3486F0D9d2488B")
if err != nil {
	fmt.Printf("err = %v", err)
	return
}
fmt.Printf("%v", p)
```

#### Get ERC20 token transfers

```go
import (
    "fmt"
    "github.com/AlchemistsLab/govalent"
    "github.com/AlchemistsLab/govalent/class_a"
)

params := class_a.TransferParams{
    ContractAddress: "0x8a0C542bA7bBBab7cF3551fFcc546CdC5362d2a1",
}
p, err := govalent.ClassA().ERCTokenTransfers("56", "0xb1b3f0e569a19E407cEb7bFAEA3486F0D9d2488B", params)
if err != nil {
    fmt.Printf("err = %v", err)
    return
}
fmt.Printf("%v", p)
```

#### Get Log Events by contract address.

```go
import (
    "fmt"
    "github.com/AlchemistsLab/govalent"
    "github.com/AlchemistsLab/govalent/class_a"
)

params := class_a.LogEventsParams{
    StartingBlock: "9601459",
    EndingBlock: "9999800",
}
p, err := govalent.ClassA().LogEventsByContract("1", "0xc0da01a04c3f3e0be433606045bb7017a7323e38", params)
if err != nil {
    fmt.Printf("err = %v", err)
    return
}
fmt.Printf("%v", p)
```

### Class B endpoints

#### Get Sushiswap address exchange liquidity transactions

```go
import (
	"fmt"
	"github.com/AlchemistsLab/govalent"
	"github.com/AlchemistsLab/govalent/class_b"
)

func main() {
	govalent.APIKey = "ckey_95c202a743e24270bc7f1706c4c"
	params := class_b.SushiSwapActsParams{
		Swaps: true,
	}
	acts, err := govalent.ClassB().SushiSwapActs("137", "0x4121dD930B15742b6d2e89B41284A79320bb8503", params)
	if err != nil {
		fmt.Printf("err = %v", err)
		return
	}
	fmt.Printf("Get sushiswap act for address: %v, at:%v\n", acts.Address, acts.UpdatedAt)
	for _, a := range acts.Items {
		fmt.Printf("Act[%v]: %v. At:%v\n", a.Act, a.Description, a.ActAt)
	}
}
```

#### TODO