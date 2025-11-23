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

## Research
- https://www.youtube.com/watch?v=0M_kIqhwbFo&t=15s

### Dictionary
- Exact operations: insert/delete/give the item mapped to specific key, or report it doesn't exist.

#### Direct-access table
The simple approach:

- store items in array
- indexed by key

If you're keys are integers this can just work. 

Negatives:
- All possibles values are reserved in memory (and running time of initialization)
- Keys may not be integers

### Prehashing
The solution for non-integer keys.

Maps keys to non-negative integers.

In theory, keys are finite and discrete. In the end everything can be expessed as a 
sequence of bits.

Ideally, no collisions: `prehash(x) == prehash(y)`, only if `x == y`.
Ideally, deterministic: so, at all times: `prehash(x) == prehash(x)`

### Hashing
The solution for being a memory hog.

Reduces universe of all keys `u` (integers)
down to a reasonable size `m` for the table.

Practically reducing the size of table roughly to the amount of items
it stores. Unlike the direct-access table, there low overhead, since
we're not reserving memory for all keys, regardless of usage. `m` should be
near `n`, the number of items stored.

This also means two keys might map to the same hash, `hash(x) == hash(y) && x != y`.
Instances of this are called collisions.

### Chaining
Dealing with collisions.

#### List
A key maps to list, instead of an actual value. Storing the key with the value to check
if that's the right value for our input key.

Worst case is that all keys are mapped to the same location, requiring iterating to do
a lookup, i.e. `O(n)`. That's why it should nicely distribute items over the bounded space.

### Simple uniform hashing
- **Uniformity**: Each key is equally likely to be hashed to any slot of the table
- **Independence**, independently of where other keys hashing.


Expected length of a chain for `n` keys, `m` slots is `n / m`. Also known as load factor.

- Being equally likely to be hashed to any slow, meaning the keys will be evenly distributed
- The amount of slots is a constant (for now) `m`.
- So, equally distributing keys over the fixed amount of slots, i.e. `n / m`.


As long as `m = O(n)`, lookups will be `O(1)`.


### Hash functions
#### Division method
`h(k) == k mod m`

Simple, and has problems.

#### Multiplication method
`h(k) == [(a*k) mod 2^w] >> (w-r)`

`m=2^r`

- So, we have key `k`, which is `w` bits long.
- We take some number `a`, think of it as a random integer, and we do binary multiplication (no clue what that is)
- Now we have a two words long result `2w`
- We extract the right part of the word using `mod 2^w`
- And the we shift the part to the right, to extract a `r` part of that righter part of the word.

Why does this work (see video attached), you're mixing up results of the multiplication, and then taking the most
random part: the middle.

#### Universal hashing
`h(k) = [(ak+b) mod p] mod m`

- `a` is random between `0` and `p - 1`
- `b` is random between `0` and `p - 1`
- `p` is a prime number bigger than the size of the universe `u`

For worst case keys `k1 != k2`, the probability of `h(k1) == h(k2)` is `1/m`.


## Action
Over my head...

So, decomposition time:

- Let's start with the dead simple division method for a hashing function
- Make collision resolving algorithm independently (not easily possible, since it affects the data structure?)
- Implement a simple hash table implementation
- Make it unit tested
- Figure out prehashing
- Implement universal hashing algorithm or something more smart and recognized
- Document choice
