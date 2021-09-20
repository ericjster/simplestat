package stat

import "math"

// Stat maintains an online collection of summary statistics. By "online" we
// mean each value is added once as in a stream, and there is only an O(1) cost.
type Stat struct {
	n    int
	min  float64
	max  float64
	sum  float64
	sum2 float64
}

// NewStat returns a new Stat struct.
func NewStat() *Stat {
	return &Stat{
		n:    0,
		min:  math.MaxFloat64,
		max:  -math.MaxFloat64,
		sum:  0.0,
		sum2: 0.0,
	}
}

// Add adds a new value to the summary statistics.
func (s *Stat) Add(x float64) {
	s.n++
	s.sum += x
	s.sum2 += x * x
	if s.min > x {
		s.min = x
	}
	if s.max < x {
		s.max = x
	}
}

// AddMany add many values to the summary statistics.
func (s *Stat) AddMany(xs ...float64) {
	for _, x := range xs {
		s.Add(x)
	}
}

// GetCount returns the count of observed values.
func (s *Stat) GetCount() int {
	return s.n
}

// GetSum returns the sum of observed values.
func (s *Stat) GetSum() float64 {
	return s.sum
}

// GetMean returns the mean (average) of observed values.
func (s *Stat) GetMean() float64 {
	if s.n <= 0 {
		return 0.0
	}
	return s.sum / float64(s.n)
}

// GetMean returns the minimum of observed values.
func (s *Stat) GetMin() float64 {
	if s.n <= 0 {
		return 0.0
	}
	return s.min
}

// GetMean returns the maximum of observed values.
func (s *Stat) GetMax() float64 {
	if s.n <= 0 {
		return 0.0
	}
	return s.max
}

// GetVarP returns the population variance.
func (s *Stat) GetVarP() float64 {
	// This is the straightforward method, in Wikipedia called "Naive
	// algorithm". See also Welford's online algorithm:
	// https://en.wikipedia.org/wiki/Algorithms_for_calculating_variance#Welford's_online_algorithm
	if s.n <= 0 {
		return 0.0
	}
	n := float64(s.n)
	return (s.sum2 - (s.sum*s.sum)/n) / (n)
}

// GetVarS returns the sample variance.
func (s *Stat) GetVarS() float64 {
	if s.n <= 1 {
		return 0.0
	}
	n := float64(s.n)
	return (s.sum2 - (s.sum*s.sum)/n) / (n - 1)
}

// GetStdDevP returns the population standard deviation.
func (s *Stat) GetStdDevP() float64 {
	return math.Sqrt(s.GetVarP())
}

// GetStdDevS returns the sample standard deviation.
func (s *Stat) GetStdDevS() float64 {
	return math.Sqrt(s.GetVarS())
}
