from typing import TypeVar

T = TypeVar("T", int, float)


def abs_value(x: T) -> T:
    return abs(x)
