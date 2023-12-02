digit_map = {
    "one": "one1one",
    "two": "two2two",
    "three": "three3three",
    "four": "four4four",
    "five": "five5five",
    "six": "six6six",
    "seven": "seven7seven",
    "eight": "eight8eight",
    "nine": "nine9nine",
}


def read_input():
    with open('input/day1_input_part2.txt', 'r') as f:
        return [str(line) for line in f]


def calculate_calibration_value(input_str: str):

    first_digit = ""
    last_digit = ""

    for char in input_str:
        if char.isdigit():
            if not first_digit:
                first_digit = char
            last_digit = char

    if first_digit == "" or last_digit == "":
        return 0

    return int(first_digit + last_digit)


def replace_words_with_digits(input_str: str):
    for word in digit_map.keys():
        input_str = input_str.replace(word, digit_map.get(word))
    return input_str


def day1_part2():
    input = read_input()
    calibration_value = 0
    for str in input:
        new_str = replace_words_with_digits(str)
        calibration_value += calculate_calibration_value(new_str)
    print(calibration_value)


if __name__ == '__main__':
    day1_part2()
