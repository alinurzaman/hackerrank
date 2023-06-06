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
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

function climbingLeaderboard(ranked, player) {
  // Write your code here
  const rank = [];
  let listWithoutDuplicates = [...new Set(ranked)];
  let index = listWithoutDuplicates.length - 1;
  for (let i = 0; i < player.length; i++) {
    for (let j = index; j >= 0; j--) {
      if (player[i] >= listWithoutDuplicates[0]) {
        rank.push(1);
        break;
      } else if (player[i] < listWithoutDuplicates[j]) {
        rank.push(j + 2);
        index = j;
        break;
      } else if (player[i] == listWithoutDuplicates[j]) {
        rank.push(j + 1);
        index = j;
        break;
      }
    }
  }

  return rank;
}

function main() {
  const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

  const rankedCount = parseInt(readLine().trim(), 10);

  const ranked = readLine()
    .replace(/\s+$/g, "")
    .split(" ")
    .map((rankedTemp) => parseInt(rankedTemp, 10));

  const playerCount = parseInt(readLine().trim(), 10);

  const player = readLine()
    .replace(/\s+$/g, "")
    .split(" ")
    .map((playerTemp) => parseInt(playerTemp, 10));

  const result = climbingLeaderboard(ranked, player);

  ws.write(result.join("\n") + "\n");

  ws.end();
}
