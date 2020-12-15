package messages

// SeismographReading encapsultes the data recorded by a sisemograph.
// This includes the scale value (from 1 and -1), the location,
// and the date recorded.
type SeismographReading struct {
	ID       string `json:"id"`
	Scale    int    `json:"scale"`
	Location string `json:"location"`
	Time     string `json:"time"`
}
