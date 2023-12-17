# Search Algorithms in Go

## Overview

This repository is dedicated to the implementation of various search algorithms in the Go programming language. The goal of this project is to provide a comprehensive collection of search algorithms, showcasing their implementation details and usage in Go. These implementations are intended to serve as a learning resource for those interested in understanding how search algorithms work and how they can be implemented in Go.

## Algorithms

Below is a list of the search algorithms currently implemented in this repository:

- **Linear Search**: A simple search algorithm that checks every element in a list until the desired element is found or the list ends.
- **Jump Search**: An optimized version of the linear search, which works on sorted lists and jumps ahead by fixed steps before performing a linear search within a smaller block.
- **Binary Search**: An efficient search algorithm for sorted lists, which repeatedly divides the list in half to locate the target element.
- **Exponential Search**: A search algorithm particularly useful for unbounded or infinite lists. It works by finding a range where the element might be (by doubling the index) and then performing a binary search within that range.
- **Ternary Search**: A divide-and-conquer search algorithm that splits the search interval into three parts. It is efficient for finding the maximum or minimum in a unimodal function or for searching a specific value in a sorted array.
- **Depth-First Search (DFS)**: An algorithm for traversing or searching tree or graph data structures. It starts at the root (or an arbitrary node) and explores as far as possible along each branch before backtracking.
- **Breadth-First Search (BFS)**: Another algorithm for traversing or searching tree or graph structures. It starts at the root (or an arbitrary node) and explores the neighbor nodes first, before moving to the next level of neighbors.

## Usage

To use any of the search algorithms in this repository, first clone the repository to your local machine:

```bash
git clone https://github.com/aaronsatko/go-search