const fs = require('fs');
const readline = require('readline');

const pathToFile = './input.txt';

let elves = [0,0,0];

function main() {
  console.log('processing file')
  const stream = fs.createReadStream(pathToFile);

  const rl = readline.createInterface({
    input: stream,
    crlfDelay: Infinity
  });

  let calories = 0;
  rl.on('line', function (line) {
    if (line.trim() === "") {
      const smallestCalories = Math.min(...elves);

      if (calories > smallestCalories) {
        elves[elves.indexOf(smallestCalories)] = calories;
      }
      
      calories = 0;
      return;
    }

    calories += +line;
  });

  rl.once("close", () => {
    console.log('end of file reached')
    console.log('top 3 elves have a total of =', elves.reduce((acc, current) => acc+current, 0))
  })
}

main();