import numpy as np
import matplotlib.pyplot as plt
import pandas as pd

df = pd.read_csv('points.csv', header=None)
xy = np.array(df)
xx, yy = xy[0], xy[1]

plt.plot(xx, yy, '.')
plt.show()


