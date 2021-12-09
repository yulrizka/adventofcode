import fileinput
import unittest

data = fileinput.input("../../input/day8").readline().strip()
w = 25
h = 6

rows = [data[i:i + w] for i in range(0, len(data), w)]
layers = [rows[i:i + h] for i in range(0, len(rows), h)]


# print(layers)

def part1():
    mem = {}
    for l in layers:
        l1 = ''.join(l)
        mem[l1] = l1.count('0')

    # print(mem)
    img = min(mem, key=lambda x: mem[x])

    ans = img.count('1') * img.count('2')
    return ans


def combine(bottom, up):
    result = ''
    # if up == '0':
        # pass
    for i in range(w):
        # print(up, i)
        if up[i] == '2':
            result += bottom[i]
        else:
            result += up[i]
    return result


def part2():
    # print(layers)
    img = layers[-1]

    for i, layer in enumerate(reversed(layers)):
        if i == 0:
            continue

        # merge image with layer
        for j, pixels in enumerate(layer):
            img[j] = combine(img[j], pixels)

    for i in img:
        print(i.replace('0', ' '))


part2()


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 0

    def test2(self):
        ans = part2()
        print(ans)
