function gcd(a, b) {
  if (!Number.isInteger(a) || !Number.isInteger(b)) {
    throw new Error(
      `GCD arguments must be integers, got: a=${a} (${typeof a}), b=${b} (${typeof b})`
    );
  }
  a = Math.abs(a);
  b = Math.abs(b);
  while (b !== 0) {
    [a, b] = [b, a % b];
  }
  return a;
}
