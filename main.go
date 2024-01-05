package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type measurements struct {
	min       float64
	max       float64
	mean      float64
	count     float64
	aggregate float64
}

func main() {

	start := time.Now()
	file, err := os.Open("measurements.txt")
	if err != nil {
		log.Fatal(err)
	}

	weatherStations := make(map[string]*measurements)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		i := strings.Index(line, ";")

		name := line[0:i]

		temperature, err := strconv.ParseFloat(line[i+1:], 64)
		if err != nil {
			panic(err)
		}

		if ws, ok := weatherStations[name]; ok {
			ws.aggregate += temperature
			ws.count++
			ws.mean = ws.aggregate / ws.count

			if ws.min < temperature {
				ws.min = temperature
			}

			if ws.max > temperature {
				ws.max = temperature
			}
			continue
		}

		weatherStations[name] = &measurements{
			aggregate: temperature,
			min:       temperature,
			max:       temperature,
			count:     1,
		}
	}

	stationNames := []string{}

	for k := range weatherStations {
		stationNames = append(stationNames, k)
	}

	slices.Sort(stationNames)

	for _, name := range stationNames {

		station := weatherStations[name]

		fmt.Printf("%s:%.1f/%.1f/%.1f\n", name, station.min, station.mean, station.max)

	}

	fmt.Println(time.Since(start))
}
