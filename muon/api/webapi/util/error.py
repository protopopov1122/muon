from enum import Enum


class WebApiErrorCode(Enum):
    UnknownError = 0
    InvalidResource = 1
    RenderingError = 2


class WebApiError(BaseException):
    def __init__(self, code: WebApiErrorCode, data = None):
        self._code = code
        self._data = data

    def error_code(self)->WebApiErrorCode:
        return self._code

    def data(self)->str:
        return self._data
