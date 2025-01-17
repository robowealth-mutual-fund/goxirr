package goxirr

import (
	"fmt"
	"testing"
	"time"
)

func ExampleXirr() {
	firstDate := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	t1 := Transaction{
		Date: firstDate,
		Cash: -100,
	}
	t2 := Transaction{
		Date: firstDate.Add(time.Hour * 24 * 365),
		Cash: 112,
	}

	tas := Transactions{t1, t2}
	fmt.Println(Xirr(tas, WithRound(2)))
	// Output: 12
}

func TestXirr(t *testing.T) {
	type args struct {
		transactions []Transaction
	}

	var case1, case2, case3, case4 args
	case1.transactions = append(case1.transactions, Transaction{
		Date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -100,
	}, Transaction{
		Date: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: 200,
	})

	case2.transactions = append(case2.transactions, Transaction{
		Date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -100,
	}, Transaction{
		Date: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: 100,
	})

	case3.transactions = append(case3.transactions, Transaction{
		Date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -100,
	}, Transaction{
		Date: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: 112,
	})

	case4.transactions = append(case4.transactions, Transaction{
		Date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -10000,
	}, Transaction{
		Date: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -10000,
	}, Transaction{
		Date: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: -10000,
	}, Transaction{
		Date: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
		Cash: 40000,
	})

	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "100%", args: case1, want: 100.00},
		{name: "0%", args: case2, want: 0.00},
		{name: "12%", args: case3, want: 12.00},
		{name: "15.09%", args: case4, want: 15.08},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Xirr(tt.args.transactions, WithRound(2), WithGuess(0.05)); got != tt.want {
				t.Errorf("Xirr() = %v, want %v", got, tt.want)
			}
		})
	}
}
