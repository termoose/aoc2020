package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readLines(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var result []string
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		result = append(result, line)
	}

	return result
}

type fieldData struct {
	ids map[int]struct{}
}

type ticket struct {
	values []int
}

type notes struct {
	fields map[string]fieldData
	mine   ticket
	nearby []ticket
}

// row: 6-11 or 33-44
func parseField(line string) (string, fieldData) {
	result := fieldData{
		ids: make(map[int]struct{}),
	}

	var ranges []int
	var rgx = regexp.MustCompile(`^(.+): (.+)-(.+) or (.+)-(.+)`)
	rs := rgx.FindStringSubmatch(line)

	for i := 2; i < len(rs); i++ {
		val, _ := strconv.Atoi(rs[i])
		ranges = append(ranges, val)
	}

	for i := ranges[0]; i <= ranges[1]; i++ {
		result.ids[i] = struct{}{}
	}
	for i := ranges[2]; i <= ranges[3]; i++ {
		result.ids[i] = struct{}{}
	}

	return rs[1], result
}

func parse(data []string) notes {
	row := 0
	result := notes{
		fields: make(map[string]fieldData),
	}

	// parse fields
	for ; row < len(data); row++ {
		line := data[row]
		if line == "" {
			break
		}

		name, field := parseField(data[row])
		result.fields[name] = field
	}

	row += 2 // skip header
	for ; row < len(data); row++ {
		line := data[row]
		if line == "" {
			break
		}
		ids := strings.Split(line, ",")
		for _, id := range ids {
			val, _ := strconv.Atoi(id)
			result.mine.values = append(result.mine.values, val)
		}
	}

	row += 2 // skip header
	for ; row < len(data); row++ {
		line := data[row]
		if line == "" {
			break
		}

		ids := strings.Split(line, ",")
		var t ticket
		for _, id := range ids {
			val, _ := strconv.Atoi(id)
			t.values = append(t.values, val)
		}
		result.nearby = append(result.nearby, t)
	}

	return result
}

func inFields(val int, m map[string]fieldData) (string, bool) {
	for f, v := range m {
		_, ok := v.ids[val]
		if ok {
			return f, true
		}
	}

	return "", false
}

func notInFields(val int, m map[string]fieldData) []string {
	var result []string
	for field, v := range m {
		_, ok := v.ids[val]

		if !ok {
			result = append(result, field)
		}
	}

	return result
}

func (n notes) errorRate() int {
	rate := 0
	for _, near := range n.nearby {
		for _, val := range near.values {
			_, ok := inFields(val, n.fields)

			if !ok {
				rate += val
			}
		}
	}

	return rate
}

func (n *notes) removeInvalid() {
	var valid []ticket
	for _, near := range n.nearby {
		nrValidFields := 0
		for _, val := range near.values {
			_, ok := inFields(val, n.fields)

			if ok {
				nrValidFields++
			}
		}

		if nrValidFields == len(near.values) {
			valid = append(valid, near)
		}
	}

	n.nearby = valid
}

func getHead(m map[int]struct{}) int {
	for key, _ := range m {
		return key
	}

	return 0
}

func allUnique(m map[string]map[int]struct{}) bool {
	for _, f := range m {
		if len(f) > 1 {
			return false
		}
	}

	return true
}

func main() {
	data := readLines("input.txt")

	n := parse(data)
	rate := n.errorRate()
	fmt.Printf("Result A: %d\n", rate)

	// field -> [index]
	fieldPositions := make(map[string]map[int]struct{})
	for field, _ := range n.fields {
		positions := make(map[int]struct{})
		for i := 0; i < len(n.mine.values); i++ {
			positions[i] = struct{}{}
		}
		fieldPositions[field] = positions
	}

	n.removeInvalid()

	for _, near := range n.nearby {
		for idx, t := range near.values {
			notFields := notInFields(t, n.fields)

			for _, notField := range notFields {
				delete(fieldPositions[notField], idx)
			}
		}
	}

	fields := make(map[int]string)
	for {
		if allUnique(fieldPositions) {
			break
		}

		for f, fieldPos := range fieldPositions {
			if len(fieldPos) == 1 {
				unique := getHead(fieldPos)
				fields[unique] = f

				for _, pos := range fieldPositions {
					delete(pos, unique)
				}
			}
		}
	}

	resultB := 1
	for i, t := range n.mine.values {
		rowName := fields[i]
		if strings.Contains(rowName, "departure") {
			resultB *= t
		}
	}

	fmt.Printf("Result B: %d\n", resultB)
}
