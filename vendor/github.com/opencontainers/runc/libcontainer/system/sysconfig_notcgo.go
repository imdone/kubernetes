// +build !cgo windows

package system

func GetClockTicks() int {
	// TODO figure out a better alternative for platforms where we're missing cgo id:3200 gh:3215
	//
	// TODO Windows. This could be implemented using Win32 QueryPerformanceFrequency(). id:2778 gh:2793
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms644905(v=vs.85).aspx
	//
	// An example of its usage can be found here.
	// https://msdn.microsoft.com/en-us/library/windows/desktop/dn553408(v=vs.85).aspx

	return 100
}
