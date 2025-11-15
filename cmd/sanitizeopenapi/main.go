package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	var inPath string
	var outPath string
	flag.StringVar(&inPath, "in", "", "input OpenAPI spec path")
	flag.StringVar(&outPath, "out", "", "output spec path (defaults to input path)")
	flag.Parse()

	if inPath == "" {
		log.Fatal("-in is required")
	}
	if outPath == "" {
		outPath = inPath
	}

	data, err := os.ReadFile(inPath)
	if err != nil {
		log.Fatalf("read input spec: %v", err)
	}

	var doc interface{}
	if err := yaml.Unmarshal(data, &doc); err != nil {
		log.Fatalf("parse yaml: %v", err)
	}

	doc = convert(doc)
	stats := &sanitizeStats{}
	doc = sanitize(doc, stats)

	outData, err := yaml.Marshal(doc)
	if err != nil {
		log.Fatalf("encode yaml: %v", err)
	}

	if err := os.WriteFile(outPath, outData, 0o644); err != nil {
		log.Fatalf("write sanitized spec: %v", err)
	}

	log.Printf("Removed %d invalid boolean format declarations", stats.removed)
}

type sanitizeStats struct {
	removed int
}

func sanitize(value interface{}, stats *sanitizeStats) interface{} {
	switch v := value.(type) {
	case map[string]interface{}:
		if t, ok := v["type"].(string); ok && t == "boolean" {
			if f, ok := v["format"].(string); ok && strings.EqualFold(f, "boolean") {
				delete(v, "format")
				stats.removed++
			}
		}
		for key, val := range v {
			v[key] = sanitize(val, stats)
		}
		return v
	case []interface{}:
		for i, val := range v {
			v[i] = sanitize(val, stats)
		}
		return v
	default:
		return value
	}
}

func convert(value interface{}) interface{} {
	switch v := value.(type) {
	case map[string]interface{}:
		for key, val := range v {
			v[key] = convert(val)
		}
		return v
	case map[interface{}]interface{}:
		m := make(map[string]interface{}, len(v))
		for key, val := range v {
			m[fmt.Sprint(key)] = convert(val)
		}
		return m
	case []interface{}:
		for i, val := range v {
			v[i] = convert(val)
		}
		return v
	default:
		return value
	}
}
