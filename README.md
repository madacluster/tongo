# TONGO

[![Maintainability](https://api.codeclimate.com/v1/badges/295ba9c21a6d24345654/maintainability)](https://codeclimate.com/github/madacluster/tongo/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/295ba9c21a6d24345654/test_coverage)](https://codeclimate.com/github/madacluster/tongo/test_coverage)
## Usage

<<<<<<< HEAD

```bash

=======
```bash
>>>>>>> dae6c224e3f1b1735813ad5b8d48bfcfa295c67f
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