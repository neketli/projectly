package usecase

import (
	"context"
	"math"
	"projectly-server/internal/domain/team/entity"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

func (u *teamUseCase) GetStatisticData(ctx context.Context, teamID int) ([]entity.StatisticData, error) {
	projects, err := u.repo.GetStatisticData(ctx, teamID)
	if err != nil {
		u.logger.Error("project - usecase - GetStatisticData - u.repo.GetStatisticData: %s", err.Error())
		return nil, err
	}

	if len(projects) <= 3 {
		for i := range projects {
			projects[i].Cluster = i
		}
		return projects, nil
	}

	matrix := mat.NewDense(len(projects), 2, nil)
	for i, d := range projects {
		matrix.Set(i, 0, float64(d.CompletedTasksCount))
		matrix.Set(i, 1, d.AvgTaskDuration)
	}

	normalizedMatrix := normalize(matrix)
	clusters := kMeans(normalizedMatrix, 3, 100)

	for i, cluster := range clusters {
		projects[i].Cluster = cluster
	}

	return projects, nil
}

// kMeans performs k-means clustering on the given data.
func kMeans(data *mat.Dense, k, maxIter int) []int {
	// Get the dimensions of the data.
	rows, cols := data.Dims()
	if rows == 0 || cols == 0 {
		return nil
	}

	// Initialize clusters and centroids.
	// clusters is a slice of length rows, where each element is the index of the cluster
	// that the corresponding row of data belongs to.
	// centroids is a matrix with k rows and cols columns, where each row is the centroid
	// of a cluster.
	clusters := make([]int, rows)
	centroids := mat.NewDense(k, cols, nil)

	// Initialize centroids by assigning each centroid to a random data point.
	for i := 0; i < k; i++ {
		centroids.SetRow(i, mat.Row(nil, i, data))
	}

	// Run k-means algorithm.
	for iter := 0; iter < maxIter; iter++ {
		// Assign each data point to a cluster.
		for i := 0; i < rows; i++ {
			minDist := math.MaxFloat64
			for j := 0; j < k; j++ {
				// Calculate the distance between the current data point and the jth centroid.
				dist := floats.Distance(mat.Row(nil, i, data), mat.Row(nil, j, centroids), 2)
				if dist < minDist {
					minDist = dist
					// Assign the data point to the cluster with the minimum distance.
					clusters[i] = j
				}
			}
		}

		// Update centroids.
		for j := 0; j < k; j++ {
			clusterPoints := make([][]float64, 0)
			for i := 0; i < rows; i++ {
				// Get all the data points that belong to the jth cluster.
				if clusters[i] == j {
					clusterPoints = append(clusterPoints, mat.Row(nil, i, data))
				}
			}
			if len(clusterPoints) > 0 {
				// Calculate the new centroid of the jth cluster.
				newCentroid := make([]float64, cols)
				for _, point := range clusterPoints {
					// Add the coordinates of each point to the new centroid.
					floats.Add(newCentroid, point)
				}
				// Calculate the mean of the coordinates of all the points in the cluster.
				floats.Scale(1/float64(len(clusterPoints)), newCentroid)
				// Set the new centroid of the jth cluster.
				centroids.SetRow(j, newCentroid)
			}
		}
	}
	return clusters
}

// normalize data in the given matrix.
func normalize(data *mat.Dense) *mat.Dense {
	rows, cols := data.Dims()
	if rows == 0 || cols == 0 {
		return mat.NewDense(rows, cols, nil)
	}

	normalized := mat.NewDense(rows, cols, nil)
	for i := 0; i < cols; i++ {
		col := mat.Col(nil, i, data)
		mean, std := stat.MeanStdDev(col, nil)

		// If the standard deviation is 0, the column is constant and we
		// don't need to normalize it.
		if std == 0 || math.IsNaN(std) {
			normalized.SetCol(i, col)
			continue
		}

		// Normalize the column.
		for j := range col {
			col[j] = (col[j] - mean) / std
		}
		normalized.SetCol(i, col)
	}

	return normalized
}
