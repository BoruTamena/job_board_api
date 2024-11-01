package logger

type CfgOptions struct {
	// options
	LevelSplit bool // default false
	// a min log level. default info-lvl
	Atomiclevel Level
	// log output path.default stdout
	Out_path string
	// url or folder path to write internal error to.
	// default stderr
	Err_path string
	// enable development mode, defaul value is false
	Dev Mode
}
