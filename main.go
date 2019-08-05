package main

//Import the packages we need
import (
	"fmt"
	"os"
	"io"
	"strconv"
	"time"

	"github.com/sensu/sensu-go/types"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	
	"github.com/spf13/cobra"
)

//Set up some variables. Most notably, warning and critical as time durations
var (
	stdin   *os.File
)

//Start our main function
func main() {
	rootCmd := configureRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

//Set up our flags for the command. Note that we have time duration defaults for warning & critical
func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sensu-go-memory-checks",
		Short: "The Sensu Go check for system memory usage",
		RunE:  run,
	}

	cmd.Flags().StringVarP(&cpu,
		"cpu",
		"c",
		true,
		"Shows the info for system cpu.")
	
	cmd.Flags().StringVarP(&diskS,
		"disk",
		"d",
		true,
		"Shows the info for system disk.")

	cmd.Flags().StringVarP(&docker,
		"disk",
		"d",
		false,
		"Shows the info for running Docker containers.")
		
	cmd.Flags().StringVarP(&host,
		"host",
		"h",
		true,
		"Shows host info.")

	cmd.Flags().StringVarP(&memory,
		"memory",
		"m",
		true,
		"Shows the info for system memory.")
		
	cmd.Flags().StringVarP(&network,
		"network",
		"n",
		true,
		"Shows the info for system network.")
	
		
	return cmd
}

func run(cmd *cobra.Command, args []string) error {

	if len(args) != 0 {
		_ = cmd.Help()
		return fmt.Errorf("invalid argument(s) received")
	}

	if stdin == nil {
		stdin = os.Stdin
	}
	
	event := &types.Event{}
	
	if cpu == true {
		return cpuInfo(event)
	}
	if disk == true {
		return diskInfo(event)
	}
	if docker == true {
		return dockerInfo(event)
	}
	if host == true {
		return hostInfo(event)
	}
	if memory == true {
		return memInfo(event)
	}
	if network == true {
		return netInfo(event)
	}

}

//Here we start the meat of what we do.
func cpuInfo(event *types.Event) error {
	
	//Let's set up some error handling
	if err != nil {
		msg := fmt.Sprintf("Failed to determine disk info %s", err.Error())
		io.WriteString(os.Stdout, msg)
		os.Exit(2)
	}
	
	//Let's set up some vars
	interval := time.Millisecond * 300
	cpuStat, _ := cpu.InfoStat()
	cpuPct, _ := cpu.Percent(interval, false)

	//Setting up our message to print some info about CPU Percent
	msg := fmt.Sprintf("Current CPU Utilization: %.2f\%", cpuPct)
	
	//Writing msg to stdout
	io.WriteString(os.Stdout, msg)
	
	//Exiting with an OK
	os.Exit(0)
		
	return nil
}

//This function gathers some informational bits about the host that the check runs on
func hostInfo(event *types.Event) error {
	
	//Let's set up some error handling
	if err != nil {
		msg := fmt.Sprintf("Failed to determine host info %s", err.Error())
		io.WriteString(os.Stdout, msg)
		os.Exit(2)
	}
	
	//Let's set up some vars...
	hostStat, _ := host.Info()
	uptime, _ := host.Uptime()
	uptimeSecs := time.Duration(uptime)*time.Second

	//Getting our message about our host info ready to print
	msg := fmt.Sprintf("Hostname: %s\nOS: %s\nPlatform: %s\nUptime: %d seconds", hostStat.Hostname, hostStat.OS, hostStat.Platform, int64(uptimeSecs.Seconds()))
	
	//Writing msg to stdout
	io.WriteString(os.Stdout, msg)
	
	//Exiting with an OK
	os.Exit(0)
}

//This function gathers some infor about the disk(s) on a system
func diskInfo(event *types.Event) error {

	//Let's set up some error handling
	if err != nil {
		msg := fmt.Sprintf("Failed to determine disk info %s", err.Error())
		io.WriteString(os.Stdout, msg)
		os.Exit(2)
	}
	
	//Let's set up some vars
	diskStat, _ := disk.Usage()

	//Setting up our message to print some info about disks
	msg := fmt.Sprintf()
	
	//Writing msg to stdout
	io.WriteString(os.Stdout, msg)
	
	//Exiting with an OK
	os.Exit(0)
	
}

func dockerInfo(event *types.Event) error {

	//Let's set up some error handling
	if err != nil {
		msg := fmt.Sprintf("Failed to determine docker info %s", err.Error())
		io.WriteString(os.Stdout, msg)
		os.Exit(2)
	}
	
	//Let's set up some vars
	dockerStat, _ := docker.CgroupDockerStat()

	//Setting up our message to print some info about disks
	msg := fmt.Sprintf()
	
	//Writing msg to stdout
	io.WriteString(os.Stdout, msg)
	
	//Exiting with an OK
	os.Exit(0)
	
}

func memInfo(event *types.Event) error {

	//Let's set up some error handling
	if err != nil {
		msg := fmt.Sprintf("Failed to determine memory info %s", err.Error())
		io.WriteString(os.Stdout, msg)
		os.Exit(2)
	}
	
	//Let's set up some vars
	vmStat, _ := mem.VirtualMemoryStat()

	//Setting up our message to print some info about network interfaces
	msg := fmt.Sprintf()
	
	//Writing msg to stdout
	io.WriteString(os.Stdout, msg)
	
	//Exiting with an OK
	os.Exit(0)
}

func netInfo(event *types.Event) error {

	//Let's set up some error handling
	if err != nil {
		msg := fmt.Sprintf("Failed to determine network info %s", err.Error())
		io.WriteString(os.Stdout, msg)
		os.Exit(2)
	}
	
	//Let's set up some vars
	netStat, _ := net.Interfaces()

	//Setting up our message to print some info about network interfaces
	msg := fmt.Sprintf()
	
	//Writing msg to stdout
	io.WriteString(os.Stdout, msg)
	
	//Exiting with an OK
	os.Exit(0)
}