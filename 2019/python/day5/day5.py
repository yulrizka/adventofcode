import fileinput
import unittest

data = [int(i) for i in fileinput.input("../../input/day5").readline().strip().split(',')]


class IntComp:
    mem = []
    ip = 0
    input = ''

    def __init__(self, nums):
        self.mem = nums.copy()

    def set_input(self, v):
        self.input = v

    def params(self, i, mode):
        addr = self.ip + i
        v = self.mem[addr]
        if mode == '0':
            return self.mem[v]  # position mode
        else:
            return v  # immediate mode

    def run(self):
        out = ''
        while True:
            print('>', self.ip)
            num = self.mem[self.ip]
            num_str = str(num)
            mode = ''
            if num < 10:
                op = num
                mode = '000'
            else:
                op = int(num_str[-2:])
                mode = num_str[-3::-1]  # remove op & reverse
                if len(mode) < 3:
                    mode += '0' * (3 - len(mode))

            print("ip:{} num:{}, op:{}, mode:{}".format(self.ip, num, op, mode))

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
                    v = self.input
                    self.mem[p] = v
                    self.ip += 2
                case 4:
                    p = self.params(1, mode)
                    v = self.mem[p]
                    print("{} > {}".format(p, v))
                    out = v
                    self.ip += 2
                case 5:
                    p1 = self.params(1, mode[0])
                    p2 = self.params(2, mode[1])
                    if p1 > 0:
                        self.ip = p2
                    else:
                        self.ip += 3
                case 6:
                    p1 = self.params(1, mode[0])
                    p2 = self.params(2, mode[1])
                    if p1 == 0:
                        self.ip = p2
                    else:
                        self.ip += 3
                case 7:
                    p1 = self.params(1, mode[0])
                    p2 = self.params(2, mode[1])
                    p3 = self.params(3, '1')
                    self.mem[p3] = 1 if p1 < p2 else 0
                    self.ip += 4
                case 8:
                    p1 = self.params(1, mode[0])
                    p2 = self.params(2, mode[1])
                    p3 = self.params(3, '1')
                    self.mem[p3] = 1 if p1 == p2 else 0
                    self.ip += 4
                case n:
                    print("invalid op", op, "self.ip:", self.ip)
                    raise Exception('invalid op', op)


def part1():
    c = IntComp(data)
    c.set_input(1)
    return c.run()


def part2():
    c = IntComp(data)
    c.set_input(5)
    return c.run()


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 16209841

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 8834787

    def run_sample(self, data, input, want):
        sample = [int(i) for i in data.split(',')]
        c = IntComp(sample)
        c.set_input(input)
        ans = c.run()
        print(ans)
        assert ans == want

    def test_sample(self):
        self.run_sample('3,9,8,9,10,9,4,9,99,-1,8', 8, 1)
        self.run_sample('3,9,8,9,10,9,4,9,99,-1,8', 7, 0)

        self.run_sample('3,9,7,9,10,9,4,9,99,-1,8', 7, 1)
        self.run_sample('3,9,7,9,10,9,4,9,99,-1,8', 8, 0)

        self.run_sample('3,3,1108,-1,8,3,4,3,99', 8, 1)
        self.run_sample('3,3,1108,-1,8,3,4,3,99', 7, 0)

        self.run_sample('3,3,1107,-1,8,3,4,3,99', 8, 0)
        self.run_sample('3,3,1107,-1,8,3,4,3,99', 7, 1)
