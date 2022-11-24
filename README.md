# Glogger
Glogger helps to write log information

## Usage
Example of usage:
```
import glogger "test/Glogger"

...

glog, err := glogger.Init("log.txt") // init log file (created if not exist)
if err != nil { 
  log.Fatal(err)
}

glog.Info("Тестовая информация, которая запишется в лог-файл")
glog.Error(errors.New("Ошибка которая запишется в лог-файл"))
	
glog.Close()
```
