import sys
import unittest

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day3'
with open(infile) as f:
    data = f.read().strip()


# print(data)


def prio(x):
    i = ord(x)
    if i > 90:
        return i - 96
    else:
        return i - 38


def part1():
    total = 0
    for line in data.split('\n'):
        mid = int(len(line) / 2)
        a, b = line[:mid], line[mid:]

        for x in b:
            if x in a:
                total += prio(x)
                break

    return total


def part2():
    total = 0
    rows = data.split('\n')
    i = 0
    while i < len(rows):
        a, b, c = rows[i], rows[i + 1], rows[i + 2]
        i += 3
        for x in a:
            if (x in b) and (x in c):
                total += prio(x)
                break

    return total


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 8018, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 2518, f'got {ans}'
