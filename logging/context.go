package logging

// ContextLogger - Struct with one field representing the "context" from which logging
//	is occurring. For each struct method, call the logging function with the corresponding
//	name and pass this struct's "context" as the first parameter.
type ContextLogger struct {
	context string
}

// GetContextLogger - Instantiates a ContextLogger with the given "context", and
//	returns it. This is used such that "context" can remain private throughout
//	the lifetime of the ContextLogger.
func GetContextLogger(context string) ContextLogger {
	return ContextLogger{context}
}

//Printf print the INFO log
func (l *ContextLogger) Printf(s string, a ...interface{}) {
	ContextPrintf(l.context, s, a...)
}

//ErrorPrintf print the ERROR log
func (l *ContextLogger) ErrorPrintf(s string, a ...interface{}) {
	ErrorPrintf(l.context, s, a...)
}

//ServicePrintf print the INFO log for service context
func (l *ContextLogger) ServicePrintf(s string, a ...interface{}) {
	ServicePrintf(l.context, s, a...)
}

//ServiceErrorPrintf print the ERROR log for service context
func (l *ContextLogger) ServiceErrorPrintf(s string, a ...interface{}) {
	ServiceErrorPrintf(l.context, s, a...)
}

//EventPrintf print the INFO log for event context
func (l *ContextLogger) EventPrintf(s string, a ...interface{}) {
	EventPrintf(l.context, s, a...)
}

//EventErrorPrintf print the ERROR log for event context
func (l *ContextLogger) EventErrorPrintf(s string, a ...interface{}) {
	EventErrorPrintf(l.context, s, a...)
}

// ServiceContextLogger - Struct with one field representing the "context" from which
//	logging is occurring. For each struct method, we attach the "SERVICE" context and
//	pass this struct's "context" as the first parameter.
type ServiceContextLogger struct {
	context string
}

// GetServiceContextLogger - Instantiates a ServiceContextLogger with the given "context",
//	and returns it. This is used such that "context" can remain private throughout
//	the lifetime of the ServiceContextLogger.
func GetServiceContextLogger(context string) ServiceContextLogger {
	return ServiceContextLogger{context}
}

func (l *ServiceContextLogger) Printf(s string, a ...interface{}) {
	ServicePrintf(l.context, s, a...)
}

func (l *ServiceContextLogger) ErrorPrintf(s string, a ...interface{}) {
	ServiceErrorPrintf(l.context, s, a...)
}

// EventContextLogger - Struct with one field representing the "context" from which
//	logging is occurring. For each struct method, we attach the "EVENT" context and
//	pass this struct's "context" as the first parameter.
type EventContextLogger struct {
	context string
}

// GetEventContextLogger - Instantiates an EventContextLogger with the given "context",
//	and returns it. This is used such that "context" can remain private throughout
//	the lifetime of the EventContextLogger.
func GetEventContextLogger(context string) EventContextLogger {
	return EventContextLogger{context}
}

func (l *EventContextLogger) Printf(s string, a ...interface{}) {
	EventPrintf(l.context, s, a...)
}

func (l *EventContextLogger) ErrorPrintf(s string, a ...interface{}) {
	EventErrorPrintf(l.context, s, a...)
}

// ControllerContextLogger - Struct with one field representing the "context" from which
//	logging is occurring. For each struct method, we attach the "EVENT" context and
//	pass this struct's "context" as the first parameter.
type ControllerContextLogger struct {
	context string
}

// GetControllerContextLogger - Instantiates an ControllerContextLogger with the given "context",
//	and returns it. This is used such that "context" can remain private throughout
//	the lifetime of the ControllerContextLogger.
func GetControllerContextLogger(context string) ControllerContextLogger {
	return ControllerContextLogger{context}
}

func (l *ControllerContextLogger) Printf(s string, a ...interface{}) {
	ControllerPrintf(l.context, s, a...)
}

func (l *ControllerContextLogger) ErrorPrintf(s string, a ...interface{}) {
	ControllerErrorPrintf(l.context, s, a...)
}
