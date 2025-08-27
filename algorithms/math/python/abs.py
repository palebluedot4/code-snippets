from typing import TypeVar

T = TypeVar("T", int, float)


def abs_value(x: T) -> T:
    return abs(x)


def _manual_abs_value(x: T) -> T:
    return x if x >= 0 else -x
