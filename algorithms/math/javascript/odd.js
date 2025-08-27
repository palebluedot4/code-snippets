function isOdd(x) {
  const type = typeof x;
  if (type === "number") {
    if (!Number.isInteger(x)) {
      throw new Error(
        `isOdd() argument must be an integer, got: ${x} (${type})`
      );
    }
    return x % 2 !== 0;
  }
  if (type === "bigint") {
    return x % 2n !== 0n;
  }
  throw new Error(
    `isOdd() argument must be number or bigint, got: ${x} (${type})`
  );
}

function isOddBitwise(x) {
  const type = typeof x;
  if (type === "number") {
    if (!Number.isInteger(x)) {
      throw new Error(
        `isOddBitwise() argument must be an integer, got: ${x} (${type})`
      );
    }
    return (x & 1) !== 0;
  }
  if (type === "bigint") {
    return (x & 1n) !== 0n;
  }
  throw new Error(
    `isOddBitwise() argument must be number or bigint, got: ${x} (${type})`
  );
}
