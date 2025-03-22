package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/secrets/v2/op"
	"os"
	"strings"
)

const (
	version = "0.0.1"
)

func main() {
	emptyString := func(v string) bool {
		return strings.TrimSpace(v) == ""
	}
	terminateIfTrue := func(condition bool, exitCode int, message string) {
		if condition {
			fmt.Printf("%s\n", message)
			os.Exit(exitCode)
		}
	}
	versionFlag := flag.Bool("version", false, "Show version")
	listVaults := flag.Bool("listVaults", false, "List the current vaults")
	listItems := flag.Bool("listItems", false, "list the item from a named vault")
	getItem := flag.Bool("getItem", false, "get item from a vault")
	vaultName := flag.String("vault", "", "specifies the name of the vault")
	itemName := flag.String("item", "", "specifies the name of the vault item")
	flag.Parse()

	terminateIfTrue(*versionFlag, 0, version)
	terminateIfTrue(*listItems && emptyString(*vaultName), 0, "missing vault name")
	terminateIfTrue(*getItem && (emptyString(*vaultName) || emptyString(*itemName)), 0, "missing item name or vault name")

	if err := op.Authenticate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *listVaults {
		out, err := op.ListVault()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(out)
	}

	if *listItems {
		out, err := op.ListItems(*vaultName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(out)
	}

	if *getItem {
		out, err := op.GetItem(*vaultName, *itemName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(out)
	}
}
