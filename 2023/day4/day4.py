

INPUT_FILE = "input.txt"


class Card:
    def __init__(self, id: int, numbers: list[int], winning_numbers: list[int]):
        self.id = id
        self.numbers = numbers
        self.winning_numbers = winning_numbers
        self.won_copies = 1
        self.won_copies_indexes = []

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

    def set_won_copies_indexes(self) -> list[int]:
        index = 1
        for number in self.numbers:
            if number in self.winning_numbers:
                self.won_copies_indexes.append(self.id + index)
                index += 1
        return self.won_copies_indexes

    def get_won_copies_indexes(self) -> list[int]:
        return self.won_copies_indexes


def read_data() -> dict[int, Card]:
    cards = {}
    with open(INPUT_FILE) as f:
        lines = f.read().splitlines()
        for line in lines:
            card = Card.from_string(line)
            cards[card.id] = card
    return cards


def part1(cards: dict[int, Card]):
    total_score = 0
    for card in cards.values():
        total_score += card.get_score()
    print("Part 1:", total_score)


def part2(cards: dict[int, Card]):
    for card in cards.values():
        card.set_won_copies_indexes()
        indexes = card.get_won_copies_indexes()
        for index in indexes * card.won_copies:
            cards[index].won_copies += 1

    total_cards_won = sum(card.won_copies for card in cards.values())
    print("Part 2:", total_cards_won)


if __name__ == "__main__":
    cards = read_data()
    part1(cards)
    part2(cards)
