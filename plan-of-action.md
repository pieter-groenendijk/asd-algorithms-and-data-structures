# Plan of Action
## Generic constraints

- **Programming language: Golang**, I just like it and it can go relatively 
low-level, has generics, built-in testing library. Slices will have to be
ignored to make implementing data structures and algorithms more fair.
- **Project structure: One golang application**, the size is relatively small,
and it's nice to have accessible automated testing (one command), instead of
it being spread out. We'll use a packages to define public and private
definitions.
- **Documentation & elaboration**: **Pandoc**, embedding the actual code statements
with a `\input`. Not completely sure that will work with great with code formatting.
So **need to check that before committing**.
- **Testing**: **Built-in Golang testing library**. Golang has a very capable
built-in testing library for both functionality as performance testing. 100%
coverage. Performance testing shall be done using benchmark testing, to make
comparison directly possible, and still usable in the future.

## Choice of algorithms & data structures

| Category | Choice |
| --- | ------ |
| Lists | Dynamic array (ArrayList) & linked list |
