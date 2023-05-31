"use strict";

const fs = require("fs");

process.stdin.resume();
process.stdin.setEncoding("utf-8");

let inputString = "";
let currentLine = 0;

process.stdin.on("data", function (inputStdin) {
  inputString += inputStdin;
});

process.stdin.on("end", function () {
  inputString = inputString.split("\n");

  main();
});

function readLine() {
  return inputString[currentLine++];
}

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

function pickingNumbers(a) {
  // Write your code here
  a.sort(function (a, b) {
    return a - b;
  });

  let max = 0;
  let total = 0;
  let diffOne = false;
  for (let i = 0; i < a.length - 1; i++) {
    if (a[i + 1] - a[i] == 0) {
      total++;
    } else if (a[i + 1] - a[i] == 1 && !diffOne) {
      total++;
      diffOne = true;
    } else {
      total = 0;
      diffOne = false;
    }

    if (total > max) max = total;
  }
  return max + 1;
}

function main() {
  const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

  const n = parseInt(readLine().trim(), 10);

  const a = readLine()
    .replace(/\s+$/g, "")
    .split(" ")
    .map((aTemp) => parseInt(aTemp, 10));

  const result = pickingNumbers(a);

  ws.write(result + "\n");

  ws.end();
}
