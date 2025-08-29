export function reverseString(s: string): string {
  return Array.from(s).reverse().join("");
}

export function reverseArray<T>(arr: readonly T[]): T[] {
  return arr.toReversed();
}
