"use strict";

const fs = require("fs");

process.stdin.resume();
process.stdin.setEncoding("utf-8");

let inputString = "";
let currentLine = 0;

process.stdin.on("data", (inputStdin) => {
  inputString += inputStdin;
});

process.stdin.on("end", (_) => {
  inputString = inputString
    .trim()
    .split("\n")
    .map((str) => str.trim());

  main();
});

function readLine() {
  return inputString[currentLine++];
}

/*
 * Complete the getMoneySpent function below.
 */
function getMoneySpent(keyboards, drives, b) {
  /*
   * Write your code here.
   */
  let total = -1;
  keyboards.sort(function (a, b) {
    return a - b;
  });
  drives.sort(function (a, b) {
    return a - b;
  });
  if (keyboards[0] + drives[0] > b) return total;
  if (keyboards[keyboards.length - 1] + drives[drives.length - 1] <= b)
    return keyboards[keyboards.length - 1] + drives[drives.length - 1];
  for (let i = keyboards.length - 1; i >= 0; i--) {
    for (let j = drives.length - 1; j >= 0; j--) {
      if (keyboards[i] + drives[j] <= b && keyboards[i] + drives[j] > total)
        total = keyboards[i] + drives[j];
    }
  }

  return total;
}

function main() {
  const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

  const bnm = readLine().split(" ");

  const b = parseInt(bnm[0], 10);

  const n = parseInt(bnm[1], 10);

  const m = parseInt(bnm[2], 10);

  const keyboards = readLine()
    .split(" ")
    .map((keyboardsTemp) => parseInt(keyboardsTemp, 10));

  const drives = readLine()
    .split(" ")
    .map((drivesTemp) => parseInt(drivesTemp, 10));

  /*
   * The maximum amount of money she can spend on a keyboard and USB drive, or -1 if she can't purchase both items
   */

  let moneySpent = getMoneySpent(keyboards, drives, b);

  ws.write(moneySpent + "\n");

  ws.end();
}
