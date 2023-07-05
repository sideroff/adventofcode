const fs = require('fs');
const readline = require('readline');

const pathToFile = './input.txt';

const ROCK = "rock";
const PAPER = "paper";
const SCISORS = "scisors";

// bigger index wins over smaller one
const ORDER = [SCISORS, ROCK, PAPER];

const LOSE = "X";
const DRAW = "Y";
const WIN = "Z";

const RESULTS = [DRAW, WIN, LOSE];


const guide = {
  A: ROCK,
  B: PAPER,
  C: SCISORS,

  X: ROCK,
  Y: PAPER,
  Z: SCISORS,
}

const scores = {
  [ROCK]: 1,
  [PAPER]: 2,
  [SCISORS]: 3,

  [LOSE]: 0,
  [DRAW]: 3,
  [WIN]: 6,
}


function main() {
  console.log('processing file')
  const stream = fs.createReadStream(pathToFile);

  const rl = readline.createInterface({
    input: stream,
    crlfDelay: Infinity
  });

  let score = 0;
  rl.on('line', function (line) {
    const [opponent, me] = line.trim().split(" ");

    const opponentPlay = guide[opponent];
    const myPlay = guide[me];

    const opponentPlayIndex = ORDER.indexOf(opponentPlay);
    const myPlayIndex = ORDER.indexOf(myPlay);

    const result = RESULTS[((myPlayIndex - opponentPlayIndex) + 3) % 3];
    
    score += scores[myPlay] + scores[result];
  });

  rl.once("close", () => {
    console.log('end of file reached')
    console.log('total score', score)
  })
}

main();