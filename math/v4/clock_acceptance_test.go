package clock_test

import (
	"bytes"
	"encoding/xml"
	"github.com/LightBulbfromSpace/kld-go-learning/math/v4"
	"testing"
	"time"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		want Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}

	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clock.SVGWriter(&b, tc.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(tc.want, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v , in the SVG lines %+v", tc.want, svg.Line)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		want Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 70},
		},
		{
			simpleTime(0, 30, 0),
			Line{150, 150, 150, 230},
		},
	}

	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clock.SVGWriter(&b, tc.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(tc.want, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v , in the SVG lines %+v", tc.want, svg.Line)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		want Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 100},
		},
		{
			simpleTime(6, 0, 0),
			Line{150, 150, 150, 200},
		},
	}

	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clock.SVGWriter(&b, tc.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(tc.want, svg.Line) {
				t.Errorf("Expected to find the hour hand line %+v , in the SVG lines %+v", tc.want, svg.Line)
			}
		})
	}
}

func containsLine(line Line, lines []Line) bool {
	for _, l := range lines {
		if l == line {
			return true
		}
	}
	return false
}

func simpleTime(hours, minutes, secs int) time.Time {
	return time.Date(1970, time.September, 31, hours, minutes, secs, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:01:05")
}
