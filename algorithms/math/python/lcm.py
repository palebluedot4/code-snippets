import math

from .gcd import _manual_gcd


def lcm(a: int, b: int) -> int:
    return math.lcm(a, b)


def _manual_lcm(a: int, b: int) -> int:
    if a == 0 or b == 0:
        return 0
    absA = abs(a)
    absB = abs(b)
    return (absA // _manual_gcd(absA, absB)) * absB
