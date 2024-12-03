

INPUT_FILE = "input-test.dat"


def read_seed_block(input: str) -> list[str]:
    seed_block = []
    for line in input.splitlines():
        if ":" in line:
            continue
        seed_block.append(line)
    return seed_block


def read_data(input_file: str) -> list[list[str]]:
    with open(input_file) as f:
        data = f.read()

    blocks = data.split('\n\n')
    result = [read_seed_block(block) for block in blocks]

    # First line of seeds is empty
    result.remove([])

    return result


def extract_seeds(input: str) -> list[int]:
    seeds = []
    with open(input) as f:
        data = f.read()
    for line in data.splitlines():
        if line.startswith("seeds:"):
            line = line.removeprefix("seeds:").strip()
            seeds = [int(x) for x in line.split(" ")]
            break
    return seeds


def get_seed_pairs(seeds: list[int]) -> list[tuple[int, int]]:
    pairs = []
    for i in range(0, len(seeds), 2):
        pairs.append((seeds[i], seeds[i + 1]))
    return pairs


def travel(seed_blocks: list[list[str]], seed: int) -> int:
    for seed_block in seed_blocks:
        seed = process_seed_block(seed_block, seed)
    return seed


def process_seed_block(seed_block: list[str], seed: int) -> int:
    for line in seed_block:
        corr_num = get_corresponding_num(line, seed)
        if corr_num != -1:
            return corr_num
    return seed


def get_corresponding_num(line: str, inp_n: int) -> int:

    source_range_start = int(line.split()[1])
    destination_range_start = int(line.split()[0])
    range = int(line.split()[2])

    if inp_n < source_range_start or inp_n > source_range_start + range - 1:
        return -1

    return destination_range_start + inp_n - source_range_start


def part1(seed_blocks: list[list[str]], seeds: list[int]):
    min = -1
    for seed in seeds:
        loc = travel(seed_blocks, seed)
        if min == -1 or loc < min:
            min = loc
    print("Part 1:", min)


def reduce(seed_blocks: list[list[int]]):
    ranges = set()
    for block in seed_blocks:
        for line in block:
            source_start = int(line.split()[1])
            destination_start = int(line.split()[0])
            range_ = int(line.split()[2])
            ranges.add((source_start, source_start + range_, destination_start))
    return list(ranges)


# Unfortunatelly, too slow for the real input
def part2(seed_blocks: list[list[int]], seed_pairs: list[tuple[int, int]]):
    ranges = reduce(seed_blocks)
    ranges.sort()

    min = -1
    visited = set()
    for seed_pair in seed_pairs:
        end = seed_pair[0] + seed_pair[1] - 1
        for i in range(seed_pair[0], end + 1):
            if i in visited:
                continue
            loc = travel(seed_blocks, i)
            if min == -1 or loc < min:
                min = loc
            visited.add(i)
    
    print("Part 2:", min)


if __name__ == "__main__":
    seed_blocks = read_data(INPUT_FILE)
    seeds = extract_seeds(INPUT_FILE)
    seed_pairs = get_seed_pairs(seeds)
    part1(seed_blocks, seeds)
    part2(seed_blocks, seed_pairs)
