package makeheap

import (
	"bufio"
	"errors"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestHeapBuilder_Heapify(t *testing.T) {
	h := &HeapBuilder{
		maxSize: 5,
		size:    5,
		data:    []int{5, 4, 3, 2, 1},
		swaps:   make([]swap, 0),
	}
	h.Heapify()
	swapsGot := h.swaps
	swapsWanted := []swap{{to: 4, from: 1}, {to: 1, from: 0}, {to: 3, from: 1}}
	if !reflect.DeepEqual(swapsWanted, swapsGot) {
		t.Logf("Expected swaps %v got %v", swapsWanted, swapsGot)
		t.FailNow()
	}

}

func TestHeapBuilder_HeapifyFromData(t *testing.T) {
	tests := []struct {
		name       string
		inputPath  string
		outputPath string
	}{
		{name: "04", inputPath: "tests/04", outputPath: "04.a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, err := initHeapBuilder(t, tt.inputPath)
			if err != nil {
				t.Fatalf("Failed to parse test case file %s\n", tt.inputPath)
			}
			h.Heapify()
			got, _ := readOutput(t, tt.outputPath)
			if reflect.DeepEqual(h.swaps, got) {
				t.Logf("Swaps are not same wantLength")
				t.FailNow()
			}
		})
	}
}

func readOutput(t *testing.T, outputPath string) ([]swap, error) {
	t.Helper()
	f, err := os.Open(outputPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	l1, _ := reader.ReadString('\n')
	size, err := strconv.Atoi(strings.TrimSpace(l1))
	if err != nil {
		return nil, err
	}
	swaps := make([]swap, size)
	var l2 string
	for i := 0; i < size; i++ {
		l2, err = reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		sl := strings.Split(strings.TrimSpace(l2), " ")
		from, _ := strconv.Atoi(sl[0])
		if err != nil {
			return nil, err
		}
		to, _ := strconv.Atoi(sl[1])
		if err != nil {
			return nil, err
		}
		swaps = append(swaps, swap{to: to, from: from})
	}
	return swaps, nil

}

func initHeapBuilder(t *testing.T, filePath string) (*HeapBuilder, error) {
	t.Helper()
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	l1, _ := reader.ReadString('\n')
	size, err := strconv.Atoi(strings.TrimSpace(l1))
	if err != nil {
		return nil, err
	}
	data := make([]int, size)
	l2, _ := reader.ReadString('\n')
	ints := strings.Split(l2, " ")
	if len(ints) != size {
		return nil, errors.New("input size mismatch")
	}
	for i := 0; i < size; i++ {
		entry, err := strconv.Atoi(strings.TrimSpace(ints[i]))
		if err != nil {
			log.Panic("input is not an integer")
		}
		data[i] = entry
	}
	return NewBuildHeap(data), nil

}
