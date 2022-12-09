import sys
import unittest

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day9'
# infile = sys.argv[1] if len(sys.argv)>1 else '../../input/day9.sample'
with open(infile) as f:
    data = f.read().strip()
# print(data)

head_move = {
    "U": (0, -1),
    "R": (1, 0),
    "D": (0, 1),
    "L": (-1, 0),
}

tail_move = {
    "R": {
        # x y
        (0, 0): (0, 0),
        (0, -1): (0, 0),
        (1, -1): (0, 0),
        (1, 0): (0, 0),
        (1, 1): (0, 0),
        (0, 1): (0, 0),
        (-1, 1): (1, -1),
        (-1, 0): (1, 0),
        (-1, -1): (1, 1),
    },
    "L": {
        # x y
        (0, 0): (0, 0),
        (0, -1): (0, 0),  # 1
        (1, -1): (-1, 1),  # 2
        (1, 0): (-1, 0),  # 3
        (1, 1): (-1, -1),  # 4
        (0, 1): (0, 0),  # 5
        (-1, 1): (0, 0),  # 6
        (-1, 0): (0, 0),  # 7
        (-1, -1): (0, 0),  # 8
    },
    "U": {
        # x y
        (0, 0): (0, 0),
        (0, -1): (0, 0),  # 1
        (1, -1): (0, 0),  # 2
        (1, 0): (0, 0),  # 3
        (1, 1): (-1, -1),  # 4
        (0, 1): (0, -1),  # 5
        (-1, 1): (1, -1),  # 6
        (-1, 0): (0, 0),  # 7
        (-1, -1): (0, 0),  # 8
    },
    "D": {
        # x y
        (0, 0): (0, 0),
        (0, -1): (0, 1),  # 1
        (1, -1): (-1, 1),  # 2
        (1, 0): (0, 0),  # 3
        (1, 1): (0, 0),  # 4
        (0, 1): (0, 0),  # 5
        (-1, 1): (0, 0),  # 6
        (-1, 0): (0, 0),  # 7
        (-1, -1): (1, 1),  # 8
    }
}


def calculate_tail_pos(dir, head, tail):
    tail_map = tail_move[dir]

    pos = (tail[0] - head[0], tail[1] - head[1])
    move = tail_map[pos]
    new_tail = (tail[0] + move[0], tail[1] + move[1])

    return new_tail


def calculate_head_pos(dir, head):
    hm = head_move[dir]
    return head[0] + hm[0], head[1] + hm[1]


def part1():
    head, tail = (0, 0), (0, 0)
    visited = {}
    for line in data.split('\n'):
        dir, n = line.split(' ')

        for _ in range(int(n)):
            tail = calculate_tail_pos(dir, head, tail)
            visited[tail] = True

            head = calculate_head_pos(dir, head)

    return len(visited)


def part2():
    ...


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 5710, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 0, f'got {ans}'
