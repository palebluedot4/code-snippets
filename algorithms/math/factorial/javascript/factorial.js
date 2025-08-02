function factorialRecursive(n) {
  if (!Number.isInteger(n) || n < 0) {
    throw new Error(
      `Factorial argument must be a non-negative integer, got ${n}`
    );
  }
  const calculate = (x) => {
    if (x === 0) {
      return 1n;
    }
    return BigInt(x) * calculate(x - 1);
  };
  return calculate(n);
}
