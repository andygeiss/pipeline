package pipeline

import (
	"encoding/csv"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

// Pipeline provides an easy-to-use interface to work with data structures in context of data science and machine learning.
type Pipeline struct {
	data  proto.Message
	err   error
	mutex sync.Mutex
}

// Data returns the underlying raw data which is currently in the pipeline.
func (p *Pipeline) Data() proto.Message {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.data
}

// Error returns the current error state of the pipeline.
func (p *Pipeline) Error() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.err
}

// Evaluate tries to guess a specific output by using the data and a given input.
func (p *Pipeline) Evaluate(in interface{}, fn func(in interface{}, data proto.Message) (err error)) *Pipeline {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.err != nil {
		return p
	}
	// Use an external function for evaluation.
	p.err = fn(in, p.data)
	return p
}

// Gather retrieves raw pipeline from an external resource and saves that into a single file.
func (p *Pipeline) Gather(sourceUrl, targetFile string) *Pipeline {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.err != nil {
		return p
	}
	// Get the raw data from a given sourceURL.
	res, err := http.Get(sourceUrl)
	if err != nil {
		p.err = err
		return p
	}
	defer res.Body.Close()
	// Read the content body of the response received.
	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		p.err = err
		return p
	}
	// Write the content into a file.
	p.err = ioutil.WriteFile(targetFile, raw, 0644)
	return p
}

// Load reads a protobuf file and pass the raw bytes into a given function fn.
func (p *Pipeline) Load(protoFilename string, fn func(raw []byte) (out proto.Message, err error)) *Pipeline {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.err != nil {
		return p
	}
	// Read the protobuf file into memory.
	pb, err := ioutil.ReadFile(protoFilename)
	if err != nil {
		p.err = err
		return p
	}
	// Replace the current data with the content read previously.
	p.data, p.err = fn(pb)
	return p
}

// Organize reads a CSV file and provides serialization to a protobuf format via a given function fn.
func (p *Pipeline) Organize(csvFilename string, fn func(records [][]string) (out proto.Message, err error)) *Pipeline {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.err != nil {
		return p
	}
	// Open a raw CSV file.
	file, err := os.Open(csvFilename)
	if err != nil {
		p.err = err
		return p
	}
	defer file.Close()
	// Read the CSV content line by line.
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		p.err = err
		return p
	}
	// Use an external function to transform the records into a protobuf format.
	p.data, p.err = fn(records)
	return p
}

// Save writes a protobuf structure from memory into a specific file.
func (p *Pipeline) Save(protoFilename string) *Pipeline {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.err != nil {
		return p
	}
	// Serialize the internal data structure.
	out, err := proto.Marshal(p.data)
	if err != nil {
		p.err = err
		return p
	}
	// Write the serialized protobuf to the file.
	p.err = ioutil.WriteFile(protoFilename, out, 0644)
	return p
}

// Transform translates a given protobuf message into another protobuf message by using a specific function.
func (p *Pipeline) Transform(fn func(in proto.Message) (out proto.Message, err error)) *Pipeline {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.err != nil {
		return p
	}
	// Use an external function to transform the internal data structure to another protobuf format.
	p.data, p.err = fn(p.data)
	return p
}

// Validate calculates the quality of a model by using the data and a given input.
func (p *Pipeline) Validate(in interface{}, fn func(in interface{}, data proto.Message) (err error)) *Pipeline {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.err != nil {
		return p
	}
	// Use an external function for prediction.
	p.err = fn(in, p.data)
	return p
}
