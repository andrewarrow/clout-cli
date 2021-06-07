package main

import (
	"clout/display"
	"clout/keys"
	"clout/models"
	"clout/network"
	"clout/session"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func HandleBalances(argMap map[string]string) {
	if argMap["words"] != "" {
		StatsOnWords()
		return
	}
	m := session.ReadAccounts()
	for username, s := range m {
		fmt.Println("")
		pub58, _ := keys.ComputeKeysFromSeed(session.SeedBytes(s))
		fmt.Println(username, pub58)
		user := session.Pub58ToUser(pub58)
		points := user.ProfileEntryResponse.CoinEntry.CreatorBasisPoints
		total := user.ProfileEntryResponse.CoinEntry.CoinsInCirculationNanos

		fmt.Printf("  %s %.02f\n", display.LeftAligned("BalanceNano", 20), float64(user.BalanceNanos)/1000000.0)
		fmt.Printf("  %s %.02f\n", display.LeftAligned("MarketCap", 20), user.ProfileEntryResponse.MarketCap())
		fmt.Printf("  %s %d\n", display.LeftAligned("Points", 20), points)
		fmt.Printf("  %s %s\n", display.LeftAligned("Price", 20), display.OneE9(user.ProfileEntryResponse.CoinPriceBitCloutNanos))

		sort.SliceStable(user.UsersWhoHODLYou, func(i, j int) bool {
			return user.UsersWhoHODLYou[i].BalanceNanos >
				user.UsersWhoHODLYou[j].BalanceNanos
		})
		for i, friend := range user.UsersWhoHODLYou {
			coins := float64(friend.BalanceNanos) / 1000000000.0
			username := friend.ProfileEntryResponse.Username
			if username == "" {
				username = "anonymous"
			}
			fmt.Printf("  %s %0.6f %0.6f\n",
				display.LeftAligned(username, 30), coins,
				float64(friend.BalanceNanos)/float64(total))
			if username == "anonymous" {
				fmt.Println(friend.HODLerPublicKeyBase58Check)
			}

			if i > 9 {
				break
			}
		}
	}
	fmt.Println("")
}

func StatsOnWords() {
	words := `betray BC1YLhKnJ2ajDNenRJaddF8777JBFx4XRXiuJnUdeyamDbELceF1rJo
chuckle BC1YLigwgGih65vUnpHHA2QTeDUYEcgGpj6f3qFKWN3vwdacQi82xGt
embark BC1YLiGAJEv7aKPhTDMrCZUw7snZVhurSCfbzwGTGt9N8rSFsEibzvN
erupt BC1YLfmkDQwchrXoKKqxed21L8HdKkZNLevQ7vvwixjR8L49cZkJDJz
oblige BC1YLhJnSGAZkP9BVeCvhLQpy1oKN8MxRg413uvVZ9c5h3cvSK9ty9U
situate BC1YLjJKthX8fGigKDFi9PvFe2ZDPByVrHbhTZBFzH23s1cypvYzgA6
stumble BC1YLgSdQqSRKLnSUgY3AL1yauQPYSu1CjFUwr9DrkqSGfPYMqJDn5b
crumble BC1YLhgzgXwFWWVhHQdAQnk25vQZvfQh8ePDDasivb7eby57fxrDgN9
enrich BC1YLgdpjLNf96dCvBpa9X9eTTdMDxreTs6Z5sWC2b4vQ1L1SAsmeEP
inflict BC1YLhYtvoRp67jp6PF9rVaVekYRuwByLMLTVBYHgSQLEyUSdB3XZD9
inherit BC1YLjNYBu8TfDg9mUF55595kgEGSRSF5cUnABBazDCr8uDRKh9Z949
unveil BC1YLh4a8fE845Lb6bGttiqbtYGGr9ZUgEatCZD1Pxa2pUx1bZYvJjS
enact BC1YLgfL44oJ9G9vV97wBUyGfCsFmpbhRm4cD9rN3hKP5355BCXH8ew
uphold BC1YLicrZdwMv1kA2XqvuLyikQ1BMWDWKv2LA7HFFuweNcCSB8i2F3Y
vanish BC1YLh2jM74p7KgPoJNkidKWANepshK9eh9DPXikfUsjUchSKzVShoh
defy BC1YLfoGzuJ6AZevMW9ME4PqGfuNu48SVdaEc76byBC6mpD12JYXnU7
demise BC1YLgKX4Ghku4uyQCJSF6FvbPznoxG4zpA6krhzjWN4sUCsuBDJKFG
dizzy BC1YLgQLdx5C7Gacctz9eBWt5zbqLdUbEBjoTEmRZzzFxNoTkZCkmh7
enlist BC1YLhvPCJXtFsfTT4hkrmzFLcAnjGLqhfK7uRNbY7rEdR293rRUjTi
eyebrow BC1YLhDDjp5Q8vBFDezRgpiET3x51xCNvfBQPY8pvXdtxHJzJC3GutR
gasp BC1YLiMdUG17JzSu5Fsz7vvyCQpmgboG9zfXiXmdAEwR7XuNQcsKMGQ
imitate BC1YLh3wUFTNqZsgxPiENgxV1cgptXMvDYwNyA7fHc4sQHUiSRVMef7
bleak BC1YLg3BaGRa9FMiWodB8K7PNrQz9tcKERybmTp9TVhJfZConLCo4QD
blouse BC1YLgRSFnhoi8UFoZNNPUQCCAWnTVwNDyn7qFN3kPiGmtshM5nYTbC
divert BC1YLgD1FTWWWQSjfWf1ZnxCJFEN9PP2WUdpocZQSY3nZVXMT9umPAm
evoke BC1YLgC2EaUe57SP86naqzoQWUMNuFDkvjeUVUR8yWQ3h6xCxieCXYW
grunt BC1YLic3M4SrXXvcRTQRY3yVbnxnwuB3M3RKELW6KMoJZeJEPd4xsmA
inhale BC1YLh2CAdb8iyH1cXrRBsMQRoKSjLg6Y2bvzDSTCex4hKNeUnUE1Ec
unfold BC1YLhdeafGA5A6s7scBwjzRY6Pq8vANpbijQYEZDrpwrXPd1SPh1tN
churn BC1YLgfah2sLSLGMimh4NvZDCnRxBtNhYWbueFvuQ7sCxbbFfyVuUVG
frown BC1YLh53wKpN8Do3trket9mMFhbCohobTYfGWXZ1pDKVcHcRcr4rbtR
ostrich BC1YLj7E4kaLAnN9SEUfzFEP69A4TocyoJ823Zp37VPRb8QsRBZGwGK
scatter BC1YLiCo6prb6M3xELpRbHUAtQvNAegcr2GHg1Z9LYDL52cZrbctHmr
cupboard BC1YLg4RaS4WcV2sYgn3gppgVQqviXCHQf8CXfy34khfuSqSq22hndU
mammal BC1YLhPUDQ7biXQjUUem2EPCE7rG8NhvTCozdrFM5xNaAADjZDLgnEX
reopen BC1YLjWn4ZryRMqqJG32QVBp21nnEy7gq5BY7B2MkJSjrtGVgkMw64F
topple BC1YLgPxFU6JWGTvM3M2YBzyP7C4tgYaSLjxX5USSsWyFqUBJX6JVMg
clog BC1YLjGW9cdFTNfuJhqsLGAm1euda9CKTEZpGJL1AehvtTLcd7SEvxS
drastic BC1YLhoDCQmfYBiYMrdypXU8NY5y3S3mzdGUSJHcyrjCk2ZSRoG4JS9
shiver BC1YLg7QooELwa4T2MXjsnryiz9ZyF62vs6d29AKzbjzvvQiewA7H7N`

	list := []string{}
	for _, line := range strings.Split(words, "\n") {
		tokens := strings.Split(line, " ")
		list = append(list, tokens[1])
	}

	js := network.GetManyUsersStateless(list)

	var us models.UsersStateless
	json.Unmarshal([]byte(js), &us)
	fields := []string{"username", "cap", "holders", "price"}
	sizes := []int{20, 10, 10, 10}
	display.Header(sizes, fields...)
	sort.SliceStable(us.UserList, func(i, j int) bool {
		return us.UserList[i].ProfileEntryResponse.CoinPriceBitCloutNanos >
			us.UserList[j].ProfileEntryResponse.CoinPriceBitCloutNanos
	})
	for _, u := range us.UserList {
		marketCap := u.ProfileEntryResponse.MarketCap()
		display.Row(sizes, u.ProfileEntryResponse.Username,
			fmt.Sprintf("%.02f", marketCap),
			len(u.UsersWhoHODLYou),
			display.OneE9(u.ProfileEntryResponse.CoinPriceBitCloutNanos))
	}
}
