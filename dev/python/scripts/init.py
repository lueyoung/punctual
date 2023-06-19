#!/usr/bin/env python

import os
import sys
from optparse import OptionParser
import logging  
import logging.handlers  
import time
from cassandra.cluster import Cluster
 
def set_logger():
    name = sys.argv[0].strip().split(r"/")[-1].strip().split('.')[0]
    LOG_DIR=options.LD
    LOG_FILE = os.path.join(LOG_DIR,"%s.log" % name)
    logger = logging.getLogger("%s" % name)

    logger.setLevel(logging.INFO)
    if options.LL.upper() in ["DEBUG",]:
        logger.setLevel(logging.DEBUG)       
    elif options.LL.upper() in ["INFO",]:
        logger.setLevel(logging.INFO) 
    elif options.LL.upper() in ["WARN","WARNING"]:
        logger.setLevel(logging.WARN)  
    elif options.LL.upper() in ["ERROR",]:
        logger.setLevel(logging.ERROR)  
    elif options.LL.upper() in ["CRITICAL",]:
        logger.setLevel(logging.CRITICAL)  
    else:
        pass
    
    #fh = logging.FileHandler(LOG_FILE)
    fh = logging.handlers.RotatingFileHandler(LOG_FILE, maxBytes = 1024*1024, backupCount = 5) 
    ch = logging.StreamHandler()  
    
    #fmt = "%(asctime)s-%(name)s-%(levelname)s-%(message)s-[%(filename)s:%(lineno)d]"
    fmt = '%(asctime)s - %(name)s - %(filename)s:%(lineno)s - %(levelname)s - %(message)s'
    formatter = logging.Formatter(fmt)

    fh.setFormatter(formatter)   
    ch.setFormatter(formatter)

    logger.addHandler(fh)   
    logger.addHandler(ch)

    return logger

def parse_opts(parser):
    parser.add_option("-i","--ip",action="store",type="string",dest="ip",default="127.0.0.1",help="IP of Cassandra")
    parser.add_option("-p","--port",action="store",type="string",dest="port",default="9042",help="Port of Cassandra")
    parser.add_option("-k","--keyspace",action="store",type="string",dest="keyspace",default="",help="Keyspace of Cassandra")
    parser.add_option("-t","--table",action="store",type="string",dest="table",default="",help="Table of Cassandra")
    parser.add_option("--ll",action="store",type="string",dest="LL",default="INFO",help="the log level")
    parser.add_option("--log_dir",action="store",type="string",dest="LD",default=r"/tmp",help="the dir to store log")
    parser.add_option("--cll",action="store",type="string",dest="CLL",default="INFO",help="the console log level")
    parser.add_option("--fll",action="store",type="string",dest="FLL",default="DEBUG",help="the file log level")
    (options,args) = parser.parse_args()

    return options

# mk global var: options & logger
options = parse_opts(OptionParser(usage="%prog [options]"))
logger = set_logger()

def timer(func):
    def wrapper(*args,**kwargs):
        t1 = time.time()
        ret = func(*args,**kwargs)
        t2 = time.time()
        logger.debug("Elapsed: %s sec." % (str(t2-t1)))
        return ret
    return wrapper

@timer
def main():
    ips = options.ip.split(',')
    cluster = Cluster(ips)
    session = cluster.connect()
    cmd = "CREATE KEYSPACE %s IF NOT EXISTS WITH REPLICATION = { \'class\' : \'SimpleStrategy\', \'replication_factor\' : 3 }" % options.keyspace
    logger.info(cmd)
    #session.execute("CREATE KEYSPACE %s IF NOT EXISTS with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 }" % options.keyspace)
    session.execute(cmd)
    session.execute("CREATE TABLE %s.%s(key text, value text, PRIMARY KEY(key)) IF NOT EXISTS" % options.keyspace)
    cluster.shutdown()

if __name__ == "__main__":
    main()
