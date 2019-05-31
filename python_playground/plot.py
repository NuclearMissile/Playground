import numpy as np
import matplotlib.pyplot as plt

x = np.mgrid[-2 * np.pi:2 * np.pi:0.02]


def show_plot(fn, order, name):
    y1 = 0
    for n in range(1, order, 1):
        b = fn(n, x)
        y1 = b + y1
    plt.plot(x, y1, linewidth=0.6)
    plt.title(f'n={order}')
    plt.xlabel('Time')
    plt.ylabel('Amplitude')
    plt.savefig(f'{name}_n={order}.png')
    plt.show()


def f1(n, xx):
    return 2 * (-1) ** (n + 1) / n * np.sin(n * xx)


def f2(n, xx):
    return ((2 * (-1) ** (n + 1) + 2) / (n * np.pi)) * np.sin(2 * n * xx)


for i in range(3, 15, 2):
    show_plot(f1, i, 'f1')
    show_plot(f2, i, 'f2')
