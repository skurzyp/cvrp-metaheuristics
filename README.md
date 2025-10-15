# cVRP Metaheuristics
## How to start?
1. Clone repository
```
git clone https://github.com/skurzyp/cvrp-metaheuristics.git
```

2. Navigate to directory
```
cd cvrp-metaheuristics
```

3. Instal dependencies
```
go mod tidy
```

4. Running the Solver

You can run the solver with different algorithms and parameters. Here are some examples:

```bash
# Run with Genetic Algorithm (default)
go run ./cmd/solver/solver.go \
  -file ./data/toy.vrp \
  -pop 100 \
  -gen 100 \
  -elitism 5 \
  -mutation 0.1 \
  -crossover 0.7 \
  -tournament 5 \
  -algorithm ga

# Run with Simulated Annealing
go run ./cmd/solver/solver.go \
  -file ./data/toy.vrp \
  -algorithm sa \
  -temp 1000.0 \
  -minTemp 0.001 \
  -cooling 0.995 \
  -innerLoop 100

# Run with Greedy Algorithm
go run ./cmd/solver/solver.go \
  -file ./data/toy.vrp \
  -algorithm greedy \
  -start 0
```

5. Running Benchmarks

The benchmark tool allows you to compare different algorithms and measure their performance:

```bash
# Run with default parameters
go run ./cmd/benchmark/benchmark.go \
  -file ./data/a-n32-k5.vrp \
  -out ./results/algorithm/summary.csv \
  -runs 10

# Run with custom GA and SA parameters
go run ./cmd/benchmark/benchmark.go \
  -file ./data/a-n32-k5.vrp \
  -out ./results/benchmark/custom-params.csv \
  -runs 20 \
  -randomRuns 5000 \
  -pop 300 \
  -gen 10000 \
  -mutation 0.1 \
  -crossover 0.95 \
  -tournament 15 \
  -temp 2000.0 \
  -cooling 0.9999
```

### Solver CLI Parameters
| Flag          | Type    | Default                | Description                                        |
| ------------- | ------- | ---------------------- | -------------------------------------------------- |
| `-file`       | string  | `./data/toy.vrp` | Path to the VRP instance file                      |
| `-algorithm`  | string  | `ga`                  | Algorithm to use (`ga`, `sa`, or `greedy`)         |
| `-pop`        | int     | 200                   | Population size (GA only)                          |
| `-gen`        | int     | 10000                 | Maximum number of generations (GA only)            |
| `-elitism`    | int     | 10                    | Number of elite solutions retained (GA only)       |
| `-mutation`   | float64 | 0.05                  | Mutation rate (0-1, GA only)                      |
| `-crossover`  | float64 | 0.9                   | Crossover rate (0-1, GA only)                     |
| `-tournament` | int     | 10                    | Tournament size for parent selection (GA only)     |
| `-start`      | int     | 0                     | Starting node ID (Greedy only)                     |
| `-temp`       | float64 | 2000.0                | Initial temperature (SA only)                      |
| `-minTemp`    | float64 | 0.001                 | Minimum temperature (SA only)                      |
| `-cooling`    | float64 | 0.999               | Cooling rate (0-1, SA only)                       |
| `-innerLoop`  | int     | 2000                  | Iterations per temperature step (SA only)          |

### Benchmark CLI Parameters
| Flag          | Type    | Default                | Description                                        |
| ------------- | ------- | ---------------------- | -------------------------------------------------- |
| `-file`       | string  | `./data/toy.vrp` | Path to the VRP instance file                      |
| `-out`        | string  | `./results/summary.csv`| Path to CSV output file                           |
| `-runs`       | int     | 10                     | Number of repetitions for GA and SA                |
| `-randomRuns` | int     | 10000                 | Number of random solution generations              |
| `-pop`        | int     | 200                   | Population size (GA only)                          |
| `-gen`        | int     | 10000                 | Maximum generations (GA only)                      |
| `-elitism`    | int     | 10                    | Number of elite solutions retained (GA only)       |
| `-mutation`   | float64 | 0.05                  | Mutation rate (0-1, GA only)                      |
| `-crossover`  | float64 | 0.90                  | Crossover rate (0-1, GA only)                     |
| `-tournament` | int     | 10                    | Tournament size for parent selection (GA only)     |
| `-temp`       | float64 | 2000.0                | Initial temperature (SA only)                      |
| `-minTemp`    | float64 | 0.001                 | Minimum temperature (SA only)                      |
| `-cooling`    | float64 | 0.999                | Cooling rate (0-1, SA only)                       |
| `-innerLoop`  | int     | 2000                  | Iterations per temperature step (SA only)          |

## TODO
#### Major
- [x] implement genetics operators for GA
- [X] implement simulated anealing
- [X] implement benchmarks
- [ ] conduct experiments and create charts

#### Minor
- [ ] rename mutation to Pm and add Px as `Probability of Crossing`