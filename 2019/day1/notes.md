# Day 1: The Tyranny of the Rocket Equation

https://adventofcode.com/2019/day/1


## filter only ok
```rust
    let result: i64 = f.lines()
        .filter_map(|l| l.parse::<i64>().ok())
        .sum();
```

got error
```
error[E0599]: no method named `parse` found for type `std::result::Result<std::string::String, std::io::Error>` in the current scope
 --> src/main.rs:9:27
  |
9 |         .filter_map(|l| l.parse::<i64>().ok())
  |                           ^^^^^ method not found in `std::result::Result<std::string::String, std::io::Error>`
```

solution
```rust
    let result: i64 = f.lines()
        .filter_map(Result::ok) 
        .filter_map(|l| l.parse::<i64>().ok())
        .sum();
```

## Using Successors iterator

https://doc.rust-lang.org/std/iter/fn.successors.html

    Creates a new iterator where each successive item is computed based on the preceding one.

This is quite useful for getting solution where the current value is calculation based on the previous value

## Unused import because mod is part of test

When running the main I got

```
warning: unused import: `crate::calc_fuel`
  --> src/main.rs:22:9
   |
22 |     use crate::calc_fuel;
   |         ^^^^^^^^^^^^^^^^
   |
   = note: `#[warn(unused_imports)]` on by default
```

I believe because it treated the test module as use.

solution
```rust

#[cfg(test)]
mod tests {
    ...
}
```

https://doc.rust-lang.org/nightly/rust-by-example/testing/unit_testing.html
> Most unit tests go into a tests mod with the #[cfg(test)] attribute. Test functions are marked with the #[test] attribute.

https://doc.rust-lang.org/nightly/rust-by-example/attribute.html

    Attribute: An attribute is metadata applied to some module, crate or item. This metadata can be used to/for:
    - conditional compilation of code
    - set crate name, version and type (binary or library)
    - disable lints (warnings)
    - enable compiler features (macros, glob imports, etc.)
    - link to a foreign library
    - mark functions as unit tests
    - mark functions that will be part of a benchmark
>

## Benchmarking in rust


https://doc.rust-lang.org/1.2.0/book/benchmark-tests.html

```rust
#[cfg(test)]
mod tests {
    use super::*;
    use test::Bencher;

    #[bench]
    fn bench_add_two(b: &mut Bencher) {
        ...
    }
}
```

Couldn't get the standard benchmark in rust works, got error

```
use of unstable library feature 'test': `bench` is a part of custom test frameworks which are unstable
```

After googling found that the feature needs a nightly build. I don't want to go to that direction,

Instead, looked into [Criterion library](https://bheisler.github.io/criterion.rs/book/getting_started.html)

## Crates and module

To use criterion, the project needed to be structured differently. Need to understand how rust is structured.

https://doc.rust-lang.org/1.2.0/book/crates-and-modules.html

## References

* [Advent of Code in Rust: lesson leaned](https://gendignoux.com/blog/2019/08/25/rust-advent-of-code.html)
* [Solution with iterator](https://github.com/hashedone/advent-of-code-2019-rust/blob/master/day1-2/src/main.rs)

