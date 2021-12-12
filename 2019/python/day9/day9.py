import collections
import fileinput
import unittest

ins = [int(i) for i in fileinput.input("../../input/day9").readline().split(',')]
# ins = [int(i) for i in fileinput.input("../../input/day9-sample").readline().split(',')]
data = {}
for i, v in enumerate(ins):
    data[i] = v

# print(data)

RUNNABLE = 0
WAITING_INPUT = 1
DONE = 99


class IntComp:

    def __init__(self, nums, pid=0):
        self.mem = []
        self.ip = 0
        self.input = collections.deque([])
        self.output = collections.deque([])
        self.state = RUNNABLE
        self.pid = pid
        self.mem = nums.copy()
        self.relative_base = 0

    def runnable(self):
        if self.state == WAITING_INPUT and len(self.input) > 0:
            self.state = RUNNABLE

        return self.state == RUNNABLE

    def add_input(self, v):
        self.input.append(v)

    def memget(self, addr):
        if addr not in self.mem:
            self.mem[addr] = 0
        return self.mem[addr]

    def params(self, i, mode):
        a = self.ip + i
        v = self.memget(a)

        # mode 0:reference 1:immediate 2:relative
        mode = mode[i-1]
        match mode:
            case '0':
                return self.memget(v)
            case '1':
                return v
            case '2':
                return self.memget(self.relative_base + v)
            case _:
                raise Exception("unknown params mode", mode)

    def addr(self, i, mode):
        a = self.ip + i
        v = self.memget(a)

        # mode 0:reference 1:immediate 2:relative
        mode = mode[i-1]
        match mode:
            case '0':
                return v
            case '1':
                raise Exception("unexpected mode 1 for addr")
            case '2':
                return self.relative_base + v
            case _:
                raise Exception("unknown params mode", mode)

    def run(self):
        out = ''
        while True:
            # print('-', self.ip)
            num = self.memget(self.ip)
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

            # print("ip:{} num:{}, op:{}, mode:{}".format(self.ip, num, op, mode))

            match op:
                case 99:
                    self.state = DONE
                    return out
                case 1:
                    p1 = self.params(1, mode)
                    p2 = self.params(2, mode)
                    addr = self.addr(3, mode)
                    self.mem[addr] = p1 + p2
                    self.ip += 4
                case 2:
                    p1 = self.params(1, mode)
                    p2 = self.params(2, mode)
                    addr = self.addr(3, mode)
                    self.mem[addr] = p1 * p2
                    self.ip += 4
                case 3:
                    p = self.addr(1, mode)
                    # v = input('input: ')

                    # block if no input
                    if len(self.input) == 0:
                        self.state = WAITING_INPUT
                        return

                    v = self.input.popleft()
                    self.mem[p] = v
                    self.ip += 2
                case 4:
                    v = self.params(1, mode)
                    # v = self.mem[p]
                    print("{} > {}".format(self.pid, v))
                    out = v
                    self.output.append(v)
                    self.ip += 2
                case 5:
                    p1 = self.params(1, mode)
                    p2 = self.params(2, mode)
                    if p1 > 0:
                        self.ip = p2
                    else:
                        self.ip += 3
                case 6:
                    p1 = self.params(1, mode)
                    p2 = self.params(2, mode)
                    if p1 == 0:
                        self.ip = p2
                    else:
                        self.ip += 3
                case 7:
                    p1 = self.params(1, mode)
                    p2 = self.params(2, mode)
                    p3 = self.addr(3, mode)
                    self.mem[p3] = 1 if p1 < p2 else 0
                    self.ip += 4
                case 8:
                    p1 = self.params(1, mode)
                    p2 = self.params(2, mode)
                    p3 = self.addr(3, mode)
                    self.mem[p3] = 1 if p1 == p2 else 0
                    self.ip += 4
                case 9:
                    p1 = self.params(1, mode)
                    self.relative_base += p1
                    self.ip += 2

                case _:
                    print("invalid op", op, "self.ip:", self.ip)
                    raise Exception('invalid op', op)


def part1():
    c = IntComp(data)
    c.input.append(1)
    c.run()
    assert c.state == DONE
    return c.output.pop()


def part2():
    c = IntComp(data)
    c.input.append(2)
    c.run()
    assert c.state == DONE
    return c.output.pop()


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 3780860499

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 33343
