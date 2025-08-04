function fibonacciRecursive(n) {
  if (!Number.isInteger(n) || n < 0) {
    throw new Error(
      `Fibonacci argument must be a non-negative integer, got ${n}`
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
