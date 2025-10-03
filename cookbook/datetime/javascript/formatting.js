/**
 * @param {*} value
 * @returns {boolean}
 */
function isValidDate(value) {
  return value instanceof Date && !isNaN(value);
}
