export function isEven(x) {
  const type = typeof x;
  if (type === "number") {
    if (!Number.isInteger(x)) {
      throw new TypeError(`isEven() argument must be integer, got: ${type}`);
    }
    return x % 2 === 0;
  }
  if (type === "bigint") {
    return x % 2n === 0n;
  }
  throw new TypeError(
    `isEven() argument must be number or bigint, got: ${type}`
  );
}

export function isEvenBitwise(x) {
  const type = typeof x;
  if (type === "number") {
    if (!Number.isInteger(x)) {
      throw new TypeError(
        `isEvenBitwise() argument must be integer, got: ${type}`
      );
    }
    return (x & 1) === 0;
  }
  if (type === "bigint") {
    return (x & 1n) === 0n;
  }
  throw new TypeError(
    `isEvenBitwise() argument must be number or bigint, got: ${type}`
  );
}
