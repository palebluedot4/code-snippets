function gcd(a, b) {
  if (!Number.isInteger(a) || !Number.isInteger(b)) {
    throw new TypeError(
      `gcd() arguments must be integers, got: a=${typeof a}, b=${typeof b}`
    );
  }
  let absA = Math.abs(a);
  let absB = Math.abs(b);
  while (absB !== 0) {
    [absA, absB] = [absB, absA % absB];
  }
  return absA;
}
