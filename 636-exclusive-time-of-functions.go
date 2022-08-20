package main

import (
	"container/list"
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
    times := make([]int, n)
    var sp = list.New()
    last_ts := 0

    for _, log := range logs {
        parts := strings.Split(log, ":")
        id, _ := strconv.Atoi(parts[0])
        op := parts[1]
        ts, _ := strconv.Atoi(parts[2])

        switch op {
        case "start":
            if e := sp.Front(); e != nil {
                f := e.Value.(int)
                times[f] += ts-last_ts
            }
            last_ts = ts
            sp.PushFront(id)

        case "end":
            f := sp.Remove(sp.Front())
            if f != id {
                panic("stack corrupted")
            }
            times[id] += ts - last_ts + 1
            last_ts = ts+1
        }
    }

	return times
}

func exclusiveTime2(n int, logs []string) []int {
    times := make([]int, n)
    var stack [500]int
    sp := 0
    
    last_ts := 0

    for _, log := range logs {
        parts := strings.Split(log, ":")
        id, _ := strconv.Atoi(parts[0])
        op := parts[1]
        ts, _ := strconv.Atoi(parts[2])

        switch op {
        case "start":
            if sp > 0 {
                times[stack[sp-1]] += ts - last_ts
            }
            last_ts = ts
            stack[sp] = id
            sp++

        case "end":
            f := stack[sp-1]
            if f != id {
                panic("stack corrupted")
            }
            sp--
            times[id] += ts - last_ts + 1
            last_ts = ts+1
        }
    }

	return times
}
