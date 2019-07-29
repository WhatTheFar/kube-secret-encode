# kube-secret-encode

## Installation

```bash
$ go get -u github.com/WhatTheFar/kube-secret-encode/cmd/kube-secret-encode
```

## Usage

```bash
$ kube-secret-encode < your-secrets.yaml > base64secrets.yaml
```

For example,

```bash
$ kube-secret-encode < test/secrets.yaml > base64secrets.yaml
```
