package examples

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode"

	"github.com/heli0dus/functionalgo/src/collections/hashSet"
	"github.com/heli0dus/functionalgo/src/stream"
)

func split(s string) []string {
	return strings.Split(s, " ")
}

func appendRune(s string, r rune) string {
	return s + string(r)
}

func removePunctuationStream(w string) string {
	s := stream.AsStream([]rune(w))
	s = s.
		Filter(unicode.IsLetter).
		Reduce(appendRune, "")
	res, _ := stream.As[string](s)
	return res
}

func removePunctuationClassic(w string) string {
	var res []rune
	for _, r := range w {
		if unicode.IsLetter(r) {
			res = append(res, r)
		}
	}
	return string(res)
}

func SolveWithStreams(lines []string, w1 string, w2 string) int {
	s := stream.AsStream(lines)
	s = s.
		FlatMap(split).
		Fmap(removePunctuationStream).
		Fmap(strings.ToLower).
		DropWhile(func(s string) bool { return s != w1 }).
		TakeWhile(func(s string) bool { return s != w2 })
	res, err := stream.AsSlice[string](s)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		os.Exit(1)
	}
	return len(hashSet.FromSlice(res))
}

func SolveClassic(lines []string, w1 string, w2 string) int {
	var words []string
	for _, l := range lines {
		words = slices.Concat(words, split(l))
	}
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(removePunctuationClassic(words[i]))
	}
	from := 0
	for ; from < len(words); from++ {
		if words[from] == w1 {
			break
		}
	}
	to := from + 1
	for ; to < len(words); to++ {
		if words[to] == w2 {
			break
		}
	}
	return len(hashSet.FromSlice(words[from:to]))
}

var TextLines = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse sit amet vehicula arcu. Curabitur et arius",
	"sapien. Vestibulum quis orci quam. Sed vel erat sollicitudin, commodo urna at, rhoncus ipsum.",
	"Praesent sit amet leo vitae lacus efficitur venenatis vel in nibh. Etiam hendrerit, ex sit amet tempus",
	"interdum, dui magna vestibulum dui, id placerat magna leo sit amet felis. Aenean porttitor, lacus",
	"nec lobortis pulvinar, odio lacus viverra justo, at interdum sapien est vitae mi. Maecenas nec aliquam",
	"mauris, sed sodales dolor. Maecenas at metus elit. Aenean ullamcorper sem vel malesuada dignissim. Nulla",
	"lacus justo, ultrices vel sollicitudin sit amet, maximus sodales purus.",
	"Proin porta sit amet turpis quis dapibus. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac",
	"turpis egestas. Nam dictum dapibus turpis vitae mattis. Cras finibus convallis risus in maximus. Quisque suscipit",
	"leo sed lorem porta, vel volutpat lectus posuere. Integer eu malesuada elit, non volutpat ligula.",
	"Proin non nulla eget mi placerat porttitor at sit amet ligula. Sed at sem in mi gravida loborti",
	"Duis maximus lacus diam, a ultrices mi mollis in. Maecenas elementum ante et eros sagittis tempor.",
	"Aenean vitae eros magna. Maecenas ultrices molestie nunc, nec imperdiet augue molestie nec. Donec aliquam porta",
	"est eu scelerisque. Suspendisse sapien neque, fringilla vel vestibulum eu, tempus vitae sapien. Phasellus porta",
	"lorem justo, ut facilisis magna vestibulum sit amet. Nulla at leo ac velit maximus fermentum ut nec",
	"libero. Curabitur purus lorem, convallis et urna eget, interdum accumsan quam. Duis ultricies",
	"dolor in luctus maximus. Phasellus quis consectetur dolor, a commodo nisi. Cras vel mauris enim.",
	"Maecenas at metus scelerisque, sollicitudin neque vitae, rhoncus odio. Quisque id quam vel ipsum consequat rhoncu",
	"Curabitur egestas est vitae sem malesuada, a fermentum nulla scelerisque. Sed et faucibus sem. Fusce ut",
	"arcu maximus, egestas ante sed, ornare lacus. In porttitor commodo ex. Quisque ac mi pharetr",
	"luctus leo in, pretium nisi. Cras quis nisi auctor, ullamcorper tellus et, sodales tortor.",
	"Duis porta in orci ac egestas. Nullam sit amet malesuada felis. Duis dignissim congue cursus. Praesent",
	"vel posuere elit. Pellentesque et purus sit amet enim fermentum iaculis at vulputate dolor. Sed pulvinar feugiat",
	"tincidunt. Praesent mattis eget velit sit amet faucibus. Maecenas vel erat porta turpis suscipit placerat. Integer",
	"nulla nisi, egestas vel maximus id, sodales gravida eros. Sed dolor ipsum, egestas et libero",
	"blandit, pretium congue turpis. Cras urna felis, fringilla at ultricies id, vehicula eu risus.",
	"Praesent tristique arcu elit, eget rhoncus nisl eleifend vitae.",
}

var W1 = "proin"
var W2 = "curabitur"
