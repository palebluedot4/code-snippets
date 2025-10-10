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
 * @param {Date} date
 * @param {string} [timeZone="UTC"]
 * @returns {string}
 * @throws {TypeError}
 * @throws {RangeError}
 */
export function toDateTimeString(date, timeZone = "UTC") {
  if (!isValidDate(date)) {
    throw new TypeError(
      `toDateTimeString() argument must be valid Date object, got: ${typeof date}`
    );
  }
  const options = {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
    hour12: false,
    timeZone: timeZone,
  };
  try {
    // The "sv-SE" locale conveniently provides the desired "YYYY-MM-DD HH:MM:SS" format.
    const formatter = new Intl.DateTimeFormat("sv-SE", options);
    return formatter.format(date);
  } catch (e) {
    if (e instanceof RangeError) {
      throw new RangeError(
        `toDateTimeString() argument timeZone must be valid IANA time zone string, got: ${timeZone}`
      );
    }
    throw e;
  }
}

/**
 * @param {*} value
 * @returns {boolean}
 */
function isValidDate(value) {
  return value instanceof Date && !isNaN(value);
}
