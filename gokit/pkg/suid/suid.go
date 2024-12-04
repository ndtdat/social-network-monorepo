package suid

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	bitLenTime      = 39                               // bit length of time
	bitLenSequence  = 8                                // bit length of sequence number
	bitLenMachineID = 63 - bitLenTime - bitLenSequence // bit length of machine id
	maskSequence    = uint16(1<<bitLenSequence - 1)
	maxRandomValue  = 1 << bitLenTime
	minInitSleep    = 10
	maxInitSleep    = 30
)

type SUID struct {
	mutex       *sync.Mutex
	startTime   int64
	elapsedTime int64
	sequence    uint16
	machineID   byte
}

func NewSUID(startTime time.Time, machineID byte) (*SUID, error) {
	return &SUID{
		mutex:       new(sync.Mutex),
		startTime:   toSUIDTime(time.Date(2023, 6, 6, 6, 6, 6, 6, time.UTC)),
		elapsedTime: 0,
		sequence:    maskSequence,
		machineID:   machineID,
	}, nil
}

func (s *SUID) NextID() (uint64, error) {
	const maskSequence = uint16(1<<bitLenSequence - 1)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	current := currentElapsedTime(s.startTime)
	if s.elapsedTime < current {
		s.elapsedTime = current
		s.sequence = 0
	} else { // s.elapsedTime >= current
		s.sequence = (s.sequence + 1) & maskSequence
		if s.sequence == 0 {
			s.elapsedTime++
			overtime := s.elapsedTime - current
			time.Sleep(sleepTime(overtime))
		}
	}

	return s.toID()
}

const suidTimeUnit = 1e7 // nsec, i.e. 10 msec

func toSUIDTime(t time.Time) int64 {
	return t.UTC().UnixNano() / suidTimeUnit
}

func currentElapsedTime(startTime int64) int64 {
	return toSUIDTime(time.Now()) - startTime
}

func sleepTime(overtime int64) time.Duration {
	return time.Duration(overtime*suidTimeUnit) -
		time.Duration(time.Now().UTC().UnixNano()%suidTimeUnit)
}

func (s *SUID) toID() (uint64, error) {
	if s.elapsedTime >= maxRandomValue {
		return 0, errors.New("over the time limit")
	}

	return uint64(s.elapsedTime)<<(bitLenSequence+bitLenMachineID) |
		uint64(s.sequence)<<bitLenMachineID |
		uint64(s.machineID), nil
}

func elapsedTime(id uint64) uint64 {
	return id >> (bitLenSequence + bitLenMachineID)
}

func SequenceNumber(id uint64) uint64 {
	const maskSequence = uint64((1<<bitLenSequence - 1) << bitLenMachineID)

	return id & maskSequence >> bitLenMachineID
}

func MachineID(id uint64) uint64 {
	const maskMachineID = uint64(1<<bitLenMachineID - 1)

	return id & maskMachineID
}

func Decompose(id uint64) map[string]uint64 {
	msb := id >> 63
	time := elapsedTime(id)
	sequence := SequenceNumber(id)
	machineID := MachineID(id)

	return map[string]uint64{
		"id":         id,
		"msb":        msb,
		"time":       time,
		"sequence":   sequence,
		"machine-id": machineID,
	}
}

func privateIPv4(interfaceAddrs InterfaceAddrs) (net.IP, error) {
	as, err := interfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}

	return nil, fmt.Errorf("no private IP addresss")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func lower16BitPrivateIP(interfaceAddrs InterfaceAddrs) (uint16, error) {
	ip, err := privateIPv4(interfaceAddrs)
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}
