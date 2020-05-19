package manager

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/runtime"
	"github.com/micro/go-micro/v2/store"
	"github.com/micro/micro/v2/internal/namespace"
)

type eventType string

const (
	eventTypeCreated eventType = "created"
	eventTypeUpdated eventType = "updated"
	eventTypeDeleted eventType = "deleted"
)

var (
	// eventTTL is the duration events will perist in the store before expiring
	eventTTL = time.Minute * 10
	// eventPollFrequency is the max frequency the manager will check for new events in the store
	eventPollFrequency = time.Minute
)

// eventPrefix is prefixed to the key for event records
const eventPrefix = "event/"

// event is the object written to the store
type event struct {
	ID      string                 `json:"id"`
	Type    eventType              `json:"type"`
	Service *runtime.Service       `json:"service"`
	Options *runtime.CreateOptions `json:"options"`
}

// key the event should be written to in the store
func (e *event) Key() string {
	return eventPrefix + e.ID
}

// publishEvent will write the event to the global store and immediately process the event
func (m *manager) publishEvent(eType eventType, srv *runtime.Service, opts *runtime.CreateOptions) error {
	e := &event{
		ID:      uuid.New().String(),
		Type:    eType,
		Service: srv,
		Options: opts,
	}

	bytes, err := json.Marshal(e)
	if err != nil {
		return err
	}

	record := &store.Record{
		Key:    e.Key(),
		Value:  bytes,
		Expiry: eventTTL,
	}

	if err := m.options.Store.Write(record); err != nil {
		return err
	}

	go m.processEvent(e.ID)
	return nil
}

// watchEvents polls the store for events periodically and processes
// them if they have not already done so
func (m *manager) watchEvents() {
	ticker := time.NewTicker(eventPollFrequency)

	for {
		<-ticker.C

		// get the keys of the events
		eventKeys, err := m.options.Store.List(store.ListPrefix(eventPrefix))
		if err != nil {
			logger.Warn("Error listing events: %v", err)
			continue
		}

		// loop through every event
		for _, key := range eventKeys {
			m.processEvent(key)
		}
	}
}

// processEvent will take an event key, verify it hasn't been consumed
// and then execute it. We pass the key and not the ID since the global
// store and the memory store use the same key prefix so there is not point
// stripping and then re-prefixing.
func (m *manager) processEvent(key string) {
	// check to see if the event has been processed before
	if _, err := m.cache.Read(key); err != store.ErrNotFound {
		return
	}

	// lookup the event
	recs, err := m.options.Store.Read(key)
	if err != nil {
		logger.Warnf("Error finding event %v: %v", key, err)
		return
	}
	var ev *event
	if err := json.Unmarshal(recs[0].Value, &ev); err != nil {
		logger.Warnf("Error unmarshaling event %v: %v", key, err)
	}

	// determine the namespace
	ns := ev.Options.Namespace
	if len(ns) == 0 {
		ns = namespace.DefaultNamespace
	}

	// log the event
	logger.Infof("Procesing %v event for service %v:%v in namespace %v", ev.Type, ev.Service.Name, ev.Service.Version, ns)

	// apply the event to the managed runtime
	switch ev.Type {
	case eventTypeDeleted:
		err = m.Runtime.Delete(ev.Service, runtime.DeleteNamespace(ns))
	case "update":
		err = m.Runtime.Update(ev.Service, runtime.UpdateNamespace(ns))
	case "create":
		err = m.Runtime.Create(ev.Service,
			runtime.CreateImage(ev.Options.Image),
			runtime.CreateType(ev.Options.Type),
			runtime.CreateNamespace(ns),
			runtime.WithArgs(ev.Options.Args...),
			runtime.WithCommand(ev.Options.Command...),
			runtime.WithEnv(m.runtimeEnv(ev.Options)),
		)
	}

	// if there was an error update the status in the cache
	if err != nil {
		logger.Warnf("Error procesing %v event for service %v:%v in namespace %v: %v,", ev.Type, ev.Service.Name, ev.Service.Version, ns, err)
		ev.Service.Metadata = map[string]string{"status": "error", "error": err.Error()}
		m.cacheStatus(ns, ev.Service)
	}

	// write to the store indicating the event has been consumed. We double the ttl to safely know the
	// event will expire before this record
	m.cache.Write(&store.Record{Key: key, Expiry: eventTTL * 2})
}

// runtimeEnv returns the environment variables which should  be used when creating a service.
func (m *manager) runtimeEnv(options *runtime.CreateOptions) []string {
	setEnv := func(p []string, env map[string]string) {
		for _, v := range p {
			parts := strings.Split(v, "=")
			if len(parts) <= 1 {
				continue
			}
			env[parts[0]] = strings.Join(parts[1:], "=")
		}
	}

	// overwrite any values
	env := map[string]string{}

	// set the env vars provided
	setEnv(options.Env, env)

	// override with vars from the Profile
	setEnv(m.options.Profile, env)

	// temp: set the auth namespace. this will be removed once
	// the namespace can be determined from certs.
	if len(options.Namespace) > 0 {
		env["MICRO_AUTH_NAMESPACE"] = options.Namespace
	}

	// create a new env
	var vars []string
	for k, v := range env {
		vars = append(vars, k+"="+v)
	}

	// setup the runtime env
	return vars
}
