import math


def lcm(a: int, b: int) -> int:
    return math.lcm(a, b)


def _lcm(a: int, b: int) -> int:
    if a == 0 or b == 0:
        return 0
    absA = abs(a)
    absB = abs(b)
    return (absA // math.gcd(absA, absB)) * absB
