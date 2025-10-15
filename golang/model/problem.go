package model

type Problem struct {
	Instance Instance
	// GA parameters
	ElitismCount   int
	MaxGenerations int
	MutationRate   float32
	CrossoverRate  float32
	PopulationSize int
	TournamentSize int

	// SA parameters
	InitialTemperature float64
	MinimumTemperature float64
	CoolingRate        float64
	InnerLoop          int
}
