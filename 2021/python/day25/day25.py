import unittest

G = {}
with open('../../input/day25') as f:
# with open('../../input/day25-sample') as f:
    data = f.read().strip().split('\n')
    R = len(data)
    C = len(data[0])
    for r,row in enumerate(data):
        for c, v in enumerate(row):
            if v == '.':
                continue
            assert v == '>' or v == 'v'
            G[(r,c)] = v



def pprint(p):
    for r in range (0, R):
        for c in range (0,C):
            if (r,c) in p:
                print(p[(r,c)], end='')
            else:
                print('.', end='')
        print()

# pprint(G)

def part1():
    global G
    step = 0
    pprint(G)

    while True:
        g = {}
        step += 1
        move = 0
        for r,c in G:
            v = G[(r,c)]
            # print(r,c,v)
            if v == '>':
                cc = c + 1
                if cc >= C:
                    cc = 0
                if (r,cc) not in G:
                    g[(r,cc)] = v
                    move += 1
                else: 
                    g[(r,c)] = v
            else:
                g[(r,c)] = v
        G = g
        g = {}
        for r,c in G:
            v = G[(r,c)]

            if v == '>':
                g[(r,c)] = v

            else:
                assert v == 'v'
                rr = r + 1
                if rr >= R:
                    rr = 0 
                if (rr,c) not in G:
                    g[(rr,c)] = v
                    move += 1
                else: 
                    g[(r,c)] = v
        G = g
        print(step)
        # pprint(G)
        
        if move == 0:
            return step 

print('ans', part1())


def part2():
    ...


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 0, f'got {ans}'

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 0, f'got {ans}'
