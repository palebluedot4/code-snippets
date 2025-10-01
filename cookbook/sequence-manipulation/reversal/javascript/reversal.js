export function reverseString(s) {
  if (typeof s !== "string") {
    throw new TypeError(
      `reverseString() argument must be string, got: ${typeof s}`
    );
  }
  return Array.from(s).reverse().join("");
}
