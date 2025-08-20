function isEven(x) {
  const type = typeof x;
  if (type === "number") {
    if (!Number.isInteger(x)) {
      throw new Error(
        `isEven() argument must be an integer, got: ${x} (${type})`
      );
    }
    return x % 2 === 0;
  }
  if (type === "bigint") {
    return x % 2n === 0n;
  }
  throw new Error(
    `isEven() argument must be number or bigint, got: ${x} (${type})`
  );
}
