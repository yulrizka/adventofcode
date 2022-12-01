import fileinput
import sys
import unittest

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day1'


def calculate():
    sums = []
    total = 0
    for x in fileinput.input('../../input/day1'):
        if x == '\n':
            sums.append(total)
            total = 0
        else:
            total += int(x)

    return sums


def part1():
    return max(calculate())


def part2():
    sums = sorted(calculate(), reverse=True)
    return sum(sums[:3])


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 69501, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 202346, f'got {ans}'
