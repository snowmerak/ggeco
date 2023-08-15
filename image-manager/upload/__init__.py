import base64
import io
import json
import os

import azure.functions as func
from PIL import Image
from azure.storage.blob import BlobServiceClient
from blake3 import blake3


def main(req: func.HttpRequest) -> func.HttpResponse:
    match req.method:
        case 'PUT':
            pass
        case _:
            return func.HttpResponse(
                "Please PUT an image to resize",
                status_code=400
            )

    account_name = os.environ.get('AZURE_STORAGE_ACCOUNT')
    storage_key = os.environ.get('AZURE_STORAGE_ACCESS_KEY')

    user_id = req.params.get('user_id')
    storage_name = req.params.get('storage_name')

    try:
        thumbnail_img = Image.open(req.files['image'])
        size = req.params.get('size')
        if size is None:
            size = '64'
        size = int(size)
    except Exception as e:
        return func.HttpResponse(
            "Failed to open image: " + str(e),
            status_code=400
        )

    origin_img = thumbnail_img.copy()

    height = thumbnail_img.height
    width = thumbnail_img.width
    if height > width:
        thumbnail_img = thumbnail_img.crop((0, int((height - width) / 2), width, int((height + width) / 2)))
    else:
        thumbnail_img = thumbnail_img.crop((int((width - height) / 2), 0, int((width + height) / 2), height))
    thumbnail_img.thumbnail((size, size))

    account_url = f'https://{account_name}.blob.core.windows.net'

    try:
        with BlobServiceClient(account_url=account_url, credential=storage_key) as storage_client:
            with storage_client.get_container_client(storage_name) as container_client:
                salt = os.urandom(16)
                hashed = blake3(salt + thumbnail_img.tobytes()).hexdigest()
                origin_image_name = f'{user_id}/{hashed}.webp'
                thumbnail_image_name = f'{user_id}/{hashed}.thumb.webp'

                print(origin_image_name)
                print(thumbnail_image_name)

                buffer = io.BytesIO()
                origin_img.save(buffer, format="webp")
                container_client.upload_blob(
                    name=origin_image_name,
                    data=buffer.getvalue(),
                    overwrite=True
                )

                buffer = io.BytesIO()
                thumbnail_img.save(buffer, format="webp")
                container_client.upload_blob(
                    name=thumbnail_image_name,
                    data=buffer.getvalue(),
                    overwrite=True
                )
    except Exception as e:
        return func.HttpResponse(
            "Failed to upload image: " + str(e),
            status_code=500
        )

    response = {
        'origin_image_url': f'{account_url}/{storage_name}/{origin_image_name}',
        'thumbnail_image_url': f'{account_url}/{storage_name}/{thumbnail_image_name}'
    }

    return func.HttpResponse(
        json.dumps(response),
        status_code=200,
        mimetype="image/webp"
    )

