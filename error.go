package pj

import "github.com/pidato/pjproject-go/pjsua2"

/*
SetStatus(arg2 int)
	GetStatus() (_swig_ret int)
	SetTitle(arg2 string)
	GetTitle() (_swig_ret string)
	SetReason(arg2 string)
	GetReason() (_swig_ret string)
	SetSrcFile(arg2 string)
	GetSrcFile() (_swig_ret string)
	SetSrcLine(arg2 int)
	GetSrcLine() (_swig_ret int)
*/
type Error struct {
	Status  int
	Title   string
	Reason  string
	SrcFile string
	SrcLine int
}

func (e *Error) Error() string {
	return e.Title
}

func exec(fn func()) (out error) {
	defer func() {
		e := recover()
		if e != nil {
			switch t := e.(type) {
			case pjsua2.Error:
				e = t
				out = &Error{
					Status:  t.GetStatus(),
					Title:   t.GetTitle(),
					Reason:  t.GetReason(),
					SrcFile: t.GetSrcFile(),
					SrcLine: t.GetSrcLine(),
				}
			case error:
				out = t
			}
		}
	}()
	out = nil
	fn()
	return
}
