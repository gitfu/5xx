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
* with a start time, stop time, and one log file.
```sh
a@futronic:~$ 5xx.py -s 1493969101 -e 1493969102 20.log

   20.log ... complete     

  Between time 1493969101.000 and time 1493969102.000:
	player.vimeo.com returned 40.00%  5xx errors 
	vimeo.com returned 0.00%  5xx errors 
```
* with a start time, stop time, and two log files.
```bash
a@futronic:~$ 5xx.py -s 1493969101 -e 1493969102 20.log 20a.log
   20.log ... complete     
   20a.log ... complete     

  Between time 1493969101.000 and time 1493969102.000:
	player.vimeo.com returned 40.00%  5xx errors 
	vimeo.com returned 20.00%  5xx errors 
```
* with a start time, stop time, and several globbed files
```sh
a@futronic:~$ 5xx.py  -s 1493969101 -e 1493969102  *.data
   20a.data ... complete     
   20b.data ... complete     
   20.data ... complete     
   1.data ... complete     
   2.data ... complete     
   3.data ... complete     
   out1.data ... complete     
   out.data ... complete     

  Between time 1493969101.000 and time 1493969102.000:
	player.vimeo.com returned 33.33%  5xx errors 
	vimeo.com returned 40.00%  5xx errors 
```
* if start time is omited, it defaults to 1.0
```
a@futronic:~$ 5xx.py  -e 1493969102  1.data 2.data 3.data
   1.data ... complete     
   2.data ... complete     
   3.data ... complete     

  Between time      1.000 and time 1493969102.000:
	player.vimeo.com returned 33.33%  5xx errors 
	vimeo.com returned 40.00%  5xx errors 
```
* if stop time is omitted, it defaults to the current time
```sh
a@futronic:~$ 5xx.py  -s 1493969101 20a.log 20.log
   20a.log ... complete     
   20.log ... complete     

  Between time 1493969101.000 and time 1554663335.094:
	player.vimeo.com returned 40.00%  5xx errors 
	vimeo.com returned 20.00%  5xx errors 
```
* invalid files are skipped
```sh
a@futronic:~$  5xx.py   2009-07-12-hong_kong_phooey.jpg 20.log
  ! problem opening 2009-07-12-hong_kong_phooey.jpg, skipping  
   20.log ... complete     

  Between time      1.000 and time 1554663560.412:
	player.vimeo.com returned 40.00%  5xx errors 
	vimeo.com returned 0.00%  5xx errors 
```

