package data

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReadPoints(filename string) ([]*Point, error) {
	fin, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fin.Close()

	var points []*Point

	for lineNo := 1; true; lineNo++ {
		var p Point
		n, err := fmt.Fscanf(fin, "%f %f %f\n", &p.Pop, &p.X, &p.Y)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if n != 3 {
			log.Printf("Input line %d, parsed %d items, wanted 3\n", lineNo, n)
			continue
		}

		points = append(points, &p)
	}

	return points, nil
}
