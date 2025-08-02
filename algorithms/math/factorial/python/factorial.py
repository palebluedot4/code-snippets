import functools


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
