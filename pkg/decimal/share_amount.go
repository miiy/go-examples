package decimal

import "github.com/shopspring/decimal"

// ShareAmount 分账
// div 平均数截断：角 Truncate(1)，分 Truncate(2)
// mod 剩余数加到最后一个补齐
func ShareAmount(amount decimal.Decimal, num int) []decimal.Decimal {
	amounts := make([]decimal.Decimal, num)
	if amount.LessThanOrEqual(decimal.Zero) || num <= 0 {
		return amounts
	}
	nd := decimal.NewFromFloat(float64(num))
	div := amount.Div(nd).Truncate(1)
	mod := amount.Sub(div.Mul(nd))
	if sum := div.Mul(nd).Add(mod); !sum.Equal(amount) {
		panic("share amount error")
	}
	for i := 0; i < num; i++ {
		amounts[i] = div
		if i == num-1 {
			amounts[i] = div.Add(mod)
		}
	}
	return amounts
}
