package card

import (
	"math/rand"
	"strconv"
)

type Card = struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
}

type Service struct {
	BankName string
	Cards    []*Card
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) IssueCard(issuer string, currency string) *Card {
	card := &Card{
		Issuer:   issuer,
		Balance:  50_000_00,
		Currency: currency,
		Number:   randNumber(),
		Icon:     "...",
	}

	s.Cards = append(s.Cards, card)
	return card
}

type Transaction = struct {
	Id     int64
	Amount int64
	Date   int64
	MCC    string
	status string
	Type   string
	Status string
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, *transaction)
}

func SumByMCC(transactions []Transaction, mcc []string) (total int64) {

	for _, transaction := range transactions {
		if isMCC(transaction.MCC, mcc) {
			total += transaction.Amount
		}
	}

	return total
}

func isMCC(mcc string, mccArr []string) bool {
	for _, candidate := range mccArr {
		if candidate == mcc {
			return true
		}
	}
	return false
}

func (s *Service) SearchByNumber(number string) *Card {
	for _, card := range s.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}

func randNumber() (number string) {
	return strconv.Itoa(rand.Intn(100))
}
