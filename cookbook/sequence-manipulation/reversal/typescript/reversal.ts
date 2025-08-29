export function reverseString(s: string): string {
  return Array.from(s).reverse().join("");
}

export function reverseArray<T>(arr: readonly T[]): T[] {
  return arr.toReversed();
}

export function reverseArrayInPlace<T>(arr: T[]): void {
  arr.reverse();
}

export function manualReverseArrayInPlace<T>(arr: T[]): void {
  let left = 0;
  let right = arr.length - 1;
  while (left < right) {
    [arr[left], arr[right]] = [arr[right], arr[left]];
    left++;
    right--;
  }
}
