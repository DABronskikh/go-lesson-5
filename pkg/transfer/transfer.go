package transfer

import (
	"errors"
	"github.com/DABronskikh/go-lesson-5/pkg/card"
)

var (
	ErrNotEnoughFundsAccount = errors.New("not enough funds account")
	ErrSourceCardNotFound    = errors.New("source card not found")
	ErrTargetCardNotFound    = errors.New("target card not found")
	ErrInvalidCardNumber     = errors.New("err invalid card number")
)

type Commission struct {
	from       bool
	to         bool
	percentage float64
	minAmount  int64
}

type Service struct {
	CardSvc    *card.Service
	Commission []*Commission
}

func NewService(cardSvc *card.Service) *Service {
	return &Service{
		CardSvc: cardSvc,
	}
}

func (s *Service) IssueCommission(from, to bool, percentage float64, minAmount int64) {
	commission := &Commission{
		from:       from,
		to:         to,
		percentage: percentage,
		minAmount:  minAmount,
	}
	s.Commission = append(s.Commission, commission)
}

func (s *Service) Card2Card(from, to string, amount int64) (total int64, ok bool, err error) {
	if !card.IsValid(from) || !card.IsValid(to) {
		return 0, false, ErrInvalidCardNumber
	}

	fromCard, fromBool := s.CardSvc.FindByNumber(from)
	if !fromBool {
		return 0, false, ErrSourceCardNotFound
	}

	toCard, toBool := s.CardSvc.FindByNumber(to)
	if !toBool {
		return 0, false, ErrTargetCardNotFound
	}

	commission := s.searchCommission(fromBool, toBool)
	percentage := commission.percentage
	minAmount := commission.minAmount

	var sumCommission int64
	total = amount + sumCommission
	if percentage != 0 {
		sumCommission = int64(float64(amount) * percentage / 100)
	}

	if sumCommission < minAmount {
		sumCommission = minAmount
	}

	newBalance := fromCard.Balance - total

	if newBalance < 0 {
		return 0, false, ErrNotEnoughFundsAccount
	}

	fromCard.Balance = newBalance
	toCard.Balance += amount

	return total, true, err
}

func (s *Service) searchCommission(from, to bool) *Commission {
	for _, candidate := range s.Commission {
		if candidate.from == from && candidate.to == to {
			return candidate
		}
	}
	return &Commission{
		from:       false,
		to:         false,
		percentage: 10,
		minAmount:  100_00,
	}
}
