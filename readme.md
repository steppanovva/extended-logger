Logger extended with a log-level (info, warning, error) parameter.

```
Example:

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(Warning)
	logger.PrintInfo("Should not be printed") // because the "info" level lower in priority than "warning" level of this logger
	logger.PrintWarning("Warning")
	logger.PrintError("Error")
	logger.Println("DEBUG")
}


Result:

2022/07/13 20:24:02 WARNING: Warning
2022/07/13 20:24:02 ERROR: Error
2022/07/13 20:24:02 DEBUG
```
