package client

import (
	"encoding/binary"
	"log"
	"math"

	triton "github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/proto"
)

func GetDummyInput() *triton.ModelInferRequest {
	var inputs []*triton.ModelInferRequest_InferInputTensor
	var FeatureNames = []string{"product_type_name", "product_group_name", "colour_group_name", "department_name", "club_member_status", "price", "age"}
	for _, feature := range FeatureNames {
		dummy_input := &triton.ModelInferRequest_InferInputTensor{
			Name:     feature,
			Shape:    []int64{3, 1},
			Datatype: "FP32",
			Contents: &triton.InferTensorContents{
				Fp32Contents: []float32{0.0, 1.0, 2.0},
			},
		}
		inputs = append(inputs, dummy_input)
	}

	return &triton.ModelInferRequest{
		ModelName: "mlp-model",
		Inputs:    inputs,
	}
}

func PostProcess(modelResponse *triton.ModelInferResponse) any {
	rawOutputContents := modelResponse.GetRawOutputContents()
	outputDatatype := modelResponse.Outputs[0].GetDatatype()
	if outputDatatype == "FP32" {
		return PostprocessRawFloat32(rawOutputContents)
	} else {
		log.Printf("Got output datatype of: %v", outputDatatype)
		return nil
	}
}

func PostprocessRawFloat32(rawOutputContents [][]byte) []float32 {
	flattenedOutputs := rawOutputContents[0]
	outputs := []float32{}
	// Convert each []byte to a float64
	for i := 0; i < len(flattenedOutputs); i += 4 {
		// Convert 4-byte chunk to uint32
		bits := binary.LittleEndian.Uint32(flattenedOutputs[i : i+4])

		// Convert bits to float32
		floatValue := math.Float32frombits(bits)
		outputs = append(outputs, floatValue)
	}
	return outputs
}
