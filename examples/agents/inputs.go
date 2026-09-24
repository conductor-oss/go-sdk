package agents

// Tool argument types. tool.Func derives each tool's input schema from the
// struct: a field is advertised in snake_case, so AccountID is account_id.

type accountIn struct {
	AccountID string
}

type transferIn struct {
	FromAcct string
	ToAcct   string
	Amount   float64
}

type cityIn struct {
	City string
}

type symbolIn struct {
	Symbol string
}

type queryIn struct {
	Query string
}

type sqlIn struct {
	SQL string
}

type dataIn struct {
	Data string
}

type reportIn struct {
	Title string
	Body  string
}

type orderIn struct {
	OrderID string
}

type productIn struct {
	Product string
}

type serviceIn struct {
	ServiceName string
}

type deleteDataIn struct {
	ServiceName string
	DataType    string
}

type factorialIn struct {
	N int
}

type summaryIn struct {
	Text string
}

type checkSummaryIn struct {
	Text     string
	MinChars int
}

type customerIn struct {
	CustomerID string
}

type userIn struct {
	UserID string
}

type orderActionIn struct {
	OrderID string
	Action  string
}

type formatIn struct {
	Data map[string]any
}

type inventoryIn struct {
	ProductID string
	Warehouse string `json:"warehouse,omitempty"`
}
