import boto3
import json
import uuid
import io

s3 = boto3.client('s3')
sts = boto3.client('sts')


def handler(event, context):
    print(sts.get_caller_identity())

    fileobj = io.BytesIO(event['body'].encode())
    bucket = 'pastewin-pastes'
    key = "%s.json" % uuid.uuid4()

    s3.upload_fileobj(fileobj, bucket, key)

    return key
