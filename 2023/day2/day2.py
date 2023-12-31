from typing import Dict


class Reveal:
    def __init__(self, blue: int, green: int, red: int):
        self.blue = int(blue)
        self.green = int(green)
        self.red = int(red)


class Game:
    def __init__(self, id: int, reveals: list[Reveal]):
        self.id = id
        self.reveals = reveals

    @classmethod
    def from_str(cls, line: str):
        data = line.split(' ')
        game_id = int(data[1].split(":")[0].strip())
        reveals_data = line.removeprefix(f"Game {game_id}: ").split(';')

        reveals = []
        for reveal_data in reveals_data:
            reveal_parts = reveal_data.split(',')
            rev = Reveal(0, 0, 0)
            for reveal in reveal_parts:
                number, color = reveal.strip().split(' ')
                if color == 'blue':
                    rev.blue = int(number)
                elif color == 'green':
                    rev.green = int(number)
                elif color == 'red':
                    rev.red = int(number)
            reveals.append(rev)

        return cls(game_id, reveals)


def read_games(file_path: str) -> Dict[int, Game]:
    games = {}
    with open(file_path, 'r') as f:
        lines = f.readlines()
        for line in lines:
            game = Game.from_str(line.strip())
            games[game.id] = game
    return games


def day2_part1():
    games_data = read_games('input/day2_input.txt')
    sum = get_sum_of_possible_games(games_data, 12, 13, 14)
    print("Part 1 solution: " + str(sum))


def day2_part2():
    games_data = read_games('input/day2_input.txt')
    total_sum_of_powers = 0
    for _, game in games_data.items():
        total_sum_of_powers += get_power_of_set_of_cubes(game)
    print("Part 2 solution: " + str(total_sum_of_powers))


def get_power_of_set_of_cubes(game: Game):

    max_red = 0
    max_green = 0
    max_blue = 0

    for reveal in game.reveals:
        if reveal.red > max_red:
            max_red = reveal.red
        if reveal.green > max_green:
            max_green = reveal.green
        if reveal.blue > max_blue:
            max_blue = reveal.blue

    return max_red * max_green * max_blue


def get_sum_of_possible_games(games: Dict[int, Game], num_red: int, num_green: int, num_blue: int) -> int:
    total_sum = 0
    for game_id, game in games.items():
        valid = True
        for reveal in game.reveals:
            if reveal.blue > num_blue or reveal.green > num_green or reveal.red > num_red:
                valid = False
                break

        if valid:
            total_sum += game_id

    return total_sum


if __name__ == '__main__':
    day2_part1()
    day2_part2()
