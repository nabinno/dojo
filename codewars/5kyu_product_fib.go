package main

// ProductFib ....
func ProductFib(prod uint64) [3]uint64 {
	var n, nPlus, fiboP uint64

	nPlus = 1
	for n*nPlus < prod {
		nPlus = n + nPlus
		n = nPlus - n
	}
	if n*nPlus == prod {
		fiboP = 1
	}

	return [3]uint64{n, nPlus, fiboP}
}
