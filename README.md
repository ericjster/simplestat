Simple Stat
===========

This is a simple class to compute online statistics.

Sample usage:
```
st := stat.NewStat()
st.Add(1.0)
st.AddMany(2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
fmt.Printf("Count: %d, Min: %.3f, Max: %.3f, Mean: %.3f, StdDev: %.3f\n",
    st.GetCount(), st.GetMin(), st.GetMax(), st.GetMean(), st.GetStdDevS())

Count: 10, Min: 1.000, Max: 10.000, Mean: 5.500, StdDev: 3.028
```

License
=======

This project is licensed under the terms of the MIT license.