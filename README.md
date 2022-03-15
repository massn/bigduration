# bigduration
Go library for bigger duration than `time.duration`.  
Covering from `0 seconds` to `2.925 * 10^11 years (9223372036854775807 seconds)`.

## Usage
```
d, err := bigduration.ParseDuration("31m12s")
```
