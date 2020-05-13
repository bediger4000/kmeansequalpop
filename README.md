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

## Initialization

1. Compute the desired cluster size, n/k.
2. Initialize means, preferably with k-means++
3. Order points by the distance to their nearest cluster minus distance to the
farthest cluster (= biggest benefit of best over worst assignment)
4. Assign points to their preferred cluster until this cluster is full, then
resort remaining objects, without taking the full cluster into account anymore.

## Iteration

1. Compute current cluster means
2. For each object, compute the distances to the cluster means
3. Sort elements based on the delta of the current assignment and the best possible alternate assignment.
4. For each element by priority:
   1. For each other cluster, by element gain, unless already moved:
      1. If there is an element wanting to leave the other cluster and this swap yields and improvement, swap the two elements
      2. If the element can be moved without violating size constraints, move it
   2. If the element was not changed, add to outgoing transfer list.
5. If no more transfers were done (or max iteration threshold was reached), terminate

This isn't terribly well-specified.
I'll have to make decisions about various alternatives.
Also, my points have a population, or a weight associated with them.
I'll have to account for that.

### Iteration Notes

1. Compute current cluster means: have to include population in
calculating the moment of each point around the X or Y axis.
