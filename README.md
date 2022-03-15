# bigduration
Go library for bigger duration than `time.duration`.  
Covering from 0 seconds to around 2.925 * 10^11 years (`math.MaxInt64` seconds).

## Usage
```
d, err := bigduration.ParseDuration("31m12s")
```
