import functools


def fibonacci(n: int) -> int:
    if n < 0:
        raise ValueError(
            f"fibonacci_iterative() argument must be a non-negative integer, got {n}"
        )
    if n <= 1:
        return n
    a, b = 0, 1
    for _ in range(2, n + 1):
        a, b = b, a + b
    return b


@functools.lru_cache(maxsize=None)
def fibonacci_recursive(n: int) -> int:
    if n < 0:
        raise ValueError(
            f"fibonacci_recursive() argument must be a non-negative integer, got {n}"
        )
    return n if n <= 1 else fibonacci_recursive(n - 1) + fibonacci_recursive(n - 2)
