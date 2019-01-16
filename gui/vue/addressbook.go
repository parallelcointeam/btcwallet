package vue

import (
	"encoding/json"
	"fmt"

	"github.com/parallelcointeam/mod/gui/jdb"
)

// var VAB []AddressBook

type AddressBookLabel struct {
	Label   string `json:"label"`
	Address string `json:"address"`
}
type AddressBook struct {
	AddressBookLabel []AddressBookLabel `json:"label"`
}

func (ab *AddressBook) AddressBookData() {
	addressbooks, err := jdb.JDB.ReadAll("addressbook")
	if err != nil {
		fmt.Println("Error", err)
	}

	for _, f := range addressbooks {
		var addressbook AddressBookLabel
		if err := json.Unmarshal([]byte(f), &addressbook); err != nil {
			fmt.Println("Error", err)
		}
		ab.AddressBookLabel = append(ab.AddressBookLabel, addressbook)
	}
	fmt.Println("Ersssssssssssssssssssssssssssror", ab.AddressBookLabel)

}
func (ab *AddressBookLabel) AddressBookLabelWrite(label, address string) {
	ab.Label = label
	ab.Address = address
	jdb.JDB.Write("addressbook", ab.Label, ab)
	fmt.Println("Ersssssssssssssssssssssssssssror", ab)

}
