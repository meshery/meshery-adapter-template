package adaptor

// Operation holds the informormation for list of operations
type Operation struct {
	Type       string            `json:"type,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

// Operations hold a map of Operation objects
type Operations map[string]*Operation

// ApplyOperation applies the operation on <adaptor-name>
func (h *handler) ApplyOperation() error {
	// Implementation for this method goes here.
	return nil
}

// ListOperations lists the operations available
func (h *handler) ListOperations() (Operations, error) {
	// Needs implementation
	operations := make(Operations)
	err := h.config.Operations(&operations)
	if err != nil {
		return nil, err
	}
	return operations, nil
}
