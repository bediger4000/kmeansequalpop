CENTROIDS =  centroids/types.go centroids/initial.go centroids/assign.go \
			 centroids/center.go centroids/count.go centroids/leavers.go
DATA = data/types.go data/read.go

kmeans: cmd/kmeans/kmeans.go $(CENTROIDS) $(DATA)
	go build -o kmeans cmd/kmeans/kmeans.go

genrand: cmd/genrand/genrand.go
	go build -o genrand cmd/genrand/genrand.go

rtest: cmd/rtest/rtest.go
	go build -o rtest cmd/rtest/rtest.go

ctest: cmd/centroids/ctest.go centroids/initial.go centroids/assign.go
	go build -o ctest cmd/centroids/ctest.go

10.dat: genrand
	./genrand -p 1000 10 > 10.dat

10000.dat: genrand
	./genrand -p 1000 10000 > 10000.dat

tests: 10.dat rtest
	./rtest 10.dat > x.dat
	-diff 10.dat x.dat

clean:
	-rm -rf genrand 10.dat x.dat
	-rm -rf rtest ctest
