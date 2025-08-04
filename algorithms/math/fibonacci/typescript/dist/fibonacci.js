"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.fibonacci_recursive = fibonacci_recursive;
function fibonacci_recursive(n) {
    const num = typeof n === "bigint" ? n : BigInt(n);
    if (num < 0n) {
        throw new RangeError(`Fibonacci argument must be a non-negative integer, got ${n}`);
    }
    const calculate = (x) => {
        if (x <= 1n) {
            return x;
        }
        return calculate(x - 1n) + calculate(x - 2n);
    };
    return calculate(num);
}
//# sourceMappingURL=fibonacci.js.map