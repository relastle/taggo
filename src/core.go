package taggo

import (
	"bufio"
	"fmt"
	"os"
	"path"
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

func addTagtep(line string) string {
	return fmt.Sprintf(
		"%v%v%v",
		ColorFuncMap[tagColor](tag),
		tagDelimiter,
		line,
	)
}

func removeTagStep(line string) string {
	elms := strings.Split(line, delimiter)
	return strings.Join(elms[1:], tagDelimiter)
}

func addBasenameStep(line string) string {
	elms := strings.Split(line, delimiter)
	pathUsed := elms[basenamedIndex]
	basename := path.Base(pathUsed)
	return fmt.Sprintf(
		"%v%v%v",
		addIcon(basename),
		delimiter,
		line,
	)
}

func removeBasenameStep(line string) string {
	elms := strings.Split(line, delimiter)
	return strings.Join(elms[1:], delimiter)
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

// removeIconStep add nerd icons to `nerdIndex`th column
func removeIconStep(line string) string {
	elms := strings.Split(line, delimiter)
	for _, index := range iconIndices {
		if !(0 <= index && index < len(elms)) {
			return line
		}
		elms[index] = removeIcon(elms[index])
	}
	return fmt.Sprintf(
		"%v",
		strings.Join(elms, delimiter),
	)
}

func decorateStream() {
	mss := &multiStepStream{}
	// append `COLORIZE STEP`
	if len(colorizer) > 0 {
		mss.addStep(colorColumnStep)
	}
	// append `ICON STEP`
	if len(iconIndices) > 0 {
		mss.addStep(addIconsStep)
	}
	// append `ADD BASE NAME STEP`
	if basenamedIndex >= 0 {
		mss.addStep(addBasenameStep)
	}
	// append `ADD TAG STEP`
	if tag != "" {
		mss.addStep(addTagtep)
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

// revertStream reverts lines into original ones.
// It assume that all ANSI color codes were removed
func revertStream() {
	mss := &multiStepStream{}

	// append `REMOVE TAG STEP`
	if tag != "" {
		mss.addStep(removeTagStep)
	}

	// append `REMOVE BASENAME STEP`
	if basenamedIndex >= 0 {
		mss.addStep(removeBasenameStep)
	}

	// append `REMOVE ICON STEP`
	if len(iconIndices) > 0 {
		mss.addStep(removeIconStep)
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

// MainStream is a main I/O stream of taggo
func MainStream() {
	if !revertFlag {
		decorateStream()
	} else {
		revertStream()
	}
}
