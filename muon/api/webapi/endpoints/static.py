import os
from flask import Blueprint, current_app, send_from_directory
from webapi.util.middleware import endpoint

StaticBlueprint = Blueprint('static', __name__, '/api/v1/static')

@StaticBlueprint.route('/api/v1/static/<path:path>')
@endpoint(wrap_json=False)
def serve_static(path):
    static_root = os.path.abspath(current_app.config['settings']['static']['root'])
    return send_from_directory(static_root, path)
