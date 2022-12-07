import re
import sys
import unittest
from collections import deque

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day5'
with open(infile) as f:
    data = f.read()
# print(data)


n = 9


def get_stack():
    drawing, ins = data.split('\n\n')
    stack = []
    for i in range(0, n):
        stack.append(deque())

    for line in drawing.split('\n'):
        pos = 1
        for i in range(0, 9):
            c = line[pos]
            if 'A' <= c <= 'Z':
                stack[i].appendleft(c)
            pos += 4

    # handle instruction
    p = re.compile(r'\d+')
    rows = []
    for line in ins.split('\n'):
        if line == '':
            break
        num, a, b = [int(x) for x in p.findall(line)]
        rows.append([num, a - 1, b - 1])

    return [stack, rows]


def part1():
    stack, ins = get_stack()

    # handle instruction
    for row in ins:
        num, a, b = row
        for i in range(0, num):
            x = stack[a].pop()
            stack[b].append(x)

    ans = ''
    for i in range(0, n):
        ans += stack[i][-1]

    return ans


def part2():
    stack, ins = get_stack()

    # handle instruction
    for row in ins:
        num, a, b = row
        chunk = []
        for i in range(0,num):
            chunk.insert(0, stack[a].pop())
        stack[b].extend(chunk)

    ans = ''
    for i in range(0, n):
        ans += stack[i][-1]

    return ans


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 'WCZTHTMPS', f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 'BLSGJSDTS', f'got {ans}'
