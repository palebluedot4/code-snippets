"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.reverseString = reverseString;
exports.reverseArray = reverseArray;
exports.reverseArrayInPlace = reverseArrayInPlace;
function reverseString(s) {
    return Array.from(s).reverse().join("");
}
function reverseArray(arr) {
    return arr.toReversed();
}
function reverseArrayInPlace(arr) {
    arr.reverse();
}
//# sourceMappingURL=reversal.js.map