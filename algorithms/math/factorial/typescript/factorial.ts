export function factorialRecursive(n: number | bigint): bigint {
  const num = typeof n === "bigint" ? n : BigInt(n);
  if (num < 0n) {
    throw new RangeError(
      `Factorial argument must be a non-negative integer, got ${n}`
    );
  }
  const calculate = (x: bigint): bigint => {
    if (x === 0n) {
      return 1n;
    }
    return x * calculate(x - 1n);
  };
  return calculate(num);
}
