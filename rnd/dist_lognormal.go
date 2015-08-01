// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rnd

import "math"

type LogNormal struct {
	Mu  float64 // μ: location
	Sig float64 // σ: scale

	// auxiliary
	vari2 float64 // 2 * σ²: 2 times variance
	den   float64 // σ * sqrt(2 * π)
}

// derived compute derived/auxiliary quantities
func (o *LogNormal) derived() {
	o.vari2 = 2.0 * o.Sig * o.Sig
	o.den = o.Sig * math.Sqrt(2.0*math.Pi)
}

// InitNonlog initialise lognormal distribution with non-log parameters
//  m -- non-logarithmised mean: μ
//  s -- non-logarithmised standard deviation: σ
func (o *LogNormal) InitStd(m, s float64) {
	o.Sig = math.Sqrt(math.Log(1.0 + s*s/(m*m)))
	o.Mu = math.Log(m / o.Sig)
	o.derived()
}

// Init initialises lognormal distribution
func (o *LogNormal) Init(μ, σ float64) {
	o.Mu, o.Sig = μ, σ
	o.derived()
}

// Pdf computes the probability density function @ x
func (o LogNormal) Pdf(x float64) float64 {
	if x < 1e-16 {
		return 0
	}
	return math.Exp(-math.Pow(math.Log(x)-o.Mu, 2.0)/o.vari2) / o.den / x
}

// Cdf computes the cumulative probability function @ x
func (o LogNormal) Cdf(x float64) float64 {
	if x < 1e-16 {
		return 0
	}
	return math.Erfc((o.Mu-math.Log(x))/(o.Sig*math.Sqrt2)) / 2.0
}