import collections
import unittest

with open('../../input/day20') as f:
# with open('../../input/day20-sample') as f:
    algo, *img = f.read().strip().split("\n\n")
    img = img[0].split("\n")

assert len(algo) == 512

m = set()
for r, row in enumerate(img):
    for c, v in enumerate(row):
        if v == '#':
            m.add((r, c))

R = len(img)
C = len(img[0])


def print_map(m):
    rmin = min([r for r, c in m])
    rmax = max([r for r, c in m])
    cmin = min([c for r, c in m])
    cmax = max([c for r, c in m])

    for r in range(rmin-5, rmax + 10):
        v = ''
        for c in range(cmin-5, cmax + 10):
            if (r, c) in m:
                v += '#'
            else:
                v += ' '
        print(v)
    print()


def solve(step):
    G = set(m)
    on = True

    for step in range(step):
        rmin = min([r for r, c in G])
        rmax = max([r for r, c in G])
        cmin = min([c for r, c in G])
        cmax = max([c for r, c in G])
        mm = set()
        for r in range(rmin-5, rmax+10):
            for c in range(cmin - 5, cmax + 10):
                num = ''
                for dr in [-1, 0, 1]:
                    for dc in [-1, 0, 1]:
                        newp = (r + dr, c + dc)
                        if (newp in G) == on:
                            num += '1'
                        else:
                            num += '0'
                v = int(num, 2)
                # print(f'num {num}, v {v}, rule {algo[v] == "#"}, on {on}')
                assert 0 <= v <= 512
                if (algo[v] == '#') != on:
                    # print('add')
                    mm.add((r,c))

        G = set(mm)        
        on = False if on else True

    return len(G)


def part1():
    return solve(2)

def part2():
    return solve(50)


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 5571

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 17965
