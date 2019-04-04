package gotm1

// Configuration specifies configuration parameters as defined in the Tm1s.cfg
// file for an IBM Cognos TM1 server.
//
// For details about all of the parameters, refer to the comments in the
// Tm1s.cfg file and the Operations Guide.
type Configuration struct {
	ServerName                                        string
	AdminHost                                         string
	ProductVersion                                    string
	PortNumber                                        int64
	ClientMessagePortNumber                           int64
	HTTPPortNumber                                    int64
	IntegratedSecurityMode                            bool
	SecurityMode                                      string
	PrincipalName                                     string
	SecurityPackageName                               string
	ClientCAMURIs                                     []string
	WebCAMURI                                         string
	ClientPingCAMPassport                             int64
	ServerCAMURI                                      string
	AllowSeparateNandCRules                           bool
	DistributedOutputDir                              string
	DisableSandboxing                                 bool
	JobQueuing                                        bool
	ForceReevaluationOfFeedersForFedCellsOnDataChange bool
	DataBaseDirectory                                 string
	UnicodeUpperLowerCase                             bool
}

// Cube represents a single cube on a TM1 server.
type Cube struct {
	Name              string
	Rules             string
	DrillthroughRules string
	LastSchemaUpdate  string
	LastDataUpdate    string
	Attributes        map[string]string
}

// Dimension represents a single dimension on a TM1 server.
//
// A dimension is a broad grouping of descriptive data about a major aspect of
// a business, such as products, dates, or locations. Each dimension includes
// different levels of members in one or more hierarchies and an optional set
// of calculated members or special categories.
type Dimension struct {
	Name                   string
	UniqueName             string
	AllLeavesHierarchyName string
	Attributes             map[string]interface{}
}

// Chore executes one or more processes of TM1 at a user-defined frequency.
type Chore struct {
	Name          string
	StartTime     string
	DSTSensitive  bool
	Active        bool
	ExecutionMode string
	Frequency     string
	Attributes    map[string]string
}

// Process is a TurboIntegrator process that can be used to manipulate TM1 data and metadata.
type Process struct {
	Name              string
	HasSecurityAccess bool
	PrologProcedure   string
	MetadataProcedure string
	DataProcedure     string
	EpilogProcedure   string
	DataSource        map[string]interface{}
	Parameters        []map[string]string
	Variables         []interface{}
	Attributes        map[string]string
}

// Logger is one of many log options of a TM1 server.
type Logger struct {
	Name  string
	Level string
}

// Session represents a unique user session with the server.
type Session struct {
	ID      int64
	Context string
}

// Thread that can run queries concurrently on the TM1 server.
type Thread struct {
	ID          int64
	Type        string
	Name        string
	Context     string
	State       string
	Function    string
	ObjectType  string
	ObjectName  string
	RLocks      int64
	IXLocks     int64
	WLocks      int64
	ElapsedTime string
	WaitTime    string
	Info        string
}

// User on a TM1 server.
type User struct {
	Name         string
	FriendlyName string
	Password     string
	Type         string
	IsActive     bool
	Enabled      bool
}

// ProcessExecuteResult returns the result of a TurboIntegrator process execution.
type ProcessExecuteResult struct {
	ProcessExecuteStatusCode string
}
