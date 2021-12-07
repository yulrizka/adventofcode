import collections
import fileinput
import itertools
import unittest

data = [int(i) for i in fileinput.input("../../input/day7").readline().strip().split(',')]

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

    def runnable(self):
        if self.state == WAITING_INPUT and len(self.input) > 0:
            self.state = RUNNABLE

        return self.state == RUNNABLE

    def add_input(self, v):
        self.input.append(v)

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
            # print('>', self.ip)
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

            # print("ip:{} num:{}, op:{}, mode:{}".format(self.ip, num, op, mode))

            match op:
                case 99:
                    self.state = DONE
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

                    # block if no input
                    if len(self.input) == 0:
                        self.state = WAITING_INPUT
                        return

                    v = self.input.popleft()
                    self.mem[p] = v
                    self.ip += 2
                case 4:
                    p = self.params(1, mode)
                    v = self.mem[p]
                    # print("{} > {}".format(self.id, v))
                    out = v
                    self.output.append(v)
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


class Scheduler:

    def __init__(self):
        self.procs = []

    def add(self, proc):
        self.procs.append(proc)

    def run(self):
        while len(self.procs) > 0:
            proc = self.procs[0]
            if not proc.runnable():
                self.procs = self.procs[1:] + [self.procs[0]]
                continue

            proc.run()
            if not proc.state == DONE:
                self.procs = self.procs[1:] + [self.procs[0]]
            else:
                self.procs = self.procs[1:]


def part1():
    vals = []
    nums = itertools.permutations([0, 1, 2, 3, 4])
    # print(nums)
    for num in nums:
        output = 0
        for i in range(5):
            comp = IntComp(data)
            comp.add_input(num[i])
            comp.add_input(output)
            comp.run()
            output = comp.output.popleft()

        vals.append(output)

    return max(vals)


def part2():
    vals = []
    nums = itertools.permutations([5, 6, 7, 8, 9])
    for num in nums:
        comps = [
            IntComp(data, 0),
            IntComp(data, 1),
            IntComp(data, 2),
            IntComp(data, 3),
            IntComp(data, 4),
        ]

        sched = Scheduler()

        for i, v in enumerate(num):
            # setup pipe. connect process output to the next process input
            comps[i].input = comps[i - 1].output = collections.deque([])
            comps[i].add_input(v)  # settings
            sched.add(comps[i])

        comps[0].add_input(0)  # initial value

        sched.run()
        val = comps[4].output.popleft()
        vals += [val]

    ans = max(vals)
    return ans


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        # print(ans)
        assert ans == 398674

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 39431233
