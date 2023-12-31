package ewelink

import "fmt"

// Session holds the session metadata.
type Session struct {
	MobileDevice        MobileDevice
	Application         *Application
	AuthenticationToken string
	User                *User
	Configuration       *Configuration
}

func (s Session) String() string {
	return fmt.Sprintf("%#v", s)
}

func (s *Session) updateTokenAndResponse(response *AuthenticationResponse) {
	s.AuthenticationToken = response.At
	s.User = &response.User
}

// SessionOptionFunc option definition.
type SessionOptionFunc func(session *Session)

// Function to create SessionOptionFunc func to set custom mobile device.
// nolint:deadcode,unused
func withMobileDevice(device MobileDevice) SessionOptionFunc {
	return func(subject *Session) {
		subject.MobileDevice = device
	}
}

// Function to create SessionOptionFunc func to set custom application.
// nolint:deadcode,unused
func withApplication(application *Application) SessionOptionFunc {
	return func(subject *Session) {
		subject.Application = application
	}
}
