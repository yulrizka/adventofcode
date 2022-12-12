import functools
import re
import sys
import unittest
from collections import deque

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day11'
# infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day11.sample'
with open(infile) as f:
    data = f.read().strip()


# print(data)

class Monkey:
    def __init__(self, raw: str):
        lines = raw.split('\n')
        p = re.compile(r'\d+')
        for i, line in enumerate(lines):
            match i:
                case 0:
                    self.num = int(p.findall(line)[0])
                case 1:
                    self.items = deque([int(x) for x in p.findall(line)])
                case 2:
                    self.operation = line.removeprefix('  Operation: new = ').split(' ')
                case 3:
                    self.div = int(p.findall(line)[0])
                case 4:
                    self.if_true = int(p.findall(line)[0])
                case 5:
                    self.if_false = int(p.findall(line)[0])

        self.inspected = 0

    def __repr__(self):
        return f'Monkey-{self.num}'


def calculate(monkeys, worry_func, iteration):
    for round_n in range(iteration):
        for monkey in monkeys:
            while monkey.items:
                item = monkey.items.popleft()
                monkey.inspected += 1
                op, num = monkey.operation[1], monkey.operation[2]

                if num == 'old':
                    num = item
                else:
                    num = int(num)

                worry = 0
                match op:
                    case '*':
                        worry = item * num
                    case '+':
                        worry = item + num

                item = worry_func(worry)
                if item % monkey.div == 0:
                    to_monkey = monkeys[monkey.if_true]
                else:
                    to_monkey = monkeys[monkey.if_false]
                to_monkey.items.append(item)

    inspection = [x.inspected for x in monkeys]
    inspection.sort(reverse=True)

    return inspection[0] * inspection[1]


def part1():
    monkeys = []
    for monkey in data.split('\n\n'):
        monkeys.append(Monkey(monkey))

    return calculate(monkeys, lambda x: x // 3, 20)


def part2():
    monkeys = []
    for monkey in data.split('\n\n'):
        monkeys.append(Monkey(monkey))

    common_divisor = functools.reduce(lambda cd, x: cd * x, (m.div for m in monkeys))

    return calculate(monkeys, lambda x: x % common_divisor, 10000)


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 50830, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 14399640002, f'got {ans}'
