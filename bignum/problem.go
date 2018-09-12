package bignum

import (
	"math/big"
)

//Problem defines the parameters of the algorithm
type Problem struct {
	BucketA *big.Int
	BucketB *big.Int
	Desired *big.Int
}

//NewProblem creates a new Problem from parameters describing each bucket size and the desired remainder.
//
// Parameters:
// a *bignum.Int Size of bucket A
// b *bignum.Int Size of bucket B
// d *bignum.Int Desired remainder from the solution
//
// Returns:
// p *Problem a pointer to the created object
func NewProblem(a, b, d *big.Int) (p *Problem) {
	p = new(Problem)
	p.BucketA = a
	p.BucketB = b
	p.Desired = d
	return p
}

func (p *Problem) compareAandB() int {
	return p.BucketA.Cmp(p.BucketB)
}
