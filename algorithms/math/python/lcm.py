import math

from algorithms.math.python.gcd import _gcd_manual


def lcm(a: int, b: int) -> int:
    return math.lcm(a, b)


def _lcm_manual(a: int, b: int) -> int:
    if a == 0 or b == 0:
        return 0
    absA = abs(a)
    absB = abs(b)
    return (absA // _gcd_manual(absA, absB)) * absB
