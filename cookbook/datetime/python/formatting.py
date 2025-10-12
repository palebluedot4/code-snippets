import datetime
from email.utils import formatdate


def to_iso8601(dt: datetime.datetime) -> str:
    if dt.tzinfo is None:
        raise TypeError(
            "to_iso8601() argument must be timezone-aware datetime object, got naive datetime"
        )
    return dt.isoformat()


def to_datetime_string(dt: datetime.datetime) -> str:
    return dt.strftime("%Y-%m-%d %H:%M:%S")


def to_date_string(dt: datetime.datetime) -> str:
    return dt.strftime("%Y-%m-%d")


def to_time_string(dt: datetime.datetime) -> str:
    return dt.strftime("%H:%M:%S")


def to_http_header(dt: datetime.datetime) -> str:
    if dt.tzinfo is None:
        raise TypeError(
            "to_http_header() argument must be timezone-aware datetime object, got naive datetime"
        )
    return formatdate(dt.timestamp(), usegmt=True)
