export function toISO8601(date: Date): string {
  return date.toISOString();
}

export function toDateTimeString(date: Date, timeZone: string = "UTC"): string {
  const options: Intl.DateTimeFormatOptions = {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
    hour12: false,
    timeZone: timeZone,
  };
  let formatter: Intl.DateTimeFormat;
  try {
    formatter = new Intl.DateTimeFormat("sv-SE", options);
  } catch (e) {
    if (e instanceof RangeError) {
      throw new RangeError(
        `toDateTimeString() argument timeZone must be valid IANA time zone string, got: ${timeZone}`
      );
    }
    throw e;
  }
  return formatter.format(date);
}
