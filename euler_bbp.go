// This program calculates hexadecimal digits of Pi without calculating
// the previous digits using the Euler BBP formulae.
package main

import (
	"fmt"
	"math"
	"math/bits"
	"runtime"
	"sync"
	"time"
)

// Calculate a**b mod k
//
// The speed of this could be improved by converting the modulo
// operations into multiply by the inverse
func powmod(a, b, k uint64) (s uint64) {
	var hi, lo uint64
	s = 1
	for ; b > 0; b >>= 1 {
		if b&1 != 0 {
			// s = (s * a) % k
			hi, lo = bits.Mul64(s, a)
			_, s = bits.Div64(hi, lo, k)
		}
		// a = (a * a) % k
		hi, lo = bits.Mul64(a, a)
		_, a = bits.Div64(hi, lo, k)
	}
	return s
}

// Calculate a**b mod k
//
// This uses multiplication to do the divisions which should be
// quicker, but turns out not to be.
func powmodTest(a, b, k uint64) (s uint64) {
	if k <= 1 {
		return 0
	}

	// First calculate 2^64 / k which we will use to multiply in
	// order to divide by k as multiplication is a lot quicker
	// than division
	m, _ := bits.Div64(1, 0, k)

	// Multiply two numbers together then calculate mod k
	// (a * b) mod k
	// 0 <= a,b < k
	var x_hi, x_lo, q, q2, y_hi, y_lo, carry, r_hi, r_lo uint64
	mulmod := func(a, b uint64) uint64 {
		x_hi, x_lo = bits.Mul64(a, b)
		// Estimate of divisor
		// This can't overflow
		q = x_hi * m
		q2, _ = bits.Mul64(x_lo, m)
		// This can't overflow
		q += q2
		// Calculate remainder
		y_hi, y_lo = bits.Mul64(q, k)
		r_lo, carry = bits.Sub64(x_lo, y_lo, 0)
		r_hi, _ = bits.Sub64(x_hi, y_hi, carry)
		for r_hi > 0 || r_lo > k {
			r_lo, carry = bits.Sub64(r_lo, k, 0)
			r_hi -= carry
		}
		return r_lo
	}

	s = 1
	for ; b > 0; b >>= 1 {
		if b&1 != 0 {
			s = mulmod(s, a)
		}
		a = mulmod(a, a)
	}
	return s
}

// Calculate a**b mod k
// NB this will go wrong if k exceeds 2**32
func powmod32(a, b, k uint64) uint64 {
	s := uint64(1)
	for ; b > 0; b >>= 1 {
		if b&1 != 0 {
			s = (s * a) % k
		}
		a = (a * a) % k
	}
	return s
}

// Calculate a term of the BBP series mod 1
//
// 2**x / k mod 1
func term(x int64, k uint64) float64 {
	// If exponent is non-negative use modular arithmetic
	if x >= 0 {
		return float64(powmod(2, uint64(x), k)) / float64(k)
	}
	// otherwise use normal floating point arithmetic
	return math.Pow(2, float64(x)) / float64(k)
}

// Produce a term of the original BBP formula multiplied by 2**m and
// mod 1
func bbp_original(m, n uint64) float64 {
	exponent := int64(m) - 4*int64(n)
	_, s := math.Modf(+term(exponent+2, 8*n+1) +
		-term(exponent+1, 8*n+4) +
		-term(exponent, 8*n+5) +
		-term(exponent, 8*n+6))
	return s
}

// Produce a term of Euler's first BBP formula multiplied by 2**m and
// mod 1
func bbp_euler(m, n uint64) float64 {
	exponent := int64(m) - 2*int64(n)
	s := (+term(exponent+1, 4*n+1) +
		+term(exponent+1, 4*n+2) +
		+term(exponent, 4*n+3))
	if n%2 != 0 {
		s = -s
	}
	_, s = math.Modf(s)
	return s
}

