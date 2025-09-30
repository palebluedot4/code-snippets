import { reverseString } from "./reversal";

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
