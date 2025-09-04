package mapreduce

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
)

// doReduce does the job of a reduce worker: it reads the intermediate
// key/value pairs (produced by the map phase) for this task, sorts the
// intermediate key/value pairs by key, calls the user-defined reduce function
// (reduceF) for each key, and writes the output to disk.

// As with doMap(), it does so either directly or through a worker.
// doReduce() collects corresponding files from each map result , and runs the reduce function on each collection.
// This process produces nReduce result files.
func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTaskNumber int, // which reduce task this is
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	kvs := make(map[string][]string)
	for m := 0; m < nMap; m++ {
		fileName := reduceName(jobName, m, reduceTaskNumber)
		dat, err := ioutil.ReadFile(fileName)
		if err != nil {
			debug("file open fail:%s", fileName)
		} else {
			var items []KeyValue
			json.Unmarshal(dat, &items)
			for _ , item := range items {
				k := item.Key
				v := item.Value
				kvs[k] = append(kvs[k], v)
			}
		}
	}

	// create the final output file
	mergeFileName := mergeName(jobName, reduceTaskNumber)
	file, err := os.Create(mergeFileName)
	if err != nil {
		debug("file open fail:%s", mergeFileName)
	}

	// sort
	var keys []string
	for k := range kvs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	enc := json.NewEncoder(file)
	for _, key := range keys {
		enc.Encode(KeyValue{key, reduceF(key, kvs[key])})
	}
	file.Close()
}
