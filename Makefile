default:
	go run main-lorenz-gnuplotOnly.go | gnuplot --persist

rotate:
	go run main-lorenz-gnuplotOnly.go | gnuplot  -1

run:
	go run main-lorenz-gnuplotOnly.go 

installdependancies:
	# Uncomment following ling if you receive an error message refering to GO111MODULE environment variable.
	#go env -w GO111MODULE=auto
	go mod download

build:
	go build -o a.out

clean:
	rm -f a.out  my_plot.png