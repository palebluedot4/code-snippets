export function gcd(a: number, b: number): number {
  let absA = Math.abs(a);
  let absB = Math.abs(b);
  while (absB !== 0) {
    [absA, absB] = [absB, absA % absB];
  }
  return absA;
}
