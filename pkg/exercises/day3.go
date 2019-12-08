package exercises

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Day3a(r io.Reader) int {
	var dist int
	p := paths(r)
	origin := Point{0, 0}
	for _, a := range p[0].Segments {
		for _, b := range p[1].Segments {
			if a.Intersects(b) {
				d := origin.Manhattan(a.Where(b))
				if dist == 0 {
					dist = d
				}
				dist = Min(dist, d)
			}
		}
	}
	return dist
}

func Day3b(r io.Reader) int {
	var dist int
	p := paths(r)
	for _, a := range p[0].Segments {
		for _, b := range p[1].Segments {
			if a.Intersects(b) {
				point := a.Where(b)
				d := p[0].Boston(point) + p[1].Boston(point)
				if dist == 0 {
					dist = d
				}
				dist = Min(dist, d)
			}
		}
	}
	return dist
}

type Point struct {
	X int
	Y int
}

func (p Point) Manhattan(q Point) int {
	return Abs(p.X-q.X) + Abs(p.Y-q.Y)
}

type Path struct {
	Segments []Line
}

func (p Path) Boston(q Point) int {
	var length int
	for _, s := range p.Segments {
		if s.Contains(q) {
			length += Line{s.P, q}.Length()
			break
		}
		length += s.Length()
	}
	return length
}

func NewPath(t string) Path {
	var p, pos = make([]Point, 0, 20), Point{}
	p = append(p, pos)
	for _, m := range strings.Split(t, ",") {
		magnitude, _ := strconv.Atoi(m[1:])
		switch m[0] {
		case 'R':
			pos = Point{pos.X + magnitude, pos.Y}
		case 'U':
			pos = Point{pos.X, pos.Y + magnitude}
		case 'L':
			pos = Point{pos.X - magnitude, pos.Y}
		case 'D':
			pos = Point{pos.X, pos.Y - magnitude}
		}
		p = append(p, pos)
	}
	return Path{segments(p)}
}

type Line struct {
	P Point
	Q Point
}

func (l Line) Intersects(s Line) bool {
	lminx := Min(l.P.X, l.Q.X)
	lmaxx := Max(l.P.X, l.Q.X)
	lminy := Min(l.P.Y, l.Q.Y)
	lmaxy := Max(l.P.Y, l.Q.Y)
	sminx := Min(s.P.X, s.Q.X)
	smaxx := Max(s.P.X, s.Q.X)
	sminy := Min(s.P.Y, s.Q.Y)
	smaxy := Max(s.P.Y, s.Q.Y)

	return (lminx <= smaxx && lmaxx >= sminx) &&
		(lminy <= smaxy && lmaxy >= sminy)
}

func (l Line) Where(s Line) Point {
	var x, y int
	if l.P.X-l.Q.X == 0 {
		x = l.P.X
	} else {
		x = s.P.X
	}
	if l.P.Y-l.Q.Y == 0 {
		y = l.P.Y
	} else {
		y = s.P.Y
	}
	return Point{x, y}
}

func (l Line) Length() int {
	return l.P.Manhattan(l.Q)
}

func (l Line) Contains(p Point) bool {
	minx := Min(l.P.X, l.Q.X)
	maxx := Max(l.P.X, l.Q.X)
	miny := Min(l.P.Y, l.Q.Y)
	maxy := Max(l.P.Y, l.Q.Y)
	return (p.X >= minx && p.X <= maxx) && (p.Y >= miny && p.Y <= maxy)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func paths(r io.Reader) []Path {
	var p []Path
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		p = append(p, NewPath(scanner.Text()))
	}
	return p
}

func segments(path []Point) []Line {
	var line []Line
	for i := 1; i < len(path); i++ {
		line = append(line, Line{path[i-1], path[i]})
	}
	return line
}
