import base64
import os
import time

import jwt

jwtSecretKey = base64.decodebytes(os.environ.get('JWT_SECRET_KEY').encode('utf-8'))


def get_claims(token: str) -> dict:
    claims = jwt.decode(token, jwtSecretKey, algorithms=['HS512'])
    if 'expired_at' in claims and claims['expired_at'] < time.time():
        raise Exception('Token expired')
    return claims
