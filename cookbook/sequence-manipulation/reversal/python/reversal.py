from typing import TypeVar

S = TypeVar("S", str, list, tuple)
T = TypeVar("T")


def reversed_sliceable(seq: S) -> S:
    return seq[::-1]


def reverse_list_in_place(lst: list[T]) -> None:
    lst.reverse()


def _reverse_list_in_place_manual(lst: list[T]) -> None:
    left, right = 0, len(lst) - 1
    while left < right:
        lst[left], lst[right] = lst[right], lst[left]
        left += 1
        right -= 1
