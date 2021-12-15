import unittest
import collections

# with open("../../input/day14-sample") as f:
with open("../../input/day14") as f:
    input, ins = f.read().strip().split('\n\n')

rulemap = {}

for x in ins.split('\n'):
    a, b = x.split(' -> ')
    rulemap[a] = b


def part1():
    str = input
    for step in range(10):
        i = 0
        c = str[0]
        for i in range(len(str) - 1):
            c = c[0:-1]
            ch = str[i] + str[i + 1]
            c += str[i] + rulemap[ch] + str[i + 1]
        str = c
        freq = collections.Counter(str)

    vmin = min(freq, key=lambda x: freq[x])
    vmax = max(freq, key=lambda x: freq[x])
    return freq[vmax] - freq[vmin]


def part2():
    freq = collections.Counter(map(str.__add__, input, input[1:]))
    chars = collections.Counter(input)
    for step in range(40):
        for (a, b), c in freq.copy().items():
            x = rulemap[a + b]
            freq[a + b] -= c
            freq[a + x] += c
            freq[x + b] += c
            chars[x] += c

    return (max(chars.values()) - min(chars.values()))




class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 3411

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 7477815755570
