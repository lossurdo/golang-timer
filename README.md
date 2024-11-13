# golang-timer

Linux-like timer command for Windows.

## Overview

`golang-timer` is a simple command-line tool written in Go that mimics the functionality of the Linux `time` command for Windows. It measures the duration of a command's execution and displays the elapsed time.

## Usage

1. **Install Go**:
   - Ensure you have Go installed on your system. You can download and install it from the [official Go website](https://golang.org/dl/).

2. **Build the Project**:
   - Clone the repository.
   - Open a terminal and navigate to the project directory.
   - Run `go build -o timer.exe`.

3. **Run a Command with Timer**:
   - Open a terminal in the directory containing `timer.exe`.
   - Execute a command using `timer.exe` followed by the command you want to time.

   ```shell
   timer.exe <command> [args]
   ```

   Replace `<command>` with the command you want to measure, and `[args]` with any arguments needed for the command.

4. **Add timer.exe to PATH**:
   - For convenience, add `timer.exe` to your system's PATH environment variable.
   - This allows you to run `timer.exe` from any directory without specifying its full path.

## Example

```shell
timer.exe curl https://httpbin.org/ip
1.5725139s        <<<<<<---- time it took to complete
```

This will execute the `curl` command and display the time it took to complete.

## License

This project is licensed under the MIT License.
