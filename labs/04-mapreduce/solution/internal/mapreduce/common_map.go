package mapreduce

import (
	"encoding/json"
	"hash/fnv"
	"io/ioutil"
	"sync"
)
var hashes = make(map[string]int)
var lock = sync.RWMutex{}
// doMap does the job of a map worker: it reads one of the input files
// (inFile), calls the user-defined map function (mapF) for that file's
// contents, and partitions the output into nReduce intermediate files.
//The master considers each input file one map task, and makes a call to doMap() at least once for each task.
//Each call to doMap() reads the appropriate file, calls the map function on that file's contents, and produces nReduce files for each map file.
//Thus, after all map tasks are done, the total number of files will be the product of the number of files given to map (nIn) and nReduce.
func doMap(
	jobName string, // the name of the MapReduce job
	mapTaskNumber int, // which map task this is
	inFile string,
	nReduce int, // the number of reduce task that will be run ("R" in the paper)
	mapF func(file string, contents string) []KeyValue,
) {
	dat, err := ioutil.ReadFile(inFile)
	if err != nil {
		debug("file open fail:%s", inFile)
	} else {
		kvs := mapF(inFile, string(dat))
		partitions := make([][]KeyValue, nReduce)

		for _ , kv:= range kvs {
			r := int(ihash(kv.Key)) % nReduce
			partitions[r] = append(partitions[r], kv)
		}

		for i := range partitions {
			j, _ := json.Marshal(partitions[i])
			f := reduceName(jobName, mapTaskNumber, i)
			ioutil.WriteFile(f, j, 0644)
		}
	}
}

func ihash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
