export function isEven(x: number | bigint): boolean {
  if (typeof x === "bigint") {
    return x % 2n === 0n;
  }
  return x % 2 === 0;
}
