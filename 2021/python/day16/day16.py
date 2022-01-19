import unittest

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
            print(grp, v)

        print(f'read % 4 {read %4}, i:{i}')
        if read % 4 != 0:
            i += (read %4) - 1

        v = int(bs, 2)
        print(f'>> literal: {v} i:{i}')
    else:
        # operator
        len_id = s[i]
        i += 1
    

        if len_id == '0':
            total_len = int(s[i:i+15], 2)
            i += 15

            last = i + total_len

            print(f'type 0: total_len={total_len}')
            
            while i < last:
                print(f'last: {last} , i:{i}')
                i = parse(s, i)

        else:
            # 11 bit are number sub package
            total_pkg = int(s[i:i+11], 2)
            i += 11
            print(f'type not 0: total_pkg={total_pkg}')

            for x in range(0,total_pkg):
                i = parse(s, i)
                

    return i


def part1():
    global total_version
    parse(hexToBin(data), 0)
    print(f'total version {total_version}')

part1()

def part2():
    ...


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
        assert ans == 0

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 0
