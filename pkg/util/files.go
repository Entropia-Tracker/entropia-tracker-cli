package util

import (
  "bufio"
  "github.com/fsnotify/fsnotify"
  "io"
  "os"
  "path/filepath"
  "time"
)

// ReadFile reads/watches a file and calls the callback function for each row.
func ReadFile(path string, fromStart, watchFile bool, callback func(string)) error {

  // Open file handle for reading
  handle, err := os.Open(path)
  if err != nil {
    return err
  }
  defer handle.Close()

  // Read initial size
  stat, err := os.Stat(path)
  if err != nil {
    return err
  }
  previousLogSize := stat.Size()

  // Setup reader
  reader := bufio.NewReader(handle)

  if !fromStart { // Seek to end of file if neccessary
    handle.Seek(previousLogSize, 0)
  } else { // Read all lines in the file
    for {
      line, _, err := reader.ReadLine()

      if err == io.EOF { // Reached end of logfile.
        break
      } else if err != nil { // Something bad happened
        return err
      } else if len(line) == 0 { // Empty line, ignore
        continue
      }

      // Send line to callback
      callback(string(line))
    }
  }

  // Setup watcher and watch for changes
  if watchFile {

    // Setup watcher
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
      return err
    }
    defer watcher.Close()

    // Watch the logfile directory instead of the logfile path so that
    // we can catch DELETE and CREATE events in addition to WRITE.
    directory := filepath.Dir(path)
    err = watcher.Add(directory)
    if err != nil {
      return err
    }

    for {
      select {
      case event, ok := <-watcher.Events:
        if !ok || filepath.Base(event.Name) != "chat.log" {
          // Only care about events on "chat.log"
          continue

        } else if event.Op&fsnotify.Remove == fsnotify.Remove {
          // File was removed, do nothing
          continue
        } else if event.Op&fsnotify.Rename == fsnotify.Rename {
          // File was renamed, do nothing
          continue

        } else if event.Op&fsnotify.Create == fsnotify.Create {
          // When "chat.log" is created we need to initialize the
          // handle and reader instances for the new file
          handle, err = os.Open(path)
          if err != nil {
            return err
          }

          reader = bufio.NewReader(handle)
        }

        // If the logfile became smaller, we start from the beginning again
        currentLogSize, err := getFileSize(path)
        if err != nil {
          return err
        }

        if previousLogSize > currentLogSize {

          // Handle atomic saves
          if currentLogSize == 0 {
            // Wait for a second to hopefully give enough time to fill
            // the new file.
            time.Sleep(time.Second)

            // Get the new filesize
            currentLogSize, err := getFileSize(path)
            if err != nil {
              return err
            }

            // Verify again that it indeed became smaller and
            // seek to end of file
            if previousLogSize > currentLogSize {
              handle.Seek(currentLogSize, 0)
            }

          } else {
            // Seek to end of file
            handle.Seek(currentLogSize, 0)
          }
        }
        previousLogSize = currentLogSize

        // Read all lines from current position
        for {
          line, _, err := reader.ReadLine()

          if err == io.EOF { // Reached end of logfile.
            break
          } else if err != nil { // Something bad happened
            return err
          } else if len(line) == 0 { // Empty line, ignore
            continue
          }

          // Send line to callback
          callback(string(line))
        }
      }
    }
  }

  return nil
}

// getFileSize helper
func getFileSize(path string) (int64, error) {
  stat, err := os.Stat(path)
  if err != nil {
    return 0, err
  }

  return stat.Size(), nil
}
