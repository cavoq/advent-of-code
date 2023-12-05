
INPUT_FILE = "input.txt"


class Card:
    def __init__(self, id: int, numbers: list[int], winning_numbers: list[int]):
        self.id = id
        self.numbers = numbers
        self.winning_numbers = winning_numbers

    @classmethod
    def from_string(cls, string: str):
        card_values = string.split(":")[1].strip()
        id = int(string.split(":")[0].strip().split()[1].strip())
        winning_numbers = [int(x) for x in card_values.split('|')[
            0].split() if x.strip()]
        numbers = [int(x)
                   for x in card_values.split("|")[1].split() if x.strip()]
        return cls(id, numbers, winning_numbers)

    def __repr__(self):
        return f"Card(id={self.id}, numbers={self.numbers}, winning_numbers={self.winning_numbers})"

    def get_score(self):
        multiplier = 1
        score = 0
        for number in self.numbers:
            if number in self.winning_numbers:
                if multiplier == 1:
                    multiplier += 1
                    score += 1
                    continue
                score *= multiplier
        return score


def read_data() -> list[Card]:
    with open(INPUT_FILE) as f:
        lines = f.read().splitlines()
    return [Card.from_string(line) for line in lines]


def part1(cards: list[Card]):
    total_score = 0
    for card in cards:
        total_score += card.get_score()
    print("Part 1:", total_score)


if __name__ == "__main__":
    cards = read_data()
    score = part1(cards)
