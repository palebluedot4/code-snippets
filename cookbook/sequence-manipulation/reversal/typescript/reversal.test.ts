import { reverseArray, reverseString } from "./reversal";

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
      input: "TypeScript ✅",
      expected: "✅ tpircSepyT",
    },
    { name: "palindrome string", input: "madam", expected: "madam" },
  ];

  test.each(testCases)(
    "should correctly reverse: $name",
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
      name: "mixed type elements array",
      input: [1, "two", 3.0],
      expected: [3.0, "two", 1],
    },
    {
      name: "null and undefined elements array",
      input: [1, null, "2", undefined],
      expected: [undefined, "2", null, 1],
    },
  ];

  test.each(testCases)(
    "should reverse an $name without modifying the original",
    ({ input, expected }) => {
      const originalArray = [...input];

      const actual = reverseArray(input);

      expect(actual).toEqual(expected);
      expect(input).toEqual(originalArray);
    }
  );
});
