"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.gcd = gcd;
function gcd(a, b) {
    let absA = Math.abs(a);
    let absB = Math.abs(b);
    while (absB !== 0) {
        [absA, absB] = [absB, absA % absB];
    }
    return absA;
}
//# sourceMappingURL=gcd.js.map