import fileinput
import unittest

# Player 1 starting position: 4
# Player 2 starting position: 8
# players = [4-1, 8-1]

# Player 1 starting position: 6
# Player 2 starting position: 1
players = [6 - 1, 1 - 1]

# players = [0, 0]

scores = [0, 0]


def part1():
    current = 0
    dice = 0
    roll = 0
    winner = 0
    while True:
        move = dice + 1 + dice + 2 + dice + 3
        dice += 3

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


mem = {}


def solve(players, score, cp, turns, move, limit, depth):
    key = (players, score, cp, turns, move)

    if cp == 0 and score[0] >= limit:
        return [1, 0]
    elif cp == 1 and score[1] >= limit:
        return [0, 1]

    if key in mem:
        return mem[key]

    pp = list(players)
    score = list(score)
    pp[cp] += move
    pp[cp] %= 10
    turns += 1

    if turns == 3:
        score[cp] += pp[cp] + 1  # record score (base 0)
        if score[cp] >= limit:
            return [1, 0] if cp == 0 else [0, 1]

        cp = 1 if cp == 0 else 0
        turns = 0

    ans = [0, 0]
    for x in range(1, 4):
        s = solve(tuple(pp), tuple(score), cp, turns, x, limit, depth + 1)
        ans[0] += s[0]
        ans[1] += s[1]

    # print(f'{key} -> {ans}')
    mem[key] = ans

    return ans


def part2():
    ans = [0, 0]
    for x in range(1, 4):
        wins = solve(tuple(players), (0, 0), 0, 0, x, 21, 0)
        ans[0] += wins[0]
        ans[1] += wins[1]

    return max(ans[0], ans[1])


class TestSum(unittest.TestCase):
    def test1(self):
        ans = part1()
        assert ans == 929625, f"got {ans}"

    def test2(self):
        ans = part2()
        assert ans == 175731756652760, f"got {ans}"
