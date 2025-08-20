"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isOdd = isOdd;
function isOdd(x) {
    if (typeof x === "bigint") {
        return x % 2n !== 0n;
    }
    return x % 2 !== 0;
}
//# sourceMappingURL=odd.js.map