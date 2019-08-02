import hashlib
from matplotlib import pyplot as plt
import numpy as np
import time

np.random.seed(42)
LOOP_TIMES = 10
sizes = [0.5, 1, 2, 4, 8, 16, 32]
blocks = [np.random.bytes(size * 1024 * 1024) for size in sizes]
blocks = zip([str(size) for size in sizes], blocks)
time_log = {}
sha256 = hashlib.sha256()

if __name__ == '__main__':
    for name, block in blocks:
        start = time.time()
        for _ in range(LOOP_TIMES):
            sha256.update(block)
            sha256.update(sha256.digest())
        end = time.time()
        time_log[name] = (end - start) / LOOP_TIMES * 1000
    for k, v in time_log.items():
        print(f'{k}MB block: {v:.2f}ms')
    plt.title('SHA256 calculation time with different block sizes')
    plt.xlabel('Block size(MB)')
    plt.ylabel('Time(ms)')
    plt.plot([float(size) for size in time_log.keys()], list(time_log.values()))
    plt.show()
