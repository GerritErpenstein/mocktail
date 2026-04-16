package a

import (
	"bytes"
	"context"
	"iter"
	"time"

	"golang.org/x/mod/module"
)

type Pineapple interface {
	Hello(bar Water) string
	World() string
	Goo() (string, int, Water)
	Coo(context.Context, string, Water) Water
	Noo(context.Context) string
}

type Coconut interface {
	Boo(src *bytes.Buffer) time.Duration
	Doo(src time.Duration) time.Duration
	Foo(st Strawberry) string
	Goo(st string) Strawberry
	Hoo(string, int, Water)
	Joo(string, int, Water) (string, int)
	Koo(src string) (dst string)
	Loo(st string, values ...int) string
	Too(src string) time.Duration
	Voo(src *module.Version) time.Duration
	Yoo(st string) interface{}
	Zoo(st interface{}) string
	Moo(fn func(st, stban Strawberry) Pineapple) string
	Noo(ar [][2]string) string
	Poo(str struct{ name string }) string
}

type Water struct{}

type Strawberry interface {
	Bar(string) int
}

type Orange interface {
	Juice() <-chan struct{}
}

// Lime is a type alias for time.Duration.
type Lime = time.Duration

type Mango interface {
	Squeeze(dur Lime) string
}

type Guava interface {
	Iterate() iter.Seq2[string, error]
	List() iter.Seq[Water]
}
