"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isOdd = isOdd;
exports.isOddBitwise = isOddBitwise;
function isOdd(x) {
    if (typeof x === "bigint") {
        return x % 2n !== 0n;
    }
    return x % 2 !== 0;
}
function isOddBitwise(x) {
    if (typeof x === "bigint") {
        return (x & 1n) !== 0n;
    }
    return (x & 1) !== 0;
}
//# sourceMappingURL=odd.js.map