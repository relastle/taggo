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
		ColorFuncMap[tagColor](tag),
		tagDelimiter,
		line,
	)
}

// colorColumnStep colors `index`th column
// with a given color
func colorColumnStep(line string) string {
	elms := strings.Split(line, delimiter)
	for index, color := range colorizer {
		if !(0 <= index && index < len(elms)) {
			continue
		}
		elms[index] = ColorFuncMap[color](elms[index])
	}
	return fmt.Sprintf(
		"%v",
		strings.Join(elms, delimiter),
	)
}

// addIconsStep add nerd icons to `nerdIndex`th column
func addIconsStep(line string) string {
	elms := strings.Split(line, delimiter)
	for _, index := range iconIndices {
		if !(0 <= index && index < len(elms)) {
			return line
		}
		elms[index] = addIcon(elms[index])
	}
	return fmt.Sprintf(
		"%v",
		strings.Join(elms, delimiter),
	)
}

// MainStream is a main I/O stream of taggo
func MainStream() {
	mss := &multiStepStream{}
	if len(colorizer) > 0 {
		mss.addStep(colorColumnStep)
	}
	if len(iconIndices) > 0 {
		mss.addStep(addIconsStep)
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
