import unittest

f = open("../../input/day4", 'r')
seq = [int(i) for i in f.readline().strip().split(',')]

n = 5  # num col & row

boards = []


class Board:
    b = []
    step = 0

    def __init__(self, b):
        self.b = b

    def mark(self, x):
        for row in range(n):
            for col in range(n):
                if self.b[row][col][0] == x:
                    self.b[row][col] = (x, True)

    def check(self):
        # check row
        for row in range(n):
            win = True
            for col in range(n):
                if not self.b[row][col][1]:
                    win = False
            if win:
                return True

        # check row
        for col in range(n):
            win = True
            for row in range(n):
                if not self.b[row][col][1]:
                    win = False
            if win:
                return True

    def score(self, x):
        # check row
        score = 0
        for row in range(n):
            for col in range(n):
                if not self.b[row][col][1]:
                   score += self.b[row][col][0]

        return score * x


while True:
    # read empty line
    line = f.readline()
    if not line:
        break

    board = [[0] * n for i in range(n)]
    for row in range(n):
        line = f.readline().strip().split(' ')
        line = list(filter(lambda x: x.strip() != '', line))
        nums = [int(i) for i in line]
        for col in range(n):
                v = nums[col]
                board[row][col] = (v, False)

    boards += [Board(board)]


def part1():
    for x in seq:
        for b in boards:
            b.mark(x)
            if b.check():
                return b.score(x)

    raise Exception("no solution found")


def part2():
    score_board = {}
    step = 0
    for x in seq:
        step += 1
        for b in boards:
            if b in score_board:
                continue

            b.mark(x)
            if b.check():
                score_board[b] = (step, b.score(x))

    winner = max(score_board, key=lambda k: score_board[k][0])
    return score_board[winner][1]


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 44736

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 8834787
