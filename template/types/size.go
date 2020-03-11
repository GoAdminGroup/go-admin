package types

import "strconv"

// Extra small screen / phone
// xs: 0

// Small screen / phone
// sm: 576px

// Medium screen / tablet
// md: 768px

// Large screen / desktop
// lg: 992px

// Extra large screen / wide desktop
// xl: 1200px

type S map[string]string

func Size(sm, md, lg int) S {
	var s = make(S)
	if sm > 0 && sm < 13 {
		s["sm"] = strconv.Itoa(sm)
	}
	if md > 0 && md < 13 {
		s["md"] = strconv.Itoa(md)
	}
	if lg > 0 && lg < 13 {
		s["lg"] = strconv.Itoa(lg)
	}
	return s
}

func (s S) LG(lg int) S {
	if lg > 0 && lg < 13 {
		s["lg"] = strconv.Itoa(lg)
	}
	return s
}

func (s S) XS(xs int) S {
	if xs > 0 && xs < 13 {
		s["xs"] = strconv.Itoa(xs)
	}
	return s
}

func (s S) XL(xl int) S {
	if xl > 0 && xl < 13 {
		s["xl"] = strconv.Itoa(xl)
	}
	return s
}

func (s S) SM(sm int) S {
	if sm > 0 && sm < 13 {
		s["sm"] = strconv.Itoa(sm)
	}
	return s
}

func (s S) MD(md int) S {
	if md > 0 && md < 13 {
		s["md"] = strconv.Itoa(md)
	}
	return s
}

func SizeXS(xs int) S {
	var s = make(S)
	if xs > 0 && xs < 13 {
		s["xs"] = strconv.Itoa(xs)
	}
	return s
}

func SizeXL(xl int) S {
	var s = make(S)
	if xl > 0 && xl < 13 {
		s["xl"] = strconv.Itoa(xl)
	}
	return s
}

func SizeSM(sm int) S {
	var s = make(S)
	if sm > 0 && sm < 13 {
		s["sm"] = strconv.Itoa(sm)
	}
	return s
}

func SizeMD(md int) S {
	var s = make(S)
	if md > 0 && md < 13 {
		s["md"] = strconv.Itoa(md)
	}
	return s
}

func SizeLG(lg int) S {
	var s = make(S)
	if lg > 0 && lg < 13 {
		s["lg"] = strconv.Itoa(lg)
	}
	return s
}
