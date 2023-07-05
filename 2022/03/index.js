const fs = require('fs');
const readline = require('readline');

const pathToFile = './input.txt';

function getStringIntersection(t, o, breakOnFirstMatch) {

  let intersection = "";
  for (let i = 0; i < t.length; i++) {
    const c = t[i];
    if (o.includes(c)) { 
      intersection+=c;
      if (breakOnFirstMatch) break;
    }
  }

  console.log('returning ',breakOnFirstMatch, intersection)
  return intersection;
}

function main() {
  console.log('processing file')
  const stream = fs.createReadStream(pathToFile);

  const rl = readline.createInterface({
    input: stream,
    crlfDelay: Infinity
  });

  let sum = 0;

  rl.on('line', function (line) {
    let ranges = line.split(',');

    let [t, o] = ranges.reduce(r => r.)

  });

  rl.once("close", () => {
    console.log('end of file reached')
    console.log('total sum ', sum)
  })
}

main();