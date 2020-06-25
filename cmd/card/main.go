package main

import (
	"fmt"
	"github.com/DABronskikh/go-lesson-5/pkg/card"
	"github.com/DABronskikh/go-lesson-5/pkg/transfer"
)

func main() {
	nonameBank := card.NewService("Noname bank")
	visa := nonameBank.IssueCard("Visa", "RUB")
	master := nonameBank.IssueCard("MasterCard", "RUB")
	fmt.Println(visa, master)
	fmt.Println(nonameBank)

	nonameBankTransfer := transfer.NewService(nonameBank)
	nonameBankTransfer.IssueCommission(true, true, 0, 0)
	nonameBankTransfer.IssueCommission(true, false, 0.5, 10_00)
	nonameBankTransfer.IssueCommission(false, false, 1.5, 30_00)

	fmt.Println(nonameBankTransfer)
	total, ok, err := nonameBankTransfer.Card2Card(visa.Number, "0002", 1000_00)
	fmt.Println(total, ok, err)

	fmt.Println(card.IsValid("4276 1600 0000 0000"))
}
