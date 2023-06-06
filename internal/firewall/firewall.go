package firewall

import "github.com/google/nftables"

const TableName = "catflap-filter"

func Init(debug bool) error {
	table := createTable()
	_ = table
	return nil
}

func createTable() *nftables.Table {
	return &nftables.Table{
		Name:   TableName,
		Use:    1, // ?
		Family: nftables.TableFamilyINet,
	}
}
