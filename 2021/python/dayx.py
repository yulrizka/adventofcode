import fileinput
import unittest

data = [int(i) for i in fileinput.input("../../input/day").readline().split(',')]
print(data)


def part1():
    ...


def part2():
    ...


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 0

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 0
