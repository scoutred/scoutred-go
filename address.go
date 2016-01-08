package scoutred

import (
	"fmt"
	"strings"
)

type Address struct {
	Number        *int64  `json:"number"`
	PreDirection  *string `json:"preDirection"`
	Street        *string `json:"street"`
	Suffix        *string `json:"suffix"`
	PostDirection *string `json:"postDirection"`
	Fraction      *string `json:"fraction"`
	Unit          *string `json:"unit"`
	Postal        *string `json:"postal"`
	Jurisdiction  *string `json:"jurisdiction"`
	State         *string `json:"state"`
	Country       *string `json:"country"`
}

//	returns a formatted address as a single string
func (this Address) Format() (addr string) {
	if this.Number != nil {
		//	if 0, don't format (likely raw land)
		if *this.Number != 0 {
			//	convert from a float64, to an int then cast to a string
			addr += fmt.Sprintf("%v", int(*this.Number)) + " "
		}
	}

	if this.Fraction != nil {
		addr += *this.Fraction + " "
	}

	if this.PreDirection != nil {
		addr += *this.PreDirection + " "
	}

	if this.Street != nil {
		//	for addresses that start with a "0" we need to
		//	pop it off. i.e 08th street
		if strings.HasPrefix(*this.Street, "0") {
			street := *this.Street
			//	pop off the first char
			*this.Street = street[1:]
		}
		addr += *this.Street + " "
	} else {
		addr += "No address"
	}

	if this.Suffix != nil {
		addr += *this.Suffix + " "
	}

	if this.PostDirection != nil {
		addr += *this.PostDirection + " "
	}

	if this.Unit != nil {
		addr += "#" + *this.Unit + " "
	}

	//	clean up trailing space if it exists
	addr = strings.TrimSpace(addr)

	//	format our address with a captial first letter of each word
	addr = strings.Title(strings.ToLower(addr))

	return
}
