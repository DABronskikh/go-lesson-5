package transfer

import (
	"github.com/DABronskikh/go-lesson-5/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	testBank := card.NewService("Test bank")
	card1 := testBank.IssueCard("Visa", "RUB")
	card1.Number = "5106 2100 0000 0007"
	card2 := testBank.IssueCard("MasterCard", "RUB")
	card2.Number = "5106 2100 0000 0049"

	testBankTransfer := NewService(testBank)
	testBankTransfer.IssueCommission(true, true, 0, 0)
	testBankTransfer.IssueCommission(true, false, 0.5, 10_00)
	testBankTransfer.IssueCommission(false, false, 1.5, 30_00)

	testBankTransfer1 := testBankTransfer
	testBankTransfer2 := testBankTransfer
	testBankTransfer3 := testBankTransfer
	testBankTransfer4 := testBankTransfer
	testBankTransfer5 := testBankTransfer
	testBankTransfer6 := testBankTransfer
	testBankTransfer7 := testBankTransfer
	testBankTransfer8 := testBankTransfer
	testBankTransfer9 := testBankTransfer

	type fields struct {
		CardSvc    *card.Service
		Commission []*Commission
	}
	type args struct {
		from   string
		to     string
		amount int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantOk    bool
		wantErr   error
	}{
		{
			name: "Карта своего банка -> Карта своего банка (денег достаточно)",
			fields: fields{
				CardSvc:    testBankTransfer1.CardSvc,
				Commission: testBankTransfer1.Commission,
			},
			args: args{
				from:   testBankTransfer1.CardSvc.Cards[0].Number,
				to:     testBankTransfer1.CardSvc.Cards[1].Number,
				amount: 1000_00,
			},
			wantTotal: 1000_00,
			wantOk:    true,
			wantErr:   nil,
		},
		{
			name: "Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields: fields{
				CardSvc:    testBankTransfer2.CardSvc,
				Commission: testBankTransfer2.Commission,
			},
			args: args{
				from:   testBankTransfer2.CardSvc.Cards[0].Number,
				to:     testBankTransfer2.CardSvc.Cards[1].Number,
				amount: 100_000_00,
			},
			wantTotal: 0,
			wantOk:    false,
			wantErr:   ErrNotEnoughFundsAccount,
		},
		{
			name: "Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields: fields{
				CardSvc:    testBankTransfer3.CardSvc,
				Commission: testBankTransfer3.Commission,
			},
			args: args{
				from:   testBankTransfer3.CardSvc.Cards[0].Number,
				to:     "4276 1600 0000 0043",
				amount: 100_00,
			},
			wantTotal: 0,
			wantOk:    false,
			wantErr:   ErrTargetCardNotFound,
		},
		{
			name: "Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields: fields{
				CardSvc:    testBankTransfer4.CardSvc,
				Commission: testBankTransfer4.Commission,
			},
			args: args{
				from:   testBankTransfer4.CardSvc.Cards[0].Number,
				to:     "4276 1600 0000 0043",
				amount: 100_000_00,
			},
			wantTotal: 0,
			wantOk:    false,
			wantErr:   ErrTargetCardNotFound,
		},
		{
			name: "Карта чужого банка -> Карта своего банка",
			fields: fields{
				CardSvc:    testBankTransfer5.CardSvc,
				Commission: testBankTransfer5.Commission,
			},
			args: args{
				from:   "4276 1600 0000 0043",
				to:     testBankTransfer5.CardSvc.Cards[0].Number,
				amount: 1000_00,
			},
			wantTotal: 0,
			wantOk:    false,
			wantErr:   ErrSourceCardNotFound,
		},
		{
			name: "Карта чужого банка -> Карта чужого банка",
			fields: fields{
				CardSvc:    testBankTransfer6.CardSvc,
				Commission: testBankTransfer6.Commission,
			},
			args: args{
				from:   "4276 1600 0000 0043",
				to:     "4276 1600 0000 0050",
				amount: 100_00,
			},
			wantTotal: 0,
			wantOk:    false,
			wantErr:   ErrSourceCardNotFound,
		},
		{
			name: "Задача №2 – Card Structure (проверка по префиксу карты)",
			fields: fields{
				CardSvc:    testBankTransfer7.CardSvc,
				Commission: testBankTransfer7.Commission,
			},
			args: args{
				from:   testBankTransfer7.CardSvc.Cards[0].Number,
				to:     "5106 2100 0000 0049",
				amount: 100_00,
			},
			wantTotal: 100_00,
			wantOk:    true,
			wantErr:   nil,
		},
		{
			name: "Задача №3 – Алгоритм Луна (номер валидный)",
			fields: fields{
				CardSvc:    testBankTransfer8.CardSvc,
				Commission: testBankTransfer8.Commission,
			},
			args: args{
				from:   testBankTransfer8.CardSvc.Cards[0].Number,
				to:     testBankTransfer8.CardSvc.Cards[1].Number,
				amount: 100_00,
			},
			wantTotal: 100_00,
			wantOk:    true,
			wantErr:   nil,
		},
		{
			name: "Задача №3 – Алгоритм Луна (номер НЕ валидный)",
			fields: fields{
				CardSvc:    testBankTransfer9.CardSvc,
				Commission: testBankTransfer9.Commission,
			},
			args: args{
				from:   testBankTransfer9.CardSvc.Cards[0].Number,
				to:     "5106 2100 0000 00XX",
				amount: 100_00,
			},
			wantTotal: 0,
			wantOk:    false,
			wantErr:   ErrInvalidCardNumber,
		},
	}
	for _, tt := range tests {
		s := &Service{
			CardSvc:    tt.fields.CardSvc,
			Commission: tt.fields.Commission,
		}
		gotTotal, gotOk, gotErr := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
		if gotTotal != tt.wantTotal {
			t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
		}
		if gotOk != tt.wantOk {
			t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
		}
		if gotErr != tt.wantErr {
			t.Errorf("Card2Card() gotOk = %v, want %v", gotErr, tt.wantErr)
		}
	}
}
