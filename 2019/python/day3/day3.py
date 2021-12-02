import fileinput
import unittest

reader = fileinput.input("../../input/day3")
wire1 = [i for i in reader.readline().split(',')]
wire2 = [i for i in reader.readline().split(',')]


def manhattan(point):
    return sum(map(abs, point))


def path(wire):
    x, y = 0, 0
    visited = {}
    step = 0
    for s in wire:
        d = s[0]
        count = int(s[1:])
        for i in range(count):
            step += 1
            match d:
                case 'U':
                    y += 1
                case 'D':
                    y += -1
                case 'L':
                    x += -1
                case 'R':
                    x += 1
            if (x, y) not in visited:
                visited[x, y] = step

    return visited


w1 = path(wire1)
w2 = path(wire2)
w1_set = set(w1)
w2_set = set(w2)
intersection = w1_set & w2_set


def part1():
    return min(list(map(manhattan, intersection)))


def part2():
    distance = [w1[point] + w2[point] for point in intersection]
    return min(distance)


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 375

    def test2(self):
        assert part2() == 14746
