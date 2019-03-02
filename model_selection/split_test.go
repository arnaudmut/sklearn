package modelselection

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func ExampleKFold() {
	X := mat.NewDense(6, 1, []float64{1, 2, 3, 4, 5, 6})
	subtest := func(shuffle bool) {
		randomState := RandomState(7)
		fmt.Println("shuffle", shuffle)
		kf := &KFold{NSplits: 3, Shuffle: shuffle, RandomState: &randomState}
		for sp := range kf.Split(X, nil) {
			fmt.Printf("%#v\n", sp)
		}

	}
	subtest(false)
	subtest(true)
	// Output:
	// shuffle false
	// modelselection.Split{TrainIndex:[]int{0, 1, 2, 4}, TestIndex:[]int{5, 3}}
	// modelselection.Split{TrainIndex:[]int{4, 5, 2, 3}, TestIndex:[]int{0, 1}}
	// modelselection.Split{TrainIndex:[]int{4, 1, 2, 3}, TestIndex:[]int{5, 0}}
	// shuffle true
	// modelselection.Split{TrainIndex:[]int{5, 2, 4, 1}, TestIndex:[]int{0, 3}}
	// modelselection.Split{TrainIndex:[]int{4, 0, 1, 2}, TestIndex:[]int{3, 5}}
	// modelselection.Split{TrainIndex:[]int{4, 5, 2, 1}, TestIndex:[]int{0, 3}}

}
