import fileinput
import unittest
import math

data = [line.strip('\n') for line in fileinput.input("../../input/day3")]

digits = 12


def part1(nums):
    common = [0] * digits
    for line in nums:
        for i in range(digits):
            digit = line[i]
            if digit == '1':
                common[i] += 1

    result, inv = '', ''
    for i in range(digits):
        if common[i] > math.ceil(len(data) / 2 + 1):
            result += '1'
            inv += '0'
        else:
            result += '0'
            inv += '1'

    return result, inv


def part2(nums, useCommon):
    nums = nums[:]
    for i in range(digits):
        if len(nums) == 1:
            return nums[0]

        ones = 0
        for line in nums:
            if line[i] == '1':
                ones += 1

        if useCommon:
            keep_one = ones >= math.ceil(len(nums) / 2)
        else:
            keep_one = ones < math.ceil(len(nums) / 2)

        if keep_one:
            filtered = list(filter(lambda x: (x[i] == '1'), nums))
        else:
            filtered = list(filter(lambda x: (x[i] == '0'), nums))
        nums = filtered

    if len(nums) > 1:
        raise Exception("more than 1 solution")

    return nums[0]


class TestSum(unittest.TestCase):

    def test1(self):
        v, inv = part1(data)
        assert int(v, 2) * int(inv, 2) == 4103154

    def test2(self):
        oxygen = int(part2(data, True), 2)
        co2 = int(part2(data, False), 2)
        ans = oxygen * co2
        assert ans == 4245351
