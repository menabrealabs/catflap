package firewall

import "github.com/google/nftables"

func createTable() *nftables.Table {
	return &nftables.Table{
		Name:   "catflap",
		Use:    1,
		Family: nftables.TableFamilyINet,
	}
}
