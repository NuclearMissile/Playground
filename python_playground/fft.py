import numpy as np
import matplotlib.pyplot as plt

j = complex(0, 1)

ranges = [np.arange(r) for r in [8, 64, 1024]]

signal1 = [3 * np.sin(r) + 4.5 * np.cos(r) for r in ranges]


def fft(xt: np.ndarray) -> np.ndarray:
    N = len(xt)
    if N < 2:
        return xt
    even, odd = fft(xt[0::2]), fft(xt[1::2])
    T = np.exp(-2j * np.pi * np.arange(N // 2) / N) * odd
    return np.concatenate([even + T, even - T])


if __name__ == '__main__':
    fig = plt.figure(figsize=(10, 18))
    f = [fft(x) for x in signal1]
    for i in range(3):
        plt.subplot(3, 1, i + 1)
        plt.plot(ranges[i], f[i])
    plt.show()
