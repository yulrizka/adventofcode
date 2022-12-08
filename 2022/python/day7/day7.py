import sys
import unittest

infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day7'
# infile = sys.argv[1] if len(sys.argv) > 1 else '../../input/day7.sample'
with open(infile) as f:
    data = f.read().strip()


# print(data)

def parse():
    wd = []
    folders = {}
    parsing_dir = False
    for line in data.split('\n'):
        if parsing_dir:
            if not line.startswith('$'):
                if line.startswith('dir'):
                    continue  # not worry about 'dir xxx'
                else:
                    size, name = line.split(' ')
                    size = int(size)

                    for i in range(0, len(wd)):
                        full_path = '/'.join(wd[:i + 1])
                        if full_path not in folders:
                            folders[full_path] = 0
                        folders[full_path] += size

        if line.startswith('$ cd'):
            path = line.removeprefix("$ cd ")
            if path == '..':
                wd.pop()
            else:
                wd.append(path)
            continue

        if line.startswith('$ ls'):
            parsing_dir = True
            continue

    return folders


def part1():
    folders = parse()
    total = 0
    for name, size in folders.items():
        if size < 100000:
            total += size

    return total


def part2():
    folders = parse()

    root = folders['/']
    avail = 70000000

    sizes = []
    for name, size in folders.items():
        sizes.append(size)
    sizes = sorted(sizes)
    for x in sizes:
        new_size = root - x
        free = avail - new_size
        # print(root, x , new_size, free)
        if free >= 30000000:
            return x


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 1513699, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 7991939, f'got {ans}'
