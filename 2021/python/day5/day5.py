import fileinput
import unittest

data = [[[int(y) for y in x.split(',')] for x in line.strip('\n').split(' -> ')] for line in
        fileinput.input("../../input/day5")]


def solve(part):
    visited = {}
    for coord in data:
        start, end = coord
        x1, y1 = start[0], start[1]
        x2, y2 = end[0], end[1]

        if x1 != x2 and y1 != y2:
            if part == 1:
                continue

        if x1 < x2:
            xstep = 1
        elif x1 > x2:
            xstep = -1
        else:
            xstep = 0

        if y1 < y2:
            ystep = 1
        elif y1 > y2:
            ystep = -1
        else:
            ystep = 0

        start = (x1, y1)
        visited[start] = visited.get(start, 0) + 1
        while x1 != x2 or y1 != y2:
            x1 += xstep
            y1 += ystep
            start = (x1, y1)
            visited[start] = visited.get(start, 0) + 1

    cross = list(filter(lambda x: visited[x] > 1, visited))
    return len(cross)


def part1():
    return solve(1)


def part2():
    return solve(2)


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 5576

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 8834787
