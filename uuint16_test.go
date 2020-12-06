package uuint16

import(
  "sync"
  "testing"
)

func TestRent(t *testing.T) {
  // rents all values
  var i int
  for i = 0; i <= 65535; i++ {
    val, err := Rent()
    if err != nil { t.Fail() }
    if val != uint16(i) {
      t.Logf("Failed! %v is not %v", i, val)
      t.Fail()
    }
  }

  if len(p.Available) != 0 {
    t.Logf("Failed! did not rent all values %v", len(p.Available))
    t.Fail()
  }

  if len(p.InUse) != 65536 {
    t.Logf("Failed! all values are not in use %v", len(p.InUse))
    t.Fail()
  }

  _, err := Rent()
  if err != ErrorNoneAvailable {
    t.Logf("%v is not %v", err.Error(), ErrorNoneAvailable.Error())
    t.Fail()
  }
}

func TestReturn(t *testing.T) {
  // returns all values
  var i int
  for i = 0; i <= 65535; i++ {
    Return(uint16(i))
  }

  if len(p.InUse) != 0 {
    t.Logf("Failed! did not return all values %v", len(p.InUse))
    t.Fail()
  }

  if len(p.Available) != 65536 {
    t.Logf("Failed! all values are not available %v", len(p.Available))
    t.Fail()
  }
}

func rentReturnTen(t *testing.T, wg *sync.WaitGroup) {
  wg.Add(1)
  defer wg.Done()
  vals := []uint16{}
  var i int

  for i = 0; i < 10; i++ {
    val, err := Rent()
    if err != nil { t.Fail() }
    vals = append(vals, val)
  }

  for i = 0; i < 10; i++ {
    Return(vals[i])
  }

  vals = nil
}

func TestRentReturnLock(t *testing.T) {
  wg := new(sync.WaitGroup)
  go rentReturnTen(t, wg)
  go rentReturnTen(t, wg)
  go rentReturnTen(t, wg)
  go rentReturnTen(t, wg)
  go rentReturnTen(t, wg)
  go rentReturnTen(t, wg)
  go rentReturnTen(t, wg)
  go rentReturnTen(t, wg)
  go rentReturnTen(t, wg)
  go rentReturnTen(t, wg)
  wg.Wait()
}

func TestDoubleReturn(t *testing.T) {
  val, err := Rent()
  if err != nil { t.Fail() }
  Return(val)
  before := len(p.Available)
  Return(val)
  if before != len(p.Available) {
    t.Logf("value was returned twice")
    t.Fail()
  }
}
