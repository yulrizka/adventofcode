import re
import unittest

# with open('../../input/day22') as f:
with open('../../input/day22-sample') as f:
    data = f.read().strip().split("\n")

p = re.compile(r'(\w+) .=(.*),.=(.*),.=(.*)')

ins = []

X = set()
Y = set()
Z = set()

for line in data:
    m = p.match(line)
    switch = m.group(1) == 'on'
    x1, x2 = m.group(2).split('..')
    y1, y2 = m.group(3).split('..')
    z1, z2 = m.group(4).split('..')
    x1, x2, y1, y2, z1, z2 = int(x1), int(x2), int(y1), int(y2), int(z1), int(z2)
    ins.append((switch, x1, x2, y1, y2, z1, z2))

    # build a coordinate compressions
    x1, x2 = min(x1, x2), max(x1, x2)
    y1, y2 = min(y1, y2), max(y1, y2)
    z1, z2 = min(z1, z2), max(z1, z2)

    X.add(x1)
    X.add(x2)
    Y.add(y1)
    Y.add(y2)
    Z.add(z1)
    Z.add(z2)


def compress(c):
    coord = {}
    c = sorted(c)

    llen = {}

    for i, v in enumerate(c):
        coord[v] = i
        if i < len(c) - 1:
            llen[i] = c[i + 1] - v
        else:
            llen[i] = None

    return coord, llen


xc, xl = compress(X)
yc, yl = compress(Y)
zc, zl = compress(Z)


print(f'xc {xc}')
print(f'yc {yc}')
print(f'zc {zc}')
print(f'xl {xl}')
print(f'yl {yl}')
print(f'zl {zl}')


def solve(part):
    on = set()
    for step,i in enumerate(ins):
        print(step, i)
        turn_on, x1, x2, y1, y2, z1, z2 = i
        if part == 1:
            if x1 < -50 or y1 < -50 or z1 < -50:
                continue
            if x2 > 50 or y2 > 50 or z2 > 50:
                continue

        for x in range(xc[x1], xc[x2]):
            for y in range(yc[y1], yc[y2]):
                for z in range(zc[z1], zc[z2]):
                    print(f'x {x}, y {y}, z {z}')
                    print(f'x1 {x1}, y1 {y1}, z1 {z1}')
                    if turn_on:
                        # print('add', x, y, z)
                        on.add((x, y, z))
                    elif (x, y, z) in on:
                        on.remove((x, y, z))

    # print(on)
    # calculate on cube
    ans = 0
    for x, y, z in on:
        assert x in xl, f'{x}, {xl}'
        l1 = xl[x]
        l2 = yl[y]
        l3 = zl[z]

        ans += l1 * l2 * l3

    return ans


def part1():
    return solve(1)


def part2():
    return solve(2)


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 615700, f'{ans}'

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 0, f'{ans}'
