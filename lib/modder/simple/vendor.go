package simple

func (m *Modder) Vendor() error {
	return m.Load(".")
}

