"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.absValue = absValue;
function absValue(x) {
    if (typeof x === "bigint") {
        return (x < 0n ? -x : x);
    }
    return Math.abs(x);
}
//# sourceMappingURL=abs.js.map