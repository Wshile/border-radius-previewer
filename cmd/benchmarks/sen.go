
// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/sen"
)

func senParse(b *testing.B) {
	j, _ := ioutil.ReadFile(filename)
	var sample []byte
	if data, err := (&oj.Parser{}).Parse(j); err == nil {
		sample = []byte(sen.String(data, &sen.Options{Indent: 2}))
	} else {
		panic(err)
	}
	b.ResetTimer()
	p := &sen.Parser{}
	for n := 0; n < b.N; n++ {
		if _, err := p.Parse(sample); err != nil {
			panic(err)
		}
	}
}

func senParseReuse(b *testing.B) {
	j, _ := ioutil.ReadFile(filename)
	var sample []byte
	if data, err := (&oj.Parser{}).Parse(j); err == nil {
		sample = []byte(sen.String(data, &sen.Options{Indent: 2}))
	} else {
		panic(err)
	}
	b.ResetTimer()
	p := &sen.Parser{Reuse: true}
	for n := 0; n < b.N; n++ {
		if _, err := p.Parse(sample); err != nil {
			panic(err)
		}
	}
}

func senTokenize(b *testing.B) {
	sample, _ := ioutil.ReadFile(filename)
	b.ResetTimer()
	h := oj.ZeroHandler{}
	t := sen.Tokenizer{}
	for n := 0; n < b.N; n++ {
		if err := t.Parse(sample, &h); err != nil {
			panic(err)
		}
	}
}

func senTokenizeLoad(b *testing.B) {
	t := sen.Tokenizer{}
	h := oj.ZeroHandler{}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if err := t.Load(f, &h); err != nil {
			panic(err)
		}
	}
}

func senParseReader(b *testing.B) {
	var p sen.Parser
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if _, err = p.ParseReader(f); err != nil {
			panic(err)
		}
	}
}

func senParseReaderReuse(b *testing.B) {
	p := sen.Parser{Reuse: true}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if _, err = p.ParseReader(f); err != nil {
			panic(err)
		}
	}
}

func senUnmarshalPatient(b *testing.B) {
	sample, _ := ioutil.ReadFile(patFilename)
	p := sen.Parser{Reuse: true}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {