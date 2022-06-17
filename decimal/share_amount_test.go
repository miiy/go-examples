package decimal

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestShareReturnItemAmt(t *testing.T) {
	t.Log("0, 0", ShareAmount(decimal.NewFromFloat(0), 0))
	t.Log("1, 0", ShareAmount(decimal.NewFromFloat(1), 0))
	t.Log("0, 1", ShareAmount(decimal.NewFromFloat(0), 1))
	t.Log("1, 2", ShareAmount(decimal.NewFromFloat(1), 2))
	t.Log("2, 1", ShareAmount(decimal.NewFromFloat(2), 1))
	t.Log("3, 2", ShareAmount(decimal.NewFromFloat(3), 2))
	t.Log("3, 3", ShareAmount(decimal.NewFromFloat(3), 3))
	t.Log("4, 3", ShareAmount(decimal.NewFromFloat(4), 3))
	t.Log("4.5, 3", ShareAmount(decimal.NewFromFloat(4.5), 3))
	t.Log("45.5, 3", ShareAmount(decimal.NewFromFloat(45.5), 3))
	t.Log("10, 2", ShareAmount(decimal.NewFromFloat(10), 2))
}
