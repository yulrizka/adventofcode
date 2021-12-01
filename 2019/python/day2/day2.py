import unittest

with open("../../input/day2") as f:
    lines = f.read()
    data = [int(line) for line in lines.split(',')]


class IntComp:
    initial_memory = []
    mem = []
    ip = 0
    noun = 0
    verb = 0

    def __init__(self, nums):
        self.initial_memory = nums.copy()
        self.mem = nums.copy()

    def initialise(self, noun, verb):
        self.mem[1] = noun
        self.mem[2] = verb

    def reset(self):
        self.mem = self.initial_memory.copy()
        self.ip = 0

    def run(self):
        while True:
            op = self.mem[self.ip]
            if op == 99:
                break
            elif op == 1:
                v1 = self.mem[self.mem[self.ip + 1]]
                v2 = self.mem[self.mem[self.ip + 2]]
                v3 = self.mem[self.ip + 3]
                self.mem[v3] = v1 + v2
            elif op == 2:
                v1 = self.mem[self.mem[self.ip + 1]]
                v2 = self.mem[self.mem[self.ip + 2]]
                v3 = self.mem[self.ip + 3]
                self.mem[v3] = v1 * v2
            else:
                print("invalid op", op, "self.ip:", self.ip)
                raise Exception('invalid op', op)
            self.ip += 4
        return self.mem[0]


def part2():
    c = IntComp(data)
    for x in range(90):
        for y in range(90):
            c.reset()
            c.initialise(x, y)
            if c.run() == 19690720:
                return 100 * x + y


class TestSum(unittest.TestCase):
    def test_sample(self):
        nums = [int(x) for x in "1,9,10,3,2,3,11,0,99,30,40,50".split(",")]
        c = IntComp(nums)
        assert c.run() == 3500

    def test_part1(self):
        c = IntComp(data)
        c.initialise(12, 2)
        assert c.run() == 6568671

    def test_part2(self):
        c = IntComp(data)
        assert part2() == 3951
