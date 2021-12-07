import fileinput
import unittest

data = [int(i) for i in fileinput.input("../../input/day7").readline().split(',')]


def part1():
    fuels = {}
    for p in range(max(data)):
        count = 0
        for x in data:
            count += abs(p - x)
        fuels[p] = count

    v = min(fuels, key=lambda x: fuels[x])
    return fuels[v]


def part2():
    fuels = {}
    mem = {}
    for p in range(max(data)):
        count = 0
        for x in data:
            key = abs(p - x)
            if key in mem:
                count += mem[key]
            else:
                val = 0
                for i in range(key):
                    val += i + 1
                count += val
                mem[key] = val

        fuels[p] = count

    v = min(fuels, key=lambda x: fuels[x])
    ans = fuels[v]
    return ans


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 356922

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 100347031
