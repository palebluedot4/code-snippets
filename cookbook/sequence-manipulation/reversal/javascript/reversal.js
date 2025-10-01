export function reverseString(s) {
  if (typeof s !== "string") {
    throw new TypeError(
      `reverseString() argument must be string, got: ${typeof s}`
    );
  }
  return Array.from(s).reverse().join("");
}

export function reverseArray(arr) {
  if (!Array.isArray(arr)) {
    throw new TypeError(
      `reverseArray() argument must be array, got: ${typeof arr}`
    );
  }
  return arr.toReversed();
}
