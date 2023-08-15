import io
import logging

import azure.functions as func

from PIL import Image


def main(req: func.HttpRequest) -> func.HttpResponse:
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
        size = req.params.get('size')
        if size is None:
            size = '256'
        size = int(size)
    except Exception as e:
        return func.HttpResponse(
            "Failed to open image: " + str(e),
            status_code=400
        )

    height = img.height
    width = img.width
    if height > width:
        img = img.crop((0, int((height-width)/2), width, int((height+width)/2)))
    else:
        img = img.crop((int((width-height)/2), 0, int((width+height)/2), height))
    img.thumbnail((size, size))

    print("Image size: " + str(img.size))

    buffer = io.BytesIO()
    img.save(buffer, format="webp")

    return func.HttpResponse(
        buffer.getvalue(),
        status_code=200,
        mimetype="image/webp"
    )
