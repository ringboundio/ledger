package ledger

import (
	"errors"
	"time"

	"github.com/cockroachdb/apd/v3"
	"github.com/google/uuid"
)

var (
	ErrTooFewLegs           = errors.New("ledger: must have at least two legs")
	ErrLegAmountNotPositive = errors.New("ledger: leg amount must be positive")
	ErrDecimalSet           = errors.New("ledger: could not set Decimal to leg amount")
	ErrDecimalAdd           = errors.New("ledger: could not add leg amount to Decimal")
	ErrDecimalInexact       = errors.New("ledger: Decimal operation resulted in rounding")
	ErrDebitCreditMismatch  = errors.New("ledger: total debits must equal total credits")
)

// Assets = Liabilities + Equity
//
// Assets and Expenses increase with a debit and decrease with a credit.
// Liabilities, Equity, and Revenue increase with a credit and decrease with a debit.
//
// Every debit must have a corresponding credit, and the total debits must equal total credits.
type TransactionBuilder struct {
	builder Transaction
}

func NewTransactionBuilder() *TransactionBuilder {
	return &TransactionBuilder{
		builder: Transaction{
			legs: make([]leg, 0, 3),
		},
	}
}

func (this *TransactionBuilder) Stamp(time time.Time) *TransactionBuilder {
	this.builder.time = time.UTC()
	return this
}

func (this *TransactionBuilder) Tag(tag string) *TransactionBuilder {
	this.builder.tag = tag
	return this
}

func (this *TransactionBuilder) Blame(person string) *TransactionBuilder {
	this.builder.blame = person
	return this
}

func (this *TransactionBuilder) House(house string) *TransactionBuilder {
	this.builder.house = house
	return this
}

func (this *TransactionBuilder) Location(location string) *TransactionBuilder {
	this.builder.location = location
	return this
}

func (this *TransactionBuilder) Debit(account AccountDescriptor, amount float64) *TransactionBuilder {
	this.builder.legs = append(this.builder.legs, newLeg(account, false, amount))
	return this
}

func (this *TransactionBuilder) Credit(account AccountDescriptor, amount float64) *TransactionBuilder {
	this.builder.legs = append(this.builder.legs, newLeg(account, true, amount))
	return this
}

func (this *TransactionBuilder) Note(note string) *TransactionBuilder {
	this.builder.note = note
	return this
}

func (this *TransactionBuilder) Build() (Transaction, error) {
	if err := this.Error(); err != nil {
		return Transaction{uuid: uuid.Nil}, err
	}

	var (
		count = len(this.builder.legs)
		legs  = make([]leg, 0, count)
	)

	for _, leg := range this.builder.legs {
		legs = append(legs, leg)
	}

	transaction := Transaction{
		uuid:     uuid.New(),
		time:     time.Now().UTC(),
		tag:      this.builder.tag,
		blame:    this.builder.blame,
		house:    this.builder.house,
		location: this.builder.location,
		legs:     legs,
		note:     this.builder.note,
	}

	if !this.builder.time.IsZero() {
		transaction.time = this.builder.time
	}

	return transaction, nil
}

func (this *TransactionBuilder) Error() error {
	if len(this.builder.legs) < 2 {
		return ErrTooFewLegs
	}

	var debit, credit, pending apd.Decimal

	for _, leg := range this.builder.legs {
		if leg.amount <= 0 {
			return ErrLegAmountNotPositive
		}

		if _, err := pending.SetFloat64(leg.amount); err != nil {
			return errors.Join(ErrDecimalSet, err)
		}

		var (
			condition apd.Condition
			err       error
		)

		if leg.credit {
			condition, err = apd.BaseContext.Add(&credit, &credit, &pending)
		} else {
			condition, err = apd.BaseContext.Add(&debit, &debit, &pending)
		}

		if err != nil {
			return errors.Join(ErrDecimalAdd, err)
		}

		if condition.Inexact() {
			return ErrDecimalInexact
		}
	}

	if debit.Cmp(&credit) != 0 {
		return ErrDebitCreditMismatch
	}

	return nil
}

type Transaction struct {
	uuid         uuid.UUID
	time         time.Time
	tag          string
	blame, house string
	location     string
	legs         []leg
	note         string
}

func (this Transaction) UUID() uuid.UUID {
	return this.uuid
}

func (this Transaction) IsZero() bool {
	return this.uuid == uuid.Nil
}

type leg struct {
	account AccountDescriptor
	credit  bool
	amount  float64
}

func newLeg(account AccountDescriptor, credit bool, amount float64) leg {
	return leg{
		account: account,
		credit:  credit,
		amount:  amount,
	}
}
