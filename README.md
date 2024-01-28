# DontPad File Uploader

Recently I was watching a video about a guy who start uploading files to youtube, so I thought it would be cool to have a script to upload files to dontpad.
Why not?

## table of contents

- [Description](#description)
- [Architecture / Pipeline](#architecture--pipeline)
- [Getting Started](#getting-started)
- [tech stack](#tech-stack)
- [Config](#config)
- [Endpoints](#endpoints)

## Description

This API will upload a file to dontpad.com, you can use it to upload images, videos, pdfs, etc.
All files will be encrypted and checksumed before being uploaded and download.

## Architecture / Pipeline

![image](https://github.com/publi0/dontpad-storage/assets/14155185/f1839154-9623-4196-a3b9-9cab3b870873)


## Getting Started

Start by running the go server:

```bash
go run cmd/main.go
```

Then you can upload a file using curl:

```bash
curl -X PUT -F "file=@/Users/publio/Downloads/IMG_20190818_123456.jpg" http://localhost:8080/upload
```

## Endpoints:

- PUT /files
  - Description: Upload a file
  - Body: multipart/form-data
  - Params: file
- GET /files
  - Description: List all files
  - Response: json
  - Params: none
- GET /files/{id}
  - Description: Get a file
  - Response: file
  - Params: file id
- DELETE /files/{id}
  - Description: Delete a file
  - Response: json
  - Params: file id

> Obs: Inside the `/requests` folder you can find a insomnia collection with all requests.

## Tech stack

- Go
- Sqlite

## Config

- You can change the encryption key in the config file (/resources/key.txt)
