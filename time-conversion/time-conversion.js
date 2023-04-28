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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

function timeConversion(s) {
  // Write your code here
  let hours = s.substring(0, 2);
  let minutesAndSeconds = s.substring(2, 8);
  let meridiem = s.substring(s.length - 2);

  let finalHours = hours;
  if (meridiem === "PM" && hours !== "12")
    finalHours = String(parseInt(hours, 10) + 12);
  else if (meridiem === "AM" && hours === "12") finalHours = "00";

  return finalHours + minutesAndSeconds;
}

function main() {
  const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

  const s = readLine();

  const result = timeConversion(s);

  ws.write(result + "\n");

  ws.end();
}
