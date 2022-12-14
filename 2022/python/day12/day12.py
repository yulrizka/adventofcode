import sys
import unittest
from heapq import heappop, heappush

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

rows = data.split('\n')
Y = len(rows)
X = len(rows[0])


def solve(root=None):
    # init distance
    target = (0, 0)
    distances = {}
    for y in range(Y):
        for x in range(X):
            p = (x, y)
            distances[p] = float('inf')
            match rows[y][x]:
                case 'S':
                    if root is None:
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

        for n in neighbours:
            nx, ny = x + n[0], y + n[1]
            if not (0 <= nx < X and 0 <= ny < Y):
                continue

            nv = rows[ny][nx]
            nv = 'z' if nv == 'E' else nv
            vv = rows[y][x]
            vv = 'a' if vv == 'S' else vv

            if ord(nv) - ord(vv) > 1:  # character comparison
                continue

            neigh = (nx, ny)
            d = distances[node] + 1  # weight is 1
            if distances[neigh] <= d:
                continue
            parent[neigh] = node
            heappush(queue, (d, neigh))
            distances[neigh] = d

    return distances[target]


def part1():
    return solve(None)


def part2():
    min_val = float('inf')
    for y in range(Y):
        for x in range(X):
            if rows[y][x] != 'a':
                continue

            v = solve((x, y))
            min_val = min(min_val, v)

    return min_val


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 520, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 0, f'got {ans}'