// Produce a term of Euler's second BBP formula multiplied by 2**m
// and mod 1
func bbp_euler2(m, n uint64) float64 {
	exponent := int64(m) - 30*int64(n) - 26
	s := (+term(exponent+27, 20*n+1) +
		+term(exponent+25, 12*n+1) +
		+term(exponent+25, 10*n+1) +
		+term(exponent+24, 20*n+3) +
		+term(exponent+22, 6*n+1) +
		-term(exponent+20, 60*n+15) +
		-term(exponent+19, 10*n+3) +
		-term(exponent+18, 20*n+7) +
		-term(exponent+15, 12*n+5) +
		+term(exponent+15, 20*n+9) +
		+term(exponent+12, 30*n+15) +
		+term(exponent+12, 20*n+11) +
		-term(exponent+10, 12*n+7) +
		-term(exponent+9, 20*n+13) +
		-term(exponent+7, 10*n+7) +
		-term(exponent+5, 60*n+45) +
		+term(exponent+3, 20*n+17) +
		+term(exponent+2, 6*n+5) +
		+term(exponent+1, 10*n+9) +
		+term(exponent, 12*n+11) +
		+term(exponent, 20*n+19))
	if n%2 != 0 {
		s = -s
	}
	_, s = math.Modf(s)
	return s
}

type bbpFormula func(m, n uint64) float64

// Calculate the m-th hex digit of pi using the formula passed in.
//
// bits_per_iteration is used to calculate the stopping point and
// terms is used to estimate the error.
func bbp(m uint64, name string, formula bbpFormula, bits_per_iteration int, terms int) {
	t0 := time.Now()
	// python floats have this many bits of precision
	precision := 53
	// this means
	//assert 1.0 + 2.0 **(-precision) == 1.0
	s := 0.0
	// start from the m-th hex digit counting from 1 not 0
	mh := 4 * (m - 1)

	// Calculate the number of iterations needed
	iterations := (mh+uint64(precision))/uint64(bits_per_iteration) + 1

	// Create worker routines
	workers := runtime.NumCPU()
	// work unit to range from [n0, n1)
	type work struct {
		n0, n1 uint64
	}
	in := make(chan work, workers)
	out := make(chan float64, workers)
	chunkSize := uint64(1e5) // iterations for each worker
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range in {
				s := 0.0
				for n := job.n0; n < job.n1; n++ {
					term := formula(mh, n)
					_, s = math.Modf(s + term)
					// fmt.Println(n, s, term)
				}
				out <- s
			}
		}()
	}
	// pump the work units in
	go func() {
		for n := uint64(0); n < iterations+1; n += chunkSize {
			job := work{n0: n, n1: n + chunkSize}
			if job.n1 > iterations+1 {
				job.n1 = iterations + 1
			}
			in <- job
		}
		// Signal to workers they are finished
		close(in)
		// Wait for workers to exit
		wg.Wait()
		// Signal to receiver that it is finished
		close(out)
	}()
	// Receive the results
	for result := range out {
		_, s = math.Modf(s + result)
	}
	if s < 0 {
		s += 1.0
	}
	dt := time.Since(t0)
	// Estimate the accuracy
	//
	// We've done n * terms additions. We might expect to have
	// an error of 1/2 * (2**-precision) per addition here which would
	// make a crude estimate of the maximum error to be:
	max_error := float64(iterations) * float64(terms) / 2 * math.Pow(2, -float64(precision))
	ok_bits := -math.Log2(max_error)
	ok_digits := int(ok_bits / 4)
	digits := int(s * math.Pow(2, float64(4*ok_digits)))
	fmt.Printf("%10s hex digits of pi from digit %12d %0*x in %v\n", name, m, ok_digits, digits, dt)
}

func main() {
	for m := uint64(10); m <= 1e11; m *= 10 {
		bbp(m, "original", bbp_original, 4, 4)
		bbp(m, "euler", bbp_euler, 2, 3)
		bbp(m, "euler2", bbp_euler2, 30, 21)
	}
}
