import numpy as np
import math

import matplotlib.pyplot as plt


for N in range(8, 11):
    n = np.array(range(N))
    x = np.cos(0.5 * math.pi * n)
    d = np.fft.fft(x)
    fig = plt.figure(figsize=(10, 10))
    plt.scatter(n, np.real(d))
    plt.title(f'N = {N}')
    plt.xlabel('n')
    plt.ylabel('DFT')
    plt.savefig(f'DFT_N={N}.png')
    plt.show()
