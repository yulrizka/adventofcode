import queue
import unittest

# with open("../../input/day15-sample") as f:
with open("../../input/day15") as f:
    raw = [[int(x) for x in y] for y in f.read().strip().split("\n")]


def wrap(x):
    while x > 9:
        x = x - 9
    return x


data2 = raw.copy()
for i in range(4):
    row = list(map(lambda x: list(map(lambda y: wrap(y + (i + 1)), x)), raw))
    data2 += row

for i, current_row in enumerate(data2):
    rr = current_row.copy()
    for j in range(4):
        row = list(map(lambda y: wrap(y + (j + 1)), current_row))
        rr += row
    data2[i] = rr

nr = [-1, 0, 1, 0]
nc = [0, 1, 0, -1]


def solve(raw):
    R = len(raw)
    C = len(raw[0])

    # build vertices
    D = {}
    G = {}
    for r in range(R):
        for c in range(C):
            D[(r, c)] = float('inf')
            for dd in range(4):
                rr = r + nr[dd]
                cc = c + nc[dd]
                if 0 <= rr < R and 0 <= cc < C:
                    G[((r, c), (rr, cc))] = int(raw[rr][cc])
    D[(0, 0)] = 0

    # dijkstra
    pq = queue.PriorityQueue()
    pq.put((0, (0, 0)))

    while not pq.empty():
        (dist, current_vertex) = pq.get()

        for dd in range(4):
            rr = current_vertex[0] + nr[dd]
            cc = current_vertex[1] + nc[dd]

            if 0 <= rr < R and 0 <= cc < C:
                neighbor = (rr, cc)
                distance = G[(current_vertex, neighbor)]
                old_cost = D[neighbor]
                new_cost = D[current_vertex] + distance
                if new_cost < old_cost:
                    D[neighbor] = new_cost
                    pq.put((new_cost, neighbor))

    return D[(R - 1, C - 1)]


def part1():
    return solve(raw)


def part2():
    return solve(data2)


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 498

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 2901
