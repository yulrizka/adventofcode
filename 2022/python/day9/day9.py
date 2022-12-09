import sys
import unittest

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day9'
# infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day9.sample'
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
    #  01 02 03 04 05
    #  16 .. .. .. 06
    #  15 .. XX .. 07
    #  14 .. .. .. 08
    #  13 12 11 10 09
    #
    #  XX is position of the tails
    #  .. directly connected, we don't do anything


    (-2, -2): (-1, -1),  # 01
    (-1, -2): (-1, -1),  # 02
    (0, -2): (0, -1),    # 03
    (1, -2): (1, -1),    # 04
    (2, -2): (1, -1),    # 05

    (2, -1): (1, -1),    # 06
    (2, 0): (1, 0),      # 07
    (2, 1): (1, 1),      # 08
    (2, 2): (1, 1),      # 09

    (1, 2): (1, 1),      # 10
    (0, 2): (0, 1),      # 11
    (-1, 2): (-1, 1),    # 12
    (-2, 2): (-1, 1),    # 13

    (-2, 1): (-1, 1),    # 14
    (-2, 0): (-1, 0),    # 15
    (-2, -1): (-1, -1),  # 16
}


def calculate_tail_pos(dir, head, tail):

    pos = (head[0] - tail[0], head[1] - tail[1])
    if pos not in tail_move:
        return tail

    move = tail_move[pos]
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
            head = calculate_head_pos(dir, head)
            tail = calculate_tail_pos(dir, head, tail)
            visited[tail] = True

    return len(visited)


def part2():
    num_node = 10
    tails = [(0, 0)] * num_node
    visited = {}

    for line in data.split('\n'):
        dir, n = line.split(' ')
        n = int(n)

        for x in range(n):
            for i in range(num_node):
                if i == 0:
                    tails[0] = calculate_head_pos(dir, tails[0])
                else:
                    tails[i] = calculate_tail_pos(dir, tails[i - 1], tails[i])
            visited[tails[9]] = True

    return len(visited)


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 5710, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 2259, f'got {ans}'
