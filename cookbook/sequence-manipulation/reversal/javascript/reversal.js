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

export function reverseArrayInPlace(arr) {
  if (!Array.isArray(arr)) {
    throw new TypeError(
      `reverseArrayInPlace() argument must be array, got: ${typeof arr}`
    );
  }
  arr.reverse();
}

export function reverseArrayInPlaceManual(arr) {
  if (!Array.isArray(arr)) {
    throw new TypeError(
      `reverseArrayInPlaceManual() argument must be array, got: ${typeof arr}`
    );
  }
  let left = 0;
  let right = arr.length - 1;
  while (left < right) {
    [arr[left], arr[right]] = [arr[right], arr[left]];
    left++;
    right--;
  }
}
