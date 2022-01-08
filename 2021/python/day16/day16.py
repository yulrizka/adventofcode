import unittest

with open('../../input/day16') as f:
    data = f.read().strip()

bitmap = {}

class BITS:
    def __init__(self, data):
        self.data = bin(int(data, 16))[2:]
        while len(self.data) != 4*len(data):
            self.data = '0'+self.data

def part1():
    b = BITS(data)



def part2():
    ...


class TestSum(unittest.TestCase):

    def test_sample(self):
        b = BITS('D2FE28')
        print(b.data)

        b = BITS('38006F45291200')
        assert b.data == "00111000000000000110111101000101001010010001001000000000"

        b = BITS('EE00D40C823060')
        assert b.data == '11101110000000001101010000001100100000100011000001100000'

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 0

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 0
