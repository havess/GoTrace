package core

import "math"

type Spectrum RGBSpectrum

type CoefficientSpectrum struct {
	nSpectrumSamples int
	c                []float64
}

func NewCoefficientSpectrum(nSpectrumSamples int, v float64) CoefficientSpectrum {
	ret := CoefficientSpectrum{nSpectrumSamples: nSpectrumSamples}
	for i := 0; i < nSpectrumSamples; i++ {
		ret.c[i] = v
	}
	return ret
}

func (self CoefficientSpectrum) Add(c2 CoefficientSpectrum) CoefficientSpectrum {
	ret := self
	for i := 0; i < self.nSpectrumSamples; i++ {
		ret.c[i] += c2.c[i]
	}
	return ret
}

func (self CoefficientSpectrum) Subtract(c2 CoefficientSpectrum) CoefficientSpectrum {
	ret := self
	for i := 0; i < self.nSpectrumSamples; i++ {
		ret.c[i] -= c2.c[i]
	}
	return ret
}

func (self CoefficientSpectrum) MultiplyCS(c2 CoefficientSpectrum) CoefficientSpectrum {
	ret := self
	for i := 0; i < self.nSpectrumSamples; i++ {
		ret.c[i] *= c2.c[i]
	}
	return ret
}

func (self CoefficientSpectrum) MultiplyF(f float64) CoefficientSpectrum {
	ret := self
	for i := 0; i < self.nSpectrumSamples; i++ {
		ret.c[i] *= f
	}
	return ret
}

func (self CoefficientSpectrum) Divide(c2 CoefficientSpectrum) CoefficientSpectrum {
	ret := self
	for i := 0; i < self.nSpectrumSamples; i++ {
		ret.c[i] /= c2.c[i]
	}
	return ret
}

func (self CoefficientSpectrum) Equal(c2 CoefficientSpectrum) bool {
	for i := 0; i < self.nSpectrumSamples; i++ {
		if self.c[i] != c2.c[i] {
			return false
		}
	}
	return true
}

func (self CoefficientSpectrum) Negate(c2 CoefficientSpectrum) CoefficientSpectrum {
	ret := self
	for i := 0; i < self.nSpectrumSamples; i++ {
		ret.c[i] *= -1
	}
	return ret
}

func (self CoefficientSpectrum) Sqrt() CoefficientSpectrum {
	ret := self
	for i := 0; i < self.nSpectrumSamples; i++ {
		ret.c[i] = math.Sqrt(self.c[i])
	}
	return ret
}

func (self CoefficientSpectrum) Pow(p float64) CoefficientSpectrum {
	ret := self
	for i := 0; i < self.nSpectrumSamples; i++ {
		ret.c[i] = math.Pow(self.c[i], p)
	}
	return ret
}

func (self CoefficientSpectrum) Exp() CoefficientSpectrum {
	ret := self
	for i := 0; i < self.nSpectrumSamples; i++ {
		ret.c[i] = math.Pow(math.E, self.c[i])
	}
	return ret
}

func (self CoefficientSpectrum) IsBlack() bool {
	for i := 0; i < self.nSpectrumSamples; i++ {
		if self.c[i] != 0. {
			return false
		}
	}
	return true
}

func LerpSPD(c1, c2 CoefficientSpectrum, t float64) CoefficientSpectrum {
	return c1.MultiplyF(1 - t).Add(c2.MultiplyF(t))
}

type RGBSpectrum struct {
}
