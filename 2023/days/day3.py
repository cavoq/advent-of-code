import re
from itertools import chain


def part1(data, width, height, numberList):
    numbers = [n for n in numberList
               if any([is_symbol(data[srow][scol]) for srow, scol in get_adjacent(height, width, n)])]
    print("Part 1:", sum(x[3] for x in numbers))


def part2(data, width, height, numberList):
    stars = [(row, col) for row, line in enumerate(data)
             for col, ch in enumerate(line) if ch == '*']
    numberDict = {n: get_adjacent(height, width, n) for n in numberList}
    starDict = {star: [n for n, v in numberDict.items() if star in v]
                for star in stars}
    print("Part 2:", sum([x[0][3] * x[1][3]
          for x in [v for v in starDict.values() if len(v) == 2]]))


def extract_numbers(line, row):
    return [(row, x[1], len(x[0]), int(x[0]))
            for x in zip(re.findall(r'\d+', line), [x.start() for x in re.finditer(r'\d+', line)])]


def get_adjacent(height, width, number):
    adjacentCoords = [(number[0] - 1, number[1] - 1 + x) for x in range(number[2]+2)] + \
        [(number[0] + 1, number[1] - 1 + x) for x in range(number[2]+2)] + \
        [(number[0], number[1] - 1), (number[0], number[1] + number[2])]
    return [x for x in adjacentCoords if x[0] >= 0 and x[0] < height and x[1] >= 0 and x[1] < width]


def is_symbol(ch): return not ch.isdigit() and ch != '.'


def read_data():
    with open("input.txt") as f:
        return f.read().splitlines()


if __name__ == "__main__":
    data = read_data()
    width, height = len(data[0]), len(data)
    numberList = list(chain(*[extract_numbers(line, row)
                      for row, line in enumerate(data)]))
    part1(data, width, height, numberList)
    part2(data, width, height, numberList)
