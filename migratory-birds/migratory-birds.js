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
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

function migratoryBirds(arr) {
  // Write your code here
  arr.sort(function (a, b) {
    return a - b;
  });
  let bird = arr[0];
  let currentBird = bird;
  let total = 1;
  let currentTotal = total;
  for (let i = 1; i < arr.length; i++) {
    if (currentBird != arr[i]) {
      currentBird = arr[i];
      if (currentTotal > total) total = currentTotal;
      currentTotal = 1;
    } else {
      currentTotal++;
    }
    if (currentTotal > total) bird = currentBird;
  }
  return bird;
}

function main() {
  const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

  const arrCount = parseInt(readLine().trim(), 10);

  const arr = readLine()
    .replace(/\s+$/g, "")
    .split(" ")
    .map((arrTemp) => parseInt(arrTemp, 10));

  const result = migratoryBirds(arr);

  ws.write(result + "\n");

  ws.end();
}
