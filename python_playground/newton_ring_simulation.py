import matplotlib
import matplotlib.pyplot as plt
import numpy as np
import math

PI = np.pi


def lambda2RGB(wavelength, gamma=0.8) -> (float, float, float):
    if 380 <= wavelength <= 440:
        attenuation = 0.3 + 0.7 * (wavelength - 380) / (440 - 380)
        R = ((440 - wavelength) / (440 - 380)) * attenuation
        G = 0.0
        B = 1.0 * attenuation
    elif 440 <= wavelength <= 490:
        R = 0.0
        G = (wavelength - 440) / (490 - 440)
        B = 1.0
    elif 490 <= wavelength <= 510:
        R = 0.0
        G = 1.0
        B = (510 - wavelength) / (510 - 490)
    elif 510 <= wavelength <= 580:
        R = (wavelength - 510) / (580 - 510)
        G = 1.0
        B = 0.0
    elif 580 <= wavelength <= 645:
        R = 1.0
        G = (645 - wavelength) / (645 - 580)
        B = 0.0
    elif 645 <= wavelength <= 750:
        attenuation = 0.3 + 0.7 * (750 - wavelength) / (750 - 645)
        R = 1.0 * attenuation
        G = 0.0
        B = 0.0
    else:
        R = 0.0
        G = 0.0
        B = 0.0
    return (int(_ ** gamma * 255) for _ in [R, G, B])


def rgb2str(R, G, B):
    return '#%02x%02x%02x' % (R, G, B)


def render(R=1.5, lam=380e-9, size=5e-3, N=801):
    if N % 2 == 0:
        N += 1
    size_per_pixel = size / N
    matrix = np.ndarray((N, N))

    def r(i, j):
        return math.sqrt((i - N // 2) ** 2 + (j - N // 2) ** 2) * size_per_pixel

    def genCmap():
        r, g, b = lambda2RGB(lam * 1e9)
        start, mid, end = '#000000', rgb2str(r // 2, g // 2, b // 2), rgb2str(r, g, b)
        return matplotlib.colors.LinearSegmentedColormap.from_list('', [start, mid, end])

    for i in range(N):
        for j in range(N):
            matrix[i][j] = math.sin(PI * r(i, j) ** 2 / R / lam) ** 2

    plt.figure(figsize=(8, 8))
    plt.imshow(matrix, cmap=genCmap())
    plt.title(f'lambda={int(lam * 1e9)}nm, size={size * 1e3}mm, R={R}m')
    plt.savefig(f'{int(lam * 1e9)}_{int(size * 1e3)}_{R}.png')


if __name__ == '__main__':
    for l in np.linspace(380, 750, 5):
        render(lam=l * 1e-9)

    for r in np.linspace(0.5, 2.5, 5):
        render(lam=555 * 1e-9, R=r)
