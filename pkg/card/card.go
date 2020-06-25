package card

import (
	"math/rand"
	"strconv"
	"strings"
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

func (s *Service) FindByNumber(number string) (*Card, bool) {
	for _, c := range s.Cards {
		if c.Number == number {
			return c, true
		}
	}

	// если карта не найдена, но ее номер начинается с "5106 21" создадим произвольную
	//с нужным префиксом для сервиса и вернем
	if strings.HasPrefix(number, "5106 21") {
		c := s.IssueCard("MasterCard", "RUB")
		return c, true
	}
	return nil, false
}

func (s *Service) FindById(id int64) (*Card, bool) {
	for _, c := range s.Cards {
		if c.Id == id {
			return c, true
		}
	}
	return nil, false
}

func randNumber() string {
	num := strconv.Itoa(rand.Intn(99))
	num = "5106 21" + num + " 0000 0000"

	return num
}

func IsValid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	numbersStr := strings.Split(number, "")

	start := len(numbersStr) % 2
	sum := 0

	for i, v := range numbersStr {
		number, err := strconv.Atoi(v)
		if err != nil {
			return false
		}

		if i%2 == start {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}
		sum += number
	}

	if sum%10 != 0 {
		return false
	}

	return true
}
