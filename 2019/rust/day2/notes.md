# Day 2: 1202 Program Alarm 

https://adventofcode.com/2019/day/2


## Using enum
```rust
enum OpCode {
    Add, Multiply, Halt
}

impl From<usize> for OpCode {
    fn from(item: usize) -> Self {
        match item {
            1 => Add,
            2 => Multiply,
            99 => Halt,
            _ => unreachable!()
        }
    }
}
```

Implement Form so that later on we can call `into()` function to convert the type `<usize>` to enum

```rust
        let opcode = code[sp].into();
```

## use usize which is how array is addressed 

https://doc.rust-lang.org/std/primitive.usize.html

If im using i64 for example to address array, I need to convert it like

```rust
let i:i64 = 1;
code[i as usize];
```

to address the value.  Else will get an error like:

```
error[E0277]: the type `[i64]` cannot be indexed by `i64`
  --> src/lib.rs:19:19
   |
19 |         let op1 = code[pos_op1];
   |                   ^^^^^^^^^^^^^ slice indices are of type `usize` or ranges of `usize`
```

## References

* [Advent of Code in Rust: lesson leaned](https://gendignoux.com/blog/2019/08/25/rust-advent-of-code.html)
* [Solution with enum](https://dev.to/mdenchev/the-aim-for-elegant-rust-advent-of-code-2019-problem-2-1jig)

