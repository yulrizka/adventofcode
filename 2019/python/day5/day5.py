import fileinput
import unittest

data = [int(i) for i in fileinput.input("../../input/day5").readline().strip().split(',')]


class IntComp:
    mem = []
    ip = 0

    def __init__(self, nums):
        self.mem = nums.copy()

    def params(self, i, mode):
        addr = self.ip + i
        v = self.mem[addr]
        if mode == '0':
            # position mode
            return self.mem[v]
        else:
            # immediate mode
            return v

    def run(self):
        out = ''
        while True:
            num = self.mem[self.ip]
            num_str = str(num)
            mode = ''
            if num < 10:
                op = num
                mode = '000'
            else:
                op = int(num_str[-2:])
                mode = num_str[-3::-1]  # remove op & reverse
                if len(mode) < 2 and op in (1, 2):
                    mode += '0' * (2-len(mode))

            match op:
                case 99:
                    return out
                case 1:
                    p1 = self.params(1, mode[0])
                    p2 = self.params(2, mode[1])
                    addr = self.params(3, '1')
                    self.mem[addr] = p1 + p2
                    self.ip += 4
                case 2:
                    p1 = self.params(1, mode[0])
                    p2 = self.params(2, mode[1])
                    addr = self.params(3, '1')
                    self.mem[addr] = p1 * p2
                    self.ip += 4
                case 3:
                    p = self.params(1, '1')
                    # v = input('input: ')
                    v = 1
                    self.mem[p] = v
                    self.ip += 2
                case 4:
                    p = self.params(1, mode)
                    v = self.mem[p]
                    print("{} > {}".format(p, v))
                    self.ip += 2
                    out = v
                case n:
                    print("invalid op", op, "self.ip:", self.ip)
                    raise Exception('invalid op', op)


def part1():
    c = IntComp(data)
    return c.run()


def part2():
    return 0


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 16209841

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 757
