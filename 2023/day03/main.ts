const text = await Deno.readTextFile("example.txt");

// Define the type for a 2D grid of charactersZ
type Grid = string[][];

// Check if a character is a number (0-9)
const isNumber = (char: string): boolean => {
  return char >= "0" && char <= "9";
};

// Check if a character is a symbol (not a number and not a dot)
const isSymbol = (char: string): boolean => {
  return char !== "." && !isNumber(char);
};

// Get the list of adjacent cells (including diagonal)
const hasAdjacentSymbol = (
  grid: Grid,
  row: number,
  col: number,
  numberLength: number
): boolean => {
  const directions = [
    [-1, -1],
    [-1, 0],
    [-1, 1], // top-left, top, top-right
    [0, -1],
    [0, 1], // left,         right
    [1, -1],
    [1, 0],
    [1, 1], // bottom-left, bottom, bottom-right
  ];

  // Check adjacent cells for each part of the multi-digit number
  for (let offset = 0; offset < numberLength; offset++) {
    for (const [dx, dy] of directions) {
      const newRow = row + dx;
      const newCol = col + offset + dy;

      // Check if the new position is within the bounds of the grid
      if (
        newRow >= 0 &&
        newRow < grid.length &&
        newCol >= 0 &&
        newCol < grid[0].length
      ) {
        // Return true if an adjacent symbol is found
        if (isSymbol(grid[newRow][newCol])) {
          return true;
        }
      }
    }
  }

  return false;
};

// Function to calculate the sum of part numbers adjacent to symbols
const sumAdjacentNumbers = (grid: Grid): number => {
  let totalSum = 0;
  const rows = grid.length;

  for (let row = 0; row < rows; row++) {
    let col = 0;
    while (col < grid[row].length) {
      // If we find a digit, extract the full multi-digit number
      if (isNumber(grid[row][col])) {
        let numberStr = "";

        // Extract the full number
        while (col < grid[row].length && isNumber(grid[row][col])) {
          numberStr += grid[row][col];
          col++;
        }

        // Parse the number
        const number = parseInt(numberStr, 10);

        // Check if the number has any adjacent symbols
        if (
          hasAdjacentSymbol(grid, row, col - numberStr.length, numberStr.length)
        ) {
          totalSum += number;
        }
      } else {
        // Move to the next cell
        col++;
      }
    }
  }

  return totalSum;
};

const engineSchematic: Grid = text.split("\n").map((t) => t.split(""));

const result = sumAdjacentNumbers(engineSchematic);
console.log(result);
