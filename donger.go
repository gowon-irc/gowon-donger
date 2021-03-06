package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type dongers map[string][]string

func (d dongers) randomCat(r *rand.Rand) string {
	keys := make([]string, len(d))

	i := 0
	for k := range d {
		keys[i] = k
		i++
	}

	return keys[r.Intn(len(keys))]
}

func (d *dongers) random(r *rand.Rand) string {
	cat := d.randomCat(r)
	return d.randomFromCat(r, cat)
}

func (d dongers) randomFromCat(r *rand.Rand, cat string) string {
	if _, prs := d[cat]; prs == false {
		return fmt.Sprintf("No %s category", cat)
	}

	ri := r.Intn(len(d[cat]))

	return d[cat][ri]
}

func donger(args string) (msg string, err error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if len(args) > 0 {
		cat := strings.Split(args, " ")[0]
		return dongerCollection.randomFromCat(r, cat), nil
	}

	return dongerCollection.random(r), nil
}
