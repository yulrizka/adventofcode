import fileinput
import unittest

# Player 1 starting position: 4
# Player 2 starting position: 8
# players = [4-1, 8-1]

# Player 1 starting position: 6
# Player 2 starting position: 1
players = [6-1, 1-1]

scores = [0,0]

def part1():
    current = 0
    dice = 0
    roll = 0
    winner = 0
    while True:
        move = dice+1 + dice+2 + dice +3
        dice  += 3
            
        players[current] += move 
        players[current] = players[current] % 10
        scores[current] += players[current] + 1

        # print(f'{current +1} move {move}, dice {dice}, space {players[current]+1}, scores {scores[current]}')

        if scores[current] >= 1000:
            looser = current + 1 % 2
            break

        if current == 0:
            current = 1
        else:
            current = 0

    ans = scores[looser] * dice
    return ans



part1()

def part2():
    ...


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        assert ans == 0, f'got {ans}'

    def test2(self):
        ans = part2()
        assert ans == 0, f'got {ans}'
