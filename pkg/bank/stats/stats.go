package stats

import (
	"github.com/hsmo1501/bankapp/pkg/bank/types"
)

func Avg(payments []types.Payment)  types.Money{
	 summ := types.Money(0)
	 for _, payment := range payments{
		
		 if payment.Amount > 0 {
			 summ += summ + payment.Amount
		 }
	 }
	average := summ / types.Money(len(payments))
	return average 
}