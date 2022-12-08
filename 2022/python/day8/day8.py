import sys
import unittest

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day8'
# infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day8.sample'
with open(infile) as f:
    data = f.read().strip()


# print(data)

class Cell:
    def __init__(self, v, top, right, bottom, left, visible=False):
        self.v = v
        self.top = top
        self.right = right
        self.bottom = bottom
        self.left = left
        self.visible = visible


def part1():
    lines = data.split('\n')
    dim = len(lines)

    mmap = {}

    # scan left to right
    for y in range(dim):
        mx = 0
        for x in range(dim):
            v = int(lines[y][x])
            cell = Cell(v, 0, 0, 0, mx)
            mx = max(mx, v)
            mmap[(x, y)] = cell

    # scar right to left
    for y in range(0, dim):
        mx = 0
        for x in range(dim - 1, -1, -1):
            v = int(lines[y][x])
            mmap[(x, y)].right = mx
            mx = max(mx, v)

    # scan top to bottom
    for x in range(0, dim):
        mx = 0
        for y in range(dim):
            v = int(lines[y][x])
            mmap[(x, y)].top = mx
            mx = max(mx, v)

    # scan from bottom to top
    for x in range(dim):
        mx = 0
        for y in range(dim - 1, -1, -1):
            v = int(lines[y][x])
            mmap[(x, y)].bottom = mx
            mx = max(mx, v)

    total = 0

    for k, v in mmap.items():
        if k[0] == 0 or k[0] == dim - 1 or k[1] == 0 or k[1] == dim - 1:
            total += 1
            v.visible = True
            continue

        if v.v > v.top or v.v > v.left or v.v > v.bottom or v.v > v.right:
            total += 1
            v.visible = True

    return total


def part2():
    lines = data.split('\n')
    dim = len(lines)

    mm = 0
    for y in range(dim):
        for x in range(dim):
            bound = [0, dim - 1]
            if y in bound or x in bound:
                continue

            top = 0
            bottom = 0
            left = 0
            right = 0

            v = lines[y][x]

            # look to the top
            yy = y
            while yy - 1 >= 0:
                yy -= 1
                top += 1
                if lines[yy][x] >= v:
                    break

            # look to bottom
            yy = y
            while yy + 1 < dim:
                yy += 1
                bottom += 1
                if lines[yy][x] >= v:
                    break

            # look to right
            xx = x
            while xx + 1 < dim:
                xx += 1
                right += 1
                if lines[y][xx] >= v:
                    break

            # look to left
            xx = x
            while xx - 1 >= 0:
                xx -= 1
                left += 1
                if lines[y][xx] >= v:
                    break

            score = top * bottom * left * right
            mm = max(mm, score)

    return mm


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 1829, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 291840, f'got {ans}'
