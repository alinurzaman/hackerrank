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
 * Complete the 'serviceLane' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY cases
 */

function serviceLane(width, cases) {
  // Write your code here
  const max = [];
  cases.forEach((c) => {
    const slices = width.slice(c[0], c[1] + 1);
    max.push(Math.min(...slices));
  });
  return max;
}

function main() {
  const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

  const firstMultipleInput = readLine().replace(/\s+$/g, "").split(" ");

  const n = parseInt(firstMultipleInput[0], 10);

  const t = parseInt(firstMultipleInput[1], 10);

  const width = readLine()
    .replace(/\s+$/g, "")
    .split(" ")
    .map((widthTemp) => parseInt(widthTemp, 10));

  let cases = Array(t);

  for (let i = 0; i < t; i++) {
    cases[i] = readLine()
      .replace(/\s+$/g, "")
      .split(" ")
      .map((casesTemp) => parseInt(casesTemp, 10));
  }

  const result = serviceLane(width, cases);

  ws.write(result.join("\n") + "\n");

  ws.end();
}
