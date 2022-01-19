import unittest
from functools import reduce
from operator import mul

with open('../../input/day16') as f:
    data = f.read().strip()

bitmap = {}


def hexToBin(h):
    v = bin(int(h, 16))[2:]
    while len(v) % 4 != 0:
         v = '0' + v
    return v

total_version = 0

def parse(s, i):
    global total_version

    print(f'i:{i} parse {s[i:]},')
    version = int(s[i:i+3], 2)
    i += 3
    total_version += version

    typ = int(s[i:i+3], 2)
    i += 3

    print(f'version: {version}, typ: {typ}')

    if typ == 4:
        # literal value
        done = False
        bs = ''
        read = 0
        while not done:
            grp = s[i]
            i += 1

            v = s[i:i+4]
            bs += v
            i += 4

            read = 5

            done = grp == '0'

        if read % 4 != 0:
            i += (read %4) - 1

        v = int(bs, 2)
        print(f'>> literal: {v} i:{i}')

        return i, v
    else:
        # operator
        len_id = s[i]
        i += 1
    
        nums = []
        if len_id == '0':
            total_len = int(s[i:i+15], 2)
            i += 15

            last = i + total_len

            print(f'type 0: total_len={total_len}')
            
            while i < last:
                i, v  = parse(s, i)
                nums.append(v)

        else:
            # 11 bit are number sub package
            total_pkg = int(s[i:i+11], 2)
            i += 11
            # print(f'type not 0: total_pkg={total_pkg}')

            for x in range(0,total_pkg):
                i, v = parse(s, i)
                nums.append(v)
        
        match typ:
            case 0: # sum
                return i, sum(nums)
            case 1: # product
                return i, reduce(mul, nums)
            case 2: # min
                return i, min(nums)
            case 3: # max
                return i, max(nums)
            case 5: # gt
                return i, 1 if nums[0] > nums[1] else 0
            case 6: # lt
                return i, 1 if nums[0] < nums[1] else 0
            case 7: # eq
                return i, 1 if nums[0] == nums[1] else 0
            

    return i


def part1():
    global total_version
    i, _ = parse(hexToBin(data), 0)
    return total_version

def part2():
    i, ans = parse(hexToBin(data), 0)
    return ans



class TestSum(unittest.TestCase):

    def test_sample(self):
        b = hexToBin('D2FE28')
        assert b == "110100101111111000101000", f'got {b}'
    
        b = hexToBin('38006F45291200')
        assert b == "00111000000000000110111101000101001010010001001000000000", f'got {b}'
    
        b = hexToBin('EE00D40C823060')
        assert b == "11101110000000001101010000001100100000100011000001100000", f'got {b}'
        

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 895

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 1148595959144
