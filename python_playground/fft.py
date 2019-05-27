import numpy as np

j = complex(0, 1)


def fft(xt: np.ndarray) -> np.ndarray:
    N = len(xt)
    if N < 2:
        return xt
    even, odd = fft(xt[0::2]), fft(xt[1::2])
    T = np.exp(-2j * np.pi * np.array(range(N // 2)) / N) * odd
    return np.concatenate([even + T, even - T])


print(' '.join(str(abs(f)) for f in fft(np.array([1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0]))))
