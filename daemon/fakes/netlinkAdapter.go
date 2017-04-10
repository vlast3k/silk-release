// This file was generated by counterfeiter
package fakes

import (
	"net"
	"sync"

	"github.com/vishvananda/netlink"
)

type NetlinkAdapter struct {
	LinkSetUpStub        func(netlink.Link) error
	linkSetUpMutex       sync.RWMutex
	linkSetUpArgsForCall []struct {
		arg1 netlink.Link
	}
	linkSetUpReturns struct {
		result1 error
	}
	linkSetUpReturnsOnCall map[int]struct {
		result1 error
	}
	LinkAddStub        func(netlink.Link) error
	linkAddMutex       sync.RWMutex
	linkAddArgsForCall []struct {
		arg1 netlink.Link
	}
	linkAddReturns struct {
		result1 error
	}
	linkAddReturnsOnCall map[int]struct {
		result1 error
	}
	LinkSetHardwareAddrStub        func(netlink.Link, net.HardwareAddr) error
	linkSetHardwareAddrMutex       sync.RWMutex
	linkSetHardwareAddrArgsForCall []struct {
		arg1 netlink.Link
		arg2 net.HardwareAddr
	}
	linkSetHardwareAddrReturns struct {
		result1 error
	}
	linkSetHardwareAddrReturnsOnCall map[int]struct {
		result1 error
	}
	AddrAddScopeLinkStub        func(link netlink.Link, addr *netlink.Addr) error
	addrAddScopeLinkMutex       sync.RWMutex
	addrAddScopeLinkArgsForCall []struct {
		link netlink.Link
		addr *netlink.Addr
	}
	addrAddScopeLinkReturns struct {
		result1 error
	}
	addrAddScopeLinkReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NetlinkAdapter) LinkSetUp(arg1 netlink.Link) error {
	fake.linkSetUpMutex.Lock()
	ret, specificReturn := fake.linkSetUpReturnsOnCall[len(fake.linkSetUpArgsForCall)]
	fake.linkSetUpArgsForCall = append(fake.linkSetUpArgsForCall, struct {
		arg1 netlink.Link
	}{arg1})
	fake.recordInvocation("LinkSetUp", []interface{}{arg1})
	fake.linkSetUpMutex.Unlock()
	if fake.LinkSetUpStub != nil {
		return fake.LinkSetUpStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkSetUpReturns.result1
}

func (fake *NetlinkAdapter) LinkSetUpCallCount() int {
	fake.linkSetUpMutex.RLock()
	defer fake.linkSetUpMutex.RUnlock()
	return len(fake.linkSetUpArgsForCall)
}

func (fake *NetlinkAdapter) LinkSetUpArgsForCall(i int) netlink.Link {
	fake.linkSetUpMutex.RLock()
	defer fake.linkSetUpMutex.RUnlock()
	return fake.linkSetUpArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) LinkSetUpReturns(result1 error) {
	fake.LinkSetUpStub = nil
	fake.linkSetUpReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetUpReturnsOnCall(i int, result1 error) {
	fake.LinkSetUpStub = nil
	if fake.linkSetUpReturnsOnCall == nil {
		fake.linkSetUpReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkSetUpReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkAdd(arg1 netlink.Link) error {
	fake.linkAddMutex.Lock()
	ret, specificReturn := fake.linkAddReturnsOnCall[len(fake.linkAddArgsForCall)]
	fake.linkAddArgsForCall = append(fake.linkAddArgsForCall, struct {
		arg1 netlink.Link
	}{arg1})
	fake.recordInvocation("LinkAdd", []interface{}{arg1})
	fake.linkAddMutex.Unlock()
	if fake.LinkAddStub != nil {
		return fake.LinkAddStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkAddReturns.result1
}

func (fake *NetlinkAdapter) LinkAddCallCount() int {
	fake.linkAddMutex.RLock()
	defer fake.linkAddMutex.RUnlock()
	return len(fake.linkAddArgsForCall)
}

func (fake *NetlinkAdapter) LinkAddArgsForCall(i int) netlink.Link {
	fake.linkAddMutex.RLock()
	defer fake.linkAddMutex.RUnlock()
	return fake.linkAddArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) LinkAddReturns(result1 error) {
	fake.LinkAddStub = nil
	fake.linkAddReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkAddReturnsOnCall(i int, result1 error) {
	fake.LinkAddStub = nil
	if fake.linkAddReturnsOnCall == nil {
		fake.linkAddReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkAddReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetHardwareAddr(arg1 netlink.Link, arg2 net.HardwareAddr) error {
	fake.linkSetHardwareAddrMutex.Lock()
	ret, specificReturn := fake.linkSetHardwareAddrReturnsOnCall[len(fake.linkSetHardwareAddrArgsForCall)]
	fake.linkSetHardwareAddrArgsForCall = append(fake.linkSetHardwareAddrArgsForCall, struct {
		arg1 netlink.Link
		arg2 net.HardwareAddr
	}{arg1, arg2})
	fake.recordInvocation("LinkSetHardwareAddr", []interface{}{arg1, arg2})
	fake.linkSetHardwareAddrMutex.Unlock()
	if fake.LinkSetHardwareAddrStub != nil {
		return fake.LinkSetHardwareAddrStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkSetHardwareAddrReturns.result1
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrCallCount() int {
	fake.linkSetHardwareAddrMutex.RLock()
	defer fake.linkSetHardwareAddrMutex.RUnlock()
	return len(fake.linkSetHardwareAddrArgsForCall)
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrArgsForCall(i int) (netlink.Link, net.HardwareAddr) {
	fake.linkSetHardwareAddrMutex.RLock()
	defer fake.linkSetHardwareAddrMutex.RUnlock()
	return fake.linkSetHardwareAddrArgsForCall[i].arg1, fake.linkSetHardwareAddrArgsForCall[i].arg2
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrReturns(result1 error) {
	fake.LinkSetHardwareAddrStub = nil
	fake.linkSetHardwareAddrReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrReturnsOnCall(i int, result1 error) {
	fake.LinkSetHardwareAddrStub = nil
	if fake.linkSetHardwareAddrReturnsOnCall == nil {
		fake.linkSetHardwareAddrReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkSetHardwareAddrReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) AddrAddScopeLink(link netlink.Link, addr *netlink.Addr) error {
	fake.addrAddScopeLinkMutex.Lock()
	ret, specificReturn := fake.addrAddScopeLinkReturnsOnCall[len(fake.addrAddScopeLinkArgsForCall)]
	fake.addrAddScopeLinkArgsForCall = append(fake.addrAddScopeLinkArgsForCall, struct {
		link netlink.Link
		addr *netlink.Addr
	}{link, addr})
	fake.recordInvocation("AddrAddScopeLink", []interface{}{link, addr})
	fake.addrAddScopeLinkMutex.Unlock()
	if fake.AddrAddScopeLinkStub != nil {
		return fake.AddrAddScopeLinkStub(link, addr)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addrAddScopeLinkReturns.result1
}

func (fake *NetlinkAdapter) AddrAddScopeLinkCallCount() int {
	fake.addrAddScopeLinkMutex.RLock()
	defer fake.addrAddScopeLinkMutex.RUnlock()
	return len(fake.addrAddScopeLinkArgsForCall)
}

func (fake *NetlinkAdapter) AddrAddScopeLinkArgsForCall(i int) (netlink.Link, *netlink.Addr) {
	fake.addrAddScopeLinkMutex.RLock()
	defer fake.addrAddScopeLinkMutex.RUnlock()
	return fake.addrAddScopeLinkArgsForCall[i].link, fake.addrAddScopeLinkArgsForCall[i].addr
}

func (fake *NetlinkAdapter) AddrAddScopeLinkReturns(result1 error) {
	fake.AddrAddScopeLinkStub = nil
	fake.addrAddScopeLinkReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) AddrAddScopeLinkReturnsOnCall(i int, result1 error) {
	fake.AddrAddScopeLinkStub = nil
	if fake.addrAddScopeLinkReturnsOnCall == nil {
		fake.addrAddScopeLinkReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addrAddScopeLinkReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.linkSetUpMutex.RLock()
	defer fake.linkSetUpMutex.RUnlock()
	fake.linkAddMutex.RLock()
	defer fake.linkAddMutex.RUnlock()
	fake.linkSetHardwareAddrMutex.RLock()
	defer fake.linkSetHardwareAddrMutex.RUnlock()
	fake.addrAddScopeLinkMutex.RLock()
	defer fake.addrAddScopeLinkMutex.RUnlock()
	return fake.invocations
}

func (fake *NetlinkAdapter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
