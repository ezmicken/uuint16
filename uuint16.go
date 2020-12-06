package uuint16

import "sync"
import "errors"

type pool struct {
  Current        uint16
  Available      []uint16
  InUse          []uint16
  Mtx            *sync.Mutex
}

var p pool = pool{
  0,
  []uint16{},
  []uint16{},
  new(sync.Mutex),
}

var ErrorNoneAvailable error = errors.New("No ids available from uuint16")

func Rent() (uint16, error) {
  p.Mtx.Lock()
  defer p.Mtx.Unlock()
  l := len(p.Available)

  if p.Current == 65535 && l == 0 {
    return 0, ErrorNoneAvailable
  }

  buffer := l
  if (buffer < 10 && p.Current != 65535) {
    for i := 0; i < 10; p.Current++ {
      p.Available = append(p.Available, p.Current)
      if p.Current == 65535 {
        break
      }
      i++
    }
  }

  val := p.Available[0]
  p.Available = append([]uint16{}, p.Available[1:]...)
  p.InUse = append(p.InUse, val)
  return val, nil
}

func Return(val uint16) {
  p.Mtx.Lock()
  defer p.Mtx.Unlock()

  p.Available = append(p.Available, val)

  b := p.InUse[:0]
  for _, x := range p.InUse {
    if x != val {
      b = append(b, x)
    }
  }
  p.InUse = b
}
