import collections
import fileinput
import unittest

data = [[int(x) for x in i.strip()] for i in fileinput.input("../../input/day11")]
# data = [[int(x) for x in i.strip()] for i in fileinput.input("../../input/day11-sample2")]
# print(data)

dy = [-1, -1, -1, 0, 1, 1, 1, 0]
dx = [-1, 0, 1, 1, 1, 0, -1, -1]

R = len(data)
C = len(data[0])


def pp(nums, step):
    print("step", step)
    for row in nums:
        print(str(row))
    print()


def solve(part):
    ldata = data.copy()
    flash = 0
    # pp(0)
    go = True
    i = 0
    while go:
        i += 1
        q = collections.deque()
        seen = {}
        for r in range(R):
            for c in range(C):
                ldata[r][c] += 1

                if ldata[r][c] > 9:
                    q.append((r, c))
        while q:
            p = q.pop()
            if p in seen:
                continue

            seen[p] = True
            flash += 1

            # go trough neighbor
            for j in range(len(dy)):
                yy, xx = p[0] + dy[j], p[1] + dx[j]
                if 0 <= yy < R and 0 <= xx < C:
                    if (yy, xx) not in seen:
                        ldata[yy][xx] += 1

                    if ldata[yy][xx] > 9:
                        q.append((yy, xx))

            ldata[p[0]][p[1]] = 0

        if i == 100 and part == 1:
            return flash

        # part 2
        finished = True
        for rr in ldata:
            for v in rr:
                if v != 0:
                    finished = False

        if finished:
            return i


def part1():
    return solve(1)


def part2():
    return solve(2)


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 1669

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 351
