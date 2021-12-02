import fileinput
import unittest

data = [line.strip('\n') for line in fileinput.input("../../input/day2")]


def part1():
    p = (0, 0)
    for line in data:
        dir, count = line.split(' ')
        count = int(count)
        if dir == 'up':
            p = (p[0], p[1] - count)
        elif dir == 'down':
            p = (p[0], p[1] + count)
        elif dir == 'forward':
            p = (p[0] + count, p[1])
    return p[0] * p[1]


def part2():
    p = (0, 0)
    aim = 0
    for line in data:
        dir, count = line.split(' ')
        count = int(count)
        if dir == 'up':
            aim -= count
        elif dir == 'down':
            aim += count
        elif dir == 'forward':
            p = (p[0] + count, p[1] + (aim * count))
    return p[0] * p[1]



class TestSum(unittest.TestCase):

    def test1(self):
        assert part1() == 1936494

    def test2(self):
        assert part2() == 1936494
