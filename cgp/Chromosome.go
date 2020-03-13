package cgp

import (
	"math/rand"
	"strconv"
	"strings"
)

type Chromosome []int

func GenerateChromosome(maxValues *[]int) *Chromosome {
	size := len(*maxValues)
	chromosome := make([]int, size)
	for i := 0; i < size; i++ {
		chromosome[i] = rand.Intn((*maxValues)[i])
	}

	ch := Chromosome(chromosome)
	return &ch
}

func LoadChromosome(str string) *Chromosome {
	genes := strings.Split(str, "-")
	size := len(genes)
	chromosome := make([]int, size)
	for index, gene := range genes {
		chromosome[index], _ = strconv.Atoi(gene)
	}

	ch := Chromosome(chromosome)
	return &ch
}

func (chromosome Chromosome) Mutate(maxValues *[]int, mutationProbability float32) *Chromosome {
	size := len(*maxValues)
	newChromosome := make([]int, size)

	for i := 0; i < size; i++ {
		if rand.Float32() < mutationProbability {
			chromosome[i] = rand.Intn((*maxValues)[i])
		} else {
			newChromosome[i] = chromosome[i]
		}
	}

	ch := Chromosome(newChromosome)
	return &ch
}

func (chromosome Chromosome) ToString() string {
	size := len(chromosome)
	str := ""
	for i := 0; i < size; i++ {
		str += strconv.Itoa(chromosome[i]) + "-"
	}
	str = str[:len(str)-1]

	return str
}
