export function fibonacci_recursive(n: number | bigint): bigint {
  const num = typeof n === "bigint" ? n : BigInt(n);
  if (num < 0n) {
    throw new RangeError(
      `Fibonacci argument must be a non-negative integer, got ${n}`
    );
  }
  const calculate = (x: bigint): bigint => {
    if (x <= 1n) {
      return x;
    }
    return calculate(x - 1n) + calculate(x - 2n);
  };
  return calculate(num);
}
