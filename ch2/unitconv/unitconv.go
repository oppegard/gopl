package unitconv

import "fmt"

type Meter float64
type Foot float64
type Kilogram float64
type Pound float64

func (m Meter) String() string    { return fmt.Sprintf("%.2f m", m) }
func (f Foot) String() string     { return fmt.Sprintf("%.2f ft", f) }
func (k Kilogram) String() string { return fmt.Sprintf("%.2f kg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%.2f lb", p) }

func MToF(m Meter) Foot { return Foot(m * 3.28) }
func FToM(f Foot) Meter { return Meter(f / 3.28) }

func KToP(k Kilogram) Pound { return Pound(k * 2.2046226218) }
func PToK(p Pound) Kilogram { return Kilogram(p * (1 / 2.2046226218)) }
