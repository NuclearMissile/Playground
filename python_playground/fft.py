import numpy as np
import matplotlib.pyplot as plt

j = complex(0, 1)
PI = np.pi

ranges = [np.arange(r) / r for r in [8, 128, 1024]]

signal1 = [2 + 3 * np.sin(2 * PI * 50 * r) + 4 * np.cos(2 * PI * 30 * r) for r in ranges]


def fft(xt: np.ndarray) -> np.ndarray:
    N = len(xt)
    if N < 2:
        return xt
    even, odd = fft(xt[0::2]), fft(xt[1::2])
    T = np.exp(-2j * PI * np.arange(N // 2) / N) * odd
    return np.concatenate([even + T, even - T])


def fft_abs(xt):
    return np.abs(fft(xt))


if __name__ == '__main__':
    fig = plt.figure(figsize=(15, 15))
    f = [fft_abs(x) for x in signal1]
    f_np = [np.abs(np.fft.fft(x)) for x in signal1]
    plt.suptitle('x(t) = 2+3sin(2PI*50t)+4cos(2PI*30t)')

    for i in range(3):
        plt.subplot(3, 3, i * 3 + 1)
        plt.title(f'N={len(ranges[i])}, original signal')
        plt.plot(ranges[i], signal1[i])
        plt.subplot(3, 3, i * 3 + 2)
        plt.title(f'N={len(ranges[i])}, by fft_abs')
        plt.plot(ranges[i], f[i])
        plt.subplot(3, 3, i * 3 + 3)
        plt.title(f'N={len(ranges[i])}, by np.fft.fft')
        plt.plot(ranges[i], f_np[i])
    plt.show()
