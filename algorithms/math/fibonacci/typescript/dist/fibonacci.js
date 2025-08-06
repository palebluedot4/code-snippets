"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.fibonacci = fibonacci;
exports.fibonacci_recursive = fibonacci_recursive;
function fibonacci(n) {
    const num = typeof n === "bigint" ? n : BigInt(n);
    if (num < 0n) {
        throw new RangeError(`Fibonacci argument must be a non-negative integer, got ${n}`);
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