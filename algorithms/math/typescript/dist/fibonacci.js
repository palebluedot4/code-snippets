"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.fibonacci = fibonacci;
exports.fibonacciRecursive = fibonacciRecursive;
function fibonacci(n) {
    const num = typeof n === "bigint" ? n : BigInt(n);
    if (num < 0n) {
        throw new RangeError(`fibonacci() argument must be a non-negative integer, got ${n}`);
    }
    if (num < 1n) {
        return num;
    }
    let a = 0n;
    let b = 1n;
    for (let i = 2n; i <= num; i++) {
        [a, b] = [b, a + b];
    }
    return b;
}
function fibonacciRecursive(n) {
    const num = typeof n === "bigint" ? n : BigInt(n);
    if (num < 0n) {
        throw new RangeError(`fibonacciRecursive() argument must be a non-negative integer, got ${n}`);
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