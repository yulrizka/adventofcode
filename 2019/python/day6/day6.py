import fileinput
import unittest

data = [i.strip().split(')') for i in fileinput.input("../../input/day6-sample2")]

orbits = {}
for x in data:
    orbits.setdefault(x[1], []).append(x[0])


def paths(x, n=0):
    for y in orbits[x]:
        if y == 'COM':
            return n + 1
        return paths(y, n + 1)


def part1():
    count = 0
    for x in orbits:
        count += paths(x)
    return count


def part2():
    # path YOU
    you = {}
    san = {}
    node = orbits['YOU'][0]
    assert node

    count = 0
    while node != 'COM':
        count += 1
        you[node] = count
        node = orbits[node][0]

    node = orbits['SAN'][0]
    assert node

    count = 0
    while node != 'COM':
        count += 1
        san[node] = count
        node = orbits[node][0]

    distance = []
    # find min distance
    for i in san:
        if i in you:
            distance.append(san[i] + you[i])

    return min(distance) - 2  # -2 minus -1 for earch to reach the same spot


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 204521

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 407
