package p

func noParams() {}

func intParams(a int, b int) {
	if a == 3 {
		// do nothing
	}
}

func pointerParams(a *int, b *int) {
	if a == nil {
		// do nothing
	}
}

func anyParamNoNilComparison(a any) {}

func anyParamNilComparison(a any) {
	if a == nil { // want "\"a\" is an interface and must not be compared to nil. Check your interface usage or compare against reflect.Valueof"
		// do nothing
	}
}

func anyParamNilComparisonReversed(a any) {
	if nil == a { // want "\"a\" is an interface and must not be compared to nil. Check your interface usage or compare against reflect.Valueof"
		// do nothing
	}
}
