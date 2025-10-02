function fibonacci(n) {
  if (!Number.isInteger(n)) {
    throw new TypeError(
      `fibonacci() argument must be integer, got: ${typeof n}`
    );
  }
  if (n < 0) {
    throw new RangeError(
      `fibonacci() argument must be non-negative integer, got ${n}`
    );
  }
  if (n <= 1) {
    return BigInt(n);
  }
  let a = 0n;
  let b = 1n;
  for (let i = 2; i <= n; i++) {
    [a, b] = [b, a + b];
  }
  return b;
}

function fibonacciRecursive(n) {
  if (!Number.isInteger(n)) {
    throw new TypeError(
      `fibonacciRecursive() argument must be integer, got: ${typeof n}`
    );
  }
  if (n < 0) {
    throw new RangeError(
      `fibonacciRecursive() argument must be non-negative integer, got ${n}`
    );
  }
  const calculate = (x) => {
    if (x <= 1) {
      return BigInt(x);
    }
    return calculate(x - 1) + calculate(x - 2);
  };
  return calculate(n);
}
