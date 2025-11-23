# Hashing (for hash tables)
## Assumptions
Below are my thoughts before literature review; based of what I know

An ideal hashing function maps input:

- **evenly spread output**; every underlying index in the hash table needs to have the same chance of being used
- **deterministic**; any input value always leads to the same output -> making lookups possible
- **type agnostic**; any input ideally may be used (Although maybe type safe)
- **fast**; time complexity of `O(1)` is nice, but that one shouldn't be an hour.
- **resource-efficient**; space complexity should be low (not educated enough, to give an estimate of space complexity)

**Does it matter if key hashes can be traced back to the key values?** Probably not, hashes are used to map unbounded space
to (superficially) bounded space. They're not supposed to be secure or something.

## Expectations
- MUST: A independent hash table implementation
- MUST: A independent smart hashing algorithm of a hash table
- MUST: Unit tested
- SHOULD: A collision resolving algorithm being independently functional of a hash table
- COULD: A independent stupid hashing algorithm of a hash table for comparison
