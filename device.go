package uaparser

var iPad *itemSpec = &itemSpec{
    name: "iPad",
    mustContains: []string{"iPad"},
}

var iPhone *itemSpec = &itemSpec{
    name: "iPhone",
    mustContains: []string{"iPhone"},
}

var iPod *itemSpec = &itemSpec{
    name: "iPod",
    mustContains: []string{"iPod"},
}

var mac *itemSpec = &itemSpec{
    name: "Macintosh",
    mustContains: []string{"Macintosh"},
}

var pc *itemSpec = &itemSpec{
    name: "PC",
    mustContains: []string{"Windows"},
    mustNotContains: []string{"Windows Phone"},
}

var _DEVICES []*itemSpec = []*itemSpec {
    iPad,
    iPhone,
    iPod,
    mac,
    pc,
}
