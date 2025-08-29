from typing import List, Sequence, TypeVar

T = TypeVar("T")


def reverse_string(s: str) -> str:
    return s[::-1]


def reverse_list(lst: List[T]) -> List[T]:
    return lst[::-1]


def reversed_sequence(seq: Sequence[T]) -> List[T]:
    return list(seq[::-1])


def reverse_list_in_place(lst: List[T]) -> None:
    lst.reverse()


def _manual_reverse_list_in_place(lst: List[T]) -> None:
    left, right = 0, len(lst) - 1
    while left < right:
        lst[left], lst[right] = lst[right], lst[left]
        left += 1
        right -= 1
