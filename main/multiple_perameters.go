package main

import (
	"sync"
)

var (
	cache = &sync.Pool{
		New: func() interface{} {
			return &option{sex: 0}
		},
	}
)

type inter interface {
	getName(string2 string) string
}

type Option func(*option)

func (*option) getName() string {
	return "w"
}

type option struct {
	sex    int
	age    int
	height int
	weight int
	hobby  string
}

func (o *option) reset() {
	o.sex = 0
	o.age = 0
	o.height = 0
	o.weight = 0
	o.hobby = ""
}

func getOption() *option {

	return cache.Get().(*option)
}

func releaseOption(opt *option) {
	opt.reset()
	cache.Put(opt)
}

// WithSex setup sex, 1=female 2=male
func WithSex(sex int) Option {
	return func(opt *option) {
		opt.sex = sex
	}
}

// WithAge setup age
func WithAge(age int) Option {
	return func(opt *option) {
		opt.age = age
	}
}

// WithHeight set up height
func WithHeight(height int) Option {
	return func(opt *option) {
		opt.height = height
	}
}

// WithWeight set up weight
func WithWeight(weight int) Option {
	return func(opt *option) {
		opt.weight = weight
	}
}

// WithHobby set up Hobby
func WithHobby(hobby string) Option {
	return func(opt *option) {
		opt.hobby = hobby
	}
}
