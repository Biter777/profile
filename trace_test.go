package profile_test

import "github.com/biter777/profile"

func ExampleTraceProfile() {
	// use execution tracing, rather than the default cpu profiling.
	defer profile.Start(profile.TraceProfile).Stop()
}
