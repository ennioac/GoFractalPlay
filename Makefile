default:
	go run main-lorenz.go 

run:
	go run main-lorenz.go > out.csv

rungnuplot:
	go run main-lorenz-gnuplotOnly.go | gnuplot --persist

installdependancies:
	# Uncomment following ling if you receive an error message refering to GO111MODULE environment variable.
	#go env -w GO111MODULE=auto
	go mod download

build:
	go build -o a.out

clean:
	rm -f a.out  my_plot.png