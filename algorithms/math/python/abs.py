from typing import TypeVar, cast

T = TypeVar("T", int, float)


def abs_value(x: T) -> T:
    return cast(T, abs(x))


def _abs_value_manual(x: T) -> T:
    return x if x >= 0 else -x
