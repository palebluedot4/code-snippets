"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isEven = isEven;
function isEven(x) {
    if (typeof x === "bigint") {
        return x % 2n === 0n;
    }
    return x % 2 === 0;
}
//# sourceMappingURL=even.js.map