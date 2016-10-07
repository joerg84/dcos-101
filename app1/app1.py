import redis
import time

r = redis.StrictRedis(host='redis.marathon.l4lb.thisdcos.directory', port=6379, db=0)
if r.ping():
       	print("Redis Connected. Total number of keys:", len(r.keys()))
else:
       	print("Could not connect to redis")
time.sleep(5)

