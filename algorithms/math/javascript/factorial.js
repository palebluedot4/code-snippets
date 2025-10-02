/**
 * @param {number} n
 * @returns {bigint}
 * @throws {TypeError}
 * @throws {RangeError}
 */
export function factorial(n) {
  if (!Number.isInteger(n)) {
    throw new TypeError(
      `factorial() argument must be integer, got: ${typeof n}`
    );
  }
  if (n < 0) {
    throw new RangeError(
      `factorial() argument must be non-negative integer, got ${n}`
    );
  }
  if (n === 0) {
    return 1n;
  }
  const nBigInt = BigInt(n);
  let result = 1n;
  for (let i = 2n; i <= nBigInt; i++) {
    result *= i;
  }
  return result;
}

/**
 * @param {number} n
 * @returns {bigint}
 * @throws {TypeError}
 * @throws {RangeError}
 */
export function factorialRecursive(n) {
  if (!Number.isInteger(n)) {
    throw new TypeError(
      `factorialRecursive() argument must be integer, got: ${typeof n}`
    );
  }
  if (n < 0) {
    throw new RangeError(
      `factorialRecursive() argument must be non-negative integer, got ${n}`
    );
  }
  const calculate = (x) => {
    if (x === 0) {
      return 1n;
    }
    return BigInt(x) * calculate(x - 1);
  };
  return calculate(n);
}
