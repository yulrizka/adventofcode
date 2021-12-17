#!/usr/bin/env bash

if [ "$#" -ne 3 ]; then
  echo "use $0 [python] [year] [day]"
  exit 1
fi

year=$2
day=$3

function getInput() {
  aocd $day $year > "$year/input/day$day"
}

function python() {
  getInput
  mkdir -p "$year/python/day$day"
  cat <<EOF > "$year/python/day$day/day$day.py"
import fileinput
import unittest


with open('../../input/day$day') as f:
    data = f.read().strip()
print(data)


def part1():
    ...


def part2():
    ...


class TestSum(unittest.TestCase):

    def test1(self):
        ans = part1()
        print(ans)
        assert ans == 0

    def test2(self):
        ans = part2()
        print(ans)
        assert ans == 0
EOF

  echo "python"
}

case $1 in
  "python")
    python
    ;;

  *)
    echo "$1 not known"
    exit 1
    ;;

esac

