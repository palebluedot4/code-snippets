"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.factorialIterative = factorialIterative;
exports.factorialRecursive = factorialRecursive;
function factorialIterative(n) {
    const num = typeof n === "bigint" ? n : BigInt(n);
    if (num < 0n) {
        throw new RangeError(`Factorial argument must be a non-negative integer, got ${n}`);
    }
    if (num === 0n) {
        return 1n;
    }
    let result = 1n;
    for (let i = 2n; i <= num; i++) {
        result *= i;
    }
    return result;
}
function factorialRecursive(n) {
    const num = typeof n === "bigint" ? n : BigInt(n);
    if (num < 0n) {
        throw new RangeError(`Factorial argument must be a non-negative integer, got ${n}`);
    }
    const calculate = (x) => {
        if (x === 0n) {
            return 1n;
        }
        return x * calculate(x - 1n);
    };
    return calculate(num);
}
//# sourceMappingURL=factorial.js.map