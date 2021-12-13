import unittest

# with  open("../../input/day13-sample") as f:
with  open("../../input/day13") as f:
    points, ins = f.read().split("\n\n")
points = [[int(x) for x in i.split(',')] for i in points.split("\n")]
ins = [i.split('=') for i in ins.strip().replace("fold along ", "").split("\n")]


def solve(part):
    data = points.copy()
    it = 0
    done = False
    for axis, line in ins:
        line = int(line)
        if done:
            break
        it += 1
        for i, p in enumerate(data):
            if axis == 'x':
                x, y = p
                if x > line:
                    x = line - (x - line)
            if axis == 'y':
                x, y = p
                if y > line:
                    y = line - (y - line)
            data[i] = [x, y]

        clean = []
        for x in data:
            if x not in clean:
                clean.append(x)
        data = clean

        if part == 1 and it == 1:
            done = True

    return data


def part1():
    return len(solve(1))


def part2():
    points = solve(2)
    print(points)
    arr = [[' '] * 50 for i in range(50)]
    for x in points:
        arr[x[1]][x[0]] = 'X'

    for x in arr:
        print(''.join(x))


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 664

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 0
