"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.gcd = gcd;
function gcd(a, b) {
    a = Math.abs(a);
    b = Math.abs(b);
    while (b !== 0) {
        [a, b] = [b, a % b];
    }
    return a;
}
//# sourceMappingURL=gcd.js.map