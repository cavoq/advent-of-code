# /usr/bin/env python3

FILE_IN = "input.dat"


def read_time_distances(file_path: str) -> list[tuple[int, int]]:
    time_distance_pairs = []
    with open(file_path, 'r') as file:
        lines = file.readlines()
        time_values = list(map(int, lines[0].split()[1:]))
        distance_values = list(map(int, lines[1].split()[1:]))

        for time, distance in zip(time_values, distance_values):
            time_distance_pairs.append((time, distance))

    return time_distance_pairs


def read_as_whole_race(file_path: str) -> tuple[int, int]:
    with open(file_path, 'r') as file:
        lines = file.readlines()
        time = int("".join(lines[0].split()[1:]))
        distance = int("".join(lines[1].split()[1:]))
        return time, distance


def distance(velocity, time): return velocity * time


def part1(time_distances: list[tuple[int, int]]):
    total_ways = []
    for time, record_distance in time_distances:
        i = 0
        for t in range(1, time):
            s = distance(t, time - t)
            if s > record_distance:
                i += 1
                continue
        total_ways.append(i)

    res = 1
    for record in total_ways:
        res *= record

    print("Part 1:", res)


def part2(time, record_distance):
    ways = 0
    for t in range(1, time):
        s = distance(t, time - t)
        if s > record_distance:
            ways += 1
            continue
        
    print("Part 2:", ways)


if __name__ == "__main__":
    data = read_time_distances(FILE_IN)
    part1(data)
    data_p2 = read_as_whole_race(FILE_IN)
    part2(*data_p2)
