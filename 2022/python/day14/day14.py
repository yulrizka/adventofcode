import unittest
import sys

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day14'
# infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day14.sample'
with open(infile) as f:
    data = f.read().strip()

# print(data)

dir = [
    (-1, 1),
    (0, 1),
    (1, 1)
]


def add(p1, p2):
    return p1[0] + p2[0], p1[1] + p2[1]


def part1():
    map = {}
    for line in data.split('\n'):
        start = None
        for pair in line.split(' -> '):
            x, y = pair.split(',')
            x, y = int(x), int(y)
            point = (x, y)
            if start is None:
                xx, yy = x, y
                map[point] = 'x'
                start = point
                continue

            sx, sy = start
            if x == sx:
                diff = 1 if y < sy else -1
                for dy in range(y, sy, diff):
                    map[(x, dy)] = 'x'
            elif y == sy:
                diff = 1 if x < sx else -1
                for dx in range(x, sx, diff):
                    map[(dx, y)] = 'x'
            start = point

    maxy = 0
    for x in map:
        maxy = max(maxy, x[1])

    sand = None
    num_sand = 0
    while True:
        if sand is None:
            sand = (500, 0)
            num_sand += 1

        # print(sand)
        if sand[1] > maxy:
            break

        # check bottom direction
        bl, bb, br = add(sand, dir[0]), add(sand, dir[1]), add(sand, dir[2])

        if bb not in map:
            sand = bb
            continue
        else:
            if bl not in map:
                sand = bl
                continue
            else:
                if br not in map:
                    sand = br
                    continue
                else:
                    map[sand] = 'o'
                    print('stable', sand)
                    # stable
                    sand = None
                    continue

    return num_sand -1


def part2():
    ...


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 0, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 0, f'got {ans}'
