import datetime


def to_iso8601(dt: datetime.datetime) -> str:
    if dt.tzinfo is None:
        raise TypeError(
            "to_iso8601() argument must be timezone-aware datetime object, got naive datetime"
        )
    return dt.isoformat()


def to_datetime_string(dt: datetime.datetime) -> str:
    return dt.strftime("%Y-%m-%d %H:%M:%S")
