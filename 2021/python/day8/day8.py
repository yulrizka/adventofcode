import fileinput
import itertools
import unittest

data = [i.strip().split(' | ') for i in fileinput.input("../../input/day8")]


def part1():
    count = 0
    for line in data:
        word = line[1].split()
        count += len(list(filter(lambda x: len(x) in [2, 4, 3, 7], word)))
    return count


def part2():
    comb = {}
    # generate all possible solution
    for w in itertools.permutations('abcdefg'):
        v = [
            w[0] + w[1] + w[2] + w[4] + w[5] + w[6],          # 0
            w[2] + w[5],                                      # 1
            w[0] + w[2] + w[3] + w[4] + w[6],                 # 2
            w[0] + w[2] + w[3] + w[5] + w[6],                 # 3
            w[1] + w[2] + w[3] + w[5],                        # 4
            w[0] + w[1] + w[3] + w[5] + w[6],                 # 5
            w[0] + w[1] + w[3] + w[4] + w[5] + w[6],          # 6
            w[0] + w[2] + w[5],                               # 7
            w[0] + w[1] + w[2] + w[3] + w[4] + w[5] + w[6],   # 8
            w[0] + w[1] + w[2] + w[3] + w[5] + w[6],          # 9
        ]
        v = [''.join(sorted(x)) for x in v]  # sort 'becdafg' -> 'abcdefg'
        comb[''.join(w)] = v

    total = 0
    for line in data:
        # sort the word again but not the whole list, position matters
        digits = [''.join(sorted(x)) for x in line[0].split()]
        output = [''.join(sorted(x)) for x in line[1].split()]

        for c in comb:
            sets = comb[c]
            if set(sets) == set(digits):
                num = ''
                for o in output:
                    num += str(sets.index(o))

                total += (int(num))
                break

    return total


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 504

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 1073431
