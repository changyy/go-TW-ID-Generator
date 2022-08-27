package main

import (
    "fmt"
    "time"
    "flag"
    "sync"
    "os"
    "runtime"
)

var (
    firstValue map[string]int
    verifiedIDCount int = 0
    mux sync.Mutex
    start time.Time
    taskReportCount = flag.Int("report", 1000000, "report")
    outputDir = flag.String("output", "/tmp", "output dir")
    inputID = flag.String("verify", "", "verify ID")
    cpu = flag.Int("cpu", 1, "using multicore processor")
)

func buildFirstValue() {
    firstValue = map[string]int {
        "A": 10, "B": 11, "C": 12, "D": 13, "E": 14, "F": 15, "G": 16, "H": 17, "J": 18, "K": 19, "L": 20, "M": 21, "N": 22,
        "P": 23, "Q": 24, "R": 25, "S": 26, "T": 27, "U": 28, "V": 29, "X": 30, "Y": 31, "W": 32, "Z": 33, "I": 34, "O": 35,
    }
}

func verifyTWROCIDRule(id string) (bool) {
    if len(id) != 10 {
        return false
    }
    if id[0] < 'A' && id[0] > 'Z' {
        return false
    }
    rawID := fmt.Sprintf( "%d%s", firstValue[ string(id[0]) ], id[1:])
    //fmt.Println(rawID)

    val := int( (rawID[0] - '0') + (rawID[10] - '0') )

    j := 1
    for i := 9 ; i > 0 ; i-- {
        //fmt.Println("val: ", val, ", next:", int(rawID[j] - '0'), ", calc:", i)
        val += int(rawID[j] - '0') * i
        j++
    }
    //fmt.Println(val)

    return val % 10 == 0
}

func writeResult(taskID int, in map[string]bool) {
    filename := fmt.Sprintf("%s/ids-%d.txt", *outputDir, taskID) 
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    for id := range in {
        if _, err = f.WriteString(id+"\n"); err != nil {
            panic(err)
        }
    }
}

func generateID(taskID int, firstWord map[string]int) {
    taskVerifiedIDCount := 0
    output := make(map[string]bool)
    for word := range firstWord {
        for i1 := 1 ; i1 <= 2 ; i1 ++ {
            for i2 := 0 ; i2 < 10 ; i2 ++ {
                for i3 := 0 ; i3 < 10 ; i3 ++ {
                    for i4 := 0 ; i4 < 10 ; i4 ++ {
                        for i5 := 0 ; i5 < 10 ; i5 ++ {
                            for i6 := 0 ; i6 < 10 ; i6 ++ {
                                for i7 := 0 ; i7 < 10 ; i7 ++ {
                                    for i8 := 0 ; i8 < 10 ; i8 ++ {
                                        for i9 := 0 ; i9 < 10 ; i9 ++ {
                                            id := fmt.Sprintf("%s%d%d%d%d%d%d%d%d%d", word, i1, i2, i3, i4, i5, i6, i7, i8, i9)
                                            if verifyTWROCIDRule(id) {
                                                taskVerifiedIDCount ++
                                                output[id] = true
                                                if taskVerifiedIDCount % *taskReportCount == 0 { 
                                                    elapsed := time.Since(start)
                                                    mux.Lock()
                                                    verifiedIDCount += taskVerifiedIDCount
                                                    mux.Unlock()
                                                    var m runtime.MemStats
                                                    runtime.ReadMemStats(&m)
                                                    fmt.Printf("taskID: %d, taskCount: %10d, totalCount: %10d, timeCost: %s, mem: %.2v MB\n", taskID, taskVerifiedIDCount, verifiedIDCount, elapsed, m.Alloc / 1024/ 1024)
                                                    writeResult(taskID, output)
                                                    // reset
                                                    taskVerifiedIDCount = 0
                                                    output = make(map[string]bool)
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    if taskVerifiedIDCount > 0 {                                                
        writeResult(taskID, output)
    }
    mux.Lock()
    verifiedIDCount += taskVerifiedIDCount
    mux.Unlock()
}

func main() {
    flag.Parse()
    start = time.Now()
    buildFirstValue() 

    if *inputID != "" {
        testResult := verifyTWROCIDRule( *inputID )
        elapsed := time.Since(start)
        if testResult {
            fmt.Println(*inputID, ": PASS, time cost:", elapsed)
        } else {
            fmt.Println(*inputID, ": FAIL, time cost:", elapsed)
        }
        return
    }
    useCore := *cpu
    if useCore < 0 {
        useCore = 1
    }

    chunks := make([]map[string]int, useCore)
    for i := 0 ; i < useCore ; i++ {
        chunks[i] = make(map[string]int)
    }
    chunkSize := len(firstValue)/useCore
    if chunkSize * useCore < len(firstValue) {
        chunkSize ++
    }
    chunkIndex := 0
    for w := range firstValue {
        if len(chunks[chunkIndex]) == chunkSize && chunkIndex + 1 < chunkSize {
            chunkIndex++
        }
        chunks[chunkIndex][w] = firstValue[w]
    }

    wg := new(sync.WaitGroup)
    wg.Add(useCore)

    fmt.Println("[INFO] using core:", useCore)
    fmt.Println("[INFO] output dir:", *outputDir)
    fmt.Println("[INFO] state report when generated IDs count:", *taskReportCount)
    for i := 0 ; i < useCore ; i++ {
        taskID := i
        taskInfo := chunks[i]
        go func() {
            defer wg.Done() 
            generateID(taskID, taskInfo)
        }()
    }
    wg.Wait()

    elapsed := time.Since(start)
    fmt.Println("Total verified IDs count:", verifiedIDCount, ", time cost:", elapsed)
}
