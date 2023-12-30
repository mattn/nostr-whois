package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/nostr-sdk"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "relay list"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func keys(m map[string]string) []string {
	a := []string{}
	for k := range m {
		a = append(a, k)
	}
	return a
}

type profile struct {
	Website     string `json:"website"`
	Nip05       string `json:"nip05"`
	Picture     string `json:"picture"`
	Lud16       string `json:"lud16"`
	DisplayName string `json:"display_name"`
	About       string `json:"about"`
	Name        string `json:"name"`
}

func main() {
	var relays arrayFlags = []string{"wss://relay.nostr.band"}
	var j bool
	flag.Var(&relays, "relay", "relays to connect")
	flag.BoolVar(&j, "json", false, "output JSON")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	for _, arg := range flag.Args() {
		var pub string
		if pp := sdk.InputToProfile(context.TODO(), arg); pp != nil {
			pub = pp.PublicKey
		} else {
			log.Printf("failed to parse pubkey from %v", arg)
			continue
		}
		ctx := context.Background()
		pool := nostr.NewSimplePool(ctx)
		ev := pool.QuerySingle(ctx, relays, nostr.Filter{
			Kinds:   []int{nostr.KindProfileMetadata},
			Authors: []string{pub},
			Limit:   1,
		})

		if ev == nil {
			log.Printf("failed to query event for %v", arg)
			continue
		}

		if j {
			fmt.Println(ev.Content)
		} else {
			var p profile
			err := json.Unmarshal([]byte(ev.Content), &p)
			if ev == nil {
				log.Fatal(err)
			}
			fmt.Printf("Pubkey: %v\n", pub)
			fmt.Printf("Name: %v\n", p.Name)
			fmt.Printf("DisplayName: %v\n", p.DisplayName)
			fmt.Printf("WebSite: %v\n", p.Website)
			fmt.Printf("Picture: %v\n", p.Picture)
			fmt.Printf("NIP-05: %v\n", p.Nip05)
			fmt.Printf("LUD-16: %v\n", p.Lud16)
			fmt.Printf("About: %v\n", p.About)
			fmt.Println()
		}
	}
}
