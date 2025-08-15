export function lcm(a: number, b: number): bigint {
  if (a === 0 || b === 0) {
    return 0n;
  }
  const absA = BigInt(Math.abs(a));
  const absB = BigInt(Math.abs(b));
  const gcd = (x: bigint, y: bigint): bigint => {
    while (y !== 0n) {
      [x, y] = [y, x % y];
    }
    return x;
  };
  return (absA / gcd(absA, absB)) * absB;
}
