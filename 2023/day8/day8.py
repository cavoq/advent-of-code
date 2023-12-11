# /usr/bin/env python3

FILE_IN = "input.dat"


class Node:
    def __init__(self, value: str, left: str, right: str):
        self.value = value
        self.left = left
        self.right = right

    @classmethod
    def from_str(cls, line: str):
        value = line.strip(" ").split()[0]
        elements = line.split('(')[1].split(')')[0].split(', ')
        elements = [elem.strip() for elem in elements]
        return cls(value, elements[0], elements[1])

    def __repr__(self):
        return f"Node(value={self.value}, left={self.left}, right={self.right})"


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


def part1(nodes: dict[str, Node], steps: str):
    current_value = "AAA"
    i, current_step = 0, steps[0]
    total_steps = 0
    while current_value != "ZZZ":
        current_node = nodes[current_value]
        if current_step == "L":
            current_value = current_node.left
        else:
            current_value = current_node.right
        if i == len(steps) - 1:
            i = 0
        else:
            i += 1
        total_steps += 1
        current_step = steps[i]
        
    print("Part 1:", total_steps)


if __name__ == "__main__":
    nodes, steps = read_nodes(FILE_IN)
    part1(nodes, steps)