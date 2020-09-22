package stats

import (
  "fmt"
  "github.com/hsmo1501/bankapp/pkg/bank/types"
)


func ExampleAvg(){
  payments := []types.Payment{
    {
      ID:2,
      Amount:53_00,
      Category: "Cat",
    },
    {
      ID:1,
      Amount:51_00,
      Category: "Cat",
    },
    {
      ID:3,
      Amount:52_00,
      Category: "Cat",
    },
  }

  fmt.Println(Avg(payments))

  //Output: 5200
}
