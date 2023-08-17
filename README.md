# pastewin
Pastebin, but you foot the bill ;)

## Installation

Before using PasteWin, make sure you have Python installed on your system. PasteWin relies on the `boto3` library to interact with Amazon Web Services (AWS) S3. To install `boto3`, you can use the following command:

```bash
pip install boto3
```

Next, you can download the Pastewin via pip:

```
pip install pastewin
```

## Configuration

To use PasteWin, you'll need to have AWS credentials set up on your machine. The credentials should be configured in the standard AWS credential file located at `~/.aws/credentials`. If you haven't set up AWS credentials before, you can follow the guide [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html).

## Usage

PasteWin provides several commands that you can use to upload and share text content. To execute any of these commands, you need to use the `pastewin` command followed by the desired operation.

### Initializing

First you need an S3 bucket, try this:

```bash
pastewin init <your-bucket-name>
```

### Uploading

To upload and get a shareable link, use the `upload` command:

```bash
pastewin upload your/file.pdf
```

## Contributing

If you encounter any issues or have ideas for improvements, feel free to open an issue or submit a pull request on the [GitHub repository](https://github.com/bb-labs/pastewin). We welcome your contributions!
---

We hope that PasteWin makes sharing content a breeze for you. If you have any questions or need assistance, don't hesitate to reach out.

Happy sharing! 🚀
