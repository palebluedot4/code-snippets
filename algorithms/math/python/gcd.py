import math


def gcd(a: int, b: int) -> int:
    return math.gcd(a, b)


def _gcd(a: int, b: int) -> int:
    if not isinstance(a, int) or not isinstance(b, int):
        raise TypeError(
            f"gcd() arguments must be integers, got: a={a} ({type(a)}), b={b} ({type(b)})"
        )
    absA = abs(a)
    absB = abs(b)
    while absB:
        absA, absB = absB, absA % absB
    return absA
