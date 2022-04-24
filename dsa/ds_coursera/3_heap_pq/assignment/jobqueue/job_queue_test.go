package jobqueue

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestJobQueue_Start(t *testing.T) {
	tasks := make([]int, 5)
	for i := 0; i < 5; i++ {
		tasks[i] = i
	}
	queue := NewJobQueue(tasks, 2)
	queue.Start()
	for _, ai := range queue.assignments {
		fmt.Printf("%d %d\n", ai.thread, ai.startTime)
	}
}

func TestJobQueue_Start20(t *testing.T) {
	tasks := make([]int, 20)
	for i := 0; i < 20; i++ {
		tasks[i] = 1
	}
	queue := NewJobQueue(tasks, 4)
	queue.Start()
	for _, ai := range queue.assignments {
		fmt.Printf("%d %d\n", ai.thread, ai.startTime)
	}
}

func TestJobQueue_Files(t *testing.T) {
	tests := []struct {
		name       string
		inputPath  string
		outputPath string
	}{
		{name: "02", inputPath: "tests/02", outputPath: "tests/02.a"},
		{name: "08", inputPath: "tests/08", outputPath: "tests/08.a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jq, err := initJobQueue(t, tt.inputPath)
			if err != nil {
				t.Fatalf("Failed to parse test case file %s\n", tt.inputPath)
			}
			jq.Start()
			want, err := readOutput(t, len(jq.tasks), tt.outputPath)
			if err != nil {
				t.Fatalf("Failed to parse test case file %s\n", tt.inputPath)
			}
			if !reflect.DeepEqual(jq.assignments, want) {
				t.Logf("Swaps are not same wantLength")
				t.FailNow()
			}
		})
	}
}

func initJobQueue(t *testing.T, path string) (*JobQueue, error) {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	l1, _ := reader.ReadString('\n')
	c := strings.Split(strings.TrimSpace(l1), " ")
	l2, _ := reader.ReadString('\n')
	ets := strings.Split(strings.TrimSpace(l2), " ")
	tc, _ := strconv.Atoi(c[1])
	tasks := make([]int, tc)
	for idx, et := range ets {
		tm, _ := strconv.Atoi(et)
		tasks[idx] = tm
	}
	threads, _ := strconv.Atoi(c[0])
	return NewJobQueue(tasks, threads), err
}

func readOutput(t *testing.T, taskLen int, outputPath string) ([]*AssignmentInfo, error) {
	t.Helper()
	ais := make([]*AssignmentInfo, taskLen)
	f, err := os.Open(outputPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	lc := 0
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, errors.New("error reading file")
		}
		c := strings.Split(strings.TrimSpace(line), " ")
		thread, _ := strconv.Atoi(c[0])
		startTime, _ := strconv.Atoi(c[1])
		ais[lc] = &AssignmentInfo{thread: thread, startTime: startTime}
		lc++
	}
	return ais, nil
}
