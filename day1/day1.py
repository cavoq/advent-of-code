def read_input():
    with open('day1_input.txt', 'r') as f:
        return [str(line) for line in f]


def calculate_calibration_value(str: str):
    for char in str:
        if char.isdigit():
            first_digit = char
            break
    for char in str[::-1]:
        if char.isdigit():
            last_digit = char
            break
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
