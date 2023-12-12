fs = require("fs");

var FILE_IN = "input.dat";

function readInput(path) {
  let arrays = [];
  let lines = fs.readFileSync(path, "utf8").split("\n");

  for (let i = 0; i < lines.length; i++) {
    let elements = lines[i].trim().split(" ").map(Number);
    arrays.push(elements);
  }

  return arrays;
}

function differenceSequence(sequence) {
  let diff = [];
  for (let i = 1; i < sequence.length; i++) {
    diff.push(sequence[i] - sequence[i - 1]);
  }
  return diff;
}

function extrapolateSequence(sequence) {
  let differences = [sequence];
  let sum = 0;

  while (!sequence.every((element) => element === 0)) {
    differences.push(differenceSequence(sequence));
    diff = differenceSequence(sequence);
    sequence = diff;
  }

  differences[differences.length - 1].push(0);
  while (differences.length > 1) {
    const last = differences[differences.length - 1];
    const secondLast = differences[differences.length - 2];
    const newElement =
      last[last.length - 1] + secondLast[secondLast.length - 1];

    differences.pop();
    differences[differences.length - 1].push(newElement);
  }

  return differences[0][differences[0].length - 1];
}

function part1() {
  let sequences = readInput(FILE_IN);
  let res = 0;
  for (let i = 0; i < sequences.length; i++) {
    let sequence = sequences[i];
    res += extrapolateSequence(sequence);
  }
  console.log("Part 1:", res);
}

part1();
