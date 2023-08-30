package main

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/relayer/v2/cmd"
)

func main() {
	cmd.Execute()
}

func init() {
	sdk.SetAddrCacheEnabled(false)
}
