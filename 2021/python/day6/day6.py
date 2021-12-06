import fileinput
import unittest


data = [int(i) for i in fileinput.input("../../input/day6").readline().split(',')]

def part1():
    input = data[:]
    for day in range(80):
        for i in range(len(input)):
            if input[i] == 0:
                input[i] = 6
                input.append(8)
            else:
                input[i] -= 1

    return len(input)


def part2():
    data2 = [str(i) for i in data]
    freq = list(map(data2.count, '012345678'))

    for day in range(256):
        zero = freq[0]
        freq = freq[1:] + [zero]
        freq[6] += zero

    return sum(freq)


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        # print(ans)
        assert ans == 350149

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 1590327954513
