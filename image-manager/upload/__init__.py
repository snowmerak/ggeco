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

    accountName = os.environ.get('AZURE_STORAGE_ACCOUNT')
    storageKey = os.environ.get('AZURE_STORAGE_ACCESS_KEY')

    userId = req.params.get('user_id')
    storageName = req.params.get('storage_name')

    try:
        img = Image.open(req.files['image'])
        size = req.params.get('size')
        if size is None:
            size = '64'
        size = int(size)
    except Exception as e:
        return func.HttpResponse(
            "Failed to open image: " + str(e),
            status_code=400
        )

    originImg = img.copy()

    height = img.height
    width = img.width
    if height > width:
        img = img.crop((0, int((height - width) / 2), width, int((height + width) / 2)))
    else:
        img = img.crop((int((width - height) / 2), 0, int((width + height) / 2), height))
    img.thumbnail((size, size))

    account_url = f'https://{accountName}.blob.core.windows.net'

    try:
        with BlobServiceClient(account_url=account_url, credential=storageKey) as storage_client:
            with storage_client.get_container_client(storageName) as container_client:
                salt = os.urandom(16)
                hashed = blake3(salt + img.tobytes()).hexdigest()
                origin_image_name = f'{userId}/{hashed}.webp'
                thumbnail_image_name = f'{userId}/{hashed}.thumb.webp'

                print(origin_image_name)
                print(thumbnail_image_name)

                buffer = io.BytesIO()
                originImg.save(buffer, format="webp")
                container_client.upload_blob(
                    name=origin_image_name,
                    data=buffer.getvalue(),
                    overwrite=True
                )

                buffer = io.BytesIO()
                img.save(buffer, format="webp")
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
        'origin_image_url': f'{account_url}/{storageName}/{origin_image_name}',
        'thumbnail_image_url': f'{account_url}/{storageName}/{thumbnail_image_name}'
    }

    return func.HttpResponse(
        json.dumps(response),
        status_code=200,
        mimetype="image/webp"
    )

