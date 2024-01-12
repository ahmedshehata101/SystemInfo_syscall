package main

import (
	"fmt"
	"os"
	"regexp"
	sc "syscall"
)

func getprocesslist() ([]string, int) {
	var re = regexp.MustCompile(`^[0-9]+$`)
	var counter = 0
	processes, err := os.ReadDir("/proc")
	if err != nil {
		panic(err)
	}
	array := make([]string, 0, len(processes))

	for _, val := range processes {
		if re.MatchString(val.Name()) {
			array = append(array, val.Name())
			counter++
		}
	}

	return array, counter
}
func getdiskstats() (int, int, int) {
	fsstat := sc.Statfs_t{}
	var total uint64
	var used uint64
	var free uint64
	err := sc.Statfs("/", &fsstat)
	if err != nil {
		panic(err)
	}
	total = (fsstat.Blocks * uint64(fsstat.Bsize)) / 1024 / 1024 / 1024
	free = (fsstat.Bfree * uint64(fsstat.Bsize)) / 1024 / 1024 / 1024
	used = total - free
	return int(total), int(free), int(used)
}

func getsysteminfo() (int, int, int, float64, float64, float64) {
	sysinfo := sc.Sysinfo_t{}

	var uptime uint64
	var totalram uint64
	var totalswap uint64
	sc.Sysinfo(&sysinfo)

	uptime = uint64(sysinfo.Uptime) / 60

	totalswap = sysinfo.Totalswap
	totalram = sysinfo.Totalram / 1024 / 1024
	loadavg1 := float64(sysinfo.Loads[0]) / 65536
	loadavg5 := float64(sysinfo.Loads[1]) / 65536
	loadavg15 := float64(sysinfo.Loads[2]) / 65536
	return int(uptime), int(totalram), int(totalswap), loadavg1, loadavg5, loadavg15

}

func main() {
	var numofprocess int
	totaldisk, freedisk, useddisk := getdiskstats()
	_, numofprocess = getprocesslist()
	uptime, totalram, totalswap, loadavg1, loadavg5, loadavg15 := getsysteminfo()
	fmt.Println("Details will be as follow:")
	fmt.Println("==========================")
	fmt.Println("Disk Space ", totaldisk, "GB")
	fmt.Println("Disk usage", useddisk, "GB")
	fmt.Println("Free Disk Space", freedisk, "GB")
	fmt.Println("Uptime", uptime, "min")
	fmt.Println("Number of Processses", numofprocess-1)
	fmt.Println("Total Swap Memory", totalswap, "MB")
	fmt.Println("Total Memory", totalram, "MB")
	fmt.Printf("Load AVg for 1 min :%.2f\n", loadavg1)
	fmt.Printf("Load AVg for 5 min :%.2f\n", loadavg5)
	fmt.Printf("Load AVg for 15 min :%.2f\n", loadavg15)
	fmt.Println("--------------------------------------------------")

}
