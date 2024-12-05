package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	before int
	after  int
}

func readInput(fileName string) ([]Rule, [][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var rule []Rule
	var sequences [][]int
	parseRules := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			parseRules = false
			continue
		}
		if parseRules {
			data := strings.Split(line, "|")
			before, _ := strconv.Atoi(data[0])
			after, _ := strconv.Atoi(data[1])
			rule = append(rule, Rule{before, after})
		} else {
			var sequence []int
			data := strings.Split(line, ",")
			for _, item := range data {
				num, _ := strconv.Atoi(item)
				sequence = append(sequence, num)
			}
			sequences = append(sequences, sequence)
		}
	}
	return rule, sequences, nil
}

func checkValidSequence(rules []Rule, sequence []int) bool {
	m := make(map[int]int)
	for index, number := range sequence {
		m[number] = index
	}
	for _, rule := range rules {
		beforeValue, beforeExist := m[rule.before]
		afterValue, afterExist := m[rule.after]
		if beforeExist && afterExist && beforeValue > afterValue {
			return false
		}
	}
	return true
}

func getMidNumber(sequence []int) int {
	return sequence[len(sequence)/2]
}

func sortFollowRule(a, b int, rules []Rule) bool {
	for _, rule := range rules {
		if rule.before == a && rule.after == b {
			return true
		}
	}
	return false
}

func main() {
	rules, sequences, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var part1 int
	var part2 int
	for _, sequence := range sequences {
		if !checkValidSequence(rules, sequence) {
			slices.SortFunc(sequence, func(a, b int) int {
				if sortFollowRule(a, b, rules) {
					return -1
				}
				if sortFollowRule(b, a, rules) {
					return 1
				}
				return 0
			})
			part2 += getMidNumber(sequence)
		} else {
			part1 += getMidNumber(sequence)
		}
	}
	fmt.Printf("part 1: %d\n", part1)
	fmt.Printf("part 2: %d\n", part2)

}
