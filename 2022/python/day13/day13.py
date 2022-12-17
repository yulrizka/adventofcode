import unittest
import sys
from functools import cmp_to_key

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day13'
# infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day13.sample'
with open(infile) as f:
    data = f.read().strip()

# print(data)

# debug = True
debug = False


def compare(a, b) -> int:
    if isinstance(a, int) and isinstance(b, int):
        if a < b:
            return -1
        elif a == b:
            return 0
        else:
            return 1
    elif isinstance(a, list) and isinstance(b, list):
        i = 0
        while i < len(a) and i < len(b):
            c = compare(a[i], b[i])
            if c == -1:
                return -1
            if c == 1:
                return 1
            i += 1
        if i == len(a) and i < len(b):
            return -1
        elif i < len(a) and i == len(b):
            return 1
        else:
            return 0
    elif isinstance(a, int) and isinstance(b, list):
        return compare([a], b)
    else:
        return compare(a, [b])


def part1():
    i = 0
    total = 0
    for pairs in data.split('\n\n'):
        i += 1
        a, b = pairs.split('\n')
        a, b = eval(a), eval(b)

        result = compare(a, b)
        if result == -1:
            total += i

    # 380 too low
    return total


def part2():
    nums = [eval(x) for x in data.replace('\n\n', '\n').split('\n')]
    nums.append([[2]])
    nums.append([[6]])
    nums = sorted(nums, key=cmp_to_key(compare))

    i = 0
    total = 1
    for x in nums:
        i += 1
        if x in [[[2]], [[6]]]:
            total *= i

    return total


class TestSum(unittest.TestCase):

    def test_solve(self):
        assert compare([1, 1, 3, 1, 1], [1, 1, 5, 1, 1]) == -1
        assert compare([[1], [2, 3, 4]], [[1], 4]) == -1
        assert compare([9], [[8, 7, 6]]) == 1
        assert compare([[4, 4], 4, 4], [[4, 4], 4, 4, 4]) == -1
        assert compare([7, 7, 7, 7], [7, 7, 7]) == 1
        assert compare([], [3]) == -1
        assert compare([[[]]], [[]]) == 1
        assert compare([1, [2, [3, [4, [5, 6, 7]]]], 8, 9], [1, [2, [3, [4, [5, 6, 0]]]], 8, 9]) == 1

    def test1(self):
        ans = part1()
        assert ans == 5808, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 22713, f'got {ans}'
