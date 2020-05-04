# k-means clustering, equal population clusters

Based on: https://elki-project.github.io/tutorial/same-size_k_means

## k-means++ initial centroid selection

Based on: https://en.wikipedia.org/wiki/K-means%2B%2B

1. Choose one center uniformly at random among the data points.
2. For each data point x, compute D(x), the distance between x and the nearest
center that has already been chosen.
3. Choose one new data point at random as a new center, using a weighted
probability distribution where a point x is chosen with probability proportional to D(x)^2.
4. Repeat Steps 2 and 3 until k centers have been chosen.
