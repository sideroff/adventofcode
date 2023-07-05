const fs = require('fs');
const readline = require('readline');

const pathToFile = './input.txt';

function readFileLineByLine(pathToFile, lineCallback, closeCallback) {
  console.log('processing file')
  const stream = fs.createReadStream(pathToFile);

  const rl = readline.createInterface({
    input: stream,
    crlfDelay: Infinity
  });

  rl.on('line', function (line) {
    lineCallback(line);
  });

  rl.once("close", () => {
    console.log('end of file reached')
    closeCallback()
  })
}

function p1() {
  let sum = 0;

  const handleLine = (line) => {
    const [[x,y],[x1,y1]] = line.split(',').map(x => x.split("-").map(y => parseInt(y)));
    if ((x<=x1 && y>=y1) || (x1<=x && y1>=y)){
      sum +=1;
    }
  }

  const handleClose = () => {
    console.log('total sum ', sum)
  }

  readFileLineByLine(pathToFile, handleLine, handleClose);
}

function p2() {
  let sum = 0;

  const handleLine = (line) => {
    const [[x,y],[x1,y1]] = line.split(',').map(x => x.split("-").map(y => parseInt(y)));
    if ((x<=x1 || y>=y1) || (x1<=x || y1>=y)){
      sum +=1;
    }
  }

  const handleClose = () => {
    console.log('total sum ', sum)
  }

  readFileLineByLine(pathToFile, handleLine, handleClose);
}

p2();