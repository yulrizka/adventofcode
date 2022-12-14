import sys
import unittest
from heapq import heappop, heappush
import curses
from time import sleep

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day12'
# infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day12.sample'
with open(infile) as f:
    data = f.read().strip()

# print(data)

neighbours = [
    (0, -1),
    (1, 0),
    (0, +1),
    (-1, 0)
]


def part1():
    rows = data.split('\n')
    Y = len(rows)
    X = len(rows[0])

    # init distance
    root, target = (0, 0), (0, 0)
    distances = {}
    for y in range(Y):
        for x in range(X):
            p = (x, y)
            distances[p] = float('inf')
            match rows[y][x]:
                case 'S':
                    root = p
                case 'E':
                    target = p

    distances[root] = 0
    queue = [(0, root)]
    parent = {}

    while len(queue):
        distance, node = heappop(queue)
        x, y = node
        if node == target:
            break
        if distance > distances[node]:
            continue

        # print("current node:", node)
        for n in neighbours:
            nx, ny = x + n[0], y + n[1]
            if not (0 <= nx < X and 0 <= ny < Y):
                continue

            nv = rows[ny][nx]
            nv = 'z' if nv == 'E' else nv
            vv = rows[y][x]
            vv = 'a' if vv == 'S' else vv


            if ord(nv) - ord(vv) > 1:  # character comparison
                # print(nv, vv, 'skip')
                continue

            neigh = (nx, ny)
            d = distances[node] + 1  # weight is 1
            if distances[neigh] <= d:
                # print(node, vv, neigh, nv, f'distance is more {d} >= {distances[neigh]}, skip')
                continue
            # print(nv, vv, 'considering ', neigh)
            parent[neigh] = node
            heappush(queue, (d, neigh))
            distances[neigh] = d

    stdscr = curses.initscr()
    curses.noecho()
    curses.cbreak()
    curses.curs_set(False)
    if curses.has_colors():
        curses.start_color()

    # 516
    # 517 too low
    print(distances[target])
    p = target
    i = 0

    paths = []

    try:
        while True:
            i += 1
            # print(i, p)
            paths.append(p)
            # stdscr.addstr(p[1], p[0], 'x')
            p = parent[p]
            if p == root:
                paths.append(p)
                # print(i, p, "done")
                # stdscr.addstr(p[1], p[0], 'E')
                break

        paths.reverse()
        for p in paths:
            x, y = p[0], p[1]
            stdscr.addstr(y, x, rows[y][x])
            stdscr.refresh()
            sleep(0.01)

    except Exception as err:
        print(err)

    sleep(1000)

    return distances[target]


def part2():
    ...


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 0, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 0, f'got {ans}'


part1()
