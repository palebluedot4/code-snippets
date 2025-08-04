import functools


@functools.lru_cache(maxsize=None)
def fibonacci_recursive(n: int) -> int:
    if not isinstance(n, int):
        raise TypeError(
            f"fibonacci_recursive() argument must be an integer, got {type(n).__name__}"
        )
    if n < 0:
        raise ValueError(
            f"fibonacci_recursive() argument must be a non-negative integer, got {n}"
        )
    return n if n <= 1 else fibonacci_recursive(n - 1) + fibonacci_recursive(n - 2)
