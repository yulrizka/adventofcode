package day22

import (
	"bufio"
	"io"
	"strconv"

	"github.com/yulrizka/adventofcode/pkg/aoc"
)

type winner int

const (
	p1Wins winner = iota
	p2Wins
)

func Part1(f io.Reader) (string, error) {
	player1, player2 := parse(f)

	// simulate the game
	for {
		if len(player1) == 0 || len(player2) == 0 {
			break
		}

		if player1[0] > player2[0] {
			player1 = append(player1[1:], player1[0], player2[0])
			player2 = player2[1:]
		} else {
			player2 = append(player2[1:], player2[0], player1[0])
			player1 = player1[1:]
		}
	}
	sum := calculateScore(player1, player2)

	return strconv.Itoa(sum), nil
}

func parse(f io.Reader) (player1 []byte, player2 []byte) {
	s := bufio.NewScanner(f)

	pt := &player1
	for s.Scan() {
		if s.Text() == "Player 1:" || s.Text() == "" {
			continue
		}
		if s.Text() == "Player 2:" {
			pt = &player2
			continue
		}
		n, err := strconv.Atoi(s.Text())
		aoc.NoError(err)
		*pt = append(*pt, byte(n))
	}
	return player1, player2
}

func calculateScore(player1 []byte, player2 []byte) int {
	var winner []byte
	if len(player1) == 0 {
		winner = player2
	} else {
		winner = player1
	}

	fact := len(winner)
	sum := 0
	for _, v := range winner {
		sum += int(v) * fact
		fact--
	}
	return sum
}

func Part2(f io.Reader) (string, error) {
	player1, player2 := parse(f)

	result1, result2, _ := play(player1, player2)

	sum := calculateScore(result1, result2)
	return strconv.Itoa(sum), nil
}

func play(player1 []byte, player2 []byte) ([]byte, []byte, winner) {
	var played = map[string]struct{}{}
loop:
	for {
		if len(player1) == 0 {
			return player1, player2, p2Wins
		}
		if len(player2) == 0 {
			return player1, player2, p1Wins
		}
		k := string(player1) + "||" + string(player2)
		if _, ok := played[k]; ok {
			// played before, player1 wins
			return player1, player2, p1Wins
		}
		played[k] = struct{}{}

		p1, p2 := player1[0], player2[0]
		if int(p1) >= len(player1) || int(p2) >= len(player2) {
			// at least one player must not have enough cards left in their deck to recurse;
			// the winner of the round is the player with the higher-value card.
			if p1 > p2 {
				player1 = append(player1[1:], player1[0], player2[0])
				player2 = player2[1:]
			} else {
				player2 = append(player2[1:], player2[0], player1[0])
				player1 = player1[1:]
			}
			continue loop
		}

		// both player has enough game for recursion
		a, b := copySlice(player1, player2)
		_, _, winner := play(a[1:p1+1], b[1:p2+1])
		if winner == p1Wins {
			player1 = append(player1[1:], player1[0], player2[0])
			player2 = player2[1:]
		} else {
			// player2 won
			player2 = append(player2[1:], player2[0], player1[0])
			player1 = player1[1:]
		}
	}
}

func copySlice(a, b []byte) (newa, newb []byte) {
	return []byte(string(a)), []byte(string(b))
}
