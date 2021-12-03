import collections
import fileinput
import unittest

ranges = fileinput.input("../../input/day4").readline().split('-')
start, end = int(ranges[0]), int(ranges[1])


def part1():
    match = 0
    for i in range(start, end+1):
        pwd = str(i)
        if list(pwd) != sorted(pwd):
            continue

        count = collections.Counter(pwd)
        for k, v in count.items():
            if v >= 2:
                match += 1
                break

    return match


def part2():
    match = 0
    for i in range(start, end+1):
        pwd = str(i)
        if list(pwd) != sorted(pwd):
            continue

        count = collections.Counter(pwd)
        for k, v in count.items():
            if v == 2:
                match += 1
                break

    return match


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 1169

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 757
