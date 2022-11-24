# Glogger
Glogger helps to write log information

## Usage
Example of usage:
```
import glogger "test/Glogger"

...

glog, err := glogger.Init("log.txt") 
if err != nil { 
  log.Fatal(err)
}

glog.Info("Information that will be written to the log file")
glog.Error(errors.New("Error that will be written to the log file"))
	
glog.Close()
```
