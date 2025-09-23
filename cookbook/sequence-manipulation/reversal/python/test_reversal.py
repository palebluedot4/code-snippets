from typing import TypeVar

import pytest
import reversal

S = TypeVar("S", str, list, tuple)


@pytest.mark.parametrize(
    "input_seq, expected",
    [
        # str
        pytest.param("", "", id="empty_str"),
        pytest.param("a", "a", id="single_character"),
        pytest.param("abcd", "dcba", id="ascii_str_with_even_length"),
        pytest.param("abcde", "edcba", id="ascii_str_with_odd_length"),
        pytest.param("你好，世界", "界世，好你", id="unicode_str"),
        pytest.param("hello world", "dlrow olleh", id="str_with_spaces"),
        pytest.param("Python ✅", "✅ nohtyP", id="str_with_emoji"),
        pytest.param("madam", "madam", id="palindrome_str"),
        # list
        pytest.param([], [], id="empty_list"),
        pytest.param([1], [1], id="single_element_list"),
        pytest.param([1, 2, 3, 4], [4, 3, 2, 1], id="even_elements_list"),
        pytest.param([1, 2, 3, 4, 5], [5, 4, 3, 2, 1], id="odd_elements_list"),
        pytest.param(["hello", "world"], ["world", "hello"], id="string_elements_list"),
        pytest.param([1, "two", 3.0], [3.0, "two", 1], id="mixed_types_list"),
        # tuple
        pytest.param((), (), id="empty_tuple"),
        pytest.param((1,), (1,), id="single_element_tuple"),
        pytest.param((1, 2, 3, 4), (4, 3, 2, 1), id="even_elements_tuple"),
        pytest.param((1, 2, 3, 4, 5), (5, 4, 3, 2, 1), id="odd_elements_tuple"),
        pytest.param(("a", "b", "c"), ("c", "b", "a"), id="string_elements_tuple"),
        pytest.param((1, "a", True), (True, "a", 1), id="mixed_types_tuple"),
    ],
)
def test_reversed_sliceable(input_seq: S, expected: S) -> None:
    original_copy = input_seq[:] if isinstance(input_seq, (list, tuple)) else input_seq

    actual = reversal.reversed_sliceable(input_seq)

    assert expected == actual
    assert type(input_seq) is type(actual)
    assert original_copy == input_seq
