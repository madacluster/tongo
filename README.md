# TONGO

[![codecov](https://codecov.io/gh/madacluster/tongo/branch/master/graph/badge.svg?token=85EDGEIMR9)](https://codecov.io/gh/madacluster/tongo) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/c2780bae303b43bf86d8771ad6d041f1)](https://www.codacy.com/gh/madacluster/tongo/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=madacluster/tongo&amp;utm_campaign=Badge_Grade)
## Usage

```bash
Vote several time to menti.com/

Usage:
  TONGO [flags]

Flags:
  -h, --help         help for TONGO
  -l, --loop int     times to echo the input (default 1)
  -u, --url string   url (required) Ex: https://www.menti.com/1ct2pwd8ba
  -v, --value int    times to echo the input (default 1)

```

### Docker

```bash
docker builder -t tongo .

docker run \
    -e TONGO_MENTI_URL=https://menti.com/sdfasdf \
    tongo
```