const matrix = [
  [0, 1, 1, 0, 1],
  [1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1],
  [0, 1, 1, 1, 0],
  [0, 1, 1, 1, 0]
];

// ===== LeetCode Solution =====
/**
 * @param {character[][]} matrix
 * @return {number}
 */
var maximalSquare = function(matrix) {
    if (!matrix || matrix.length === 0 || matrix[0].length === 0) {
        return 0;
    }

    const m = matrix.length;
    const n = matrix[0].length;

    // Build prefix sum array
    const countOne = Array(m).fill(0).map(() => Array(n).fill(0));

    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            let a = 0, b = 0, c = 0;

            if (i !== 0) {
                a = countOne[i - 1][j];
            }
            if (j !== 0) {
                b = countOne[i][j - 1];
            }
            if (i !== 0 && j !== 0) {
                c = countOne[i - 1][j - 1];
            }

            // Convert character to number (handle both string and number input)
            countOne[i][j] = a + b - c + (matrix[i][j] === '1' || matrix[i][j] === 1 ? 1 : 0);
        }
    }

    let maxSide = 0;

    // Try each position as top-left corner
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
          const maxPossibleSide = Math.min(m - i, n - j);
          console.log(maxPossibleSide)
            
            let c = 0;

            if (i !== 0 && j !== 0) {
                c = countOne[i - 1][j - 1];
            }

            // Try square sizes from largest to smallest
            for (let side = maxPossibleSide; side >= 1; side--) {
                const lowerRightRow = i + side - 1;
                const lowerRightCol = j + side - 1;

                const totalOne = countOne[lowerRightRow][lowerRightCol];

                let a = 0;
                if (i !== 0) {
                    a = countOne[i - 1][lowerRightCol];
                }

                let b = 0;
                if (j !== 0) {
                    b = countOne[lowerRightRow][j - 1];
                }

                // Count ones in the square
                const oneInSquare = totalOne - a - b + c;

                // If all cells are 1, update max
                if (oneInSquare === side * side) {
                    maxSide = Math.max(maxSide, side);
                    break;
                }
            }
        }
    }

    // Return area (side squared)
    return maxSide * maxSide;
};

console.log(maximalSquare(matrix))