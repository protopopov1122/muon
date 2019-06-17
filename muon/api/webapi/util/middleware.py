from functools import wraps
from flask import jsonify, current_app
from webapi.util.error import WebApiError, WebApiErrorCode


def endpoint(wrap_json=True):
    def internal(fn):
        @wraps(fn)
        def proc(*args, **kwargs):
            try:
                result = fn(*args, **kwargs)
                if wrap_json:
                    return jsonify({
                        'result': result
                    })
                else:
                    return result
            except WebApiError as ex:
                return jsonify({
                    'error': {
                        'code': ex.error_code().value,
                        'data': ex.data()
                    }
                })
            except BaseException as ex:
                current_app.logger.exception(ex)
                return jsonify({
                    'error': {
                        'code': WebApiErrorCode.UnknownError.value,
                        'data': {
                            'description': 'Unknown error'
                        }
                    }
                })
        return proc
    return internal
