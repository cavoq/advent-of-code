# /usr/bin/env python3

FILE_IN = "input-test.dat"


class Pipe:
    def __init__(self, char: str, x: int, y: int):
        self.distance = -1
        self.char = char
        self.directions = directions(self.char)
        self.x = x
        self.y = y

    def __repr__(self):
        return f"Pipe(char={self.char}, x={self.x}, y={self.y})"


class Map:
    def __init__(self, width: int, height: int,
                 pipes: dict[tuple[int, int], Pipe],
                 start: tuple[int, int]):
        self.pipes = pipes
        self.start = start
        self.width = width
        self.height = height

    @classmethod
    def from_file(cls, file_path: str):
        pipes = {}
        with open(file_path, 'r') as file:
            lines = file.readlines()
            for y, line in enumerate(lines):
                for x, char in enumerate(line):
                    if char == " " or char == "\n":
                        continue
                    if char == "S":
                        start = (x, y)
                    pipes[(x, y)] = Pipe(char, x, y)
        return cls(x, y, pipes, start)

    def __repr__(self):
        return f"Map(width={self.width}, height={self.height}, pipes={self.pipes})"


def directions(char: str) -> str:
    if char == "|":
        return "NS"
    if char == "-":
        return "EW"
    if char == "L":
        return "NE"
    if char == "J":
        return "NW"
    if char == "7":
        return "SW"
    if char == "F":
        return "SE"
    return ""


def connection(Pipe1: Pipe, Pipe2: Pipe, dir: int) -> bool:
    if not Pipe1.directions or not Pipe2.directions:
        return False
    if dir == -1:
        if Pipe1.directions[0] == Pipe2.directions[0]:
            return True
    if dir == 1:
        if Pipe1.directions[1] == Pipe2.directions[1]:
            return True
    return False


def next_pipe(pipes: dict[tuple[int, int], Pipe], current_pipe: Pipe, dir: str) -> Pipe:
    if dir == "N":
        next_pipe = pipes[(current_pipe.x, current_pipe.y - 1)]
        if connection(current_pipe, next_pipe):
            next_pipe.distance = current_pipe.distance + 1
            current_pipe = next_pipe
    if dir == "S":
        next_pipe = pipes[(current_pipe.x, current_pipe.y + 1)]
        print(next_pipe)
        if connection(current_pipe, next_pipe):
            print(next_pipe)
            next_pipe.distance = current_pipe.distance + 1
            current_pipe = next_pipe
    if dir == "E":
        next_pipe = pipes[(current_pipe.x + 1, current_pipe.y)]
        if connection(current_pipe, next_pipe):
            next_pipe.distance = current_pipe.distance + 1
            current_pipe = next_pipe
    if dir == "W":
        next_pipe = pipes[(current_pipe.x - 1, current_pipe.y)]
        if connection(current_pipe, next_pipe):
            next_pipe.distance = current_pipe.distance + 1
            current_pipe = next_pipe
    return current_pipe


def part1(map: Map):
    valid_pipes = ["NS", "EW", "NE", "NW", "SW", "SE"]
    current_pipe = map.pipes[map.start]
    current_pipe.distance = 0

    loop = False
    for valid_pipe in valid_pipes:
        map.pipes[map.start].directions = valid_pipe
        current_dir = valid_pipe[1]
        while not loop:
            try:
                next = next_pipe(map.pipes, current_pipe, current_dir)
            except KeyError:
                continue
            if next.distance == -1 and current_pipe.distance == 0:
                continue
            if next.distance == 0 and current_pipe.distance != 0:
                loop = True
                return current_pipe.distance
            current_pipe = next
            current_dir = next.directions[1]


# Is not working yet
if __name__ == "__main__":
    map = Map.from_file(FILE_IN)
    res = part1(map)
    print(f"Part 1: {res}")
