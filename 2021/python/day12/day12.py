import fileinput
import unittest

data = [i.strip() for i in fileinput.input("../../input/day12")]

map = {}
for p in data:
    x, y = p.split('-')
    map[x] = sorted(map.get(x, []) + [y])
    map[y] = sorted(map.get(y, []) + [x])

def traverse(x, lseen, visited):
    seen = lseen.copy()
    if x == 'end':
        return 1

    if x not in map:
        return 0

    if x != x.upper():
        seen[x] = True

    v = 0
    for p in map[x]:
        if p != p.upper():
            if p in seen:
                continue
        v += traverse(p, seen, visited + [x])
    return v


def part1():
    paths = 0
    for x in map['start']:
        seen = {'start': True}
        paths += traverse(x, seen, [])
    return paths


def traverse2(x, lseen, visited, small_visited=False):
    seen = lseen.copy()
    # print(x, "> start,{}".format(','.join(visited)), seen, small_visited)
    if x == 'end':
        # print("FINISH: start,{},end".format(','.join(visited)))
        return 1

    if x not in map:
        return 0

    if x != x.upper():
        seen[x] = seen.get(x, 0) + 1

    v = 0
    for p in map[x]:
        if p != p.upper() and p in seen:
            if not small_visited:
                if seen[p] > 2:
                    continue
            else:
                if seen[p] > 0:
                    continue

        if x != x.upper():
            if seen[x] == 2:
                small_visited = True
        v += traverse2(p, seen, visited + [x], small_visited)
    return v


def part2():
    paths = 0
    for x in map['start']:
        seen = {'start': 3}
        paths += traverse2(x, seen, [])
    return paths


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 4167

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 98441
