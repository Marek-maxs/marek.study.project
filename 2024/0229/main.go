package main

import (
    "fmt"
    "log"
    "net"
    "time"
)

func main() {
    a := []int{1,2,2,3}

    tsArr := UpdateEnergySummary(time.Now().Add(time.Hour*-3), time.Now())

    for idx := 0; idx < len(tsArr); idx++ {
        fmt.Println(tsArr[idx])
    }

    a := time.Date(2024, time.Now().Month(), 1,0,0,0,0,time.Local)
    fmt.Println(a.Hour())
}

func UpdateEnergySummary(unixStartTimeUTC,
    unixEndTimeUTC time.Time) []time.Time {
    // Get the start of the current hour in UTC.
    startCurrentHour := time.Date(
        unixStartTimeUTC.Year(),
        unixStartTimeUTC.Month(),
        unixStartTimeUTC.Day(),
        unixStartTimeUTC.Hour(),
        0,
        0,
        0,
        time.UTC,
    )

    // Get the end of the current hour in UTC.
    endCurrentHour := time.Date(
        unixEndTimeUTC.Year(),
        unixEndTimeUTC.Month(),
        unixEndTimeUTC.Day(),
        unixEndTimeUTC.Hour()+1, // Next hour
        0,
        0,
        0,
        time.UTC,
    )

    var stTimeList []time.Time
    for t := startCurrentHour; !t.After(endCurrentHour); t = t.Add(time.Hour) {
        stTimeList = append(stTimeList, t)
    }

    return stTimeList
}


func ChangeName() {
    // Part 1: create a listener
    l, err := net.Listen("tcp", ":8000")
    if err != nil {
        log.Fatalf("Error listener returned: %s", err)
    }
    defer l.Close()

    for {
        // Part 2: accept new connection
        c, err := l.Accept()
        if err != nil {
            log.Fatalf("Error to accept new connection: %s", err)
        }

        // Part 3: create a goroutine that reads and write back data
        go func() {
            log.Printf("TCP session open")
            defer c.Close()

            for {
                d := make([]byte, 1024)

                // Read from TCP buffer
                _, err := c.Read(d)
                if err != nil {
                    log.Printf("Error reading TCP session: %s", err)
                    break
                }
                log.Printf("reading data from client: %s\n", string(d))

                // write back data to TCP client
                _, err = c.Write(d)
                if err != nil {
                    log.Printf("Error writing TCP session: %s", err)
                    break
                }
            }
        }()
    }
}