import fileinput
import unittest

data = [i.strip() for i in fileinput.input("../../input/day9")]
w = len(data[0]) - 1  # base

neigh_points = [(-1, 0), (1, 0), (0, -1), (0, 1)]  # top, bottom, left, right


def neigh(row, col):
    result = []
    for n in neigh_points:
        y, x = (row + n[0], col + n[1])
        if y < 0 or x < 0 or y > 99 or x > 99:
            continue
        # result.append(int([y][x]))
        result.append((y, x))
    return result


def low_points():
    points = []
    for y in range(len(data)):
        for x in range(w + 1):
            n_values = []
            for row, col in neigh(y, x):
                n_values += [int(data[row][col])]

            v = int(data[y][x])
            if min(n_values) > v:
                # current point is smaller then the neighbours
                points.append((y, x))
    return points


def part1():
    ans = 0
    for y, x in low_points():
        ans += int(data[y][x]) + 1
    return ans


checked = {}


def basin(y, x):
    points = [(y, x)]
    checked[(y, x)] = 1

    for row, col in neigh(y, x):
        if (row, col) in checked or int(data[row][col]) >= 9:
            continue
        points += basin(row, col)

    return points


def part2():
    basins = []
    for y, x in low_points():
        basins.append(set(basin(y, x)))
    size = sorted([len(x) for x in basins], reverse=True)
    ans = size[0] * size[1] * size[2]
    return ans


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 633

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 1050192
