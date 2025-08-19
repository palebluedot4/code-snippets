import math


def gcd(a: int, b: int) -> int:
    return math.gcd(a, b)


def __gcd(a: int, b: int) -> int:
    absA = abs(a)
    absB = abs(b)
    while absB:
        absA, absB = absB, absA % absB
    return absA
