class HashMismatchError(Exception):
    def __init__(self, message: str = "digest does not match the data") -> None:
        super().__init__(message)
