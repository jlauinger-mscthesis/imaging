package imaging

import (
	"image"
	"image/color"
	"testing"
)

func TestFlipH(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"FlipH 2x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 2 * 4,
				Pix: []uint8{
					0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33,
					0x00, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0x00,
				},
			},
		},
	}
	for _, d := range td {
		got := FlipH(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func BenchmarkFlipH(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FlipH(testdataBranchesJPG)
	}
}

func TestFlipV(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"FlipV 2x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := FlipV(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func BenchmarkFlipV(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FlipV(testdataBranchesJPG)
	}
}

func TestTranspose(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"Transpose 2x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00,
					0xcc, 0xdd, 0xee, 0xff, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := Transpose(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func BenchmarkTranspose(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Transpose(testdataBranchesJPG)
	}
}

func TestTransverse(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"Transverse 2x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xcc, 0xdd, 0xee, 0xff,
					0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33,
				},
			},
		},
	}
	for _, d := range td {
		got := Transverse(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func BenchmarkTransverse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Transverse(testdataBranchesJPG)
	}
}

func TestRotate90(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"Rotate90 2x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xcc, 0xdd, 0xee, 0xff, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff,
					0x00, 0x11, 0x22, 0x33, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00,
				},
			},
		},
	}
	for _, d := range td {
		got := Rotate90(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func BenchmarkRotate90(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Rotate90(testdataBranchesJPG)
	}
}

func TestRotate180(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"Rotate180 2x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 3),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0x00,
					0x00, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
					0xcc, 0xdd, 0xee, 0xff, 0x00, 0x11, 0x22, 0x33,
				},
			},
		},
	}
	for _, d := range td {
		got := Rotate180(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func BenchmarkRotate180(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Rotate180(testdataBranchesJPG)
	}
}

func TestRotate270(t *testing.T) {
	td := []struct {
		desc string
		src  image.Image
		want *image.NRGBA
	}{
		{
			"Rotate270 2x3",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x11, 0x22, 0x33, 0xcc, 0xdd, 0xee, 0xff,
					0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00,
					0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff,
				},
			},
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33,
					0x00, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xcc, 0xdd, 0xee, 0xff,
				},
			},
		},
	}
	for _, d := range td {
		got := Rotate270(d.src)
		want := d.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", d.desc, got)
		}
	}
}

func BenchmarkRotate270(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Rotate270(testdataBranchesJPG)
	}
}

func TestRotate(t *testing.T) {
	testCases := []struct {
		desc  string
		src   image.Image
		angle float64
		bg    color.Color
		want  *image.NRGBA
	}{
		{
			"Rotate 0",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff,
					0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			0,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff,
					0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			"Rotate 90",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff,
					0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			90,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
		},
		{
			"Rotate 180",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff,
					0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			180,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			"Rotate 45",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff,
					0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			45,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 6, 6),
				Stride: 6 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x61, 0x00, 0x00, 0xff, 0x58, 0x08, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x61, 0x00, 0x00, 0xff, 0xe9, 0x16, 0x00, 0xff, 0x35, 0xca, 0x00, 0xff, 0x00, 0x30, 0x30, 0xff, 0x00, 0x00, 0x00, 0xff,
					0x61, 0x00, 0x00, 0xff, 0xe9, 0x16, 0x00, 0xff, 0x35, 0xca, 0x00, 0xff, 0x00, 0x80, 0x80, 0xff, 0x35, 0x35, 0xff, 0xff, 0x58, 0x58, 0x61, 0xff,
					0x58, 0x08, 0x00, 0xff, 0x35, 0xca, 0x00, 0xff, 0x00, 0x80, 0x80, 0xff, 0x35, 0x35, 0xff, 0xff, 0xe9, 0xe9, 0xff, 0xff, 0x61, 0x61, 0x61, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x00, 0x30, 0x30, 0xff, 0x35, 0x35, 0xff, 0xff, 0xe9, 0xe9, 0xff, 0xff, 0x61, 0x61, 0x61, 0xff, 0x00, 0x00, 0x00, 0xff,
					0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0x58, 0x58, 0x61, 0xff, 0x61, 0x61, 0x61, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			"Rotate 0x0",
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 0, 0),
				Stride: 0,
				Pix:    []uint8{},
			},
			123,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 0, 0),
				Stride: 0,
				Pix:    []uint8{},
			},
		},
		{
			"Rotate -90",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 0, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff,
				},
			},
			-90,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0xff, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			"Rotate -360*10",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 0, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0x00, 0xff, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff,
				},
			},
			-360 * 10,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 1 * 4,
				Pix: []uint8{
					0x00, 0xff, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			"Rotate -360*10 + 90",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 0, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff,
				},
			},
			-360*10 + 90,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff,
				},
			},
		},
		{
			"Rotate -360*10 + 180",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 0, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff,
				},
			},
			-360*10 + 180,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 1 * 4,
				Pix: []uint8{
					0x00, 0xff, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			"Rotate -360*10 + 270",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 0, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff,
				},
			},
			-360*10 + 270,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0xff, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			"Rotate 360*10",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 0, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0x00, 0xff, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff,
				},
			},
			360 * 10,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 1 * 4,
				Pix: []uint8{
					0x00, 0xff, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			"Rotate 360*10 + 90",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 0, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff,
				},
			},
			360*10 + 90,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff,
				},
			},
		},
		{
			"Rotate 360*10 + 180",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 0, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff,
				},
			},
			360*10 + 180,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 1 * 4,
				Pix: []uint8{
					0x00, 0xff, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff,
				},
			},
		},
		{
			"Rotate 360*10 + 270",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 0, 1),
				Stride: 1 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff,
				},
			},
			360*10 + 270,
			color.Black,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0xff, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
	}
	for _, test := range testCases {
		got := Rotate(test.src, test.angle, test.bg)
		want := test.want
		if !compareNRGBA(got, want, 0) {
			t.Errorf("test [%s] failed: %#v", test.desc, got)
		}
	}
}

func BenchmarkRotate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Rotate(testdataBranchesJPG, 30, color.Transparent)
	}
}
