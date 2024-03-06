package belajar_golang_standard_library

import (
	"fmt"
	"math/rand"
	"testing"
)

// Point represents a data point with x and y coordinates
type Point struct {
	X float64
	Y float64
}

// KMeans performs k-means clustering on a set of data points
func KMeans(points []Point, k int) [][]Point {
	centroids := initializeCentroids(points, k)
	clusters := make([][]Point, k)

	for i := 0; i < 100; i++ { // Maximum iterations
		clusters = make([][]Point, k) // Clear clusters

		// Assign each point to the nearest centroid
		for _, point := range points {
			nearest := findNearestCentroid(point, centroids)
			clusters[nearest] = append(clusters[nearest], point)
		}

		// Update centroids
		for j := 0; j < k; j++ {
			if len(clusters[j]) > 0 {
				centroids[j] = calculateCentroid(clusters[j])
			}
		}
	}

	return clusters
}

// initializeCentroids randomly selects initial centroids from data points
func initializeCentroids(points []Point, k int) []Point {
	centroids := make([]Point, k)
	for i := 0; i < k; i++ {
		centroids[i] = points[rand.Intn(len(points))]
	}
	return centroids
}

// findNearestCentroid finds the index of the nearest centroid for a given point
func findNearestCentroid(point Point, centroids []Point) int {
	nearest := 0
	minDistance := distance(point, centroids[0])

	for i, centroid := range centroids[1:] {
		dist := distance(point, centroid)
		if dist < minDistance {
			nearest = i + 1
			minDistance = dist
		}
	}

	return nearest
}

// calculateCentroid calculates the centroid of a cluster
func calculateCentroid(cluster []Point) Point {
	var sumX, sumY float64
	for _, point := range cluster {
		sumX += point.X
		sumY += point.Y
	}
	return Point{X: sumX / float64(len(cluster)), Y: sumY / float64(len(cluster))}
}

// distance calculates the Euclidean distance between two points
func distance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return dx*dx + dy*dy
}

func TestMining(t *testing.T) {
	// Generate some random data points
	var points []Point
	for i := 0; i < 50; i++ {
		points = append(points, Point{X: rand.Float64() * 100, Y: rand.Float64() * 100})
	}

	// Perform k-means clustering
	k := 3 // Number of clusters
	clusters := KMeans(points, k)

	// Print clusters
	for i, cluster := range clusters {
		fmt.Printf("Cluster %d:\n", i+1)
		for _, point := range cluster {
			fmt.Printf("(%.2f, %.2f) ", point.X, point.Y)
		}
		fmt.Println()
	}
}
