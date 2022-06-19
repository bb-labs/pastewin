import sys


def handler(event, context):
    print(event, context)

    return 'Hello from AWS Lambda using Python' + sys.version + '!'
