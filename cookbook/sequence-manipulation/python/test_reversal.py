from collections.abc import Callable
from typing import Any, TypeVar

import pytest
import reversal

S = TypeVar("S", str, list, tuple, bytes)


@pytest.mark.parametrize(
    "sequence_input, expected_output",
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
        # bytes
        pytest.param(b"", b"", id="empty_bytes"),
        pytest.param(b"a", b"a", id="single_byte"),
        pytest.param(b"abcd", b"dcba", id="multiple_bytes"),
        pytest.param(b"\x01\x02\x03", b"\x03\x02\x01", id="hex_bytes"),
    ],
)
def test_reversed_sliceable_returns_new_reversed_sequence(
    sequence_input: S, expected_output: S
) -> None:
    original_copy = (
        sequence_input[:]
        if isinstance(sequence_input, (list, tuple))
        else sequence_input
    )

    actual_output = reversal.reversed_sliceable(sequence_input)

    assert actual_output == expected_output
    assert type(sequence_input) is type(actual_output)
    assert sequence_input == original_copy


@pytest.mark.parametrize(
    "reversal_function",
    [
        pytest.param(reversal.reverse_list_in_place, id="standard_in_place"),
        pytest.param(reversal._reverse_list_in_place_manual, id="manual_in_place"),
    ],
)
@pytest.mark.parametrize(
    ("list_input", "expected_output"),
    [
        pytest.param([], [], id="empty_list"),
        pytest.param([1], [1], id="single_element_list"),
        pytest.param([1, 2, 3, 4], [4, 3, 2, 1], id="even_elements_list"),
        pytest.param([1, 2, 3, 4, 5], [5, 4, 3, 2, 1], id="odd_elements_list"),
        pytest.param(["a", "b", "c"], ["c", "b", "a"], id="string_list"),
        pytest.param([1, "two", 3.0], [3.0, "two", 1], id="mixed_types_list"),
        pytest.param([None, True, False], [False, True, None], id="boolean_none_list"),
        pytest.param([1, 2, 2, 1], [1, 2, 2, 1], id="palindrome_list"),
    ],
)
def test_in_place_reversal_functions_modify_list(
    reversal_function: Callable[[list[Any]], None],
    list_input: list[Any],
    expected_output: list[Any],
) -> None:
    actual_output = list(list_input)

    reversal_function(actual_output)

    assert actual_output == expected_output
