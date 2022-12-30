import time
import os

print(os.getpid())


path = "./test"
with open(path, mode='a+') as f:
    time.sleep(60)
    pass
