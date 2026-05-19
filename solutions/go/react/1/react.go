package react

// Define reactor, cell and canceler types here.
// These types will implement the Reactor, Cell and Canceler interfaces, respectively.

type reactor struct {
	subscribers map[*cell][]*cell
}

type cell struct {
	reactor   *reactor
	value     int
	computeFn func() int
	callbacks map[*func(int)]struct{}
}

type canceler struct {
	cell     *cell
	callback *func(int)
}

func (c *canceler) Cancel() {
	delete(c.cell.callbacks, c.callback)
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) SetValue(value int) {
	if c.value == value {
		return
	}

	initialValues := c.reactor.snapshot()

	c.value = value

	c.reactor.propagate(c)

	for cell, oldVal := range initialValues {
		if cell.Value() != oldVal {
			cell.triggerCallbacks()
		}
	}
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	c.callbacks[&callback] = struct{}{}
	return &canceler{cell: c, callback: &callback}
}

func (c *cell) triggerCallbacks() {
	for f, _ := range c.callbacks {
		(*f)(c.value)
	}
}

func New() Reactor {
	return &reactor{
		subscribers: make(map[*cell][]*cell),
	}
}

func (r *reactor) CreateInput(initial int) InputCell {
	return &cell{
		reactor:   r,
		value:     initial,
		callbacks: make(map[*func(int)]struct{}),
	}
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	depCell := dep.(*cell)
	c := &cell{
		reactor:   r,
		callbacks: make(map[*func(int)]struct{}),
		computeFn: func() int {
			return compute(dep.Value())
		},
	}
	c.value = c.computeFn()
	r.subscribers[depCell] = append(r.subscribers[depCell], c)
	return c
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	dep1Cell := dep1.(*cell)
	dep2Cell := dep2.(*cell)
	c := &cell{
		reactor:   r,
		callbacks: make(map[*func(int)]struct{}),
		computeFn: func() int {
			return compute(dep1.Value(), dep2.Value())
		},
	}
	c.value = c.computeFn()
	r.subscribers[dep1Cell] = append(r.subscribers[dep1Cell], c)
	r.subscribers[dep2Cell] = append(r.subscribers[dep2Cell], c)
	return c
}

func (r *reactor) snapshot() map[*cell]int {
	initialValues := make(map[*cell]int)
	for _, cells := range r.subscribers {
		for _, c := range cells {
			if _, exists := initialValues[c]; !exists {
				initialValues[c] = c.value
			}
		}
	}
	return initialValues
}

func (r *reactor) propagate(c *cell) {
	for _, sub := range r.subscribers[c] {
		newValue := sub.computeFn()
		if sub.value != newValue {
			sub.value = newValue
			r.propagate(sub)
		}
	}
}
