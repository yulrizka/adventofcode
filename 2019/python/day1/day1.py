import unittest
from functools import reduce
import math

with open("../../input/day1") as file:
    data = [int(line) for line in file]


def calc(n):
    return math.floor(n/3) - 2


def calc2(n):
    sum = 0
    while True:
        n = math.floor(n/3) - 2
        if n < 0:
            break
        sum += n
    return sum


class TestSum(unittest.TestCase):
    def test(self):
        assert calc(12) == 2
        assert calc(14) == 2
        assert calc(1969) == 654
        assert calc(100756) == 33583

        assert calc2(14) == 2
        assert calc2(1969) == 966
        assert calc2(100756) == 50346


print("part 1: ",  reduce(lambda x, y: x + calc(y), data, 0))
print("part 2: ",  reduce(lambda x, y: x + calc2(y), data, 0))
