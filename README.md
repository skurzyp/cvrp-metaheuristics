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

4. Run the solver with default parameters
```
go run ./cmd/cvrp/main.go
```

5. Run the solver with custom CLI parameters
```
go run ./cmd/cvrp/main.go \
  -file ./data/toy.vrp \
  -pop 10 \
  -gen 50 \
  -elitism 3 \
  -mutation 0.1 \
  -tournament 4
```

### CLI Params
| Flag          | Type    | Default                | Description                                        |
| ------------- | ------- | ---------------------- | -------------------------------------------------- |
| `-file`       | string  | `./data/mock_data.txt` | Path to the VRP instance file                      |
| `-pop`        | int     | 100                      | Population size for the Genetic Algorithm          |
| `-gen`        | int     | 100                     | Maximum number of generations                      |
| `-elitism`    | int     | 5                      | Number of elite solutions retained each generation |
| `-mutation`   | float64 | 0.1                   | Mutation rate (between 0 and 1)      
| `-crossover`   | float64 | 0.7                   | Crossover rate (between 0 and 1)                  |
| `-tournament` | int     | 5                      | Tournament size used for parent selection          |


## TODO
#### Major
- [ ] implement genetics operators for GA
- [ ] implement ant colony
- [ ] implement benchmarks
- [ ] conduct experiments and create charts

#### Minor
- [ ] rename mutation to Pm and add Px as `Probability of Crossing`