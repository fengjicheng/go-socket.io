package engineio

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"go-socket.io/engineio/session"
	"go-socket.io/engineio/transport"
	"go-socket.io/logger"
)

// Server is instance of server
type Server struct {
	corsHandlingDisabled bool
	allowedCorsOrigins   []string

	pingInterval time.Duration
	pingTimeout  time.Duration

	transports *transport.Manager
	sessions   *session.Manager

	requestChecker CheckerFunc
	connInitiator  ConnInitiatorFunc

	connChan  chan Conn
	closeOnce sync.Once

	clients  map[string]any
	shutdown chan bool
}

type HandshakeInterceptor interface {
	Intercept(query map[string]string, headers http.Header) bool
}

// NewServer returns a server.
func NewServer(opts *Options) *Server {
	return &Server{
		pingInterval:   opts.getPingInterval(),
		pingTimeout:    opts.getPingTimeout(),
		transports:     transport.NewManager(opts.getTransport()),
		requestChecker: opts.getRequestChecker(),
		connInitiator:  opts.getConnConnInitiator(),
		sessions:       session.NewManager(opts.getSessionIDGenerator()),
		connChan:       make(chan Conn, 1),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//s.cors(r)
	query := r.URL.Query()
	params, _ := json.Marshal(query)
	logger.Debug(fmt.Sprintf("request: %s", string(params)))

	//reqTransport := query.Get("transport")
	////if reqTransport != "polling" {
	////	s.sendErrorMessage(w, UnknownTransport, http.StatusBadRequest)
	////	return
	////}
	//
	////sid := query.Get("sid")
	//
	//srvTransport, ok := s.transports.Get(reqTransport)
	//if !ok {
	//	invalidTransport := fmt.Sprintf("invalid transport: %s", reqTransport)
	//	logger.Warn(invalidTransport)
	//	http.Error(w, invalidTransport, http.StatusBadRequest)
	//	return
	//}
	//
	//header, err := s.requestChecker(r)
	//if err != nil {
	//	logger.Error("request checker err", err)
	//	http.Error(w, fmt.Sprintf("request checker err: %s", err.Error()), http.StatusBadGateway)
	//	return
	//}
	//
	//for k, v := range header {
	//	w.Header()[k] = v
	//}
	//
	//sid := query.Get("sid")
	//reqSession, ok := s.sessions.Get(sid)
	//// if we can't find session in current session pool, let's create this. behaviour for new connections
	//if !ok {
	//	if sid != "" {
	//		http.Error(w, fmt.Sprintf("invalid sid value: %s", sid), http.StatusBadRequest)
	//		return
	//	}
	//
	//	transportConn, err := srvTransport.Accept(w, r)
	//	if err != nil {
	//		http.Error(w, fmt.Sprintf("transport accept err: %s", err.Error()), http.StatusBadGateway)
	//		return
	//	}
	//
	//	reqSession, err = s.newSession(r.Context(), transportConn, reqTransport)
	//	if err != nil {
	//		http.Error(w, fmt.Sprintf("create new session err: %s", err.Error()), http.StatusBadRequest)
	//		return
	//	}
	//
	//	s.connInitiator(r, reqSession)
	//}
	//
	//// try upgrade current connection
	//if reqSession.Transport() != reqTransport {
	//	//	transportConn, err := srvTransport.Accept(w, r)
	//	//	if err != nil {
	//	//		// don't call http.Error() for HandshakeErrors because
	//	//		// they get handled by the websocket library internally.
	//	//		if _, ok := err.(websocket.HandshakeError); !ok {
	//	//			http.Error(w, err.Error(), http.StatusBadGateway)
	//	//		}
	//	//		return
	//	//	}
	//	//
	//	//	reqSession.Upgrade(reqTransport, transportConn)
	//	//
	//	//	if handler, ok := transportConn.(http.Handler); ok {
	//	//		handler.ServeHTTP(w, r)
	//	//	}
	//	return
	//}
	//
	//reqSession.ServeHTTP(w, r)
}

func (s *Server) Shutdown() {
	s.shutdown <- true
}

func (s *Server) close() {
	for c := range s.clients {
		delete(s.clients, c)
	}
	s.closeOnce.Do(func() {
		close(s.connChan)
		close(s.shutdown)
	})
}

func (s *Server) newSession(context context.Context, conn transport.Conn, reqTransport string) (*session.Session, error) {
	params := transport.ConnParameters{
		PingInterval: s.pingInterval,
		PingTimeout:  s.pingTimeout,
		Upgrades:     s.transports.UpgradeFrom(reqTransport),
	}

	sid := s.sessions.NewID()
	newSession, err := session.New(conn, sid, reqTransport, params)
	if err != nil {
		return nil, err
	}

	go func(newSession *session.Session) {
		if err = newSession.InitSession(); err != nil {
			log.Println("init new session:", err)
			return
		}

		s.sessions.Add(newSession)

		s.connChan <- newSession
	}(newSession)

	return newSession, nil
}

func (s *Server) server() {
	log.Print("server started.")
	for {
		select {
		case conn := <-s.connChan:
			// todo
			//go s.serveConn(conn)
			log.Println(conn)
		case <-s.shutdown:
			s.close()
			logger.Info("server stopped.")
			return
		}
	}
}

func (s *Server) cors(r *http.Request) {
	if !s.corsHandlingDisabled {
		origin := r.Header.Get("Origin")
		sendCors := false
		for _, corsOrigin := range s.allowedCorsOrigins {
			if corsOrigin != "" && origin == corsOrigin {
				sendCors = true
				break
			}
		}
		if sendCors {
			r.Header.Add("Access-Control-Allow-Origin", origin)
			r.Header.Add("Access-Control-Allow-Credentials", "true")
			r.Header.Add("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
			r.Header.Add("Access-Control-Allow-Headers", "origin, content-type, accept")
		}
	}
}

func (s *Server) sendErrorMessage(response http.ResponseWriter, err ErrorCode, statusCode int) {
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Content-Encoding", "UTF-8")
	response.WriteHeader(statusCode)
	_ = json.NewEncoder(response).Encode(err)
}
