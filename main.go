package main

import (
  "fmt"
  "flag"
)

const Major string = ""
const Minor string = "m"

var Notes = map[int]string{
  0: "C",
  1: "C#",
  2: "D",
  3: "D#",
  4: "E",
  5: "F",
  6: "F#",
  7: "G",
  8: "G#",
  9: "A",
  10: "A#",
  11: "B",
}

type Chord struct {
  Mode string
  Tonic int
  Third int
  Fifth int
}

type Key struct {
  Pitch int
  Mode string
  I, II, III, IV, V, VI Chord
}

func NewChord(pitch int, mode string) Chord {
  chord := Chord{
    Mode: mode,
    Tonic: (pitch % 12),
    Third: (pitch + 3) % 12,
    Fifth: (pitch + 7) % 12,
  }

  if mode == Major {
    chord.Third = (pitch + 4) % 12
  }

  return chord
}

func (c Chord) Print() {
  fmt.Print(
    Notes[c.Tonic],
    c.Mode, 
    " (",
    Notes[c.Tonic],
    ", ",
    Notes[c.Third],
    ", ",
    Notes[c.Fifth],
    ")\n",
  )
}

func NewKey(pitch int, mode string) Key {
  key := Key{
    Pitch: pitch,
    Mode: mode,
  }

  if mode == Minor {
    pitch = pitch + 3 % 12
  }

  key.I = NewChord((pitch % 12), Major)
  key.II = NewChord((pitch + 2) % 12, Minor)
  key.III = NewChord((pitch + 4) % 12, Minor)
  key.IV = NewChord((pitch + 5) % 12, Major)
  key.V = NewChord((pitch + 7) % 12, Major)
  key.VI = NewChord((pitch + 9) % 12, Minor)

  return key
}

func (k Key) Print() {
  fmt.Print("Key: ", Notes[k.Pitch], k.Mode, "\n\n")

  fmt.Print("I: ")
  k.I.Print()
  fmt.Print("II: ")
  k.II.Print()
  fmt.Print("III: ")
  k.III.Print()
  fmt.Print("IV: ")
  k.IV.Print()
  fmt.Print("V: ")
  k.V.Print()
  fmt.Print("VI: ")
  k.VI.Print()
}

func main() {
  pitch := flag.Int("pitch", 0, "Pitch")
  mode := flag.String("mode", Major, "Mode")
  flag.Parse()
  NewKey(*pitch, *mode).Print()
}
