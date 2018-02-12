package main

import (
	"fmt"

	tg "github.com/galeone/tfgo"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func main() {
	root := tg.NewRoot()
	A := tg.NewTensor(root, tg.Const(root, [2][2]int32{{1, 2}, {-1, -2}}))
	x := tg.NewTensor(root, tg.Const(root, [2][1]int64{{10}, {100}}))
	b := tg.NewTensor(root, tg.Const(root, [2][1]int32{{-10}, {10}}))
	Y := A.MatMul(x.Output).Add(b.Output)
	// Please note that Y is just a pointer to A!

	// If we want to create a different node in the graph, we have to clone Y
	// or equivalently A
	Z := A.Clone()
	results := tg.Exec(root, []tf.Output{Y.Output, Z.Output}, nil, &tf.SessionOptions{})
	fmt.Println("Y: ", results[0].Value(), "Z: ", results[1].Value())
	fmt.Println("Y == A", Y == A) // ==> true
	fmt.Println("Z == A", Z == A) // ==> false
}
