package day21

import (
	"bufio"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/yulrizka/adventofcode/pkg/aoc"

	"github.com/yulrizka/rxscan"
)

func Part1(f io.Reader) (string, error) {
	_, sum := getAllergen(f)
	return strconv.Itoa(sum), nil
}

func Part2(f io.Reader) (string, error) {
	allergens, _ := getAllergen(f)

	ingredientToAllergen := map[string]string{}
	var sortedIngredient []string
	for allergen, ingredient := range allergens {
		ingredientToAllergen[ingredient] = allergen
		sortedIngredient = append(sortedIngredient, ingredient)
	}
	sort.Strings(sortedIngredient)

	var ans []string
	for _, ingredient := range sortedIngredient {
		ans = append(ans, ingredientToAllergen[ingredient])
	}

	return strings.Join(ans, ","), nil
}

var rxIngredients = regexp.MustCompile(`(.+) \(contains (.+)\)`)

func getAllergen(f io.Reader) (map[string]string, int) {
	s := bufio.NewScanner(f)
	imap := map[string][][]string{}
	ingredientFreq := map[string]int{}

	for s.Scan() {
		var is, as string
		n, err := rxscan.Scan(rxIngredients, s.Text(), &is, &as)
		aoc.NoError(err)
		if n == 0 {
			panic("not parsed")
		}
		ingredients := strings.Split(is, " ")
		allergens := strings.Split(as, ", ")
		for _, a := range allergens {
			imap[a] = append(imap[a], ingredients)
		}
		for _, i := range ingredients {
			ingredientFreq[i]++
		}
	}

	// from foods (list of ingredient) for a particular allergen, find intersection of ingredient that always
	// appear with the allergen.
	intersections := map[string]map[string]struct{}{}
	for allergen, foods := range imap {
		m := map[string]int{}
		for _, food := range foods {
			for _, ingredient := range food {
				m[ingredient]++
			}
		}
		for ingredient, count := range m {
			if count >= len(foods) {
				if _, ok := intersections[allergen]; !ok {
					intersections[allergen] = map[string]struct{}{}
				}
				intersections[allergen][ingredient] = struct{}{}
			}
		}
	}

	// apparently at this point there is one ingredient that map with one allergen. We then go trough the
	// process of elimination to figure out the other allergen

	ingredientToAllergen := map[string]string{}

loop:
	// until we process all,
	for len(intersections) != 0 {
		//  find ingredient that only have one match
		for allergen, ingredientMap := range intersections {
			if len(ingredientMap) == 1 {
				for ingredient := range ingredientMap {
					// save ingredient -> allergen
					ingredientToAllergen[ingredient] = allergen

					// delete this new ingredient from the intersection
					for _, ingredientMap := range intersections {
						delete(ingredientMap, ingredient)
					}

					// delete the known allergen from intersection
					delete(intersections, allergen)
					continue loop
				}
			}
		}
	}

	sum := 0
	for ingredient, count := range ingredientFreq {
		if _, ok := ingredientToAllergen[ingredient]; ok {
			continue // it's an allergen
		}
		sum += count
	}

	return ingredientToAllergen, sum
}
