import * as fs from 'fs';

fs.open("text.txt", 'r', function (err, f) {
    console.log(f);
  });

console.log('bonk')