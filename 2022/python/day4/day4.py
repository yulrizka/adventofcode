import sys
import unittest

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day4'
# infile = sys.argv[1] if len(sys.argv)>1 else '../../input/day4.sample'
with open(infile) as f:
    data = f.read().strip()


def part1():
    total = 0
    for line in data.split('\n'):
        a, b = line.split(',')
        ax, ay = [int(x) for x in a.split('-')]
        bx, by = [int(x) for x in b.split('-')]
        if ax <= bx <= by <= ay:
            total += 1
            continue
        if bx <= ax <= ay <= by:
            total += 1
            continue

    return total


def part2():
    total = 0
    for line in data.split('\n'):
        a, b = line.split(',')
        ax, ay = [int(x) for x in a.split('-')]
        bx, by = [int(x) for x in b.split('-')]
        if (by < ax) or (ay < bx):
            pass
        else:
            total += 1

        # 185 too low
    return total


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 500, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 815, f'got {ans}'
