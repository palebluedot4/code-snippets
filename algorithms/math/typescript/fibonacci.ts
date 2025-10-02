export function fibonacci(n: number | bigint): bigint {
  const num = typeof n === "bigint" ? n : BigInt(n);
  if (num < 0n) {
    throw new RangeError(
      `fibonacci() argument must be non-negative integer, got ${n}`
    );
  }
  if (num < 1n) {
    return num;
  }
  let a = 0n;
  let b = 1n;
  for (let i = 2n; i <= num; i++) {
    [a, b] = [b, a + b];
  }
  return b;
}

export function fibonacciRecursive(n: number | bigint): bigint {
  const num = typeof n === "bigint" ? n : BigInt(n);
  if (num < 0n) {
    throw new RangeError(
      `fibonacciRecursive() argument must be non-negative integer, got ${n}`
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
