package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClient("unix:///var/run/docker.sock", "", nil, nil) 
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if strings.HasSuffix(container.Image, "cpu-usage") {
			continue
		}
		inspect, err := cli.ContainerInspect(context.Background(), container.ID)
		if err != nil {
			panic(err)
		}
		started, err := time.Parse(time.RFC3339Nano, inspect.State.StartedAt)
		if err != nil {
			panic(err)
		}
		resp, err := cli.ContainerStats(context.Background(), container.ID, false)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var stats types.StatsJSON
		err = json.Unmarshal(content, &stats)
		if err != nil {
			panic(err)
		}
		usage := time.Duration(stats.Stats.CPUStats.CPUUsage.TotalUsage)*time.Nanosecond
		read := stats.Stats.Read
		elapsed := read.Sub(started)
		fmt.Printf("%s - CPU: %s Elapsed: %s\n", container.Names[0][1:], usage, elapsed)
	}
}
