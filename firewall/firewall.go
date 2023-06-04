package firewall

import "github.com/google/nftables"

func CreateTable() *nftables.Table {
	return &nftables.Table{
		Name:   "catflap",
		Use:    1,
		Family: nftables.TableFamilyINet,
	}
}
