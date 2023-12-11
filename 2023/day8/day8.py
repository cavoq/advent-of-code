# /usr/bin/env python3

from math import lcm

FILE_IN = "input.dat"


class Node:
    def __init__(self, value: str, left: str, right: str):
        self.value = value
        self.left = left
        self.right = right

    @classmethod
    def from_str(cls, line: str):
        value = line.strip(" ").split()[0].strip()
        elements = line.split('(')[1].split(')')[0].split(', ')
        elements = [elem.strip() for elem in elements]
        return cls(value, elements[0], elements[1])

    def __repr__(self):
        return f"Node(value={self.value}, left={self.left}, right={self.right})\n"


def read_nodes(file_path: str) -> dict[str, Node]:
    nodes = {}
    with open(file_path, 'r') as file:
        lines = file.readlines()
        steps = lines[0].strip()
        lines = lines[1:]
        for line in lines:
            if len(line.strip()) == 0:
                continue
            node = Node.from_str(line)
            nodes[node.value] = node

    return nodes, steps


def lcm_of_list(numbers):
    lcm_result = numbers[0]
    for i in range(1, len(numbers)):
        lcm_result = lcm(lcm_result, numbers[i])
    return lcm_result


def part1(nodes: dict[str, Node], steps: str, start_value: str = "AAA", end_suffix: str = "ZZZ"):
    if start_value not in nodes:
        return -1

    i = 0
    total_steps = 0

    while not start_value.endswith(end_suffix):
        current_node = nodes[start_value]
        if steps[i] == "L":
            start_value = current_node.left
        if steps[i] == "R":
            start_value = current_node.right
        if i == len(steps) - 1:
            i = 0
        else:
            i += 1
        total_steps += 1

    return total_steps


def part2(nodes: dict[str, Node], steps: str):
    starting_values = {node.value for node in nodes.values()
                       if node.value.endswith("A")}

    min_steps = []
    for value in starting_values:
        min_steps.append(
            part1(nodes, steps, start_value=value, end_suffix="Z"))

    required_steps = lcm_of_list(min_steps)
    print("Part 2:", required_steps)


if __name__ == "__main__":
    nodes, steps = read_nodes(FILE_IN)
    total_steps = part1(nodes, steps)
    print("Part 1:", total_steps)
    part2(nodes, steps)
