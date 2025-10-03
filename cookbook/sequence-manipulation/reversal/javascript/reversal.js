/**
 * @param {string} str
 * @returns {string}
 * @throws {TypeError}
 */
export function reverseString(str) {
  if (typeof str !== "string") {
    throw new TypeError(
      `reverseString() argument must be string, got: ${typeof str}`
    );
  }
  return Array.from(str).reverse().join("");
}

/**
 * @template T
 * @param {T[]} arr
 * @returns {T[]}
 * @throws {TypeError}
 */
export function reverseArray(arr) {
  if (!Array.isArray(arr)) {
    throw new TypeError(
      `reverseArray() argument must be array, got: ${typeof arr}`
    );
  }
  return arr.toReversed();
}

/**
 * @param {*[]} arr
 * @returns {void}
 * @throws {TypeError}
 */
export function reverseArrayInPlace(arr) {
  if (!Array.isArray(arr)) {
    throw new TypeError(
      `reverseArrayInPlace() argument must be array, got: ${typeof arr}`
    );
  }
  arr.reverse();
}

/**
 * @param {*[]} arr
 * @returns {void}
 * @throws {TypeError}
 */
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
