function gcd(a, b) {
  if (!Number.isInteger(a) || !Number.isInteger(b)) {
    throw new Error(
      `GCD arguments must be integers, got: a=${a} (${typeof a}), b=${b} (${typeof b})`
    );
  }
  let absA = Math.abs(a);
  let absB = Math.abs(b);
  while (absB !== 0) {
    [absA, absB] = [absB, absA % absB];
  }
  return absA;
}
