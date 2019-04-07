# 5xx
Parse log files for 5xx return codes.

#### Dependencies:
* python3

#### Installation:
```
git clone https://github.com/gitfu/5xx.git

cd 5xx

sudo install -m  0755 5xx.py /usr/local/bin/5xx.py
```

#### Usage:

```sh
# 5xx.py --help
usage: 5xx.py [-h] [-s START] [-e END] logs [logs ...]

Parse log files for 5xx http return codes.

positional arguments:
  logs                  List of log files to process

optional arguments:
  -h, --help            show this help message and exit
  -s START, --start START
                        Start time in UNIX time.
  -e END, --end END     End time in UNIX time.
```
