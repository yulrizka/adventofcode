import fileinput
import unittest
import collections

data = [i.strip() for i in fileinput.input("../../input/day10")]
# print(data)

open = '([{<'
close = ')]}>'
point = [1, 2, 3, 4]


def part1():
    count = {}
    for line in data:
        s = collections.deque()
        for c in line:
            if c in open:
                s.append(c)
            else:
                left = s.pop()
                if open.index(left) != close.index(c):
                    count[c] = count.get(c, 0) + 1

    ans = count[')'] * 3 + count[']'] * 57 + count['}'] * 1197 + count['>'] * 25137
    return ans


def part2():
    total_score = []
    for line in data:
        valid = True
        s = collections.deque()
        for c in line:
            if c in open:
                s.append(c)
            else:
                left = s.pop()
                if open.index(left) != close.index(c):
                    valid = False
                    break
        if valid:
            score = 0
            for x in reversed(s):
                idx = open.index(x)
                score = (score * 5) + point[idx]
            total_score.append(score)

    ans = sorted(total_score)[(len(total_score) // 2)]
    return ans


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 318081

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 4361305341
