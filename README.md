# cVRP Metaheuristics

## Project layout

This repository contains a set of metaheuristic solvers for the Capacitated Vehicle Routing Problem (cVRP).

- `golang/` — main Go implementation and packages.
  - `golang/algorithms/` — GA, greedy, random and simulated annealing implementations.
  - `golang/cmd/solver/` — CLI solver entrypoint used to run a single instance.
  - `golang/cmd/benchmark/` — benchmark runner used to compare algorithms and collect CSV results.
- `data/` — sample VRP instance files (.vrp) used by the solver and benchmarks.
- `results/` — CSV outputs produced by the benchmark and solver runs.
- `python/` — small utilities for generating charts from CSV results (e.g. `chart_generator.py`).

> Note: repository code is written in Go. The examples below assume you run commands from the repository root.

## How to start

1. Clone repository

```bash
git clone https://github.com/skurzyp/cvrp-metaheuristics.git
```

2. Navigate to the project directory

```bash
cd cvrp-metaheuristics
```

3. Install dependencies

```bash
cd golang
go mod tidy
cd -
```

4. Running the Solver

You can run the solver with different algorithms and parameters. Here are some examples that use the `golang/cmd` entrypoints.

```bash
# Run with Genetic Algorithm (default)
go run ./golang/cmd/solver \
  -file ./data/toy.vrp \
  -pop 100 \
  -gen 100 \
  -elitism 5 \
  -mutation 0.1 \
  -crossover 0.7 \
  -tournament 5 \
  -algorithm ga

# Run with Simulated Annealing
go run ./golang/cmd/solver \
  -file ./data/toy.vrp \
  -algorithm sa \
  -temp 1000.0 \
  -minTemp 0.001 \
  -cooling 0.995 \
  -innerLoop 100

# Run with Greedy Algorithm
go run ./golang/cmd/solver \
  -file ./data/toy.vrp \
  -algorithm greedy \
  -start 0
```

5. Running Benchmarks

The benchmark tool allows you to compare different algorithms and measure their performance:

```bash
# Run with default parameters
go run ./golang/cmd/benchmark \
  -file ./data/a-n32-k5.vrp \
  -out ./results/algorithm/summary.csv \
  -runs 10

# Run with custom GA and SA parameters
go run ./golang/cmd/benchmark \
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

### Solver CLI Parameters (high level)

Most flags are defined in the CLI entrypoints under `golang/cmd/*`. High-level flags used by the solver and benchmark are listed below. See the Go sources for exact defaults and validation.

| Flag          | Type    | Description                                        |
| ------------- | ------- | -------------------------------------------------- |
| `-file`       | string  | Path to the VRP instance file (e.g. `./data/toy.vrp`) |
| `-algorithm`  | string  | Algorithm to use (`ga`, `sa`, or `greedy`)         |
| `-pop`        | int     | Population size (GA only)                          |
| `-gen`        | int     | Maximum number of generations (GA only)            |
| `-elitism`    | int     | Number of elite solutions retained (GA only)       |
| `-mutation`   | float64 | Mutation rate (0-1, GA only)                      |
| `-crossover`  | float64 | Crossover rate (0-1, GA only)                     |
| `-tournament` | int     | Tournament size for parent selection (GA only)     |
| `-start`      | int     | Starting node ID (Greedy only)                     |
| `-temp`       | float64 | Initial temperature (SA only)                      |
| `-minTemp`    | float64 | Minimum temperature (SA only)                      |
| `-cooling`    | float64 | Cooling rate (0-1, SA only)                       |
| `-innerLoop`  | int     | Iterations per temperature step (SA only)          |

## Python utilities

The `python/` folder contains a small `chart_generator.py` script used to convert CSV results (from `results/`) into charts. It depends on standard Python plotting libraries (matplotlib, pandas). Example usage:

```bash
# install python deps (recommended in a venv)
pip install pandas matplotlib

# generate charts from CSV
python3 python/chart_generator.py --input results/algorithm/summary.csv --out charts/
```

## TODO
- conduct experiments and create charts (using `python/chart_generator.py`)
- consider renaming GA flags: `-mutation` -> `-Pm` and adding `-Px` for crossover probability
