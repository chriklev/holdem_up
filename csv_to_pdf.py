import matplotlib.cm as cm
import matplotlib.pyplot as plt
import numpy as np
import sys

file_name = "starting_hand_win_rates.csv"
data = np.genfromtxt(file_name, delimiter=",")

min_data = np.min(data)
max_data = np.max(data)
norm_data = (data-min_data)/(max_data-min_data)


def toRGB(x):
    if x >= 0.5:
        x = (x-0.5)*2
        return 1-x, 1, 1-x
    else:
        x = abs(x-0.5)*2
        return 1, 1-x, 1-x


colors = np.empty((data.shape[0], data.shape[1], 3))
for i in range(data.shape[0]):
    for j in range(data.shape[1]):
        colors[i, j] = toRGB(norm_data[i, j])

cardnames = ["2", "3", "4", "5", "6", "7", "8",
             "9", "10", "Jack", "King", "Queen", "Ace"]
nameColors = [[0.7, 0.7, 0.9] for i in range(13)]

# convert to percentage
data = np.around(data*100, 2)

plt.table(
    cellText=data,
    cellColours=colors,
    rowLabels=cardnames,
    rowColours=nameColors,
    colLabels=cardnames,
    colColours=nameColors,
    loc="center right"
)
plt.axis("off")
plt.xlabel("Same suited")
plt.savefig(file_name.split(".")[0] + ".pdf")
