# pyright: reportMissingImports=false
import boto3
import uuid
import io
import base64
from operator import itemgetter


s3 = boto3.client('s3')
sts = boto3.client('sts')


"""
Event getters
"""


def get_mime_type(event):
    return event['headers']['content-type'].split("/").pop()


def get_file_obj(event):
    return io.BytesIO(base64.b64decode(event['body']))


def get_s3_key(event):
    return "%s.%s" % (uuid.uuid4(), get_mime_type(event))


"""
Handler
"""


def handler(event, context):
    s3_key = get_s3_key(event)

    s3.upload_fileobj(
        Fileobj=get_file_obj(event),
        Bucket='pastewin-pastes',
        Key=s3_key
    )

    return s3_key
