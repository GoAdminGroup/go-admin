// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package service

import (
	"log"
)

type Service interface {
	Name() string
}

type Generator func() (Service, error)

func Register(k string, gen Generator) {
	if _, ok := services[k]; ok {
		log.Panicf("service %s has been registered", k)
	}
	services[k] = gen
}

func GetServices() List {
	var (
		l   = make(List)
		err error
	)
	for k, gen := range services {
		if l[k], err = gen(); err != nil {
			log.Panicf("service %s initialize fail, error: %v", k, err)
		}
	}
	return l
}

var services = make(Generators)

type Generators map[string]Generator

type List map[string]Service

func (g List) Get(k string) Service {
	if v, ok := g[k]; ok {
		return v
	}
	log.Panicf("service %s not found", k)
	return nil
}

func (g List) GetOrNot(k string) (Service, bool) {
	v, ok := g[k]
	return v, ok
}

func (g List) Add(k string, service Service) {
	if _, ok := g[k]; ok {
		log.Panicf("service %s exist", k)
	}
	g[k] = service
}
