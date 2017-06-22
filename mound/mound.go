package mound

import (
	"errors"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

// Mound is the main structure of
type Mound struct {
	Grid     Grid
	Turmites []*Turmite
	Default  Move
}

// CreateMound builds the main structure that contains the grid and a slice of turmites
func CreateMound(blockSize int, x int, y int, turmites []*Turmite, def Move) (Mound, error) {
	var m Mound
	for _, t := range turmites {
		if x*y < t.Location {
			return m, errors.New("turmite out of bounds")
		}
	}
	g := createGrid(blockSize, x, y)
	m = Mound{
		Grid:     g,
		Turmites: turmites,
		Default:  def,
	}
	return m, nil
}

// Next mutates the mound and moves it forward one tick
func (m *Mound) Next() {
	for _, t := range m.Turmites {
		move, ok := t.findMove(m.Grid.currentColor(*t))
		if ok == false {
			move = m.Default
		}
		m.Grid.updateColor(t.Location, move.C)
		d, l := t.move(move.T, m.Grid)
		t.Location = l
		t.Direction = d
	}
}

// numberOfDigits calculates the number of digits in a given int using a lookup table
// less expensive than the alternative of math.Log10 calls
func numberOfDigits(x int) (int, error) {
	switch {
	case x < 10:
		return 1, nil
	case x < 100:
		return 2, nil
	case x < 1000:
		return 3, nil
	case x < 10000:
		return 4, nil
	case x < 100000:
		return 5, nil
	case x < 1000000:
		return 6, nil
	default:
		return -1, errors.New("Integer size not supported")
	}
}

func zeroString(zeros int) string {
	var zeroString string
	for x := 0; x < zeros; x++ {
		zeroString = zeroString + "0"
	}
	return zeroString
}

// BatchImages creates a temporary folder and creates a sequence of png images
func (m *Mound) BatchImages(numberOfImages int, squareSize int) (string, error) {
	// Create directory to store images in before encoding
	directory := "batch" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	_, err := numberOfDigits(numberOfImages)
	if err != nil {
		return "", err
	}
	err = os.Mkdir(directory, os.FileMode(int(0777)))
	if err != nil {
		return "", err
	}
	e := png.Encoder{
		CompressionLevel: -3,
	}

	for x := 1; x <= numberOfImages; x++ {
		digits, err := numberOfDigits(x)
		if err != nil {
			return "", err
		}
		zeros := 6 - digits
		filepath := filepath.Join(directory, "color"+zeroString(zeros)+strconv.Itoa(x)+".png")
		file, err := os.Create(filepath)
		if err != nil {
			return "", err
		}
		img := m.Grid.GridToImage(squareSize)
		err = e.Encode(file, img)
		if err != nil {
			return "", err
		}
		m.Next()
	}
	return directory, nil
}

// ConvertPngToMp4 take a directory that contains a list of pngs in the format color%06d.png and converts
// the images into an mp4 using ffmpeg
func ConvertPngToMp4(directory string, output string, fps int) error {
	file, err := os.Create("batchimages" + directory + ".sh")
	if err != nil {
		return err
	}
	defer file.Close()
	// Write shell script
	script := "ffmpeg -r " + strconv.Itoa(fps) + " -f image2 -s 1920x1080 -i " + directory + "/color%06d.png -vcodec libx264 -crf 25 -pix_fmt yuv420p " + output + ".mp4"
	file.WriteString("#!bin/bash \n")
	file.WriteString(script + "\n")
	file.WriteString("rm -rf " + directory)
	file.WriteString("rm " + "batchimages" + directory + ".sh")
	err = file.Sync()
	if err != nil {
		return err
	}
	cmd := exec.Command("/bin/sh", "batchimages.sh")
	cmd.Run()
	return nil
}
