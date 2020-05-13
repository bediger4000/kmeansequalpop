package data

import (
	"fmt"
	"io"
	"log"
	"os"
)

// ReadPoints reads textual 3-column data (population, x, y coords)
// from a file named as the formal argument.
func ReadPoints(filename string) ([]*Point, int, error) {
	fin, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer fin.Close()

	var points []*Point
	var population float64

	for lineNo := 1; true; lineNo++ {
		var p Point
		n, err := fmt.Fscanf(fin, "%f %f %f\n", &p.Pop, &p.X, &p.Y)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, 0, err
		}
		if n != 3 {
			log.Printf("Input line %d, parsed %d items, wanted 3\n", lineNo, n)
			continue
		}

		population += p.Pop

		// precalculate moments around axes.
		p.Xmoment = p.X * p.Pop
		p.Ymoment = p.Y * p.Pop

		points = append(points, &p)
	}

	return points, int(population), nil
}
