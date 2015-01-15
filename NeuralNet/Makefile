# Which compiler to use
CC = g++ -std=c++11
# Any necessary compilation flags
CFLAGS = -c
# Listing any dependencies (headers) for the project
DEPS = connected_feed_forward_neural_network.hpp neuron.hpp

# This is a target. Targets get executed if any of their
# dependencies (listed after the ':') change
all: neuralnet

neuralnet: main.o neural_network.o neuron.o
	$(CC) 0_main.o neural_network.o neuron.o -o neuralnet

main.o: 0_main.cpp
	$(CC) $(CFLAGS) 0_main.cpp

neural_network.o: connected_feed_forward_neural_network.cpp
	$(CC) $(CFLAGS) connected_feed_forward_neural_network.cpp -o neural_network.o

neuron.o: neuron.cpp
	$(CC) $(CFLAGS) neuron.cpp

clean:
	rm -rf *o neuralnet