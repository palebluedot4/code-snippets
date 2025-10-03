import datetime


def to_iso8601(dt: datetime.datetime) -> str:
    return dt.isoformat()
