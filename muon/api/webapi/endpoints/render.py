import pathlib
import os
from flask import Blueprint, current_app
import mistune
from webapi.util.middleware import endpoint
from webapi.util.error import WebApiError, WebApiErrorCode


RenderBlueprint = Blueprint('renderer', __name__, '/api/v1/render')


def get_page_path(path):
    root_dir = pathlib.Path(current_app.config['settings']['render']['root']).resolve()
    extension = current_app.config['settings']['render']['extension']
    full_path = (root_dir / '.'.join([path, extension])).resolve()
    dir_path = (root_dir / path).resolve()
    if not os.path.commonpath([root_dir, full_path]) == str(root_dir):
        raise WebApiError(WebApiErrorCode.InvalidResource, {
            'description': f'Access to resource \'{path}\' denied'
        })
    if full_path.exists() and full_path.is_file():
        return full_path
    elif dir_path.exists() and dir_path.is_dir():
        return dir_path
    else:
        raise WebApiError(WebApiErrorCode.InvalidResource, {
            'description': f'Resource \'{path}\' not found'
        })


def render_file(file_path: pathlib.Path, url: str, length_limit: int = -1):
    preview_cut = current_app.config['settings']['render']['preview']['cut']
    try:
        complete_file = True
        with file_path.open() as page:
            file_content = page.read(length_limit)
            if length_limit >= 0:
                current_pos = page.tell()
                page.seek(0, 2)
                end_pos = page.tell()
                if current_pos < end_pos:
                    complete_file = False
                    file_content += preview_cut
        return {
            'url': url,
            'complete': complete_file,
            'content': mistune.markdown(file_content, escape=False),
            'fixed': False
        }
    except BaseException as ex:
        raise WebApiError(WebApiErrorCode.RenderingError, {
            'description': 'Resource rendering error'
        })


def file_name_without_extension(file: pathlib.Path):
    return file.name[:-len(file.suffix)]


def render_directory(dir_path: pathlib.Path, url: str):
    extension = current_app.config['settings']['render']['extension']
    index_name = current_app.config['settings']['render']['directory']['index']
    header_name = current_app.config['settings']['render']['directory']['header']
    footer_name = current_app.config['settings']['render']['directory']['footer']
    preview_length = current_app.config['settings']['render']['preview']['length']

    index_path = dir_path / index_name
    if index_path.exists():
        return [render_file(index_path, f'{url}/{file_name_without_extension(index_path)}')]
    else:
        entries = [render_file(file_path, f'{url}/{file_name_without_extension(file_path)}', preview_length)
            for file_path in dir_path.iterdir()
            if file_path.is_file() and file_path.suffix == f'.{extension}']
        header_path = dir_path / header_name
        footer_path = dir_path / footer_name
        if header_path.exists():
            file = render_file(header_path, '')
            file['fixed'] = True
            entries.insert(0, file)
        if footer_path.exists():
            file = render_file(footer_path, '')
            file['fixed'] = True
            entries.append(file)
        return entries


@RenderBlueprint.route('/api/v1/render/<path:path>', methods=['GET'])
@endpoint()
def api_render_file(path):
    resource_path = get_page_path(path)
    if resource_path.is_file():
        return [render_file(resource_path, path)]
    elif resource_path.is_dir():
        return render_directory(resource_path, path)
    else:
        raise WebApiError(WebApiErrorCode.InvalidResource, {
            'description': f'Resource \'{path}\' not found'
        })
