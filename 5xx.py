#!/usr/bin/env python3

'''
Parse a list of log files for 5xx http return codes
'''

import argparse
from multiprocessing import Pool
import time

# Globals

ERROR_PREFIX='5' # http return code prefix to match
CHUNK_SIZE=4700000 # read a logfile in CHUNK_SIZE chunks
POOL_SIZE=2 # number of concurrent processes

start_time=0.0
end_time=0.0


def parse_args():
    '''
    parse command line args
    '''
    parser = argparse.ArgumentParser(description='Parse log files for 5xx http return codes.')
    parser.add_argument('logs',  nargs='+',
                default=[] ,help='List of log files to process')
    parser.add_argument('-s','--start', dest='start',default=1.0,
                    type=float, help='Start time in UNIX time.')
    parser.add_argument('-e','--end', dest='end',default=time.time(),
                    type=float, help='End time in UNIX time.')
    args = parser.parse_args()
    return (args.start,args.end,args.logs)


def logfile_to_chunks(fn):
    '''
    iterate chunks of lines in a logfile
    '''
    data={'sites':{},'logfile':fn}
    print('\r'+'> starting %s '%fn,end='')
    try:
        with open(fn) as log:
            while 1:
                chunk=log.readlines(CHUNK_SIZE)
                if not chunk:
                    break
                data=chunk_to_lines(chunk,data)
                print('\r'+'> parsing %s     '%fn,end='')
        print('\r'+'   %s ... complete     '%fn,end='\n')
    except:
        print('\r'+'  ! problem opening %s, skipping  '%fn,end='\n')
    return data


def chunk_to_lines(chunk,data):
    '''
    iterate lines in a logfile chunk
    '''
    for line in chunk:
        data=line_to_values(line,data)
    return data


def line_to_values(line,data):
    '''
    split line in to values for
    timestamp, host_name, and http_code
    '''
    try:
        timestamp,_,host_name,_,http_code,_=line.split('|',5)
    except:
        print('\r'+'failed to parse values from line in %s'%data['logfile'])
        return data
    host_name=host_name.strip()
    if chk_timestamp(timestamp):
        data=chk_host_name(host_name,data)
        data['sites'][host_name]['errors'] +=chk_http_code(http_code)
        data['sites'][host_name]['total'] +=1
        return data


def chk_timestamp(timestamp):
    '''
    check if timestamp is between start_time and end_time.
    '''
    log_time=float(timestamp)
    if end_time > log_time:
        if log_time >= start_time:
            return True
    return False


def chk_host_name(host_name,data):
    '''
    check if host_name is in data, add it if not
    '''
    if host_name not in data['sites'].keys():
        data['sites'][host_name]={'total':0,'errors':0}
    return data


def chk_http_code(http_code):
    '''
    check if http_code is 5xx
    '''
    if http_code.strip()[0]==ERROR_PREFIX:
        return 1
    return 0


def mk_report_data(data_list):
    '''
    consolidate data by host_name for report
    '''
    report_data={}
    for data in data_list:
        for k,v in data['sites'].items():
            if k not in report_data.keys():
                report_data[k]=v
            else:
                report_data[k]['errors'] +=v['errors']
                report_data[k]['total'] +=v['total']
    return report_data


def shw_report(data_list):
    '''
    display aggregated data by host_name
    '''
    report_data=mk_report_data(data_list)
    print('\n  Between time %10.3f and time %10.3f:'%(start_time,end_time))
    for k,v in report_data.items():
        p= (v['errors']/v['total'])*100
        print('\t%s returned %2.2f%% 5xx errors '%(k,p))



if __name__ == '__main__':
    start_time,end_time,logs=parse_args()
    pool=Pool(POOL_SIZE)
    data_list =pool.map(logfile_to_chunks,logs)
    shw_report(data_list)
