import sys
import unittest

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day2'
with open(infile) as f:
    data = f.read().strip()
# print(data)


result = {
    ('A', 'X'): 3,
    ('A', 'Y'): 6,
    ('A', 'Z'): 0,
    ('B', 'X'): 0,
    ('B', 'Y'): 3,
    ('B', 'Z'): 6,
    ('C', 'X'): 6,
    ('C', 'Y'): 0,
    ('C', 'Z'): 3
}

shape = {
    'X': 1,
    'Y': 2,
    'Z': 3
}

strategy = {
    ('A', 'X'): 'Z',
    ('A', 'Y'): 'X',
    ('A', 'Z'): 'Y',
    ('B', 'X'): 'X',
    ('B', 'Y'): 'Y',
    ('B', 'Z'): 'Z',
    ('C', 'X'): 'Y',
    ('C', 'Y'): 'Z',
    ('C', 'Z'): 'X'
}


def part1():
    score = 0
    for x in data.split('\n'):
        a, b = x.split()
        score += result[(a, b)] + shape[b]
    return score


def part2():
    score = 0
    for x in data.split('\n'):
        a, b = x.split()
        c = strategy[(a, b)]
        score += result[(a, c)] + shape[c]
    return score


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 10994, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 0, f'got {ans}'
