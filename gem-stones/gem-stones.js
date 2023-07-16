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
 * Complete the 'gemstones' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING_ARRAY arr as parameter.
 */

function gemstones(arr) {
  let list = [];
  let first = arr[0];
  for (let i = 0; i < first.length; i++) {
    if (list.indexOf(first.charAt(i)) == -1) list.push(first.charAt(i));
  }
  for (let i = 1; i < arr.length; i++) {
    let letter = arr[i];
    list = list.filter((c) => letter.indexOf(c) != -1);
  }
  return list.length;
}

function main() {
  const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

  const n = parseInt(readLine().trim(), 10);

  let arr = [];

  for (let i = 0; i < n; i++) {
    const arrItem = readLine();
    arr.push(arrItem);
  }

  const result = gemstones(arr);

  ws.write(result + "\n");

  ws.end();
}
