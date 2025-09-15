"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.reverseString = reverseString;
exports.reverseArray = reverseArray;
exports.reverseArrayInPlace = reverseArrayInPlace;
exports.reverseArrayInPlaceManual = reverseArrayInPlaceManual;
function reverseString(s) {
    return Array.from(s).reverse().join("");
}
function reverseArray(arr) {
    return arr.toReversed();
}
function reverseArrayInPlace(arr) {
    arr.reverse();
}
function reverseArrayInPlaceManual(arr) {
    let left = 0;
    let right = arr.length - 1;
    while (left < right) {
        [arr[left], arr[right]] = [arr[right], arr[left]];
        left++;
        right--;
    }
}
//# sourceMappingURL=reversal.js.map