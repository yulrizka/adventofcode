import fileinput
import unittest

data = [line.strip('\n') for line in fileinput.input("../../input/day2")]


def part1():
    p = (0, 0)
    for line in data:
        match line.split(' '):
            case 'up', n:
                p = (p[0], p[1] - int(n))
            case 'down', n:
                p = (p[0], p[1] + int(n))
            case 'forward', n:
                p = (p[0] + int(n), p[1])
    return p[0] * p[1]


def part2():
    p = (0, 0)
    aim = 0
    for line in data:
        match line.split(' '):
            case 'up', n:
                aim -= int(n)
            case 'down', n:
                aim += int(n)
            case 'forward', n:
                p = (p[0] + int(n), p[1] + (aim * int(n)))
    return p[0] * p[1]


class TestSum(unittest.TestCase):

    def test1(self):
        assert part1() == 1936494

    def test2(self):
        assert part2() == 1997106066
