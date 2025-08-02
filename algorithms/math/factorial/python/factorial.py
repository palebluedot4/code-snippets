import functools


def factorial_iterative(n: int) -> int:
    if not isinstance(n, int):
        raise TypeError(
            f"factorial_iterative() argument must be an integer, got {type(n).__name__}"
        )
    if n < 0:
        raise ValueError(
            f"factorial_iterative() argument must be a non-negative integer, got {n}"
        )
    result = 1
    for i in range(2, n + 1):
        result *= i
    return result


@functools.lru_cache(maxsize=None)
def factorial_recursive(n: int) -> int:
    if not isinstance(n, int):
        raise TypeError(
            f"factorial_recursive() argument must be an integer, got {type(n).__name__}"
        )
    if n < 0:
        raise ValueError(
            f"factorial_recursive() argument must be a non-negative integer, got {n}"
        )
    return 1 if n == 0 else n * factorial_recursive(n - 1)
