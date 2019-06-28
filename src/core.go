package taggo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// multiStepStream is an struct
// that defines functions for multi step
// processing against each line
type multiStepStream struct {
	steps [](func(string) string)
}

func (mss *multiStepStream) addStep(step func(string) string) {
	mss.steps = append(mss.steps, step)
}

func (mss *multiStepStream) streamLine(line string) string {
	for _, step := range mss.steps {
		line = step(line)
	}
	return line
}

func simpleTagAddStep(line string) string {
	return fmt.Sprintf(
		"%v%v%v",
		colorFuncMap[colorStr](tag),
		delimiter,
		line,
	)
}

// colorColumnStep colors `index`th column
// with a given color
func colorColumnStep(line string) string {
	elms := strings.Split(line, delimiter)
	if index >= len(elms) {
		return line
	}
	elms[index] = colorFuncMap[colorStr](elms[index])
	return fmt.Sprintf(
		"%v",
		strings.Join(elms, delimiter),
	)
}

// addNerdBadgeStep add nerd badge to `nerdIndex`th column
func addNerdBadgeStep(line string) string {
	elms := strings.Split(line, delimiter)
	if nerdIndex >= len(elms) {
		return line
	}
	elms[nerdIndex] = addBadge(elms[nerdIndex])
	return fmt.Sprintf(
		"%v",
		strings.Join(elms, delimiter),
	)
}

// MainStream is a main I/O stream of taggo
func MainStream() {
	mss := &multiStepStream{}
	if index >= 0 {
		mss.addStep(colorColumnStep)
	}
	if nerdIndex >= 0 {
		mss.addStep(addNerdBadgeStep)
	}
	if tag != "" {
		mss.addStep(simpleTagAddStep)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = mss.streamLine(text)
		fmt.Println(text)
	}

	if scanner.Err() != nil {
	}
}
