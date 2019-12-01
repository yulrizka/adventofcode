use criterion::{Criterion, criterion_group, criterion_main};
use aoc201901::{get_input, solve};

pub fn criterion_benchmark(c: &mut Criterion) {
        c.bench_function("2019 day 1 part 1", |b| {
                let input = get_input();
                b.iter(|| solve(&input));
        });
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
