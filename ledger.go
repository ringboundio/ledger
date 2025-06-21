package ledger

import (
	"context"
	"errors"
	"math"

	qdb "github.com/questdb/go-questdb-client/v3"
)

var (
	ErrZeroTransaction = errors.New("ledger: transaction is zero")
)

type Ledger struct {
	ctx    context.Context
	sender qdb.LineSender
}

func NewLedger(ctx context.Context, sender qdb.LineSender) *Ledger {
	return &Ledger{
		ctx:    ctx,
		sender: sender,
	}
}

func (this *Ledger) Post(transaction Transaction) error {
	if transaction.IsZero() {
		return ErrZeroTransaction
	}

	uuid := transaction.uuid.String()

	for _, leg := range transaction.legs {
		var debit, credit float64
		if leg.credit {
			credit = leg.amount
		} else {
			debit = leg.amount
		}

		if err := this.sender.
			Table("journal").
			Symbol("transaction", uuid).
			Symbol("tag", transaction.tag).
			Symbol("blame", transaction.blame).
			Symbol("house", transaction.house).
			Symbol("location", transaction.location).
			Symbol("class", leg.account.class.name()).
			Symbol("category", leg.account.category.name()).
			Symbol("group", leg.account.group.name()).
			Symbol("container", leg.account.container).
			Symbol("item", leg.account.item).
			Symbol("id", leg.account.id).
			Int64Column("debit", cent(debit)).
			Int64Column("credit", cent(credit)).
			StringColumn("note", transaction.note).
			At(this.ctx, transaction.time); err != nil {
			return err
		}
	}

	return nil
}

func cent(positive float64) int64 {
	return int64(math.Trunc(positive * 100))
}
