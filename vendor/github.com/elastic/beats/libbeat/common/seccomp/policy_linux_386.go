package seccomp

import (
	"github.com/elastic/go-seccomp-bpf"
)

func init() {
	defaultPolicy = &seccomp.Policy{
		DefaultAction: seccomp.ActionErrno,
		Syscalls: []seccomp.SyscallGroup{
			{
				Action: seccomp.ActionAllow,
				Names: []string{
					"_llseek",
					"access",
					"brk",
					"clock_gettime",
					"clone",
					"close",
					"dup",
					"dup2",
					"epoll_create",
					"epoll_create1",
					"epoll_ctl",
					"epoll_wait",
					"exit",
					"exit_group",
					"fchdir",
					"fchmod",
					"fchown32",
					"fcntl",
					"fdatasync",
					"flock",
					"fstat64",
					"fsync",
					"ftruncate64",
					"futex",
					"getcwd",
					"getdents",
					"getdents64",
					"geteuid32",
					"getgid32",
					"getpid",
					"getppid",
					"getrandom",
					"getrusage",
					"gettid",
					"gettimeofday",
					"getuid32",
					"inotify_add_watch",
					"inotify_init1",
					"inotify_rm_watch",
					"ioctl",
					"kill",
					"lstat64",
					"madvise",
					"mincore",
					"mkdirat",
					"mmap2",
					"mprotect",
					"munmap",
					"nanosleep",
					"open",
					"openat",
					"pipe",
					"pipe2",
					"poll",
					"pread64",
					"pselect6",
					"pwrite64",
					"read",
					"readlink",
					"readlinkat",
					"rename",
					"renameat",
					"restart_syscall",
					"rt_sigaction",
					"rt_sigprocmask",
					"rt_sigreturn",
					"sched_getaffinity",
					"sched_yield",
					"sendfile64",
					"set_robust_list",
					"set_thread_area",
					"setgid32",
					"setgroups32",
					"setitimer",
					"setuid32",
					"sigaltstack",
					"socketcall",
					"stat",
					"stat64",
					"statfs64",
					"sysinfo",
					"tgkill",
					"time",
					"tkill",
					"uname",
					"unlink",
					"unlinkat",
					"wait4",
					"waitid",
					"write",
					"writev",
				},
			},
		},
	}
}