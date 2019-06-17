import os
import logging
import json
from flask import Flask
from flask_cors import CORS

from webapi.endpoints.root import RootEndpoint
from webapi.endpoints.render import RenderBlueprint
from webapi.endpoints.static import StaticBlueprint

app = Flask(__name__)
cors = CORS(app)
ENV_NAME = 'WEBAPI_SETTINGS'
SETTINGS_PATH = os.environ.get(ENV_NAME) if ENV_NAME in os.environ.keys() else 'webapi.json'
with open(SETTINGS_PATH) as settings_file:
    settings = json.load(settings_file)
app.config['settings'] = settings

if settings['logging']:
    logging.basicConfig()

app.register_blueprint(RootEndpoint)
app.register_blueprint(RenderBlueprint)
if settings['static']['serve']:
    app.register_blueprint(StaticBlueprint)

if __name__ == '__main__':
    app.run(host=settings['host']['name'], port=settings['host']['port'])
