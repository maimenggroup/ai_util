package ai_util

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

func PprofMem(suffix string) error {
	fmt.Printf("PprofMem start. suffix: [%v]\n", suffix)

	for _, name := range []string{
		"heap", "block", "goroutine", "threadcreate",
	} {
		fp, err := os.OpenFile(fmt.Sprintf("%v.out.%v", name, suffix), os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("PprofMem failed. suffix: [%v], err: [%v]", suffix, err)
		}
		if err := pprof.Lookup(name).WriteTo(fp, 1); err != nil {
			fp.Close()
			return fmt.Errorf("PprofMem failed. suffix: [%v], err: [%v]", suffix, err)
		}
		fp.Close()
	}

	fmt.Printf("PprofMem success. suffix: [%v]\n", suffix)

	return nil
}

func PprofCpu(suffix string, duration time.Duration) error {
	fmt.Printf("PprofCpu start. suffix: [%v], duration[%v]\n", suffix, duration)

	fp, err := os.OpenFile(fmt.Sprintf("pprofcpu.out.%v", suffix), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("PprofCpu failed. suffix: [%v], duration: [%v], err: [%v]", suffix, duration, err)
	}
	defer fp.Close()
	if err := pprof.StartCPUProfile(fp); err != nil {
		return fmt.Errorf("PprofCpu failed. suffix: [%v], duration: [%v], err: [%v]", suffix, duration, err)
	}
	time.Sleep(duration)
	pprof.StopCPUProfile()

	fmt.Printf("PprofCpu success. suffix: [%v], duration[%v]\n", suffix, duration)

	return nil
}

func PprofCmd(command string) string {
	fields := strings.Split(command, " ")
	cmdType := fields[0]

	switch cmdType {
	case "pporfmem":
		suffix := time.Now().Format("2006-01-02-15")
		if len(fields) >= 2 {
			suffix = fields[1]
		}
		if err := PprofMem(suffix); err != nil {
			return fmt.Sprintf("%v", err)
		} else {
			return fmt.Sprintf("pporfmem [%v] success", suffix)
		}
	case "pporfcpu":
		suffix := time.Now().Format("2006-01-02-15")
		duration := 60 * time.Second
		if len(fields) >= 2 {
			var err error
			duration, err = time.ParseDuration(fields[1])
			if err != nil {
				return fmt.Sprintf("%v", err)
			}
		}
		if len(fields) >= 3 {
			suffix = fields[2]
		}
		if err := PprofCpu(suffix, duration); err != nil {
			return fmt.Sprintf("%v", err)
		} else {
			return fmt.Sprintf("pporfcpu [%v] [%v] success", suffix, duration)
		}
	default:
		return fmt.Sprintf("invalid command[%v]", command)
	}
}
