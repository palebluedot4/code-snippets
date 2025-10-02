function lcm(a, b) {
  if (!Number.isInteger(a) || !Number.isInteger(b)) {
    throw new TypeError(
      `lcm() arguments must be integers, got: a=${typeof a}, b=${typeof b}`
    );
  }
  if (a === 0 || b === 0) {
    return 0n;
  }
  const absA = BigInt(Math.abs(a));
  const absB = BigInt(Math.abs(b));
  const gcd = (x, y) => {
    while (y !== 0n) {
      [x, y] = [y, x % y];
    }
    return x;
  };
  return (absA / gcd(absA, absB)) * absB;
}
