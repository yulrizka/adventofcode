import sys
import unittest

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day6'
with open(infile) as f:
    data = f.read().strip()


# print(data)


def solve(input_data, n):
    s = []
    i = 0
    for x in input_data:
        i += 1

        if x in s:
            ix = s.index(x)
            s = s[ix + 1:]
        s.append(x)

        # print(x, s)
        if len(s) == n:
            ans = i
            # print('ans', ans)
            return ans


def part1(input_data):
    return solve(input_data, 4)


def part2(input_data):
    return solve(input_data, 14)


class TestSum(unittest.TestCase):

    def test1(self):
        assert part1('bvwbjplbgvbhsrlpgdmjqwftvncz') == 5
        assert part1('nppdvjthqldpwncqszvftbrmjlhg') == 6
        assert part1('nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg') == 10
        assert part1('zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw') == 11
        ans = part1(data)
        assert ans == 1235, f'got {ans}'

    def test2(self):
        assert part2('mjqjpqmgbljsphdztnvjfqwrcgsmlb') == 19
        assert part2('bvwbjplbgvbhsrlpgdmjqwftvncz') == 23
        assert part2('nppdvjthqldpwncqszvftbrmjlhg') == 23
        assert part2('nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg') == 29
        assert part2('zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw') == 26
        ans = part2(data)
        assert ans == 3051, f'got {ans}'
