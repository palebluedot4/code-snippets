/**
 * @param {number|bigint} x
 * @returns {boolean}
 * @throws {TypeError}
 */
export function isOdd(x) {
  const type = typeof x;
  if (type === "number") {
    if (!Number.isInteger(x)) {
      throw new TypeError(`isOdd() argument must be integer, got: ${type}`);
    }
    return x % 2 !== 0;
  }
  if (type === "bigint") {
    return x % 2n !== 0n;
  }
  throw new TypeError(
    `isOdd() argument must be number or bigint, got: ${type}`
  );
}

/**
 * @param {number|bigint} x
 * @returns {boolean}
 * @throws {TypeError}
 */
export function isOddBitwise(x) {
  const type = typeof x;
  if (type === "number") {
    if (!Number.isInteger(x)) {
      throw new TypeError(
        `isOddBitwise() argument must be integer, got: ${type}`
      );
    }
    return (x & 1) !== 0;
  }
  if (type === "bigint") {
    return (x & 1n) !== 0n;
  }
  throw new TypeError(
    `isOddBitwise() argument must be number or bigint, got: ${type}`
  );
}
