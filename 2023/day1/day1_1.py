def read_input():
    with open('input/day1_input.txt', 'r') as f:
        return [str(line) for line in f]


def calculate_calibration_value(str: str):

    first_digit = ""
    last_digit = ""

    for char in str:
        if char.isdigit():
            if not first_digit:
                first_digit = char
            last_digit = char

    if not first_digit or not last_digit:
        return 0

    return int(first_digit + last_digit)


def day1_part1():
    input = read_input()
    calibration_value = 0
    for str in input:
        calibration_value += calculate_calibration_value(str)
    print(calibration_value)


if __name__ == '__main__':
    day1_part1()
