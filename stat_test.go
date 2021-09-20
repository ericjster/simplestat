package stat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStat(t *testing.T) {
	st := NewStat()
	assert.Equal(t, 0, st.GetCount(), "count is wrong")
	assert.Equal(t, 0, st.GetCount(), "count is wrong")
	assert.Equal(t, 0.0, st.GetSum(), "sum is wrong")
	assert.Equal(t, 0.0, st.GetMean(), "Mean is wrong")
	assert.Equal(t, 0.0, st.GetMin(), "Min is wrong")
	assert.Equal(t, 0.0, st.GetMax(), "Max is wrong")
	assert.Equal(t, 0.0, st.GetVarP(), "GetVarP is wrong")
	assert.Equal(t, 0.0, st.GetVarS(), "GetVarS is wrong")
	assert.Equal(t, 0.0, st.GetStdDevP(), "GetStdDevP is wrong")
	assert.Equal(t, 0.0, st.GetStdDevS(), "GetStdDevS is wrong")
}

func TestStat0(t *testing.T) {
	st := NewStat()
	st.Add(0.0)
	assert.Equal(t, 1, st.GetCount(), "count is wrong")
	assert.Equal(t, 0.0, st.GetSum(), "sum is wrong")
	assert.Equal(t, 0.0, st.GetMean(), "Mean is wrong")
	assert.Equal(t, 0.0, st.GetMin(), "Min is wrong")
	assert.Equal(t, 0.0, st.GetMax(), "Max is wrong")
	assert.Equal(t, 0.0, st.GetVarP(), "GetVarP is wrong")
	assert.Equal(t, 0.0, st.GetVarS(), "GetVarS is wrong")
	assert.Equal(t, 0.0, st.GetStdDevP(), "GetStdDevP is wrong")
	assert.Equal(t, 0.0, st.GetStdDevS(), "GetStdDevS is wrong")
}

func TestStat1(t *testing.T) {
	st := NewStat()
	st.Add(1.0)
	assert.Equal(t, 1, st.GetCount(), "count is wrong")
	assert.Equal(t, 1.0, st.GetSum(), "sum is wrong")
	assert.Equal(t, 1.0, st.GetMean(), "Mean is wrong")
	assert.Equal(t, 1.0, st.GetMin(), "Min is wrong")
	assert.Equal(t, 1.0, st.GetMax(), "Max is wrong")
	assert.Equal(t, 0.0, st.GetVarP(), "GetVarP is wrong")
	assert.Equal(t, 0.0, st.GetVarS(), "GetVarS is wrong")
	assert.Equal(t, 0.0, st.GetStdDevP(), "GetStdDevP is wrong")
	assert.Equal(t, 0.0, st.GetStdDevS(), "GetStdDevS is wrong")
}

func TestStat2(t *testing.T) {
	st := NewStat()
	st.AddMany(1.0, 1.0)
	// st.AddMany([]float64{1.0, 1.0})
	assert.Equal(t, 2, st.GetCount(), "count is wrong")
	assert.Equal(t, 2.0, st.GetSum(), "sum is wrong")
	assert.Equal(t, 1.0, st.GetMean(), "Mean is wrong")
	assert.Equal(t, 1.0, st.GetMin(), "Min is wrong")
	assert.Equal(t, 1.0, st.GetMax(), "Max is wrong")
	assert.Equal(t, 0.0, st.GetVarP(), "GetVarP is wrong")
	assert.Equal(t, 0.0, st.GetVarS(), "GetVarS is wrong")
	assert.Equal(t, 0.0, st.GetStdDevP(), "GetStdDevP is wrong")
	assert.Equal(t, 0.0, st.GetStdDevS(), "GetStdDevS is wrong")
}

func TestStat10(t *testing.T) {
	st := NewStat()
	st.AddMany(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	assert.Equal(t, 10, st.GetCount(), "count is wrong")
	assert.Equal(t, 55.0, st.GetSum(), "sum is wrong")
	assert.Equal(t, 5.5, st.GetMean(), "Mean is wrong")
	assert.Equal(t, 1.0, st.GetMin(), "Min is wrong")
	assert.Equal(t, 10.0, st.GetMax(), "Max is wrong")
	assert.InDelta(t, 8.2500, st.GetVarP(), 0.0001, "GetVarP is wrong")
	assert.InDelta(t, 9.1667, st.GetVarS(), 0.0001, "GetVarS is wrong")
	assert.InDelta(t, 2.8723, st.GetStdDevP(), 0.0001, "GetStdDevP is wrong")
	assert.InDelta(t, 3.0277, st.GetStdDevS(), 0.0001, "GetStdDevS is wrong")
}
