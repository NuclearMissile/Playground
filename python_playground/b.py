import numpy as np
import math

j = complex(0, 1)

N = 9

t = 1 / 4

p = j * 2 * math.pi

for k in range(N):
    print(f'k = {k}')
    a = complex(0, 0)
    for n in range(N):
        a += 0.5 * (np.exp(p * n * (t - k / N)) + np.exp(-p * n * (t + k / N)))
    print((a.real))
    print('=========')
