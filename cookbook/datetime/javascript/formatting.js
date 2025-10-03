/**
 * @param {Date} date
 * @returns {string}
 * @throws {TypeError}
 */
export function toISO8601(date) {
  if (!isValidDate(date)) {
    throw new TypeError(
      `toISO8601() argument must be valid Date object, got: ${typeof date}`
    );
  }
  return date.toISOString();
}

/**
 * @param {*} value
 * @returns {boolean}
 */
function isValidDate(value) {
  return value instanceof Date && !isNaN(value);
}
