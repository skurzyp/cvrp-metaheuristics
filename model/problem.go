package model

type Problem struct {
	Instance       Instance
	ElitismCount   int
	MaxGenerations int
	MutationRate   float32
	CrossoverRate  float32
	PopulationSize int
	TournamentSize int
}
