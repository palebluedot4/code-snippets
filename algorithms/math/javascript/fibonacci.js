function fibonacci(n) {
  if (!Number.isInteger(n) || n < 0) {
    throw new Error(
      `fibonacci() argument must be a non-negative integer, got ${n}`
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
  if (!Number.isInteger(n) || n < 0) {
    throw new Error(
      `fibonacciRecursive() argument must be a non-negative integer, got ${n}`
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
