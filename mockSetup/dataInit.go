package mocksetup

type DataModel struct {
	IP       string
	HostName string
	Active   bool
}

var Data []DataModel

// InitData function is use to initialise example data set
func InitData() {
	Data = []DataModel{
		{
			IP:       "127.0.0.1",
			HostName: "mta-prod-1",
			Active:   true,
		},
		{
			IP:       "127.0.0.2",
			HostName: "mta-prod-1",
			Active:   false,
		},
		{
			IP:       "127.0.0.3",
			HostName: "mta-prod-2",
			Active:   true,
		},
		{
			IP:       "127.0.0.4",
			HostName: "mta-prod-2",
			Active:   true,
		},
		{
			IP:       "127.0.0.5",
			HostName: "mta-prod-2",
			Active:   false,
		},
		{
			IP:       "127.0.0.6",
			HostName: "mta-prod-3",
			Active:   false,
		},
	}
}

// AddData is an exposed function for new data set
func AddData(data []DataModel) {

	//Data=append(Data, data...)		//please un-comment this line is you want to append data with existing data
	Data = data //please comment this line to append the data
}
