"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.reverseString = reverseString;
exports.reverseArray = reverseArray;
function reverseString(s) {
    return Array.from(s).reverse().join("");
}
function reverseArray(arr) {
    return arr.toReversed();
}
//# sourceMappingURL=reversal.js.map