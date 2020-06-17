# pipeline

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/pipeline)](https://goreportcard.com/report/github.com/andygeiss/pipeline)

Build your own data pipeline to gather, organize and transform data by using protobuf as an intermediate format.

## Purpose

According to [Daniel Whitenack](https://github.com/dwhitena), surveys have shown that 90% or more of a data scientist's time is spent collecting, organizing and cleansing data.
Thus, I created an extensible data pipeline to implement the following process.

## Process

    +----------+    +------------+    +-------------+    +------------+    +------------+
    |  Gather  +---->  Organize  +---->  Transform  +---->  Evaluate  +---->  Validate  |
    +----------+    +------------+    +------^------+    +------------+    +------+-----+
                                             |                                    |
                                             +------------------------------------+

We should treat data and its format as immutable, and especially if you are using data from an external ressource or team.
So at first we need to **gather** the data and **organize** its raw format to an interim one, which is portable and can be used by many tools.
Thus, I decided to use Google's protobuf for portability.

Then you can **extract features** and **transform** your data into a new format.
Finally, you **evaluate** and **validate** your model and repeat the process until you reached the final **cleaned model**,
which can be used for **scoring**.

You can also **load** and **save** the data at any step in the pipeline.

## Installation

First install the [Protobuf Compiler](https://developers.google.com/protocol-buffers/docs/downloads) and the corresponding [Protobuf Go Plugin](https://developers.google.com/protocol-buffers/docs/gotutorial)
manually or use the following command:

    make

## Usage

```go
func main() {
    p := new(pipeline.Pipeline)

    p.Gather("https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data", "data/external/iris.csv").
    Organize("data/external/iris.csv", func(records [][]string) (out proto.Message, err error) {
        return nil, nil
    }).
    Save("data/interim/iris.pb").
    Transform(func(in proto.Message) (out proto.Message, err error) {
        return nil, nil
    }).
    Save("data/processed/iris.pb").
    Evaluate("evaluate something ...", func(in interface{}, data proto.Message) error {
        return nil
    }).
    Validate("validate something ...", func(in interface{}, data proto.Message) error {
	return nil
    })
    // Handle errors - If a certain step has failed, the next steps are ignored.
    if err := p.Error(); err != nil {
        log.Fatal(err)
    }
}
```
