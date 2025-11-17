export function absValue<T extends number | bigint>(x: T): T {
  if (typeof x === "bigint") {
    return (x >= 0n ? x : -x) as T;
  }
  return Math.abs(x) as T;
}
