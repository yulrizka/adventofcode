import sys
import unittest

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day10'
# infile = sys.argv[1] if len(sys.argv)>1 else '../../input/day10.sample'
with open(infile) as f:
    data = f.read().strip()
# print(data)


cost = {
    'noop': 1,
    'addx': 2,
}


def part1():
    x = 1

    lines = data.split('\n')
    pt = -1
    cycle = 0

    check = [20, 60, 100, 140, 180, 220]

    signal = 0

    screen = []

    cmd = None
    while True:
        cycle += 1

        if cmd is None:
            pt += 1
            if pt == len(lines):
                break

            ins = lines[pt].split(' ')
            cmd = [ins, cost[ins[0]]]

        cmd[1] -= 1

        if cycle in check:
            signal += cycle * x


        if x <= (cycle % 40) <= x + 2:
            screen.append('#')
        else:
            screen.append(' ')

        if cmd[1] == 0:
            if cmd[0][0] == 'addx':
                num = int(cmd[0][1])
                x += num

            cmd = None

    for i in range(len(screen)):
        if i > 0 and i % 40 == 0:
            print()
        print(screen[i]*2, end='')

    return signal


def part2():
    ...

class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 13720, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 0, f'got {ans}'
