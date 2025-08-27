"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isEven = isEven;
exports.isEvenBitwise = isEvenBitwise;
function isEven(x) {
    if (typeof x === "bigint") {
        return x % 2n === 0n;
    }
    return x % 2 === 0;
}
function isEvenBitwise(x) {
    if (typeof x === "bigint") {
        return (x & 1n) === 0n;
    }
    return (x & 1) === 0;
}
//# sourceMappingURL=even.js.map