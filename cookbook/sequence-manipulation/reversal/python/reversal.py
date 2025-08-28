from typing import List, TypeVar

T = TypeVar("T")


def reverse_string(s: str) -> str:
    return s[::-1]


def reverse_list(lst: List[T]) -> List[T]:
    return lst[::-1]
