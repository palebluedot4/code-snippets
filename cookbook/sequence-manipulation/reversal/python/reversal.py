from typing import List, Sequence, TypeVar

T = TypeVar("T")


def reverse_string(s: str) -> str:
    return s[::-1]


def reverse_list(lst: List[T]) -> List[T]:
    return lst[::-1]


def reversed_sequence(seq: Sequence[T]) -> List[T]:
    return list(seq[::-1])
