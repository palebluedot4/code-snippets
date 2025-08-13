from typing import TypeVar

T = TypeVar("T", int, float)


def abs_value(x: T) -> T:
    return abs(x)


def _abs_value(x: T) -> T:
    if not isinstance(x, (int, float)):
        raise TypeError(
            f"abs_value() argument must be an int or float, got {type(x).__name__}"
        )
    return x if x >= 0 else -x
