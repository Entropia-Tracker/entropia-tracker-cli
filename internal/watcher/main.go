package watcher

import (
  "bufio"
  "io"
  "os"
  "time"
)

func Parse(path string, c chan string, all, watch bool) {
  handle, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  defer handle.Close()

  stat, err := os.Stat(path)
  if err != nil {
    panic(err)
  }

  previousLogSize := stat.Size()

  if !all {
    handle.Seek(previousLogSize, 0)
  }

  reader := bufio.NewReader(handle)

  for {
    // Check if log was truncated
    stat, err = os.Stat(path)
    if err != nil {
      panic(err)
    }

    currentLogSize := stat.Size()
    if previousLogSize > currentLogSize {
      handle.Seek(currentLogSize, 0)
    }
    previousLogSize = currentLogSize

    line, _, err := reader.ReadLine()

    if err == io.EOF { // End of file, wait and try again if watch is set
      if watch {
        time.Sleep(time.Second)
        continue
      } else {
        c <- "EOF"
        break
      }
    } else if err != nil { // Something bad happened
      panic(err)
    } else if len(line) == 0 { // Empty line, ignore
      continue
    }

    c <- string(line)
  }
}
