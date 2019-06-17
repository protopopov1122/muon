from flask import Blueprint, current_app, jsonify

RootEndpoint = Blueprint('root', __name__, '/api/v1')


@RootEndpoint.route('/api/v1', methods=['GET'])
def get_api_version():
    return jsonify({
        'version': current_app.config['settings']['version']
    })
