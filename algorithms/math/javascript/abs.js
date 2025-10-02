function absValue(x) {
  const type = typeof x;
  if (type === "number") {
    return Math.abs(x);
  }
  if (type === "bigint") {
    return x >= 0n ? x : -x;
  }
  throw new TypeError(
    `absValue() argument must be number or bigint, got: ${type}`
  );
}
