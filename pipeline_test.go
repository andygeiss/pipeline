package pipeline_test

import (
	"errors"
	"github.com/andygeiss/assert"
	"github.com/andygeiss/pipeline"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestPipeline_Data_Should_Return_A_Nil_Slice(t *testing.T) {
	p := new(pipeline.Pipeline)
	assert.That("data should be nil", t, p.Data(), nil)
}

func TestPipeline_Error_Should_Return_Nil(t *testing.T) {
	p := new(pipeline.Pipeline)
	assert.That("error should be nil", t, p.Error(), nil)
}

func TestPipeline_Evaluate_Should_Not_Return_An_Error(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Evaluate("nothing ...", func(in interface{}, data proto.Message) error {
		return nil
	})
	assert.That("error should be nil", t, p.Error(), nil)
}

func TestPipeline_Evaluate_Should_Return_An_Error_If_External_Function_Fails(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Evaluate("nothing ...", func(in interface{}, data proto.Message) error {
		return errors.New("failed")
	})
	assert.That("error should be returned", t, p.Error(), "failed")
}

func TestPipeline_Gather_Should_Not_Return_An_Error(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Gather("https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data", "testdata/iris.csv")
	assert.That("error should be nil", t, p.Error(), nil)
}

func TestPipeline_Organize_Should_Not_Return_An_Error(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Organize("testdata/iris.csv", func(records [][]string) (out proto.Message, err error) {
		return nil, nil
	})
	assert.That("error should be nil", t, p.Error(), nil)
}

func TestPipeline_Organize_Should_Return_An_Error_If_External_Function_Fails(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Organize("testdata/iris.csv", func(records [][]string) (out proto.Message, err error) {
		return nil, errors.New("failed")
	})
	assert.That("error should be returned", t, p.Error(), "failed")
}

func TestPipeline_Save_Should_Not_Return_An_Error(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Save("testdata/save.pb")
	assert.That("error should be nil", t, p.Error(), nil)
}

func TestPipeline_Save_Should_Return_An_Error_If_Directory_Not_Exists(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Save("not_exists/save.pb")
	assert.That("error should be returned", t, p.Error() == nil, false)
}

func TestPipeline_Load_Should_Not_Return_An_Error(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Load("testdata/save.pb", func(raw []byte) (out proto.Message, err error) {
		return nil, nil
	})
	assert.That("error should be nil", t, p.Error(), nil)
}

func TestPipeline_Load_Should_Return_An_Error_If_Directory_Not_Exists(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Load("not_exists/save.pb", func(raw []byte) (out proto.Message, err error) {
		return nil, errors.New("failed")
	})
	assert.That("error should be returned", t, p.Error() == nil, false)
}

func TestPipeline_Transform_Should_Not_Return_An_Error(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Transform(func(in proto.Message) (out proto.Message, err error) {
		return nil, nil
	})
	assert.That("error should be nil", t, p.Error(), nil)
}

func TestPipeline_Transform_Should_Return_An_Error_If_External_Function_Fails(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Transform(func(in proto.Message) (out proto.Message, err error) {
		return nil, errors.New("failed")
	})
	assert.That("error should be returned", t, p.Error(), "failed")
}

func TestPipeline_Validate_Should_Not_Return_An_Error(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Validate("nothing ...", func(in interface{}, data proto.Message) error {
		return nil
	})
	assert.That("error should be nil", t, p.Error(), nil)
}

func TestPipeline_Validate_Should_Return_An_Error_If_External_Function_Fails(t *testing.T) {
	p := new(pipeline.Pipeline)
	p.Validate("nothing ...", func(in interface{}, data proto.Message) error {
		return errors.New("failed")
	})
	assert.That("error should be returned", t, p.Error(), "failed")
}
