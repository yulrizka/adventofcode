import fileinput
import unittest

data = [int(line) for line in fileinput.input("../../input/day1")]


def count_increase(num):
    inc = 0
    for i, v in enumerate(num):
        if i == 0:
            continue
        if v > num[i - 1]:
            inc += 1
    return inc


def sum3(num):
    retval = []
    for x in range(len(num) - 2):
        retval += [num[x] + num[x + 1] + num[x + 2]]
    return retval


print("part 1: ", count_increase(data))
print("part 2: ", count_increase(sum3(data)))


class TestSum(unittest.TestCase):

    def test(self):
        sample = [int(line) for line in fileinput.input("../../input/day1.sample")]
        print(count_increase(sample))
        assert count_increase(sample) == 7

    def test2(self):
        sample = [int(line) for line in fileinput.input("../../input/day1.sample")]
        num3 = sum3(sample)
        assert count_increase(num3) == 5
