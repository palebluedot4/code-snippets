export function isOdd(x: number | bigint): boolean {
  if (typeof x === "bigint") {
    return x % 2n !== 0n;
  }
  return x % 2 !== 0;
}

export function isOddBitwise(x: number | bigint): boolean {
  if (typeof x === "bigint") {
    return (x & 1n) !== 0n;
  }
  return (x & 1) !== 0;
}
