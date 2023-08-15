import io
import logging

import azure.functions as func

from PIL import Image


def main(req: func.HttpRequest) -> func.HttpResponse:
    logging.info('Python HTTP trigger function processed a request.')

    match req.method:
        case 'PUT':
            pass
        case _:
            return func.HttpResponse(
                "Please PUT an image to resize",
                status_code=400
            )

    try:
        img = Image.open(req.files['image'])
    except Exception as e:
        return func.HttpResponse(
            "Failed to open image: " + str(e),
            status_code=400
        )

    img.thumbnail((128, 128))

    buffer = io.BytesIO()
    img.save(buffer, format="webp")

    return func.HttpResponse(
        buffer.getvalue(),
        status_code=200,
        mimetype="image/webp"
    )
