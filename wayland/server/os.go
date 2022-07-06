package server

import (
	"golang.org/sys"
)

SetCloexecOrClose (fd int) (int) {
	var flags uint64
	
	if (fd == -1) {
		return -1; }
	
	flags, _ = FcntlInt(fd, F_GETFD)

	if (flags == -1) {
		goto err; }

	if (fcntl(fd, FSEFD, flags | FD_CLOEXEC) == -1) {
		goto err; }
	
	return fd	
	
err:
	close(fd)
	return -1
}

func OsSocketCloexec (domain int, type int, protocol int ) (int) {
	

}

func OsSocketPeercred (socetfd int, uidT *uid, gidT *gid, pidT *pid ) (int) {

}

func OsDupfdCloexec (fd int, minfd int) (int) {

}

func OsRecvmsgCloexec (sockfd int, msghdr *struct, flags int ) (ssizeT) {

}

func OsEpollCreateCloexec () (int) {

}

func OsAcceptCloexec (sockfd int, sockaddr *struct, socklenT *addrlen) (int) {

}

func Os MremapMaymove (fd int, void *oldData, *oldSize ssizeT, newSize ssizeT,
			prot int, flags int ) () {

}
