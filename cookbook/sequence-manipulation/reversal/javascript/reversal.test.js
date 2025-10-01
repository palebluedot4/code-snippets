import { describe, expect, test } from "vitest";

import { reverseArray, reverseString } from "./reversal.js";

describe("reverseString", () => {
  const testCases = [
    { name: "empty string", input: "", expected: "" },
    { name: "single character", input: "a", expected: "a" },
    { name: "ASCII string with even length", input: "abcd", expected: "dcba" },
    { name: "ASCII string with odd length", input: "abcde", expected: "edcba" },
    { name: "Unicode string", input: "你好，世界", expected: "界世，好你" },
    {
      name: "string with spaces",
      input: "hello world",
      expected: "dlrow olleh",
    },
    {
      name: "string with emoji",
      input: "JavaScript ✅",
      expected: "✅ tpircSavaJ",
    },
    { name: "palindrome string", input: "madam", expected: "madam" },
  ];

  test.each(testCases)(
    "should reverse the string for: $name",
    ({ input, expected }) => {
      const actual = reverseString(input);

      expect(actual).toBe(expected);
    }
  );
});

describe("reverseArray", () => {
  const testCases = [
    { name: "empty array", input: [], expected: [] },
    { name: "single element array", input: [1], expected: [1] },
    {
      name: "even elements array",
      input: [1, 2, 3, 4],
      expected: [4, 3, 2, 1],
    },
    {
      name: "odd elements array",
      input: [1, 2, 3, 4, 5],
      expected: [5, 4, 3, 2, 1],
    },
    {
      name: "string elements array",
      input: ["a", "b", "c"],
      expected: ["c", "b", "a"],
    },
    {
      name: "mixed primitive and reference types array",
      input: [1, "two", { three: 3 }, [4], undefined, null],
      expected: [null, undefined, [4], { three: 3 }, "two", 1],
    },
  ];

  test.each(testCases)(
    "should return a new reversed array for: $name",
    ({ input, expected }) => {
      const originalArray = [...input];

      const actual = reverseArray(originalArray);

      expect(actual).toEqual(expected);
      expect(originalArray).toEqual(input);
      if (input.length > 0) {
        expect(actual).not.toBe(originalArray);
      }
    }
  );
});
