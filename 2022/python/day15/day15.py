import unittest
import sys

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day15'
# infile = sys.argv[1] if len(sys.argv)>1 else '../../input/day15.sample'
with open(infile) as f:
    data = f.read().strip()
# print(data)

B = set()
S = set()

for line in data.split('\n'):
    w = line.split(' ')
    x = int(w[2][2:-1])
    y = int(w[3][2:-1])
    bx = int(w[8][2:-1])
    by = int(w[9][2:])
    d = abs(x - bx) + abs(y - by)
    S.add((x, y, d))
    B.add((bx, by))


def valid(x, y, S):
    for sx, sy, sd in S:
        d = abs(x - sx) + abs(y - sy)
        if d <= sd:
            return False

    return True


def part1():
    y = 2000000
    # y = 10

    total = 0
    for x in range(int(-1e7), int(1e7)):
        if not valid(x, y, S) and (x, y) not in B:
            total += 1

    return total


def part2():
    for sx, sy, sd in S:

        for dx in range(sd + 2):
            dy = (sd + 1) - dx
            for signx, signy in [(-1, -1), (-1, 1), (1, -1), (1, 1)]:
                x = sx + (dx * signx)
                y = sy + (dy * signy)

                if not (0 <= x <= 4000000 and 0 <= y <= 4000000):
                    continue

                if valid(x, y, S):
                    return x * 4000000 + y


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 5688618, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 12625383204261, f'got {ans}'
