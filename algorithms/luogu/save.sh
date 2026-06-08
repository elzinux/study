#!/bin/bash

difficulty="$1"
problem="$2"

if [ -z "$difficulty" ] || [ -z "$problem" ]; then
    echo "Usage: ./save.sh <difficulty> <problem>"
    echo "Example: ./save.sh 01-red P1001"
    exit 1
fi

if [ ! -f main.cpp ]; then
    echo "Error: main.cpp not found"
    exit 1
fi

mkdir -p "$difficulty/$problem"

cp main.cpp "$difficulty/$problem/$problem.cpp"

if [ -f input.txt ]; then
    cp input.txt "$difficulty/$problem/input.txt"
fi

echo "Saved: $difficulty/$problem/$problem.cpp"