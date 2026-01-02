from typing import Protocol, runtime_checkable


class HashMismatchError(Exception):
    def __init__(self, message: str = "digest does not match the data") -> None:
        super().__init__(message)


@runtime_checkable
class Hasher(Protocol):
    def digest(self, data: bytes) -> bytes: ...
    def verify(self, digest: bytes, data: bytes) -> None: ...
