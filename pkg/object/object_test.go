package object

import (
	"GoTracing/pkg/material"
	"testing"
)

func TestSphere(t *testing.T) {
	defaultMaterial, _ := material.NewDefaultMaterial(nil)
	sphereArgs := make(map[string]string)
	sphereArgs["radius"] = "50"
	sphereArgs["center"] = "-1,-1,-1"

	_, err := NewSphere(defaultMaterial, sphereArgs)
	if err != nil {
		t.Log("Create sphere failed", err)
		t.Fail()
	}

	sphereArgs["radius"] = "-50"
	_, err = NewSphere(defaultMaterial, sphereArgs)
	if err == nil {
		t.Log("When the radius is not a positive value, create should fail", err)
		t.Fail()
	}
}

func TestRect(t *testing.T) {
	defaultMaterial, _ := material.NewDefaultMaterial(nil)
	rectArgs := make(map[string]string)
	rectArgs["width"] = "50"
	rectArgs["length"] = "50"
	rectArgs["localX"] = "1,0,0"
	rectArgs["localY"] = "0,1,0"
	rectArgs["localZ"] = "0,0,1"
	rectArgs["position"] = "0,0,0"
	_, err := NewRect(defaultMaterial, rectArgs)
	if err != nil {
		t.Log("Create sphere failed", err)
		t.Fail()
	}

	rectArgs["width"] = "50"
	rectArgs["length"] = "50"
	rectArgs["localX"] = "1,0,0"
	rectArgs["localY"] = "1,0,0"
	rectArgs["localZ"] = "1,0,0"
	_, err = NewRect(defaultMaterial, rectArgs)
	if err == nil {
		t.Log("When the local coordiate system is invalied, create should fail", err)
		t.Fail()
	}

	rectArgs["width"] = "-50"
	rectArgs["length"] = "50"
	rectArgs["localX"] = "1,0,0"
	rectArgs["localY"] = "0,1,0"
	rectArgs["localZ"] = "0,0,1"
	_, err = NewRect(defaultMaterial, rectArgs)
	if err == nil {
		t.Log("When the width is not a positive value, create should fail", err)
		t.Fail()
	}

	rectArgs["width"] = "50"
	rectArgs["length"] = "-50"
	rectArgs["localX"] = "1,0,0"
	rectArgs["localY"] = "0,1,0"
	rectArgs["localZ"] = "0,0,1"
	_, err = NewRect(defaultMaterial, rectArgs)
	if err == nil {
		t.Log("When the lenght is not a positive value, create should fail", err)
		t.Fail()
	}
}
